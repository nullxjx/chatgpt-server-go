package completion

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/nullxjx/chatgpt-server-go/cmd/chatgpt/config"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
)

const (
	StreamPrefix = "data: "
	StreamSuffix = "\n\n"
	StreamEnd    = "data: [DONE]\n\n"
)

// Dependency 是当前pkg的外部依赖。
type Dependency interface {
	Conf() *config.Config
	GetClient() *openai.Client
}

// StreamCompletion 补全流式接口
func StreamCompletion(ctx context.Context, dep Dependency, req *openai.CompletionRequest) (<-chan []byte, error) {
	stream, err := dep.GetClient().CreateCompletionStream(ctx, *req)
	if err != nil {
		return nil, err
	}

	byteChan := make(chan []byte, 4096)
	go func(byteChan chan<- []byte) {
		defer close(byteChan)
		defer stream.Close()

		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				byteChan <- []byte(StreamEnd)
				return
			}
			if err != nil {
				log.Errorf("Stream receive error: %v", err)
				byteChan <- []byte(fmt.Sprintf("error: %v", err))
				return
			}
			responseBytes, err := json.Marshal(response)
			if err != nil {
				log.Errorf("JSON marshal error: %v", err)
				byteChan <- []byte(fmt.Sprintf("marshal error: %v", err))
				return
			}
			formattedResponse := append([]byte(StreamPrefix), responseBytes...)
			formattedResponse = append(formattedResponse, []byte(StreamSuffix)...)
			byteChan <- formattedResponse
		}
	}(byteChan)
	return byteChan, nil
}
