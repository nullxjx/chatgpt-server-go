package completion

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

// Completion 补全非流式接口
func Completion(ctx context.Context, dep Dependency, req *openai.CompletionRequest) (openai.CompletionResponse, error) {
	return dep.GetClient().CreateCompletion(ctx, *req)
}
