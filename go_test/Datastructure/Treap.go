package datastructure

import (
	"go_test/Tool"
	"math/rand"
	"time"
)



type Treap struct{
	Left *Treap
	Right *Treap
	val int
	priority int
}

func (t *Treap)KeepBalance() *Treap{
	if t.GetPriority() > t.Left.GetPriority() {
		return t.RightRotation()
	}else if t.GetPriority() >= t.Right.GetPriority() {
		return t.LeftRotation()
	}
	return t
}
func (t *Treap)LeftRotation() *Treap{
	temp := t.Right
	t.Right = t.Right.Left
	temp.Left = t
	return temp
}
func (t *Treap)RightRotation() *Treap{
	temp := t.Left
	t.Left = t.Left.Right
	temp.Right = t
	return temp
}
func (t *Treap)GetPriority() int{
	if t == nil {
		return 10000000
	}
	return t.priority
}
func (t *Treap)SetPriority(k int){
	t.priority = k
}
func (t *Treap)GetVal() int{
	if t == nil {
		return 0
	}
	return t.val
}
func (t *Treap)SetVal(k int){
	t.val = k
}
func NewTreap(k int) *Treap{
	t := Treap{nil, nil, k, 1}
	return &t
}
func (t *Treap)Add(k int) *Treap{
	rand.Seed(time.Now().UnixNano())
	if t == nil {
		newNode := Treap{nil,nil, k, rand.Intn(9999999)}
		return &newNode
	}
	if k > t.val {
		t.Right = t.Right.Add(k)
	}else{
		t.Left = t.Left.Add(k)
	}
	return t.KeepBalance()
}
func (t *Treap)Print(){//输出节点
	if t != nil {	
		t.Left.Print()
		println(t.priority)
		t.Right.Print()
	}
}
func (t *Treap)GetHeight() int{
	if(t == nil){
		return 0
	}
	return Tool.Max(t.Left.GetHeight()+1, t.Right.GetHeight()+1)
}
func (t *Treap)BFS(){//输出节点
	
	for {
		print(t.Left.GetPriority())
		print(" ")
		print(t.Right.GetPriority())
	}
	
}