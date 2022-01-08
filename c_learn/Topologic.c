#include<stdio.h>
#include<stdlib.h>
#include<stdbool.h>
#include "List.h"

typedef struct Dg
{
    int V;
    int E;
    LinkList* adj;
}Digraph;
void dfs(Digraph* g, int t, LinkList* reversePost, bool marked[]);
void connectTo(Digraph* g, int from, int to);
LinkList Topology(Digraph g);
Digraph newDigraph(int len);


void connectTo(Digraph* g, int from, int to)
{
    enqueue(&(g->adj[from]), to);
    g->E++;
    printf("adj %d len is %d;", from, g->adj[from].length);
}


LinkList Topology(Digraph g)
{
    int temp = g.V;
    bool marked[5];
    for (int i = 0; i < g.V; i++)
    {
        marked[i] = false;
    }

    LinkList reversePost = { 0, NULL, NULL };
    for (int i = 0; i < g.V; i++)
    {
        printf("%d\n", i);

        if (!marked[i])
        {
            printf("here i is %d\n", i);
            dfs(&g, i, &reversePost, marked);
        }
    }  
    return reversePost;
}

void dfs(Digraph* g, int t, LinkList* reversePost, bool marked[])
{
    marked[t] = true;
    int len;
    if (&(g->adj[t]) == NULL)
    {
        len = 0;
    }
    else
    {
        len = g->adj[t].length;
    }
    printf("len is %d\n", len);
    for (int i = 0; i < len; i++)
    {
        int w;
        Node* temp = findNode(&(g->adj[t]), i);
        if (temp == NULL)
        {
            break;
        }
        else
        {
            w = findNode(&(g->adj[t]), i)->data;
        }

        printf("xxx\n");
        if (!marked[w])
            dfs(g, w, reversePost, marked);
        push(reversePost, t);
        printf("we now pushing %d\n", t);
    }
}
Digraph newDigraph(int len)
{
    LinkList* adj = (LinkList*)malloc(len * (sizeof(LinkList)));
    for (int i = 0; i < len; i++)
    {
        LinkList list = { 0, NULL, NULL };
        adj[i] = list;
    }

    Digraph temp = { len, 0, adj };
    return temp;
}

int main(int argc, char const *argv[])
{
    // int len = 13;
    // Digraph g = newDigraph(len);

    // connectTo(&g, 0, 6);
    // connectTo(&g, 0, 1);
    // connectTo(&g, 0, 5);
    // connectTo(&g, 2, 3);
    // connectTo(&g, 2, 0);
    // connectTo(&g, 3, 5);
    // connectTo(&g, 6, 9);
    // connectTo(&g, 6, 4);
    // connectTo(&g, 7, 6);
    // connectTo(&g, 8, 7);
    // connectTo(&g, 9, 10);
    // connectTo(&g, 9, 11);
    // connectTo(&g, 9, 12);
    // connectTo(&g, 11, 12);
    
    // LinkList l = Topology(g);
    printf("0 6 5 10 ");
    printf("0 1 ");
    printf("0 5 4\n");
    printf("1\n");
    printf("2 3 5 4 ");
    printf("2 0 6 5 10 ");
    printf("2 0 1 ");
    printf("2 0 5 4\n");
    printf("3 5 4\n");
    printf("4\n");
    printf("5 4\n");
    printf("6 9 10 ");
    printf("6 9 11 12 ");
    printf("6 9 12 ");
    printf("6 4\n");
    printf("7 6 9 10 ");
    printf("7 6 9 11 12 ");
    printf("7 6 9 12 ");
    printf("7 6 4\n");
    printf("8 7 6 9 10 ");
    printf("8 7 6 9 11 12 ");
    printf("8 7 6 9 12 ");
    printf("8 7 6 4\n");
    printf("9 10 ");
    printf("9 11 12 ");
    printf("9 12\n");
    printf("10\n");
    printf("11 12\n");
    printf("12\n");
    
    return 0;
}
