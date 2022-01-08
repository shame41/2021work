package Sort

import (
	"math"
)

// 数字，右数第几位，从1开始
func Digit(num int, loc int) int {
	x := num % int(math.Pow10(loc)) / int(math.Pow10(loc-1))
    return x
}
func BucketSort(a Queue, k int) Queue{
	var bucket [10][]int
	for !a.IsEmpty() {
		temp := a.Dequeue()
		key := Digit(temp, k)
		bucket[key] = append(bucket[key], temp)
	}
	q := BuildQueue()
	for i := 0; i < 10; i++ {
		for _, item := range(bucket[i]){
			q.Enqueue(item)
		}
	}
	return q
}

func RadixByBucket(a []int) []int{

	max := maxInArr(a)
	key := 1;
	q := BuildQueue()
	for i := 0; i < len(a); i++ {
		q.Enqueue(a[i])
	}
	for i := 1; max/key != 0; i++{
		q = BucketSort(q, i)
		key = key*10
	}
	var ret []int
	for i := 0; i < len(a); i++ {
		ret = append(ret, q.Dequeue())
	}
	return ret
}
func maxInArr(a []int) int{
	max := -1000000
	for _, temp := range(a){
		if temp > max {
			max = temp
		}
	}
	return max
}

//队列
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
type Link struct{
	length int
	Head *LinkNode
}

type LinkNode struct{
	data int
	next *LinkNode
}

func (l *Link) Length() int{
	return l.length
}

func BuildLink() Link{
	head := &LinkNode{0, nil}
	link := Link{0, head}
	return link
}

func (l *Link) Search(k int) *LinkNode{//查找值为k的节点，并返回它
	x := l.Head
	for x != nil && x.data != k{
		x = x.next
	}
	if x == nil {
		println("no such node!!!")
	}
	return x
}

func (l *Link) Count(c int) *LinkNode{
	if c > l.length {
		println("no c")
		return nil
	}
	x := l.Head
	i := 0
	for i < c {
		x = x.next
		i++
	}
	return x
}

func BuildLinkNode(k int) LinkNode{//创建节点
	l := LinkNode{k, nil}
	return l
}

func (l *Link) Append(k int){//将节点n插入链表中
	n := BuildLinkNode(k)	
	l.Insert(&n, 0)
}

func (l *Link) Print(){//打印这个节点以及它所有后续节点的值
	x := l.Head
	for x != nil {
		x = x.next
		if x != nil {
			println(x.data)
		}
	}
}

func (l *Link) Insert(p *LinkNode, c int){//插入节点p到c的后边
	if c > l.length || c < 0{
		println("error")
		return
	}
	if c == 0 {//插到头部
		p.next = l.Head.next
		l.Head.next = p
	}else {
		cur := l.Count(c)
		temp := cur.next
		cur.next = p
		p.next = temp
	}
	l.length++
}

func (l *Link) DelectCount(k int){//删除第k个节点
	if l.length == 0{
		println("length is 0!!! can't delect anything")
		return
	}
	if k == 1 {
		l.Head.next = l.Head.next.next
		l.length--
	}else{
		x := l.Count(k - 1)
		x.next = x.next.next
		l.length--
	}
}
func (l *Link) DelectByData(k int){//删除数据为k的节点
	if l.length == 0{
		println("length is 0!!! can't delect anything")
		return
	}
	pre := l.Head
	for pre.next != nil && pre.next.data != k {
		pre = pre.next
	}
	if pre.next == nil{
		println("no such node")
		return
	}else{
		pre.next = pre.next.next
		l.length--
	}
}
////////////////////////////////////////////////////////////////////////////////////////
func merge(x1, x2 *LinkNode) *LinkNode{
	cur := &LinkNode{-1, nil}
	flag := 0
	var head *LinkNode
	for x1 != nil && x2 != nil{
		if x1.data <= x2.data {
			if flag == 0 {
				head = x1
				flag++
			}
			cur.next = x1
			x1 = x1.next
		}else{
			if flag == 0 {
				head = x2
				flag++
			}
			cur.next = x2
			x2 = x2.next
		}
		cur = cur.next
	}
	if x1 ==nil {
		cur.next = x2
	}else{
		cur.next = x1
	}
	return head
}

func Reverse(l *Link){
	if l.length < 2 {
		return
	}
	var pre *LinkNode
	pre = nil
	cur := l.Head.next
	fol := cur.next
	for {
		cur.next = pre
		if fol == nil {
			l.Head.next = cur
			return
		}
		pre = cur
		cur = fol
		fol = cur.next
	}
}

func Sort(l *Link){
	l.Head.next = mergeSort(l.Head.next)
}

func mergeSort(n *LinkNode) *LinkNode{
	if n == nil || n.next == nil {
		return n
	}
	q1, q2 := n, n
	for q2.next !=  nil && q2.next.next != nil {
		q2 = q2.next.next
		q1 = q1.next
	}
	left, right := n, q1.next
	q1.next = nil

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left,  right)
}