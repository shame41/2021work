#include<stdio.h>
#include<math.h>
#include<stdlib.h>

typedef struct bucketitem{
    int index;
    int data;
    struct bucketitem* next;
} bucketitem;

int getRadix(int k, int r)
{
    return(k/((int)pow(10.0, (double)r))%10);
}

void link_push(struct bucketitem* src,struct bucketitem* item){

    bucketitem* pre=src;
    if (pre == NULL)
    {
        pre = item;
        item->next = NULL;
        return;
    }
    
    bucketitem* cur=src->next;

    while(cur!=NULL){
        pre=cur;
        cur=cur->next;
    }
    pre->next=item;
    item->next=cur;
}

bucketitem *link_out(struct bucketitem* src){
    bucketitem* pre=src;
    bucketitem* cur=src->next;

   if (cur != NULL)
   {    
       bucketitem *temp = pre;
       pre = cur;
       cur = cur->next;
       return(temp);
   }
   else
   {
       bucketitem *temp = pre;
       pre = NULL;
       return(temp);
   }
   
}

void swap(int a[], int i, int j)
{
    int t = a[i];
    a[i] = a[j];
    a[j] = t;
}

int Select(int a[], int s, int len)
{
    for (int i = s+1; i < len; i++)
    {
        for (int j = i; j > s; j--)
        {
            if(a[j] < a[i])
                swap(a, i, j);
        }
        
    }
}

int bucket(int a[], int len)
{
   bucketitem b[10];
   for (int i = 0; i < len; i++)
   {
       bucketitem *temp = (bucketitem*)malloc(1);
       temp->data = a[i];
       int key = getRadix(a[i], 0);
       temp->index = key;
       link_push(&b[key], temp);
   }
   //把桶中数据逐一放到队列中
    bucketitem *queue;
    int flag = 1;
    int k = 1;
    while(flag == 1)
    {   
        int flag = 0;

        for (int i = 0; i < len; i++)
        {
           int key = getRadix(a[i], k);
           if (key != 0)
           {
               flag == 1;
           }
           link_push(&b[key], link_out(queue));
        }
        k++;
    }
    for (int i = 0; i < len; i++)
    {
        bucketitem *temp = (bucketitem*)malloc(1);
        do
        {
            temp = link_out(&b[i]);
            link_push(queue, temp);
        }while(temp->next != NULL);
    }
    while(queue != NULL)
    {
        printf("%d ", queue->data);
    }
    printf("\n");
}


int main()
{
    int a[] = {8,100,50,22,15,6,1,1000,999,0};
    bucket(a, 10);
    return 0;
}