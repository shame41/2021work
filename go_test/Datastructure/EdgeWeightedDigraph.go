package datastructure

type EdgeWeightDigraph struct{
	v int;//顶点数量
	e int;//边的数量
	adj [][]Diedge//邻接表
}

func (g *EdgeWeightDigraph) newEdgeWeighDigraph(v int){
	g.v = v
	g.e = 0
	g.adj = make([][]Diedge, v)
}
func (g *EdgeWeightDigraph) getVertex() int{
	return g.v
}
func (g *EdgeWeightDigraph) getEdge() int{
	return g.e
}
func (g *EdgeWeightDigraph) addEdge(e Diedge){
	v := e.from()
	g.adj[v] = append(g.adj[v], e)
	g.e++
}