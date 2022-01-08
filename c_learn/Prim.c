#include <malloc.h>
#include <stdio.h> 
#include <string.h> 

#define graphInMatrix 
#define MAX 10 
#define INF 1226 
#define NONE '\0' 
#define LEN_EXAMPLE 7 
typedef struct _Graph {     
    char verx[MAX];     
    int edgnum;     
    int vernum;     
    int matrix[MAX][MAX]; 
}Graph,*PGraph;  
PGraph create_Graph_example()
{     
    char _verx[LEN_EXAMPLE] = {'1','2','3','4','5','6','7'};     
    int _matrix[][LEN_EXAMPLE] = {         
        { 0 ,INF, 1 , 5 ,INF,INF,INF},         
        {INF, 0 ,INF, 9 , 2 ,INF,INF},         
        { 1 ,INF, 0 , 6 ,INF,INF,INF},         
        { 5 , 9 , 6 , 0 , 8 , 2 , 3 },         
        {INF, 2 ,INF, 8 , 0 ,INF, 7 },         
        {INF,INF,INF, 2 ,INF, 0 , 4 },         
        {INF,INF,INF, 3 , 7 , 4 , 0 },     
    };     
    int i, j;     
    PGraph PG = (PGraph)malloc(sizeof(Graph));     
    PG->vernum = LEN_EXAMPLE;     
    for (i = 0; i < PG->vernum;i++)         
        PG->verx[i] = _verx[i];     
    //初始化邻接矩阵     
    for (i = 0; i < PG->vernum;i++)         
        for (j = 0; j < PG->vernum;j++)
        {             
            PG->matrix[i][j] = _matrix[i][j];         
        }     
        //统计边数     
    PG->edgnum = 0;     
    for (i = 0; i < PG->vernum;i++)         
        for (j = 0; j < PG->vernum;j++)
        {             
            if (i!=j&&PG->matrix[i][j]!=INF)                 
                PG->edgnum++;         
        }     
    PG->edgnum /= 2;     
    return PG; 
}; 
void Prim(Graph G,int start)
{     
    char prim[MAX];//结果数组     
    int index = 0;//为prim数组的索引，表示最小生成树中现有几个结点     
    int dist[MAX];//顶点到当前最小生成树的距离     
    int path[MAX];//顶点被添加时距离哪个树中节点最近     
    int i, j, k,min;     //prim算法生成树的第一个结点是给定的起始节点     
    prim[index++] = G.verx[start];     //初始化 顶点到当前最小生成树的距离dist     
    //每个顶点的dist = start顶点到该顶点的距离     
    for (i = 0; i < G.vernum; i++)         
        dist[i] = G.matrix[start][i], path[i] = start;     
    //将start顶点dist初始化为0     
    dist[start] = 0;        
    //从start顶点开始，即start已经被处理过，遇到时跳过     
    for (i = 0; i < G.vernum;i++)
    {         
        if(start==i)             
            continue;         
        k = 0;         
        min = INF;         
        //在尚未被添加到结果最小生成树的节点中找到dist最小的         
        for (j = 0; j < G.vernum;j++)
        {             
            if(dist[j]!=0&&dist[j]<min)
            {                 
                min = dist[j];                 
                k = j;             
            }         
                        
        }         
        //将该结点k添加到最小生成树中         
        prim[index++] = G.verx[k];         
        printf("%c<->%d\n",G.verx[k],path[k]+1);         
        //结点k的dist标记为0         
        dist[k] = 0;         
        //用新的被添加结点更新其他节点的dist         
        for (j = 0; j < G.vernum;j++)
        {             
            //当节点未被添加且需要更新dist时，更新dist和path             
            if(dist[j]!=0&&G.matrix[k][j]<dist[j])                 
                dist[j] = G.matrix[k][j], path[j] = k;         
        }     
    } 
}
int main()
{     
    Graph *PG = create_Graph_example();     
    Prim(*PG, 0);     
    return 0; 
}
