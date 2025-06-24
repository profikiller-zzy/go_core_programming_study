package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cloudwego/eino-ext/components/embedding/tencentcloud"
)

func main() {
	ctx := context.Background()

	// 创建 embedder 配置
	cfg := &tencentcloud.EmbeddingConfig{
		SecretID:  os.Getenv("TENCENTCLOUD_SECRET_ID"),
		SecretKey: os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		Region:    "ap-guangzhou",
	}

	// 创建 embedder
	embedder, err := tencentcloud.NewEmbedder(ctx, cfg)
	if err != nil {
		panic(err)
	}

	// 获取文本的向量表示
	embeddings, err := embedder.EmbedStrings(ctx, []string{"hello world", "bye world"})
	if err != nil {
		panic(err)
	}

	fmt.Println(len(embeddings[0]))
	fmt.Printf("Embeddings: %v\n", embeddings)
}
