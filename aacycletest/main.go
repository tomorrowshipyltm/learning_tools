package main

import (
	"fmt"
	"sync"
)

type Person struct {
	name string
	age  int
}

type a interface {
	sayHello()
}

// 引用传递 避免 值拷贝
func (p *Person) getName() (string, int) {
	return p.name, p.age
}

// 2个线程，交替打印，使用channel
func cycle() {
	ch1 := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			ch1 <- struct{}{}
			if i%2 == 0 {
				fmt.Println(i)
			}

		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			<-ch1
			if i%2 != 0 {
				fmt.Println(": ", i)
			}
		}
	}()
	wg.Wait()
}

// closure，返回函数
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// switch默认有break，不需要加
func character(s string) {
	// golang中无字符， 使用rune(int32 aliasName) 表示Unicode
	// 字符间的比较 基于unicode, 字符串的元素 都是rune(int32)
	var ch rune = 'a'
	if ch == 'a' {
		fmt.Println(ch, string(ch))
	}
	tmpStr := "information desk"
	for index, ele := range tmpStr {
		fmt.Printf("%d: %v ", index, ele)
	}
	// 2D数组[m][n]int ，m/n必须是常数，否则报错
	arr := [2]int{1, 2}

	// slice 动态数组
	m := 3
	n := 2
	arr2D := make([][]int, 3)
	for i := 0; i < m; i++ {
		arr2D[i] = make([]int, n)
		for j := 0; j < n; j++ {
			arr2D[i][j] = i * j
		}
	}
	arr2D = append(arr2D, []int{1, 3})
	fmt.Println(arr, " ", arr2D, " ", "le ", len(arr2D), cap(arr2D))

	// map
	myHash := make(map[int]string, 3)
	myHash = map[int]string{1: "hello", 2: "world"}
	for key, value := range myHash {
		fmt.Println(key, " ", value)
	}
	if ele, ok := myHash[2]; ok {
		fmt.Printf("found %s \n", ele)
	}

	switch s {
	case "he":
		fmt.Println("I'm ", s)
		break
	case "hel":
		fmt.Println("I'm ", s)
	default:
		fmt.Println("default ", s)
	}

}

//
//func doHandler(handler BaseHandler) {
//	handler.run()
//}

func main() {
	//doHandler(SynchroHandler{})
	//doHandler(ChanHandler{})

	// interface
	ren := Ren{name: "he", dir: 3.14}
	dog := Dog{name: "dog", dir: 6.8}
	measure(ren)
	measure(&dog)

	// struct
	p := Person{name: "liu", age: 7}
	name, age := p.getName()
	fmt.Println(name, " ", age, " ", &p)

	// 函数闭包，i 状态保持，输出 1、2、3
	nextInt := intSeq()
	fmt.Println(nextInt(), " ", nextInt(), " ", nextInt())
	character("hel")
	//cycle()
}
