package datastructure

type Queue Link

func BuildQueue() Queue{
	head := &LinkNode{0, nil}
	queue := Queue{0, head}
	return queue
}
func (q *Queue) Enqueue(k int){
	n := BuildLinkNode(k)	
	q.insert(&n, q.length)
}
func (q *Queue) Dequeue() int{
	if q.length == 0{
		println("length is 0!!! can't delect anything")
		return -1
	}
	temp := q.Head.next
	q.Head.next = q.Head.next.next
	q.length--
	return temp.data
}
func (q *Queue) insert(p *LinkNode, c int){//插入节点p到c的后边
	cur := q.count(c)
	temp := cur.next
	cur.next = p
	p.next = temp
	q.length++
}
func (q *Queue) count(c int) *LinkNode{
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
func (q *Queue) IsEmpty() bool{
	return q.length == 0
}