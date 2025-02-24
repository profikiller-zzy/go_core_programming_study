package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// Store 给 Monster 绑定方法 Store, 可以将一个 Monster 变量(对象),序列化后保存到文件中
func (this *Monster) Store() bool {
	//先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err =", err)
		return false
	}
	//保存到文件
	filePath := "./monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err =", err)
		return false
	}
	return true
}

// Load 给 Monster 绑定方法 Load, 可以将一个序列化的 Monster,从文件中读取，
// 并反序列化为 Monster 对象,检查反序列化，名字正确
func (this *Monster) Load() bool {
	//1. 先从文件中，读取序列化的字符串
	filePath := "./monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile err =", err)
		return false
	}
	//2.使用读取到 data []byte ,对反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("UnMarshal err =", err)
		return false
	}
	return true
}
