#include<stdio.h>
#include<stdlib.h>
#include<time.h>

int mainpoint(int a[], int len)
{
    int temp = -1;
    int tim = -1;
    for(int i = 0; i < len-1; i++)
    {
        if(a[i] == a[i+1] && tim >= 0)
        {
            temp = a[i];
            tim++ ;
        }else if(a[i] == a[i+1] && tim < 0)
        {
            tim++;
        }
        else
        {
            tim = -1;
        }
    }
    return(temp);
}


int mainpointt(int a[], int len)
{
    if(len == 1)
        return(a[0]);
	else
	{
		if(len&1)//奇数长 
		{
			int b[len/2+1];
			int j=0;
			for(int i=0;i<len-1;i+=2)
			{
				if(a[i]==a[i+1])
				   b[j++]=a[i];
			}
			b[j]=a[len-1];
			return mainpointt(b,j+1);
		}
		else
		{
			int b[len/2];
			int j=0;
			for(int i=0;i<len;i+=2)
			{
				if(a[i]==a[i+1])
					b[j++]=a[i];
			}
        	return mainpointt(b,j+1);
		}
	}
}

int main()
{
    int len = 1000;
    int a[len];
    srand((unsigned)time(NULL));
    for(int i = 0; i<len; i++)
    {
        a[i] = rand()% 200 + 1;
        printf("%d", a[i]);
    }
    printf("\n");
    clock_t start_time, end_time;
    start_time = clock();
    for(int i = 0; i < 10000; i++)
    {
        int k = mainpointt(a, len);
    }
    end_time = clock();
    printf("%lu\n", (long)(end_time - start_time));
    return(0);
}