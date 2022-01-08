package algorithms

import datastructure "go_test/Datastructure"

var mst []datastructure.Edge
func PrimLazy(g datastructure.EdgeWeightGraph) []datastructure.Edge{
	marked := make([]bool, g.GetVertex())
	pq := make(datastructure.SHeapForPrim, g.GetEdge())
	
	visit(g, 0)
	for !pq.IsEmpty(){
		e := pq.Min()

		v := e.GetVertex()
		w := e.GetOther(v)
		if !marked[v] && !marked[w] {
			continue
		}
		if !marked[v] {
			visit(g, v)
		}
		if !marked[w] {
			visit(g, w)
		}
	}
	return mst
}

func visit(g datastructure.EdgeWeightGraph, v int){
	for _, e := range(g.GetAdj(v)){
		if !marked[e.GetOther(v)] {
			mst = append(mst, e)
		}
	}
}
