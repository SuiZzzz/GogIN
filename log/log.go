package log

import (
	"io"
	"log"
	"os"
)

// 设置日期格式
func init() {
	file, _ := os.OpenFile("E:\\GoLearning\\GoGin\\log\\log.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	// 日志输出到控制台和日志文件
	multiWriter := io.MultiWriter(file, os.Stdout)
	log.SetOutput(multiWriter)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}
