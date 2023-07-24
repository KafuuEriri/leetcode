package main

import "container/list"

type MyStack struct {
	list *list.List
}

func Constructor() MyStack {
	return MyStack{list: list.New()}
}

func (this *MyStack) Push(x int) {
	this.list.PushBack(x)
}

func (this *MyStack) Pop() int {
	// 这里注意需要返回值是一个int，但list结构返回的是一个interface{}，所以我需要进行类型转换
	// 具体的可以参考https://golang.org/ref/spec#Type_assertions
	return this.list.Remove(this.list.Back()).(int)
}

func (this *MyStack) Top() int {
	return this.list.Back().Value.(int)
}

func (this *MyStack) Empty() bool {
	return this.list.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
