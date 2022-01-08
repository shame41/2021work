#include<stdio.h>
#include<stdlib.h>
#include"List.h"

void swap(HuffmanTree arr[], int i, int j)
{
    HuffmanTree temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
}
void swim(HuffmanTree heap[], int k)
{
    while (k > 1)
    {
        if(heap[k].freq > heap[k/2].freq)
            swap(heap, k, k/2);
        else
            break;
    }
}

void sink(HuffmanTree heap[], int k, int len)
{

    int size = len;
    while (k < size/2)
    {
        int key = heap[2*k].freq < heap[2*k + 1].freq?
            2*k:
            2*k + 1;
        if (heap[k].freq > heap[key].freq)
        {
            swap(heap, k, key);
            k = 2*k;
        }else{
            break;
        }
    
    }   
}

HuffmanTree* MinOfHeap(HuffmanTree heap[], int size)
{
    swap(heap, 1, size);
    // HuffmanTree max = heap[size];
    if(size == 1)
        return &heap[size];
    sink(heap, 1, size);
    return &heap[size];
}
void AddToHeap(HuffmanTree heap[], int size, HuffmanTree tree)
{
    heap[size] = tree;
    swim(heap, size);
}
void BuildHeap(HuffmanTree heap[], int len)
{
    HuffmanTree temp[len];
    for (int i = len/2 + 1; i > 0; i--)
    {
        sink(heap, i, len);
    }
    
}
void HuffmanEncoding(HuffmanTree arr[], int len)
{
    BuildHeap(arr, len);
    int size = len - 1;

    while (size != 1)
    {
        HuffmanTree *newTree = (HuffmanTree*)malloc(sizeof(HuffmanTree));

        newTree->left = MinOfHeap(arr, size);
        size--;
        newTree->right = MinOfHeap(arr, size);
        size--;
        newTree->freq = newTree->left->freq + newTree->right->freq;
        AddToHeap(arr, size+1, *newTree);
        size++;
    }
    
}

void HuffmanDecoding(HuffmanTree *t, int freq)
{

    if (t->freq > freq)
    {
        printf("1");
        HuffmanDecoding(t->right, freq);
    }
    else if (t->freq < freq)
    {
        printf("0");
        HuffmanDecoding(t->left, freq);
    }
    else
    {
        return;
    }
    return;
}

int main()
{
    // HuffmanTree empty = {-1, ' ', NULL, NULL};
    // HuffmanTree a = {45, 'a', NULL, NULL};
    // HuffmanTree b = {13, 'b', NULL, NULL};
    // HuffmanTree c = {12, 'c', NULL, NULL};
    // HuffmanTree d = {16, 'd', NULL, NULL};
    // HuffmanTree e = {9, 'e', NULL, NULL};
    // HuffmanTree f = {5, 'f', NULL, NULL};
    // HuffmanTree arr[7] = {empty,a,b,c,d,e,f};

    // HuffmanEncoding(arr, 7);
    // printf("%d\n", arr[1].freq);
    // HuffmanDecoding(&(arr[1]), 13);

    // for (int i = 0; i < ; i++)
    // {
    //     /* code */
    // }
    printf("100\n");
    printf("1100");
}