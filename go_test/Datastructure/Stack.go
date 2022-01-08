package datastructure

type Stack Link

func BuildStack() Stack{
	head := &LinkNode{0, nil}
	Stack := Stack{0, head}
	return Stack
}
func (q *Stack) Push(k int){
	n := BuildLinkNode(k)	
	q.insert(&n, q.length)
}
func (q *Stack) Pop() int{
	if q.length == 0{
		println("length is 0!!! can't delect anything")
		return -1
	}
	x := q.count(q.length - 1)
	temp := x.next
	x.next = x.next.next
	q.length--
	return temp.data
}
func (q *Stack) insert(p *LinkNode, c int){//插入节点p到c的后边
	cur := q.count(c)
	temp := cur.next
	cur.next = p
	p.next = temp
	q.length++
}
func (q *Stack) count(c int) *LinkNode{
	if c > q.length {
		println("no c")
		return nil
	}
	x := q.Head
	i := 0
	for i < c {
		x = x.next
		i++
	}
	return x
}
func (q *Stack) IsEmpty() bool{
	return q.length == 0
}