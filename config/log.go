package config

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io"
	"path/filepath"
	"runtime"
)

// tabFormatter tab数据格式化
type logFormatter struct {
	log.TextFormatter
}

// Format自定义日志输出格式
func (c *logFormatter) Format(entry *log.Entry) ([]byte, error) {
	prettyCaller := func(frame *runtime.Frame) string {
		_, fileName := filepath.Split(frame.File)
		return fmt.Sprintf("%s:%d", fileName, frame.Line)
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	b.WriteString(fmt.Sprintf("[%s] %s", entry.Time.Format(c.TimestampFormat)))
	if entry.HasCaller() {
		b.WriteString(fmt.Sprintf("[%s]", prettyCaller(entry.Caller)))
	}
	b.WriteString(fmt.Sprintf(" %s\n", entry.Message)) // 输出日志内容
	return b.Bytes(), nil
}

func setGinLog(out io.Writer) {
	gin.DefaultWriter = out
	gin.DefaultErrorWriter = out

}
