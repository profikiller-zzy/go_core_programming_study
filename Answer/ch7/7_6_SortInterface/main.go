package main

// 练习 7.8： 很多图形界面提供了一个有状态的多重排序表格插件：
// 主要的排序键是最近一次点击过列头的列，第二个排序键是第二最近点击过列头的列，等等。
// 定义一个sort.Interface的实现用在这样的表格中。比较这个实现方式和重复使用sort.Stable来排序的方式。

// Person 定义需要排列的数据
type Person struct {
	Name string
	Age  int
	Sex  bool
	ID   int
}

// Table 定义数据
type Table struct {
	Data []Person
}

// MultipleSorter 定义多路选择器
type MultipleSorter struct {
	table    *Table
	criteria []int // 记录排行的列
}

func NewMultipleSorter(table *Table) *MultipleSorter {
	return &MultipleSorter{table: table}
}

func (ms *MultipleSorter) AddCriteria(column int) { // 添加新列，有点类似于LRU算法的替换
	for index, value := range ms.criteria {
		if value == column {
			ms.criteria = append(ms.criteria[:index], ms.criteria[index+1:]...)
			break
		}
	}
	ms.criteria = append([]int{column}, ms.criteria...) // 把刚刚点击的列放在第一位
}

// 实现sort.Interface

func (ms *MultipleSorter) Len() int {
	return len(ms.table.Data)
}

func (ms *MultipleSorter) Less(i, j int) bool {
	for _, column := range ms.criteria {
		switch column {
		case 0: // Name
			if ms.table.Data[i].Name != ms.table.Data[j].Name {
				return ms.table.Data[i].Name < ms.table.Data[j].Name
			}
		case 1: // Age
			if ms.table.Data[i].Age != ms.table.Data[j].Age {
				return ms.table.Data[i].Age < ms.table.Data[j].Age
			}
		case 2: // Sex
			return ms.table.Data[i].Sex && !ms.table.Data[j].Sex // 默认把Sex = true的值排在前
		case 3:
			if ms.table.Data[i].ID != ms.table.Data[j].ID {
				return ms.table.Data[i].ID < ms.table.Data[j].ID
			}
		}
	}
	return false
}

func (ms *MultipleSorter) Swap(i, j int) {
	ms.table.Data[i], ms.table.Data[j] = ms.table.Data[j], ms.table.Data[i]
}

// 练习 7.9： 使用html/template包（§4.6）替代printTracks将tracks展示成一个HTML表格。
// 将这个解决方案用在前一个练习中，让每次点击一个列的头部产生一个HTTP请求来排序这个表格。

const TEMPL = `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Document</title>
</head>
<body>
	<table border="1" cellspacing="0">
		<tr>
			<th><a href="/?sort=Name">Name</a></th>
			<th><a href="/?sort=Age">Age</a></th>
			<th><a href="./?sort=Sex">Sex</a></th>
			<th><a href="./?sort=ID">ID</a></th>
		</tr>
		{{range .Data}}
		<tr>
			<td>{{.Name}}</td>
			<td>{{.Age}}</td>
			<td>{{.Sex}}</td>
			<td>{{.ID}}</td>
		</tr>
		{{end}}
	</table>
</body>
</html>
`

// 进行测试
func main() {
	// 构建数据
	//table := &Table{Data: []Person{
	//	{Name: "Alice", Age: 30, Sex: true, ID: 1},
	//	{Name: "Bob", Age: 25, Sex: false, ID: 2},
	//	{Name: "Charlie", Age: 35, Sex: true, ID: 3},
	//	{Name: "Diana", Age: 28, Sex: false, ID: 4},
	//	{Name: "Eve", Age: 22, Sex: true, ID: 5},
	//}}
	//ms := NewMultipleSorter(table)
	//
	//// 模拟点击列头的顺序
	//ms.AddCriteria(1) // 先按第1列排序，第一列是年龄
	//ms.AddCriteria(0) // 再按第0列排序，第零列是姓名
	//
	//sort.Sort(ms)
	//for _, row := range table.Data {
	//	fmt.Println(row)
	//}
	//
	//ms.AddCriteria(1) // 先按第1列排序，第一列是年龄
	//sort.Sort(ms)
	//for _, row := range table.Data {
	//	fmt.Println(row)
	//}
	//
	//// 构建模板
	//templ, err := template.New("sort").Parse(TEMPL)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	sort.Sort(ms)
	//	// 输出
	//	templ.Execute(w, table)
	//})
	//// 监听
	//http.ListenAndServe(":8080", nil)
}
