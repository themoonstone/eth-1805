package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// 反射实践
// 针对结构体，
// 可以通过反射遍历结构体字段，
// 调用结构体方法，
// 获取结构体标签的值

type Monster struct {
	Name 	string		`json:"name"`
	Age 	int
	Skill	string
	Heigth	float64
	ChildRen	map[string]string
	Slice		[]string
	Channel 	chan int
	Test	int
}

func SizeOf()  {
	var p []int
	// int string bool
	fmt.Println(reflect.TypeOf(unsafe.Sizeof(p)))
	fmt.Println(unsafe.Sizeof(p))
	var s uintptr
	fmt.Println(unsafe.Sizeof(s))
}
// reflect.type中结构体字段相关方法
/*

	Field(i int) StructField

	FieldByIndex(index []int) StructField

	FieldByName(name string) (StructField, bool)

	FieldByNameFunc(match func(string) bool) (StructField, bool)
*/
/*
	// Name is the field name.
	Name string
	// PkgPath is the package path that qualifies a lower case (unexported)
	// field name. It is empty for upper case (exported) field names.
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	PkgPath string

	Type      Type      // field type
	Tag       StructTag // field tag string
	Offset    uintptr   // offset within struct, in bytes
	Index     []int     // index sequence for Type.FieldByIndex
	Anonymous bool      // is an embedded field	是否嵌套
*/
// 获取结构体字段相关信息
func getMonsterFieldInfo(s interface{})  {
	monsterType := reflect.TypeOf(s)
	fmt.Printf("monsterType.Field(0) is %v\n", monsterType.Field(5))
	fmt.Printf("monsterType.FieldByIndex([]int{0}) is %v\n", monsterType.FieldByIndex([]int{0}))
	s1, _ := monsterType.FieldByName("Name")
	fmt.Printf("monsterType.FieldByName(\"Name\") is %v\n", s1)
}

// 结构体方法
func (s Monster) GetSum(n1, n2 int)int {
	return n1 + n2
}

func (s Monster) GetSub(n1, n2 int) int {
	return n1 - n2
}

/*
	Method(int) Method
	MethodByName(string) (Method, bool)
	NumMethod() int
*/
// 获取结构体方法相关信息
func GetStructMethodInfo(s interface{}) {
	monsterType := reflect.TypeOf(s)
	fmt.Printf("monsterType.Method(0) is %v\n", monsterType.Method(0))
	method, _ := monsterType.MethodByName("GetSub")
	fmt.Printf("monsterType.MethodByName(GetSub) is %v\n", method)
	fmt.Printf("monsterType.NumMethod is %v\n", monsterType.NumMethod())
}

//
func TestStruct(a interface{}) {
	// 获取类型
	typ := reflect.TypeOf(a)
	// 获取值
	val := reflect.ValueOf(a)
	// 获取类别
	kind := val.Kind()

	if kind != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	// 获取字段数量
	num := val.NumField()
	fmt.Println("字段数量:", num)

	// 获取结构体所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d is %v\n", i, val.Field(i))
		// 检测是否有标签
		tagVal := typ.Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag is %v\n", i, tagVal)
		}
	}
}
// monsterType.Field(0) is {Name  string  0 [0] false}
// monsterType.Field(0) is {Age  int  16 [1] false}
// monsterType.Field(0) is {ChildRen  map[string]string  48 [4] false}
// monsterType.Field(0) is {Slice  []int  56 [5] false}
// monsterType.Field(0) is {Channel  chan int  80 [6] false}
// monsterType.Field(0) is {Test  int  88 [7] false}
func main() {
	SizeOf()
	monster := Monster{
		Name: "黄毛精",
		Age: 200,
		Skill: "喷火",
		Heigth: 666.666,
		ChildRen: map[string]string{
			"dagou":"boy",
			"ergou":"girl",
		},
	}
	getMonsterFieldInfo(monster)
	fmt.Println("-------------------------------------")
	GetStructMethodInfo(monster)
	fmt.Println("-------------------------------------")
	TestStruct(monster)
}