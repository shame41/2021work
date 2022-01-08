package datastructure


type SHeapForHuffman []HuffmanTree

func (a SHeapForHuffman) swim(k int){
	for k > 1 {
		if a.less(k, k/2){
			swap(k, k/2)
		}else{
			break
		}
	}
}

func (a *SHeapForHuffman) sink(k int){
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

func (a *SHeapForHuffman) Min() *HuffmanTree{
	Min := (*a)[pq[1]]
	size := len(*a)
	swap(1, size)
	*a = (*a)[:size - 1]
	a.sink(1)
	return &Min
}

func (a *SHeapForHuffman) Build(arr []HuffmanTree) {
	*a = append(*a, HuffmanTree{})
}

func (a *SHeapForHuffman)less(i, j int) bool{
	return (*a)[pq[j]].GetFreq() < (*a)[pq[j+1]].GetFreq()
}
func (a *SHeapForHuffman)IsEmpty() bool{
	return len(*a)-1 == 0
}
func (a *SHeapForHuffman)Add(e HuffmanTree, k int){
	*a = append(*a, e)
	pq = append(pq, k)
	
	qp[pq[len(pq)-1]] = len(pq)-1
	a.swim(len(*a)-1)
}