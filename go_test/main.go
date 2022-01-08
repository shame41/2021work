package main

import (
	datastructure "go_test/Datastructure"
	"math/rand"
	"time"
)


func main(){
    var arr []int
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < 50; i++ {
        arr = append(arr, rand.Intn(9999))
    }
    blackredtree := datastructure.NewBRT(13)
    // treap := datastructure.NewTreap(13)
    // avl := datastructure.NewAVLT(13)

    for _, k := range(arr){
        blackredtree = blackredtree.AddNode(k)
        // treap.Add(k)
        // avl = avl.Add(k)
    }
    blackredtree.Print()
    blackredtree.DelMin()
    blackredtree.Print()
    println(blackredtree.GetHeight())

}
