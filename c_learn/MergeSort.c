#include<stdio.h>
#include<stdlib.h>
#include<time.h>

#define memorySize 400//内存大小400

void swap(int a[], int i, int j)
{
    int temp = a[i];
    a[i] = a[j];
    a[j] = temp;    
}

void sort(int arr[], int lo, int hi)
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
    sort(arr, lo, lt-1);
    sort(arr, gt+1, hi);
    
}
void QuickSort(int arr[], int len)
{
    sort(arr, 0, len-1);
}

void BuildHeap(int heap[], int len)
{
    // printf("最小堆构建完成\n");
}
int MinOfHeap(int heap[], int size)
{
    int min = 1000000;
    int key = 0;
    for (int i = 0; i < size; i++)
    {
        if (heap[i] < min)
        {
            min = heap[i];
            key = i;
        }
    }
    swap(heap, key, size - 1);
    return min;
}
void mergesort(int bigBuf[], int len)
{
    int flag = 0;
    int numOfStream = len/memorySize;//串的数量
    if (numOfStream*memorySize < len)
    {
        flag = 1;
    }
    
    int lenOfPerMemory = memorySize/(numOfStream + 1);//单个缓冲区的长度
    FILE *fp = fopen("aaa", "rb+");
    {
        int buf[memorySize];
        for (int i = 0; i < numOfStream; i++)
        {
            // printf("编号%d\n", i);
            rewind(fp);
            fseek(fp, sizeof(int)*i*memorySize, 0);
            fread(buf, sizeof(int), memorySize, fp);

            QuickSort(buf, memorySize);
            rewind(fp);
            fseek(fp, sizeof(int)*i*memorySize, 0);
            fwrite(buf, sizeof(int), memorySize, fp);


        }       
        if (flag == 1)
        {
            rewind(fp);
            fseek(fp, sizeof(int)*numOfStream*memorySize, 0);
            fread(buf, sizeof(int), len - numOfStream*memorySize, fp);

            QuickSort(buf, len - numOfStream*memorySize);

            rewind(fp);
            fseek(fp, sizeof(int)*numOfStream*memorySize, 0);
            fwrite(buf, sizeof(int), len - numOfStream*memorySize, fp);
        }
    }

    // printf("\n");
    rewind(fp);
    int buf[len];
    fread(buf, sizeof(int), len, fp);
    // for (int i = 0; i < len; i++)
    // {
    //     printf("%d ", buf[i]);
    // }
    
    //----------------------------------------------------------------第一阶段完成----------------------------------------------------------------

    int lenOfInput = numOfStream * lenOfPerMemory ;//输出缓冲区大小
    int lenOfOutput = memorySize - lenOfInput;//输入缓冲区大小
    int inputBuf[lenOfInput];//输入缓冲区
    int outputBuf[lenOfOutput];//输出缓冲区
    int exInputBuf[lenOfInput];
    
    for (int i = 0; i < numOfStream; i++)//向输入缓冲区输入数据
    {   
        // printf("读取串编号%d\n", i);
        int *arrPointer = inputBuf;
        arrPointer += i*lenOfPerMemory;
        rewind(fp);
        fseek(fp, sizeof(int)*i*memorySize, 0);
        fread(arrPointer, sizeof(int), lenOfPerMemory, fp);
      
        if (lenOfInput - lenOfPerMemory*(i + 1) < lenOfPerMemory)//如果最后input还有空位
        {
            rewind(fp);
            fseek(fp, sizeof(int)*(i*memorySize + lenOfPerMemory), 0);
            arrPointer += lenOfPerMemory;
            fread(arrPointer, sizeof(int), lenOfInput - lenOfPerMemory*(i + 1), fp);

        }    
        // printf("\n");
        // for (int i = 0; i < lenOfInput; i++)
        // {
        //     printf("%d ", inputBuf[i]);
        // }
        // printf("\n");
            
    }
    if (flag == 1)
    {
        rewind(fp);
        fseek(fp, sizeof(int)*numOfStream*memorySize, 0);
        fread(exInputBuf, sizeof(int), len - memorySize*numOfStream, fp);

    }
    

    //----------------------------------------------------------------第二阶段完成----------------------------------------------------------------
    int i = 0;
    int out = 0;
    int sizeOfHeap = lenOfInput + len - memorySize*numOfStream;
    int heap [sizeOfHeap];
    for (int i = 0; i < sizeOfHeap; i++)
    {
        if (i < lenOfInput)
        {
            heap[i] = inputBuf[i];
        }
        else
        {
            heap[i] = exInputBuf[i - lenOfInput];
        }
    }
    // printf("\n");    
    BuildHeap(heap, sizeOfHeap);
    int numOfWrite = 0;//输出缓冲区输出次数
    int numOfRead = 0;//输入缓冲区被写入次数

    FILE *temp = fopen("temp", "ab+");
    for(int j = 0; j < len; j++)
    {

        outputBuf[i] = MinOfHeap(heap, sizeOfHeap);
        sizeOfHeap--;
        // printf("%d\n", outputBuf[i]);
        // }
        
        i++;
        if (i == lenOfPerMemory)
        {
            // printf("输出缓存已满， 输出到文件\n");
            // for (int i = 0; i < lenOfPerMemory; i++)
            // {
            //     printf("%d ", outputBuf[i]);
            // }
            // printf("然后写入文件\n");

            fwrite(outputBuf, sizeof(int), lenOfPerMemory, temp);

            out += lenOfPerMemory;

            i = 0;

            numOfWrite++;
            // printf("第%d次输出成功\n", numOfWrite);
        }
        
        //如果堆(输入缓冲区)中刚好能存一个缓冲区的数字，则读入一个缓冲区长度
        if (lenOfInput + len - memorySize*numOfStream - sizeOfHeap == lenOfPerMemory)
        {
            
            // printf("输入缓存出现空串，尝试从文件中读入一个串缓冲区\n");
            for (int i = 0; i < numOfStream; i++)
            {
                int start = memorySize*i + lenOfPerMemory + (lenOfPerMemory/numOfStream)*numOfRead;
                int size = (sizeOfHeap + lenOfPerMemory/numOfStream > lenOfInput + len - memorySize*numOfStream) ? 
                            lenOfInput + len - memorySize*numOfStream - sizeOfHeap ://堆加入size数量的数后如果会比heap长
                            lenOfPerMemory/numOfStream;//正常
                if (start >= memorySize*(i + 1))
                {
                    numOfRead--;
                    break;
                }    
                if (start + size >= memorySize*(i + 1))
                {
                    int *arrPointer = heap;
                    rewind(fp);
                    fseek(fp, sizeof(int)*start, 0);
                    arrPointer += sizeOfHeap;
                    fread(arrPointer, sizeof(int), (i + 1)*memorySize - start, fp);

                    sizeOfHeap+= (i + 1)*memorySize - start;
                }    
                else
                {
                    int *arrPointer = heap;
                    rewind(fp);
                    fseek(fp, sizeof(int)*start, 0);
                    arrPointer += sizeOfHeap;
                    fread(arrPointer, sizeof(int), size, fp);

                    sizeOfHeap+= size;
                }    
                
            }    
            numOfRead++;
            // printf("第%d次读入内存成功\n", numOfRead);
    
            // printf("\n");
            // for (int i = 0; i < sizeOfHeap; i++)
            // {
            //     printf("%d ", heap[i]);
            // }
            // printf("\n");

        }

        
        
    }
    if (i > 0)
    {
        // printf("最后一次输出 ");

        fwrite(outputBuf, sizeof(int), i, temp);        

        out += i;
        // for (int j = 0; j < i; j++)
        // {
        //     printf("%d\n", outputBuf[j]);
        // }
        
        // printf("out is %d\n", out);

    }
    fclose(temp);

}
int main(int argc, char const *argv[])
{
    int len = 7500;
    int aa;
    int bigbuf[len];
    srand((unsigned)time(NULL));
    FILE *fp = fopen("aaa", "w+b");
	for (int i = 0; i < len; i++)
	{
		aa = rand() % 4000 + 1;//向文件中输出len个随机数，它们全都小于4000
        bigbuf[i] = aa;
	}
    fwrite(bigbuf, sizeof(int), len, fp);
    rewind(fp);
    fclose(fp);


    mergesort(bigbuf, len);

    printf("接下来输出结果\n");
    int aaa[len];
    FILE *temp = fopen("temp", "rb+");
    fread(aaa, sizeof(int), len, temp);
    for (int i = 0; i < len; i++)
    {
        printf("%d ", aaa[i]);
    }
    
    return 0;
}
