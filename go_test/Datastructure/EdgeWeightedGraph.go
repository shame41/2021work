package datastructure

type EdgeWeightGraph struct{
	v int;//顶点数量
	e int;//边的数量
	adj [][]Edge//邻接表
}

func (g *EdgeWeightGraph) NewEdgeWeighGraph(v int){
	g.v = v
	g.e = 0
	g.adj = make([][]Edge, v)
}
func (g *EdgeWeightGraph) GetAdj(k int) []Edge{
	return g.adj[k]
}
func (g *EdgeWeightGraph) GetVertex() int{
	return g.v
}
func (g *EdgeWeightGraph) GetEdge() int{
	return g.e
}
func (g *EdgeWeightGraph) AddEdge(e Edge){
	v := e.GetVertex()
	w := e.GetOther(v)
	g.adj[v] = append(g.adj[v], e)
	g.adj[w] = append(g.adj[w], e)
	g.e++
}