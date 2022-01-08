package datastructure

import "go_test/Tool"

var MAXHIGH int
type BRT struct{
	val int
	size int
	Left *BRT
	Right *BRT
	color bool
}
func NewBRT(val int) *BRT {
	temp := BRT{val, 1, nil, nil, false}
	return &temp
}

func (t *BRT)IsRed() bool{
	if t == nil {
		return false
	}
	return t.color
}
func (t *BRT)SetColorRed(color bool){
	t.color = color
}
func (t *BRT)GetSize() int{
	if t == nil {
		return 0
	}
	return t.size
}
func (t *BRT)rotationLeft() *BRT{
	temp := t.Right
	t.Right = t.Right.Left
	temp.Left = t
	temp.SetColorRed(t.IsRed())
	t.SetColorRed(true)

	return temp
}
func (t *BRT)rotationRight() *BRT{
	temp := t.Left
	t.Left = t.Left.Right
	temp.Right = t
	temp.SetColorRed(t.IsRed())
	t.SetColorRed(true)

	return temp
}
func (t *BRT)flipColors(){
	t.SetColorRed(!t.IsRed())
	t.Left.SetColorRed(!t.Left.IsRed())
	t.Right.SetColorRed(!t.Right.IsRed())
}
// func (t *BRT)flipflipColors(){
// 	t.color = false
// 	t.Left.color = true
// 	t.Right.color = true
// }
func (t *BRT)AddNode(k int) *BRT {
	t = t.add(k)
	t.SetColorRed(false)
	return t
}
func (t *BRT)add(k int) *BRT{
	newNode := BRT{k, 1, nil, nil, true}
	if t == nil {
		return &newNode
	}
	if k - t.val < 0{//要插入的值比当前节点小
		t.Left = t.Left.add(k)
	}else{
		t.Right = t.Right.add(k)
	}
	t = t.KeepBalance()
	t.size = t.Left.GetSize() + t.Right.GetSize() + 1
	return t
}
func (t *BRT)Print(){//输出节点
	if t != nil {	
		t.Left.Print()
		println(t.val)
		t.Right.Print()
	}
}
func (t *BRT)GetHeight() int{
	if(t == nil){
		return 0
	}
	if t.Left.IsRed(){
		return Tool.Max(t.Left.GetHeight(), t.Right.GetHeight()+1)
	}else{
		return Tool.Max(t.Left.GetHeight()+1, t.Right.GetHeight()+1)
	}
}
func (t *BRT)KeepBalance() *BRT{
	if t.Right.IsRed() && !t.Left.IsRed() {
		//出现了红色右链接，需要左旋
		t = t.rotationLeft()
	}
	if t.Left.IsRed() && t.Left.Left.IsRed(){
		//出现了连续两个红色左连接，需要右旋
		t = t.rotationRight()
	}
	if t.Left.IsRed() && t.Right.IsRed() {
		//出现了红色左连接和红色右链接，需要颜色反转
		t.flipColors()
	}
	return t
}


func (t *BRT)DelMin(){
	if !t.Left.IsRed() && !t.Right.IsRed() {
		t.SetColorRed(true)
	}
	t = t.delMin()
	if t != nil {
		t.SetColorRed(false)
	}
}
func (t *BRT)delMin() *BRT{
	if t.Left == nil {
		return nil
	}
	if !t.Left.IsRed() && !t.Left.Left.IsRed() {//t的左节点是2节点
		t = t.moveRedLeft()
	}
	t.Left = t.Left.delMin()
	return t.balance()
}
func (t *BRT)moveRedLeft() *BRT{
	t.flipColors()
	if t.Right.Left.IsRed() {//t的右节点不是2节点
		t.Right = t.Right.rotationRight()
		t = t.rotationLeft()
	}
	return t
}
func (t *BRT)balance() *BRT{
	t = t.KeepBalance()
	if !t.Right.IsRed() {//不能出现右红链接
		t = t.rotationLeft()
	}
	return t
}