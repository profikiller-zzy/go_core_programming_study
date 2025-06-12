package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	// 创建 deepseek 模型
	cm, err := deepseek.NewChatModel(ctx, &deepseek.ChatModelConfig{
		APIKey:    "sk-10f31c7c90ef46d2bd270cbaec105799",
		Model:     "deepseek-chat",
		MaxTokens: 2000,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 创建天气查询工具
	weatherTool, err := utils.InferTool("get_weather", "查询某地天气", getWeather)
	if err != nil {
		log.Fatal(err)
	}

	// 创建空气质量查询工具
	airQualityTool, err := utils.InferTool("get_air_quality", "查询某地空气质量", getAirQuality)
	if err != nil {
		log.Fatal(err)
	}

	// 创建时间查询工具
	timeInfoTool, err := utils.InferTool("get_time_info", "查询当前时间和日期信息", getTimeInfo)
	if err != nil {
		log.Fatal(err)
	}

	// 收集所有工具
	tools := []tool.BaseTool{weatherTool, airQualityTool, timeInfoTool}

	// 收集工具信息并绑定到模型
	toolInfos := make([]*schema.ToolInfo, 0, len(tools))
	for _, tool := range tools {
		info, err := tool.Info(ctx)
		if err != nil {
			log.Fatal(err)
		}
		toolInfos = append(toolInfos, info)
	}

	err = cm.BindTools(toolInfos)
	if err != nil {
		log.Fatal(err)
	}

	// 创建工具节点
	toolsNode, err := compose.NewToolNode(context.Background(), &compose.ToolsNodeConfig{
		Tools: tools,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 定义 state 结构体用于存储历史消息
	type agentState struct {
		history []*schema.Message
		steps   int // 记录步骤数
	}

	// 使用 Graph 构建 ReAct agent
	g := compose.NewGraph[[]*schema.Message, *schema.Message](
		compose.WithGenLocalState(func(ctx context.Context) *agentState {
			return &agentState{
				history: []*schema.Message{},
				steps:   0,
			}
		}),
	)

	// 添加 ChatModel 节点，使用 StatePreHandler 处理输入的消息
	err = g.AddChatModelNode("chat_model", cm,
		compose.WithStatePreHandler(func(ctx context.Context, messages []*schema.Message, state *agentState) ([]*schema.Message, error) {
			// 将输入消息添加到历史消息中
			state.history = append(state.history, messages...)
			state.steps++
			fmt.Printf("\n===== 第 %d 次模型调用(输入) =====\n", state.steps)
			fmt.Println("chat_model WithStatePreHandler")
			for _, msg := range messages {
				fmt.Printf("角色: %s\n内容: %s\n", msg.Role, msg.Content)
				// 打印工具调用结果 (如果存在)
				if msg.Role == schema.Tool {
					fmt.Printf("工具结果: %s\n", msg.Content)
				}
			}
			return state.history, nil
		}),
		compose.WithStatePostHandler(func(ctx context.Context, message *schema.Message, state *agentState) (*schema.Message, error) {
			// 将模型输出添加到历史消息中
			state.history = append(state.history, message)
			fmt.Printf("\n===== 第 %d 次模型调用(输出) =====\n", state.steps)
			fmt.Println("chat_model WithStatePostHandler")
			fmt.Printf("角色: %s\n内容: %s\n", message.Role, message.Content)
			if len(message.ToolCalls) > 0 {
				fmt.Printf("工具调用: %d 个\n", len(message.ToolCalls))
				for i, tc := range message.ToolCalls {
					fmt.Printf("  调用 %d: 工具名称: %s, 参数: %s\n", i+1, tc.Function.Name, tc.Function.Arguments)
				}
			}
			return message, nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 添加 Tools 节点，使用 StatePreHandler 获取最新的消息
	err = g.AddToolsNode("tools_node", toolsNode,
		compose.WithStatePreHandler(func(ctx context.Context, message *schema.Message, state *agentState) (*schema.Message, error) {
			// 使用最新的消息调用工具
			fmt.Printf("\n===== 工具调用 =====\n")
			fmt.Println("tools_node WithStatePreHandler")
			latestMsg := state.history[len(state.history)-1]
			fmt.Printf("调用工具的消息: %s\n", latestMsg.Content)
			return latestMsg, nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 添加边和分支
	// START -> chat_model
	err = g.AddEdge(compose.START, "chat_model")
	if err != nil {
		log.Fatal(err)
	}

	// tools_node -> chat_model (工具调用后返回到模型)
	err = g.AddEdge("tools_node", "chat_model")
	if err != nil {
		log.Fatal(err)
	}

	// chat_model -> 分支条件
	err = g.AddBranch("chat_model", compose.NewGraphBranch(func(ctx context.Context, message *schema.Message) (string, error) {
		// 检查模型输出是否有工具调用
		if len(message.ToolCalls) > 0 {
			return "tools_node", nil // 如果有工具调用，转到工具节点
		}
		return compose.END, nil // 如果没有工具调用，结束流程
	}, map[string]bool{"tools_node": true, compose.END: true}))
	if err != nil {
		log.Fatal(err)
	}

	// 编译 Graph
	agent, err := g.Compile(ctx)
	if err != nil {
		log.Fatal(fmt.Errorf("编译失败: %w", err))
	}

	// 运行示例
	fmt.Println("开始处理查询...")
	resp, err := agent.Invoke(ctx, []*schema.Message{
		{
			Role:    schema.User,
			Content: "我计划今天在北京出门，请告诉我今天北京的天气、空气质量和当前时间。根据这些信息，判断我是否需要带遮阳伞，如果需要，给出理由。最后，根据当前时间建议我什么时候出门最合适。",
		},
	})

	if err != nil {
		log.Fatalf("处理失败: %v", err)
	}

	// 输出结果
	fmt.Printf("\n最终响应:\n")
	fmt.Printf("角色: %s\n内容: %s\n", resp.Role, resp.Content)
}

// 天气查询参数
type getWeatherParams struct {
	Location string `json:"location" jsonschema:"description=地点名称"`
}

// 天气查询实现
func getWeather(_ context.Context, params *getWeatherParams) (string, error) {
	if params.Location == "北京" {
		return params.Location + " 的天气为 晴转多云 35 度，紫外线指数很高，达到级别8(很强)", nil
	}
	return params.Location + " 的天气为 晴转多云 30 度", nil
}

// 空气质量查询参数
type getAirQualityParams struct {
	Location string `json:"location" jsonschema:"description=地点名称"`
}

// 空气质量查询实现
func getAirQuality(_ context.Context, params *getAirQualityParams) (string, error) {
	if params.Location == "北京" {
		return params.Location + " 的空气质量指数为 75，空气质量良好", nil
	}
	return params.Location + " 的空气质量指数为 50，空气质量良好", nil
}

// 时间信息查询参数
type getTimeInfoParams struct {
	TimeZone string `json:"timezone" jsonschema:"description=时区名称,默认为当地时区"`
}

// 时间信息查询实现
func getTimeInfo(_ context.Context, params *getTimeInfoParams) (string, error) {
	now := time.Now()
	return fmt.Sprintf("当前时间是 %s，%d-%02d-%02d，星期%s",
		now.Format("15:04:05"),
		now.Year(), now.Month(), now.Day(),
		[]string{"日", "一", "二", "三", "四", "五", "六"}[now.Weekday()]), nil
}
