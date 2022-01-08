package datastructure

import (
	"go_test/Tool"
)

type AVLT struct{
	Left *AVLT
	Right *AVLT
	val int
	height int
}

func (t *AVLT)LeftRotation() *AVLT{
	temp := t.Right
	t.Right = t.Right.Left
	temp.Left = t
	t.height = Tool.Max(t.Left.GetHeight(), t.Right.GetHeight())+1
	temp.height = Tool.Max(temp.Left.GetHeight(), temp.Right.GetHeight())+1
	return temp
}
func (t *AVLT)RightRotation() *AVLT{
	temp := t.Left
	t.Left = t.Left.Right
	temp.Right = t
	t.height = Tool.Max(t.Left.GetHeight(), t.Right.GetHeight())+1
	temp.height = Tool.Max(temp.Left.GetHeight(), temp.Right.GetHeight())+1
	return temp
}
func (t *AVLT)KeepBalance() *AVLT{
	if t.GetHeight() <= 2 {
		return t
	}
	if t.Left.GetHeight() >= t.Right.GetHeight() + 2 && t.Left.Left.GetHeight() > t.Left.Right.GetHeight() {
		return t.RightRotation()//LL
	}else if t.Right.GetHeight() >= t.Left.GetHeight() + 2 && t.Right.Right.GetHeight() > t.Right.Left.GetHeight(){
		return t.LeftRotation()//RR
	}else if t.Left.GetHeight() >= t.Right.GetHeight() + 2 && t.Left.Left.GetHeight() < t.Left.Right.GetHeight(){
		t.Left = t.Left.LeftRotation()
		return t.RightRotation()
	}else if t.Right.GetHeight() >= t.Left.GetHeight() + 2 && t.Right.Right.GetHeight() < t.Right.Left.GetHeight(){
		t.Right = t.Right.RightRotation()
		return t.LeftRotation()
	}else{
		return t
	}
}
func (t *AVLT)Add(k int) *AVLT{
	if t == nil {
		newNode := AVLT{nil,nil, k, 1}
		return &newNode
	}
	if k > t.val {
		t.Right = t.Right.Add(k)
	}else{
		t.Left = t.Left.Add(k)
	}
	t.height = 1 + Tool.Max(t.Left.GetHeight(), t.Right.GetHeight())
	return t.KeepBalance()
}
func (t *AVLT)Print(){//输出节点
	if t != nil {	
		t.Left.Print()
		println(t.val)
		t.Right.Print()
	}
}

func (t *AVLT)GetHeight() int{
	if t == nil {
		return 0
	}
	return t.height
}
func (t *AVLT)SetHeight(k int){
	t.height = k
}
func (t *AVLT)GetVal() int{
	if t == nil {
		return 0
	}
	return t.val
}
func (t *AVLT)SetVal(k int){
	t.val = k
}
func (t *AVLT)Height() int{
	if(t == nil){
		return 0
	}
	return Tool.Max(t.Left.GetHeight()+1, t.Right.GetHeight()+1)
}
func NewAVLT(k int) *AVLT{
	t := AVLT{nil, nil, k, 1}
	return &t
}