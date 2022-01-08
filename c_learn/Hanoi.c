#include<stdio.h>


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
    Hanoi(4, 'A', 'B', 'C');
    return 0;
}