package completion

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nullxjx/chatgpt-server-go/cmd/chatgpt/config"
	"github.com/nullxjx/chatgpt-server-go/internal/chatgpt/chat"
	"github.com/nullxjx/chatgpt-server-go/internal/chatgpt/completion"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
)

type Dependency interface {
	Conf() *config.Config
	GetClient() *openai.Client
}

func Completion(ctx *gin.Context, dep Dependency) {
	var req openai.CompletionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Infof("receive completion request, stream: %v, prompt: %v", req.Stream, req.Prompt)

	if req.Stream {
		stream, err := completion.StreamCompletion(ctx.Request.Context(), dep, &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.Header("Content-Type", "text/event-stream")
		ctx.Stream(func(w io.Writer) bool {
			select {
			case <-ctx.Request.Context().Done():
				return false
			case data, ok := <-stream:
				if !ok {
					return false
				}
				if _, err := w.Write(data); err != nil {
					return false
				}
				return true
			}
		})
	} else {
		resp, err := completion.Completion(ctx.Request.Context(), dep, &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}

func Chat(ctx *gin.Context, dep Dependency) {
	var req openai.ChatCompletionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Infof("receive chat request, stream: %v, messages: %v", req.Stream, req.Messages)

	if req.Stream {
		stream, err := chat.StreamChat(ctx.Request.Context(), dep, &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.Header("Content-Type", "text/event-stream")
		ctx.Stream(func(w io.Writer) bool {
			select {
			case <-ctx.Request.Context().Done():
				return false
			case data, ok := <-stream:
				if !ok {
					return false
				}
				if _, err := w.Write(data); err != nil {
					return false
				}
				return true
			}
		})
	} else {
		resp, err := chat.Chat(ctx.Request.Context(), dep, &req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, resp)
	}
}
