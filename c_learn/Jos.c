#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define SizeOfMemory 100

struct Node
{
    int Item;
    int next;
};

struct Node Josephus[SizeOfMemory];

void Initialize();
int Culloc();
void Insert(int, int, int);
int buildList(int) ;
void CFree(int);
void PrintList(int, int);
int FindJosephus(int,int);

int main() 
{
    
    int linklist = buildList(17); // 初始化链表
    clock_t start_time, end_time;
    start_time = clock();
    // for(int i = 0; i < 10000; i++)
    // {
    //     int temp = FindJosephus(19, linklist);
    //     if(i == 0)
    //         printf("%d\n", temp);
    // }
    int temp = FindJosephus(2, linklist);
    printf("%d\n", temp);
    end_time = clock();
    printf("%lu\n", (long)(end_time - start_time));
    return 0;
}



void Initialize()
{
    int i;
    for(i = 0; i < SizeOfMemory; i++)
        if(i == SizeOfMemory - 1)
            Josephus[i].next = 0;
        else
            Josephus[i].next = i + 1;
}

int Culloc()
{
    int P;

    P = Josephus[0].next;
    Josephus[0].next = Josephus[P].next;
    
    return P;
}

void Insert(int X, int L, int P)
{
    int Address;

    Address = Culloc();
    if (Address == 0)
        printf("Out of space!\n");
    Josephus[Address].Item = X;
    Josephus[Address].next = Josephus[P].next;
    Josephus[P].next = Address;
}

int buildList(int IndexOfNode) 
{
    if(IndexOfNode>=SizeOfMemory)
    {
        printf("没有内存了！！！");
        return -1;
    }
    Initialize();
    int L = Culloc();
    Josephus[L].next = 0;
    int P = L;

    for(int i = 0; i< IndexOfNode ; ++i)
    {
        Insert(i+1,L,P);
        P = Josephus[P].next;
    }

    Josephus[P].next = Josephus[L].next;

    return L;
}

void CFree(int P)
{
    Josephus[P].next = Josephus[0].next;
    Josephus[0].next = P;
}

void PrintList(int L, int n)
{
    int p = L;
    if (!Josephus[p].next)
    {
        printf("链表为空!\n");
        return;
    }

    int P;

    P = Josephus[L].next;
    for(int i = 0 ; i< n ; ++i)
    {
        printf("%d", Josephus[P].Item);
        P = Josephus[P].next;
    }
    printf("\n");
    return;
}

int FindJosephus(int k,int L)
{
    int cur = L;
    int i;
    int prev = cur, tmp;

    if (Josephus[cur].next == Josephus[Josephus[cur].next].next) 
    {
        printf("%d\n",Josephus[Josephus[cur].next].Item);
        return 0;
    }

    while(Josephus[cur].next != Josephus[Josephus[cur].next].next)
    {
        for(i = 1, tmp = Josephus[L].next, prev = L; i < k; i++)
        {
            tmp = Josephus[cur].next;
            prev = tmp;
            tmp = Josephus[tmp].next;
        }
        Josephus[cur].next = Josephus[tmp].next;
        Josephus[prev].next = Josephus[tmp].next;
        CFree(tmp);
    }
    return Josephus[Josephus[cur].next].Item;

}

