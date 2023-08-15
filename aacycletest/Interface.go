package main

import "fmt"

// Geometry interface是方法集合，而struct是属性集合
type Geometry interface {
	area() float64
	circle() float64
}

type Ren struct {
	name string
	dir  float64
}
type Dog struct {
	name string
	dir  float64
}

// 只有实现了interface所有方法，才能用赋值给interface 类型
// 注意：如果是值传递，那么直接转；  指针传递，那么需要 &Dog 赋值给interface
func (r Ren) area() float64 {
	return r.dir * r.dir
}
func (r Ren) circle() float64 {
	return r.dir * r.dir
}
func (r *Dog) area() float64 {
	return r.dir * r.dir
}
func (r *Dog) circle() float64 {
	return r.dir * r.dir
}

func measure(geometry Geometry) {
	area := geometry.area()
	fmt.Println(area)
}
