package datastructure

type Graph struct {
	v int//顶点数
	e int//边数
	adj [][]int//邻接表
}

func  NewGraph(v int) Graph{
	g := Graph{}
	g.v = v
	g.e = 0
	g.adj = make([][]int, v+1)
	return g
}
func (g *Graph) Edge() int{
	return g.e
}
func (g *Graph) Vertex() int{
	return g.v
}
func (g *Graph) AdjOfV(v int) []int{
	return g.adj[v]
}
func (g *Graph) AddEdge(v, w int){
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.e++
}