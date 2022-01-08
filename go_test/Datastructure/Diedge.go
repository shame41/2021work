package datastructure

type Diedge struct{
	v int//边的起点
	w int//边的终点
	weight float64//边的权重
}

func (e *Diedge) newDiedge(v, w int, weight float64){
	e.w = w
	e.v = v
	e.weight = weight
}
func (e *Diedge) getWeight() float64{
	return e.weight
}
func (e *Diedge) from() int{
	return e.v
}
func (e *Diedge) to() int{
	return e.w
}

