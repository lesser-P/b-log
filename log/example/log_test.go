package example

import (
	"context"
	"github.com/lesser-P/b-log/log"
	"github.com/rs/zerolog"
	"testing"
)

func TestExampleBergLog(t *testing.T) {
	ctx := context.WithValue(context.Background(), "traceId", "test")
	log.InitBergLog("stdout", zerolog.InfoLevel)
	data := struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}{
		UserName: "berg",
		Password: "aaxks",
	}
	log.Interface(ctx, zerolog.ErrorLevel, data, "this is %s a test log ", "haha")
}
