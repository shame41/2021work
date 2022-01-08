#include<stdio.h>
#include<time.h>

long long power(int x, int n)
{
    if( 0 == n )
        return(1);
    if( 0 == n % 2 )
        return(power(x*x, n/2));
    return(power(x*x, n/2)*x);
}

long long power3(int x, int n)
{
    if(n == 0)
        return (1);
    if(n == 1)
        return (x);
    if(n == 2)
        return (x*x);
    if(n % 3 == 0)
        return (power3(x*x*x, n/3));
    if(n % 3 == 1)
        return (power3(x*x*x, n/3)*x);
    return (power3(x*x*x, n/3)*x*x);
}

int main()
{
    clock_t start_time, end_time;
    start_time = clock();
    printf("%lld\n", power(19, 5));
    end_time = clock();
    printf("%lu\n", (long)(end_time - start_time));
    return 0;
}