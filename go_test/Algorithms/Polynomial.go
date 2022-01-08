package algorithms

type PolyLink struct{
	Head *Node
	Tail *Node
	len int
}
type Node struct{
	index int
	coef int
	Next *Node
}

func (l *PolyLink) GetLen() int{
	return l.len
}
func (l *PolyLink) GetIndex(k int) int{
	if k > l.len {
		return -1
	}else if k == l.len {
		return l.Tail.index
	}else if k == 1 {
		return l.Head.Next.index
	}else if k > 1 {
		cur := l.Head
		for i := 0; i < k; i++ {
			cur = cur.Next
		}
		return cur.index
	}else{
		return -1
	}
}
func (l *PolyLink) GetCoef(k int) int{
	if k > l.len {
		return -1
	}else if k == l.len {
		return l.Tail.coef
	}else if k == 1 {
		return l.Head.Next.coef
	}else if k > 1 {
		cur := l.Head
		for i := 0; i < k; i++ {
			cur = cur.Next
		}
		return cur.coef
	}else{
		return -1
	}
}

func BuildNode(index, coef int) *Node{
	n := Node{index, coef, nil}
	return &n
}

func BuildPoly() *PolyLink{
	head := BuildNode(0, -1)
	tail := BuildNode(0, -1)
	p := PolyLink{head, tail, 0}
	return &p
}

func Mult(list1, list2 *PolyLink) *PolyLink{
	cnt := 0
	pat1 := list1.Head.Next
	pat2 := list2.Head.Next
	list := BuildPoly()
	for pat1 != nil {
		for pat2 != nil {
			cnt++
			newNode := BuildNode(pat1.index+pat2.index, pat1.coef*pat2.coef)
			if cnt == 1 {
				list.Head.Next = newNode
				list.Tail = newNode
			}else{
				list.Tail.Next = newNode
				list.Tail = list.Tail.Next
			}
			list.len++
			pat2 = pat2.Next
		}
		pat1 = pat1.Next
		pat2 = list2.Head.Next
	}
	Sort(list)
	cur := list.Head.Next
	for i := 1; i <= list.len; i++ {
		if cur.Next == nil {
			break
		}else if cur.index == cur.Next.index {
			cur.coef = cur.coef+cur.Next.coef
			cur.Next = cur.Next.Next
			list.len--
		}
		cur = cur.Next
	}
	return list
}
func Sort(l *PolyLink){
	l.Head.Next = mergeSort(l.Head.Next)
}

func mergeSort(n *Node) *Node{
	if n == nil || n.Next == nil {
		return n
	}
	q1, q2 := n, n
	for q2.Next !=  nil && q2.Next.Next != nil {
		q2 = q2.Next.Next
		q1 = q1.Next
	}
	left, right := n, q1.Next
	q1.Next = nil

	left = mergeSort(left)
	right = mergeSort(right)

	return merge(left,  right)
}
func merge(x1, x2 *Node) *Node{
	cur := BuildNode(-1,-1)
	flag := 0
	var head *Node
	for x1 != nil && x2 != nil{
		if x1.index >= x2.index {
			if flag == 0 {
				head = x1
				flag++
			}
			cur.Next = x1
			x1 = x1.Next
		}else{
			if flag == 0 {
				head = x2
				flag++
			}
			cur.Next = x2
			x2 = x2.Next
		}
		cur = cur.Next
	}
	if x1 ==nil {
		cur.Next = x2
	}else{
		cur.Next = x1
	}
	return head
}