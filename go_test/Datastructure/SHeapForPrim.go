package datastructure


type SHeapForPrim []Edge
var pq []int
var qp []int


func (a SHeapForPrim) swim(k int){
	for k > 1 {
		if a.less(k, k/2){
			swap(k, k/2)
		}else{
			break
		}
	}
}

func (a *SHeapForPrim) sink(k int){
	size := len(*a) - 1
	for k <= size/2 {
		j := 2*k
		if j < size && a.less(j, j+1){
			//如果右儿子比左儿子大
			//那么j指向左儿子
			j++
		}//此时，j就是最大的儿子
		if !a.less(k, j) {
			//如果儿子比父亲小
			break
		}
		swap(k, j)
	}
}

func (a *SHeapForPrim) Min() Edge{
	Min := (*a)[pq[1]]
	size := len(*a)
	swap(1, size)
	*a = (*a)[:size - 1]
	a.sink(1)
	return Min
}

func (a *SHeapForPrim) Build(arr []Edge) {
	*a = append(*a, Edge{})
}

func swap(i, j int){
	t := pq[i]
	pq[i] = pq[j]
	pq[j] = t
	//交换pq[j]与pq[i]
	t = qp[pq[i]]
	qp[pq[i]] = qp[pq[j]]
	qp[pq[j]] = t
	//交换qp[pq[i]]与qp[pq[j]]
}
func (a *SHeapForPrim)less(i, j int) bool{
	return (*a)[pq[j]].GetWeight() < (*a)[pq[j+1]].GetWeight()
}
func (a *SHeapForPrim)IsEmpty() bool{
	return len(*a)-1 == 0
}
func (a *SHeapForPrim)Add(e Edge, k int){
	*a = append(*a, e)
	pq = append(pq, k)
	
	qp[pq[len(pq)-1]] = len(pq)-1
	a.swim(len(*a)-1)
}