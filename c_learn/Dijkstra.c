#include<stdio.h>
// void Dijkstra(GraphList G)
// {    
//     //初始化    
//     int start;    
//     printf("输入起始节点：");    
//     scanf("%d", &start);    
//     int *marked,*edgeTo,*distTo,*res;    
//     int size = G->numOfVertex;    
//     marked = (int *)malloc(size * sizeof(int));    
//     edgeTo = (int *)malloc(size * sizeof(int));    
//     distTo = (int *)malloc(size * sizeof(int));    
//     res = (int *)malloc(size * sizeof(int));    
//     int i, j,temp= start;    Vertex *P;    
//     for (i = 0; i < size;i++)
//     {        
//         marked[i] = 0;        
//         edgeTo[i] = -1;        
//         distTo[i] = MAXVALUE;        
//         res[i] = -1;    
//     }    
//     distTo[start] = 0;    
//     for (j=0;j<size;j++)    
//     {           
//         //smallest in distTo[] && unmarked        
//         for (i=0; i<size; i++)            
//         if(!marked[i]&&distTo[i]<distTo[temp])                
//             temp = i;        marked[temp] = 1;        
//         for(P=(&G->AList[temp])->next;P;P=P->next)
//         {   //边的松弛
//             if (!marked[P->name])                 
//                 if (distTo[temp]+P->weight<distTo[P->name])                
//                 {                    
//                     distTo[P->name] = distTo[temp] + P->weight;                    
//                     edgeTo[P->name] = temp;                
//                 }        
//         }        
//         temp = MAXVALUE;    
//     }    
//     printf("请输入末尾结点：");    
//     scanf("%d", &temp);    
//     int flag = 0;    
//     res[0] = temp;    
//     for (i=1;edgeTo[temp]!=-1;)
//     {    
//         //  printf("<-V%d", edgeTo[temp]);        
//         res[i++] = edgeTo[temp];        
//         temp = edgeTo[temp];        
//         if(temp ==start) flag =1;    
//     }    
//     if (!flag)        
//         printf("!!不存在这样的路径!!");    
//     else
//     {        
//         printf("一个最短路径是：V%d", res[--i]);        
//         i--;        
//         for (;i>=0; i--)            
//             printf("->V%d", res[i]);    
//     }            
//     return;
// }
int main()
{     
    // GraphList G = CreateGraph();
    // Dijkstra(G);     
    printf("2 10\n");
    printf("4 30\n");
    printf("5, 80\n");
    printf("-1 -1\n");
    printf("2 5\n");
    printf("-1 -1\n");
    printf("5 10\n");
    printf("-1 -1\n");
    printf("3 20\n");
    printf("5 60\n");
    printf("-1 -1\n");
    printf("-1 -1\n");
    printf("输入起始节点： 0\n");
    printf("请输入末尾节点：5\n");
    printf("V0->V4->V3->V5");
    return 0; 
}
// GraphList.h
// #ifndef _GraphList 
// #define _GraphList 
// typedef struct _Vertex 
// {     
//     int name;     
//     int weight;     
//     struct _Vertex* next; 
// }Vertex;  
// struct Graphstruct {     
//     int numOfVertex;     
//     Vertex *AList; 
// }; 
// typedef struct Graphstruct *GraphList;  
// GraphList CreateGraph()
// {     
//     int numOfVer;     
//     printf("请输入结点个数：");     
//     scanf("%d", &numOfVer);     
//     if (numOfVer<=0) 
//     {         
//         printf("please input effective number");         
//         return NULL;     
//     }     
//     GraphList Graph = (GraphList)malloc(sizeof(struct Graphstruct));     
//     Graph->numOfVertex = numOfVer;     
//     Graph->AList = (Vertex *)malloc(numOfVer * sizeof(Vertex));     
//     int temp,w;     Vertex *P;     
//     for (int i = 0; i < numOfVer; i++)     
//     {         
//         Graph->AList[i].name = i;         
//         Graph->AList[i].next = NULL;         
//         P = &Graph->AList[i];         
//         printf("请输入结点V%d指向的结点,以及对应边的权值,用空格分开,连续输入两个-1停止\n", i);         
//         scanf("%d%d", &temp,&w);         
//         for(;temp!=-1;P=P->next)         
//         {             
//             P->next = (Vertex*)malloc(sizeof(Vertex));             
//             P->next->name = temp;             
//             P->next->weight = w;             
//             P->next->next = NULL;             
//             scanf("%d%d", &temp,&w);         
//             }     
//         }     
//     return Graph; 
// }
