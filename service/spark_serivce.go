package service

import (
	"GoGin/conf"
	"GoGin/serializer"
	"GoGin/types"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/goccy/go-json"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/url"
	"strings"
	"sync"
	"time"
)

type SparkService struct {
}

var SparkServiceInstance *SparkService

var SparkServiceOnce sync.Once

func GetSparkServiceInstance() *SparkService {
	SparkServiceOnce.Do(func() {
		SparkServiceInstance = &SparkService{}
	})
	return SparkServiceInstance
}

// ConnToSpark 建立连接并发起会话
func (*SparkService) ConnToSpark(req *types.AIReq) *serializer.Response {
	config := conf.Conf.Spark
	fmt.Println(config)
	// 建立websocket连接
	dialer := websocket.Dialer{HandshakeTimeout: time.Second * 5}
	conn, resp, err := dialer.Dial(assembleAuthUrl(config.HostUrl, config.ApiKey, config.ApiSecret), nil)
	defer conn.Close()
	// 错误处理
	if err != nil {
		body, err1 := io.ReadAll(resp.Body)
		if err1 != nil {
			log.Println("io.ReadAll error", err)
			return &serializer.Response{
				Code:    serializer.SystemError,
				Message: serializer.ServerError,
			}
		}
		log.Println(string(body))
		return &serializer.Response{
			Code:    resp.StatusCode,
			Message: serializer.ServerError,
		}
	}

	// 发送请求
	go func() {
		req := buildReq(config.Appid, req.Data)
		err = conn.WriteJSON(req)
	}()
	// 获取响应
	answer := ""
	for {
		_, bytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("read message err:", err)
			return &serializer.Response{
				Code:    serializer.SystemError,
				Message: serializer.ServerError,
			}
		}
		var resp serializer.SparkResp
		err = json.Unmarshal(bytes, &resp)
		if err != nil {
			log.Println("json parse to resp err:", err)
			return &serializer.Response{
				Code:    serializer.SystemError,
				Message: serializer.ServerError,
			}
		}
		// 解析数据
		if resp.Header.Code != 0 {
			// 错误码非0
			log.Println(resp.Header.Message)
			return &serializer.Response{
				Code:    serializer.SystemError,
				Message: serializer.ServerError,
			}
		}
		content := (*resp.Payload.Choices.Text)[0].Content
		answer += content
		// status != 2 表示非最后文本结果
		if resp.Payload.Choices.Status == 2 {
			// 已收到最终结果
			return &serializer.Response{
				Code:    serializer.OK,
				Data:    answer,
				Message: serializer.Succeed,
			}
		}
	}
}

// 鉴权
func assembleAuthUrl(hostUrl, apiKey, apiSecret string) string {
	ul, err := url.Parse(hostUrl)
	if err != nil {
		log.Println("spark_service.assembleAuthUrl()转url失败：", err)
	}
	//签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	//date = "Tue, 28 May 2019 09:10:42 MST"
	//参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	sign := strings.Join(signString, "\n")
	// fmt.Println(sign)
	//签名结果
	sha := hmacWithShaToBase64("hmac-sha256", sign, apiSecret)
	// fmt.Println(sha)
	//构建请求参数 此时不需要urlencoding
	authUrl := fmt.Sprintf("hmac username=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey,
		"hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	callUrl := hostUrl + "?" + v.Encode()
	return callUrl
}

func hmacWithShaToBase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}

func buildReq(appId, str string) *serializer.SparkReq {
	return &serializer.SparkReq{
		Header: &serializer.Header{
			AppId: appId,
			Uid:   "1",
		},
		Parameter: &serializer.Parameter{Chat: &serializer.Chat{
			Domain:      "generalv2",
			Temperature: 0.5,
			MaxTokens:   2048,
		}},
		Payload: &serializer.Payload{Message: &serializer.Message{
			Text: &([]serializer.Text{{
				Role:    "user",
				Content: str},
			}),
		}},
	}
}
