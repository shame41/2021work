package datastructure

type Digraph struct {
	v int//顶点数
	e int//边数
	adj [][]int//邻接表
}

func NewDigraph(v int) Digraph{
	g := Digraph{}
	g.v = v
	g.e = 0
	g.adj = make([][]int, v+1)
	return g
}
func (g *Digraph) Edge() int{
	return g.e
}
func (g *Digraph) Vertex() int{
	return g.v
}
func (g *Digraph) AdjOfV(v int) []int{
	return g.adj[v]
}
func (g *Digraph) AddEdge(v, w int){
	g.adj[v] = append(g.adj[v], w)
	g.e++
}
func (g *Diedge) GetWeight() float64{
	return g.weight
}