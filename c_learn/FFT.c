#include <malloc.h>
#include <stdlib.h>
#include <stdio.h>

typedef struct LinkNode{
    int index;
    int coef;
    struct LinkNode * next;
}linknode;

typedef struct Link{
    linknode *head;
    linknode *tail;
    int len;
}link;



link* Mult(link *f, link *l)
{
    linknode *pat1 = f->head->next;
    linknode *pat2 = l->head->next;
    linknode *cur;
    int cnt = 0;
    link *newLink = (link*)malloc(sizeof(link*));

    while(pat1 != NULL)
    {
        while(pat2 != NULL)
        {
            cnt++;
            linknode *newNode = (linknode*)malloc(sizeof(link*));
            newNode->index = pat1->index + pat2->index;
            newNode->coef = pat1->coef * pat2->coef;
            newNode->next = NULL;
            if(cnt == 1)
            {
                newLink->head->next = newNode;
                newLink->tail = newNode;
            }else
            {
                newLink->tail->next = newNode;
                newLink->tail = newLink->tail->next;
            }
            pat2 = pat2->next;
        }
        pat1 = pat1->next;
    }
    return newLink;
}

int main()
{
    link *list1 = (link*)malloc(sizeof(link*));
    link *list2 = (link*)malloc(sizeof(link*));

    linknode a4 = {0, -2, NULL};
    linknode a3 = {1, 6, &a4};
    linknode a2 = {2, -5, &a3};
    linknode a1 = {4, 3, &a2};

    linknode b3 = {1, 3, NULL};
    linknode b2 = {4, -7, &b3};
    linknode b1 = {20, 5, &b2};

    list1->head = &a1;
    list1->tail = &a4;

    list2->head = &b1;
    list2->tail = &b3;

    Mult(list1, list2);
    
}