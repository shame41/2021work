package datastructure

type Edge struct{
	v int//边的一个节点
	w int//边的另一个节点
	weight float64//边的权重
}

func (e *Edge) NewEdge(v, w int, weight float64){
	e.w = w
	e.v = v
	e.weight = weight
}
func (e *Edge) GetWeight() float64{
	return e.weight
}
func (e *Edge) GetVertex() int{
	return e.v
}
func (e *Edge) GetOtherVertex() int{
	return e.w
}
func (e *Edge) GetOther(v int) int{
	if v == e.v {
		return e.w
	}else if v == e.w {
		return e.v
	}else{
		return -1
	}
}
