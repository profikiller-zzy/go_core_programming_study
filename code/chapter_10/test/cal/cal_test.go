package cal

import "testing"

func TestAddUpper(t *testing.T) {
	// 调用要测试的函数
	res := addUpper(10)
	if res != 45 {
		t.Fatalf("AddUpper(10) 执行错误，期望值 %v, 实际值%v", 55, res)
	}
	// 如果正确
	t.Logf("AddUpper(10) 执行成功...")
}
