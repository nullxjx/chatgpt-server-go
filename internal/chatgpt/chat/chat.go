package chat

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

// Chat 对话非流式接口
func Chat(ctx context.Context, dep Dependency, req *openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error) {
	return dep.GetClient().CreateChatCompletion(ctx, *req)
}
