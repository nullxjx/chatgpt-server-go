package path

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nullxjx/chatgpt-server-go/cmd/chatgpt/app/path/api/completion"
	"github.com/nullxjx/chatgpt-server-go/cmd/chatgpt/config"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
)

type Dependency interface {
	Conf() *config.Config
	GetClient() *openai.Client
}

// RegisterHttp 注册http接口。
func RegisterHttp(app *gin.Engine, dep Dependency) error {
	app.Use(gin.Recovery())

	app.GET("/ping", func(ctx *gin.Context) {
		log.Infof("ping pong")
		ctx.String(http.StatusOK, "pong")
	})
	apiGroup := app.Group("/v1")
	apiGroup.POST("/completions", func(ctx *gin.Context) {
		completion.Completion(ctx, dep)
	})
	apiGroup.POST("/chat/completions", func(ctx *gin.Context) {
		completion.Chat(ctx, dep)
	})
	return nil
}
