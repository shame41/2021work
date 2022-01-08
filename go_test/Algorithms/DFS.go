package algorithms

import datastructure "go_test/Datastructure"

var marked []bool
var edgeTo []int

func DFS(g datastructure.Digraph, s,end int) {
	marked = make([]bool, g.Vertex()+1)
	edgeTo = make([]int, g.Vertex()+1)
	dfs(g, s)
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
func dfs(g datastructure.Digraph, s int){
	marked[s] = true
	for _, v := range(g.AdjOfV(s)){
		if !marked[v] {
			edgeTo[v] = s
			dfs(g, v)
		}
	}
}