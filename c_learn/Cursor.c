#include<stdio.h>
#include<stdlib.h>
#define max 5
typedef struct node
{
	int num;
	int next;
}node;
node cursor[max+1];
void initialize()
{
	int i;
	for(i=0;i<max+1;i++)
		cursor[i].next=(i+1)%(max+1);//完成初始化的操作。 
}
int cursor_malloc()//模仿malloc 
{
	int p=cursor[0].next;
	cursor[0].next=cursor[p].next;
	return p;	 
}
int cursor_free(int p)
{
	cursor[p].next=cursor[0].next;
	cursor[0].next=p;
}
int makeempty(int l)
{
	if(l!=0)
		cursor[l].next=0;
	return 0;
}
int isempty(int l)
{
	return cursor[l].next==0;
}
int islast(int p)
{
	return cursor[p].next==0;
}
int find_previous(int x,int l)
{
	int p=l;
	while(cursor[p].next&&cursor[cursor[p].next].num!=x)
		p=cursor[p].next;
	return p;
}
int find_x(int x,int l)
{
	int p=cursor[l].next;
	while(p&&cursor[p].num!=x)
		p=cursor[p].next;
	return p;
}
void del(int x,int l)
{
	int p=find_previous(x,l),t;
	if(p==0)
		error("no this num!\n");
	else
	{
		t=cursor[p].next;
		cursor[p].next=cursor[t].next;
		cursor_free(t);
	}
}
void insert(int x,int l,int p)
{
	int newcell=cursor_malloc();
	if(cursor[newcell].next==0)
		printf("is overflow!\n");
	cursor[newcell].num=x;
	cursor[newcell].next=cursor[p].next;
	cursor[p].next=newcell;
}
int main()
{
	int i,l=cursor_malloc();
	initialize();
	for(i=0;i<max+1;i++)
		printf("%d- -%d\n",i,cursor[i].next);
	
	insert(5,l,1);
	for(i=0;i<max+1;i++)
		printf("%d-%d-%d\n",i,cursor[i].num,cursor[i].next);
	insert(9,l,2);
	printf("_________\n");
	for(i=0;i<max+1;i++)
		printf("%d-%d-%d\n",i,cursor[i].num,cursor[i].next);
	return 0;
}
