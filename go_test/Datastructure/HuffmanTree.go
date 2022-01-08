package datastructure

type HuffmanTree struct{
	word rune
	freq int
	Left *HuffmanTree
	Right *HuffmanTree
}

func (h *HuffmanTree)GetFreq() int{
	return h.freq
}
func (h *HuffmanTree)SetFreq(k int){
	h.freq = k
}
func (h *HuffmanTree)GetWord() rune{
	return h.word
}
func (h *HuffmanTree)SetWord(k rune){
	h.word = k
}