package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nullxjx/chatgpt-server-go/cmd/chatgpt/app/path"
	"github.com/nullxjx/chatgpt-server-go/cmd/chatgpt/config"
	"github.com/sashabaranov/go-openai"
)

type App struct {
	HttpEngine   *gin.Engine
	OpenAIClient *openai.Client
	Config       *config.Config
}

func (app *App) Conf() *config.Config {
	return app.Config
}

func (app *App) GetClient() *openai.Client {
	return app.OpenAIClient
}

// initGinApplication 初始化gin
func (app *App) initGinApplication() error {
	app.HttpEngine = gin.New()
	if err := path.RegisterHttp(app.HttpEngine, app); err != nil {
		return fmt.Errorf("register http error: %w", err)
	}
	return nil
}

// New 用于新建App实例
func New(configPath string) (*App, error) {
	cfg, err := config.New(configPath)
	if err != nil {
		return nil, err
	}

	openaiConfig := openai.DefaultConfig(cfg.OpenAIKey)
	openaiConfig.BaseURL = cfg.BaseURL
	client := openai.NewClientWithConfig(openaiConfig)
	app := &App{
		Config:       cfg,
		OpenAIClient: client,
	}

	if err = app.initGinApplication(); err != nil {
		return nil, fmt.Errorf("init gin application error: %w", err)
	}
	return app, nil
}
