package algorithms

import datastructure "go_test/Datastructure"

func HuffmanEncoding(freq map[rune]int) *datastructure.HuffmanTree{
	heap := make(datastructure.SHeapForHuffman, 1)
	for key, val := range(freq){
		node := datastructure.HuffmanTree{}
		node.SetFreq(val)
		node.SetWord(key)
		heap.Add(node, val)
	}
	for (len(heap) - 2) !=0 {
		node := datastructure.HuffmanTree{}
		node.Left = heap.Min()
		node.Right = heap.Min()
		node.SetFreq(node.Left.GetFreq()+node.Right.GetFreq())
	}
	return heap.Min()
}