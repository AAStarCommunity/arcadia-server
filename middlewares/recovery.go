package middlewares

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

// GenericRecovery generic recovery
func GenericRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(&PanicExceptionRecord{})
}

// PanicExceptionRecord  panic等异常记录
type PanicExceptionRecord struct{}

func (p *PanicExceptionRecord) Write(b []byte) (n int, err error) {
	s1 := "An error occurred in the server's internal code："
	var build strings.Builder
	build.WriteString(s1)
	build.Write(b)
	errStr := build.String()
	return len(errStr), errors.New(errStr)
}
