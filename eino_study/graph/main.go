package main

import (
	"context"
	"fmt"
	"io"

	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
)

const (
	nodeOfModel          = "model"
	nodeOfPrompt         = "prompt"
	nodeOfPostProcessing = "post-processing" // 新增的节点
)

func main() {
	ctx := context.Background()
	g := compose.NewGraph[map[string]any, *schema.Message]()

	pt := prompt.FromMessages(
		schema.FString,
		schema.UserMessage("what's the weather in {location}?"),
	)

	_ = g.AddChatTemplateNode(nodeOfPrompt, pt)
	_ = g.AddChatModelNode(nodeOfModel, &mockChatModel{}, compose.WithNodeName("ChatModel"))
	_ = g.AddLambdaNode("outputConverter", compose.InvokableLambda(func(ctx context.Context, message *schema.Message) ([]*schema.Message, error) {
		fmt.Println("ChatModel output: ", message.Content)
		return []*schema.Message{message}, nil
	}))
	_ = g.AddChatModelNode(nodeOfPostProcessing, &mockPostProcessingModel{}, compose.WithNodeName("PostProcessingModel")) // 新增的节点

	_ = g.AddEdge(compose.START, nodeOfPrompt)
	_ = g.AddEdge(nodeOfPrompt, nodeOfModel)
	_ = g.AddEdge(nodeOfModel, "outputConverter") // 报错
	_ = g.AddEdge("outputConverter", nodeOfPostProcessing)
	_ = g.AddEdge(nodeOfPostProcessing, compose.END)

	r, err := g.Compile(ctx, compose.WithMaxRunSteps(10))
	if err != nil {
		panic(err)
	}

	in := map[string]any{"location": "beijing"}
	ret, err := r.Invoke(ctx, in)
	fmt.Println("invoke result: ", ret)

	// stream
	s, err := r.Stream(ctx, in)
	if err != nil {
		panic(err)
	}

	defer s.Close()
	for {
		chunk, err := s.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		fmt.Println("stream chunk: ", chunk)
	}
}

type mockChatModel struct{}

func (m *mockChatModel) Generate(ctx context.Context, input []*schema.Message, opts ...model.Option) (*schema.Message, error) {
	return schema.AssistantMessage("the weather is good", nil), nil
}

func (m *mockChatModel) Stream(ctx context.Context, input []*schema.Message, opts ...model.Option) (*schema.StreamReader[*schema.Message], error) {
	sr, sw := schema.Pipe[*schema.Message](0)
	go func() {
		defer sw.Close()
		sw.Send(schema.AssistantMessage("the weather is", nil), nil)
		sw.Send(schema.AssistantMessage("good", nil), nil)
	}()
	return sr, nil
}

func (m *mockChatModel) BindTools(tools []*schema.ToolInfo) error {
	panic("implement me")
}

type mockPostProcessingModel struct{}

func (m *mockPostProcessingModel) Generate(ctx context.Context, input []*schema.Message, opts ...model.Option) (*schema.Message, error) {
	return schema.AssistantMessage("processed response", nil), nil
}

func (m *mockPostProcessingModel) Stream(ctx context.Context, input []*schema.Message, opts ...model.Option) (*schema.StreamReader[*schema.Message], error) {
	sr, sw := schema.Pipe[*schema.Message](0)
	go func() {
		defer sw.Close()
		sw.Send(schema.AssistantMessage("processed", nil), nil)
		sw.Send(schema.AssistantMessage("response", nil), nil)
	}()
	return sr, nil
}

func (m *mockPostProcessingModel) BindTools(tools []*schema.ToolInfo) error {
	panic("implement me")
}
