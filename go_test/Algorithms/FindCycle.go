package algorithms

import datastructure "go_test/Datastructure"

var markedcycle []bool
var edgeTocycle []int
var onStack []bool
var hasCycle bool

func FindCycle(g datastructure.Digraph) bool{
	markedcycle = make([]bool, g.Vertex()+1)
	edgeTocycle = make([]int, g.Vertex()+1)
	onStack = make([]bool, g.Vertex()+1)
	hasCycle = false
	for v := 0; v < g.Vertex(); v++ {
		if !markedcycle[v] {
			dfsforcycle(g, v)
		}
	}
	return hasCycle
}

func dfsforcycle(g datastructure.Digraph, v int){
	onStack[v] = true
	markedcycle[v] = true
	for _, w := range(g.AdjOfV(v)){
		if onStack[w] {
			hasCycle = true
			return
		}else if !markedcycle[w] {
			edgeTocycle[w] = v
			dfsforcycle(g, w)
		}
	}
	onStack[v] = false
}