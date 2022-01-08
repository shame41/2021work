#pragma once
#define element WEdge
#define MAXNUM 10
typedef struct WeightEdge
{
    int v;
    int w;
    int weight;
}WEdge;
typedef struct WeightDiedge
{
    int v;
    int w;
    int weight;
}WDiedge;
typedef struct EdgeWeightGraph
{
    int v;
    int e;
    LinkList *adj[MAXNUM];
}EWG;
typedef struct EdgeWeightDigraph
{
    int v;
    int e;
    LinkList *adj[MAXNUM];
}EWD;


typedef struct Huffman
{
    int freq;
    char word;
    struct Huffman* right;
    struct Huffman* left;
}HuffmanTree;

typedef struct LinkListNode
{
    element data;
    struct LinkListNode* next;
    struct LinkListNode* pre;
}Node;
typedef struct Link
{
    int length;
    Node* head;
    Node* tail;
}LinkList;

Node* findNode(LinkList* l, int k);
void listInsert(LinkList* l, element num, int k);
Node *delectNode(LinkList* list, int k);
void push(LinkList* list, element num);//栈的操作，入栈
element *pop(LinkList* list);//栈的操作，出栈
void enqueue(LinkList* list, element num);//队列操作，入队
element *dequeue(LinkList* list);//队列操作，出队
int digit(int num, int loc);//从1开始，从左开始
// void swap(int arr[], int i, int j);
void myQuickSort(int arr[], int lo, int hi);
void QuickSort(int arr[], int len);