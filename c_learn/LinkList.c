#include<stdio.h>
#include<stdlib.h>
#include "List.h"
typedef element Huffman;

Node* findNode(LinkList* l, int k)
{
    Node* find;
    if (l->length == 0)
    {
        find = NULL;
    }
    else if (l->length > 0)
    {
        if (l->length < k || k < 0)//当要找的值在链表外
        {
            find = NULL;
        }
        else if (l->length - 1 == k)//当要找的值为尾节点
        {
            find = l->tail;
        }
        else
        {
            Node* cur = l->head;
            for (int i = 0; i < k; i++)
            {
                cur = cur->next;
            }
            find = cur;
        }
    }
    else
    {
        find = NULL;
        printf("you mean len is 0 ???");
    }

    return find;
}

void listInsert(LinkList* l, element *num, int k)
{
    Node* newNode = (Node*)malloc(sizeof(Node));
    newNode->data = num;
    newNode->pre = NULL;
    newNode->next = NULL;
    if (l->length == 0)
    {
        l->head = newNode;
        l->tail = newNode;
        l->length++;
    }
    else if (l->length > 0)
    {
        if (k == 0)//插到头部
        {
            newNode->next = l->head;
            l->head->pre = newNode;
            l->head = l->head->pre;
            l->head->pre = NULL;
        }
        else if (k == l->length)//插到尾部
        {
            l->tail->next = newNode;
            newNode->pre = l->tail;
            l->tail = l->tail->next;
            l->tail->next = NULL;
        }
        else//插到中间
        {
            Node* cur = findNode(l, k - 1);
            Node* temp = cur->next;
            cur->next = newNode;
            newNode->pre = cur;
            newNode->next = temp;
            temp->pre = newNode;
        }
        l->length++;
    }
}

Node *delectNode(LinkList* list, int k)
{
    if (list->length == 0)
    {
        printf("list is empty!");
        return NULL;
    }
    if (k >= list->length)
    {
        printf("no such node!");
        return NULL;
    }
    if (list->length == 1)
    {
        Node* target = list->head;
        list->head = NULL;
        list->tail = NULL;
        list->length--;
        return target;
    }
    else
    {
        Node* target = findNode(list, k);
        if (target->pre == NULL)//头节点
        {
            list->head = list->head->next;
            list->head->pre = NULL;
        }
        else if (target->next == NULL)//尾节点
        {
            list->tail = list->tail->pre;
            list->tail->next = NULL;
        }
        else
        {
            Node* first = target->pre;
            Node* last = target->next;
            first->next = last;
            last->pre = first;
        }
        list->length--;
        return target;
    }

}

void push(LinkList* list, element *num)//栈的操作，入栈
{
    listInsert(list, num, list->length);
}
element *pop(LinkList* list)//栈的操作，出栈
{
    return delectNode(list, list->length - 1)->data;
}
void enqueue(LinkList* list, element *num)//队列操作，入队
{
    listInsert(list, num, list->length);
}
element *dequeue(LinkList* list)//队列操作，出队
{
    return delectNode(list, 0)->data;
}