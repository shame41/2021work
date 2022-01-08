#include<stdio.h>
#include"List.h"

void swap(HuffmanTree arr[], int i, int j)
{
    HuffmanTree temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
}
void swin(HuffmanTree heap[], int k)
{
    while (k > 1)
    {
        if(heap[k] < heap[k/2])
            swap(heap, k, k/2);
        else
            break;
    }
}

void sink(HuffmanTree heap[], int k, int len)
{
    int size = len - 1;
    while (k < size/2 - 1)
    {
        int key = heap[2*k] > heap[2*k + 1]?
            2*k:
            2*k + 1;
        if (heap[k] < heap[key])
        {
            swap(heap, k, key);
            k = 2*k;
        }else{
            break;
        }
    
    }   
}

HuffmanTree MaxOfHeap(HuffmanTree heap[], int len)
{
    int max = heap[1];
    int size = len - 1;
    if(size == 1)
        return max;
    swap(heap, 1, size);
    heap[size] = NULL;
    sink(heap, 1, len - 1);
    return max;
}
void BuildHeap(HuffmanTree heap[], int len)
{
    HuffmanTree temp[len];
    for (int i = len/2 + 1; i > 0; i--)
    {
        sink(heap, i);
    }
    
}