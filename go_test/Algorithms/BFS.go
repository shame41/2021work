package algorithms

import (
	datastructure "go_test/Datastructure"
)

func BFS(g datastructure.Digraph, s, end int) {
	edgeTo := make([]int, g.Vertex()+1)
	marked := make([]bool, g.Vertex()+1)
	queue := datastructure.BuildQueue()
	queue.Enqueue(s)
	marked[s] = true
	for !queue.IsEmpty() {
		v := queue.Dequeue()
		for _, w := range(g.AdjOfV(v)){
			if !marked[w] {
				edgeTo[w] = v
				marked[w] = true
				queue.Enqueue(w)
			}
		}
	}
	stack := datastructure.BuildStack()
	for x := end; x != s; x = edgeTo[x] {
		stack.Push(x)
	}
	var path []int
	for !stack.IsEmpty(){
		path = append(path, stack.Pop())
	}
	print(s)
	for _, i := range(path) {
		print("->",i)	
	}
	println()
}