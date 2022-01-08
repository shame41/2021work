#include<stdio.h>
#include<time.h>
#include<stdlib.h>
#include"List.h"

void swap(int arr[], int i, int j)
{
    int temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
}
void QuickSort(int arr[], int len)
{
    myQuickSort(arr, 0, len-1);
}
void myQuickSort(int arr[], int lo, int hi)
{
    if(hi <= lo)
        return;
    int lt, gt, i;
    i = lo + 1;
    lt = lo;
    gt = hi;
    int v = arr[lo];
    while(i <= gt)
    {
        if(arr[i] < v)
        {
            swap(arr, lt, i);
            lt++;
            i++;
        }
        else if(arr[i] > v)
        {
            swap(arr, gt, i);
            gt--;
        }
        else
        {
            i++;
        }
    }
    myQuickSort(arr, lo, lt-1);
    myQuickSort(arr, gt+1, hi);
    
}

int main()
{
    int len = 300;
    int arr[len];
    srand((unsigned)time(NULL));
	for (int i = 0; i < len; i++)
	{
		arr[i] = rand() % 4000 + 1;//向文件中输出len个随机数，它们全都小于4000
	}
    QuickSort(arr, len);
    for (int i = 0; i < len; i++)
    {
        printf("%d ", arr[i]);
    }
    
}