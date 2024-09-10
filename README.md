# chatgpt-server-go
## 介绍
使用 [Gin](https://github.com/gin-gonic/gin) and [Go OpenAI](https://github.com/sashabaranov/go-openai) 实现的 chatgpt server，支持流式调用和非流式调用，兼容openai接口

项目采用 [golang 标准目录布局](https://github.com/golang-standards/project-layout)

## 使用指南
启动server
```bash
cd cmd/chatgpt
go run main.go
```

启动好服务之后，使用下面的命令测试
```bash
curl localhost:8080/v1/completions \
  -H "Content-Type: application/json" \
  -d '{
        "model": "codewise-7b",
        "prompt": "如何使用nginx进行负载均衡?",
        "max_tokens": 256,
        "temperature": 0.2,
        "stream": true
    }'

curl localhost:8080/v1/chat/completions \
  -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "model": "codewise-7b",
    "messages": [
        {"role": "system", "content": "You are a helpful assistant."},
        {"role": "user", "content": "如何使用nginx进行负载均衡？"}
    ],
    "max_tokens": 256,
    "temperature": 1,
    "stream": true,
    "skip_special_tokens": false
  }'
```