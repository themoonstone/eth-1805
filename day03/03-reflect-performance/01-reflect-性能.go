package _3_reflect_performance

import "reflect"

// 反射性能说明

type Data struct {
	X int
}
// 初始化结构体实例对象
var d = new(Data)
// 直接赋值
func set(x int) {
	d.X = x
}

// 通过反射进行赋值
func rset(x int) {
	v := reflect.ValueOf(d).Elem()
	f := v.FieldByName("X")
	f.Set(reflect.ValueOf(x))
}

// 结构体方法
func (x *Data) Inc() {
	x.X ++
}

// 方法正常调用
func call() {
	d.Inc()
}

func rcall() {
	v := reflect.ValueOf(d)
	mt:=v.MethodByName("Inc")
	mt.Call(nil)
}