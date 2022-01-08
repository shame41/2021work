#include<stdio.h>
#include<time.h>

int fib(int n)
{
    int first = 1;
    int second = 1;
    int third;
    for (int i = 3; i <= n; i++)
    {
        third = first + second;
        first = second;
        second = third;
    }
    return(second);
}

int lazyfib(int n)
{
    
    if(n == 1 || n == 2)
        return 1;
    return(lazyfib(n-1)+lazyfib(n-1));
}

void Hanoi(int n, char a, char b, char c)
{
    if (n == 0)
    {
        return;
    }
    
    Hanoi(n-1, a, c, b);
    printf("%c -> %c \n", a, c);
    Hanoi(n-1, b, a, c);
}

int main()
{
    clock_t start_time, end_time;
    start_time = clock();
    Hanoi(35, 'A', 'B', 'C');
    end_time = clock();
    printf("%lu\n", (long)(end_time - start_time));
    return 0;
}