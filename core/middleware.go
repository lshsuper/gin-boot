package core

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"strings"
)

func TraceID(traceIDKey string)gin.HandlerFunc  {

	if len(traceIDKey)<=0{
		traceIDKey=getTraceIDKey()
	}else{
		setTraceIDKey(traceIDKey)
	}
	return func(context *gin.Context) {
		traceId:=context.GetHeader(traceIDKey)
		if len(traceId)<=0{
			traceId=strings.ReplaceAll(uuid.NewV4().String(),"-","")
			context.Request.Header.Add(traceIDKey,traceId)
		}

	}
}
