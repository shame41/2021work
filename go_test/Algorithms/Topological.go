package algorithms

import datastructure "go_test/Datastructure"

var reverse datastructure.Stack

func dfsfortopo(g datastructure.Digraph ,s int) {
	marked = make([]bool, g.Vertex()+1)
	reverse = datastructure.BuildStack()
	dfst(g, s)
}
func dfst(g datastructure.Digraph, s int){
	marked[s] = true
	for _, v := range(g.AdjOfV(s)){
		if !marked[v] {
			dfst(g, v)
		}
	}
	reverse.Push(s)
}
func Topological(g datastructure.Digraph) [][]int{
	topo := make([][]int, g.Vertex()+1)
	if FindCycle(g) {
		println("have cycle!!!")
		return nil
	}else{
		for i := 0; i <= g.Vertex(); i++ {
			dfsfortopo(g, i)
			for !reverse.IsEmpty(){
			topo[i] = append(topo[i], reverse.Pop())
			}
		}
		return topo
	}
}
