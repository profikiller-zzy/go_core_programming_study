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
	nodeA = "node_a"
	nodeB = "node_b"
	nodeC = "node_c"
)

func main() {
	ctx := context.Background()
	g := compose.NewGraph[map[string]any, *schema.Message]()

	// Node A: Chat Model to generate some information
	systemTplA := `
你是一名信息生成助手。你的任务是生成关于审计公司的相关信息。
`

	userMsgA := `
生成关于审计公司的一些基本信息。
`

	chatModelA := prompt.FromMessages(
		schema.FString,
		schema.SystemMessage(systemTplA),
		schema.UserMessage(userMsgA),
	)

	_ = g.AddChatTemplateNode(nodeA, chatModelA)

	// Node B: Template Node to fill placeholders with output from Node A
	templateB := prompt.TemplateFunc(func(ctx context.Context, input map[string]interface{}) (string, error) {
		// 获取 Node A 的输出
		messageContent := ""
		if msg, ok := input["message"]; ok {
			messageContent = msg.(*schema.Message).Content
		}

		// 模板内容，包含占位符 {event_content}
		templateContent := `
基于以下信息，推断审计公司可能存在的运营风险：
审计公司提交的信息: {event_content}。
首先，调用 get_similar_events 工具从知识库中获取类似的事件信息。
然后，分析这些信息以推断审计公司的运营风险。
推断的运营风险应符合以下格式：

type RiskPoint struct {
    Name      string          
    RiskLevel ctype.RiskLevel 
    Reason    string          

    HistoryID int64 
}
确保你的响应包含多个 RiskPoint 条目，如果有多个风险被识别。
`

		// 替换占位符
		filledTemplate := fmt.Sprintf(templateContent, messageContent)
		return filledTemplate, nil
	})

	_ = g.AddTemplateNode(nodeB, templateB)

	// Node C: Chat Model to process the filled template
	systemTplC := `
你是一名风险管理助手。你的任务是分析提供的信息并推断审计公司的潜在运营风险。
`

	chatModelC := prompt.FromMessages(
		schema.FString,
		schema.SystemMessage(systemTplC),
		schema.UserMessage("{template_output}"),
	)

	_ = g.AddChatTemplateNode(nodeC, chatModelC)

	// Add edges between nodes
	_ = g.AddEdge(compose.START, nodeA)
	_ = g.AddEdge(nodeA, nodeB)
	_ = g.AddEdge(nodeB, nodeC)
	_ = g.AddEdge(nodeC, compose.END)

	r, err := g.Compile(ctx, compose.WithMaxRunSteps(10))
	if err != nil {
		panic(err)
	}

	in := map[string]any{}
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

// Mock implementation of a chat model
type mockChatModel struct{}

func (m *mockChatModel) Generate(ctx context.Context, input []*schema.Message, opts ...model.Option) (*schema.Message, error) {
	// 根据输入生成响应
	userMsg := input[len(input)-1].Content
	response := "这是审计公司的基本信息：..."
	return schema.AssistantMessage(response, nil), nil
}

func (m *mockChatModel) Stream(ctx context.Context, input []*schema.Message, opts ...model.Option) (*schema.StreamReader[*schema.Message], error) {
	sr, sw := schema.Pipe[*schema.Message](0)
	go func() {
		defer sw.Close()
		userMsg := input[len(input)-1].Content
		response := "这是审计公司的基本信息：..."
		sw.Send(schema.AssistantMessage(response, nil), nil)
	}()
	return sr, nil
}

func (m *mockChatModel) BindTools(tools []*schema.ToolInfo) error {
	panic("implement me")
}
