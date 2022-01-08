#include<stdio.h>
#include<string.h>
#define min(a, b) (a)<(b)?(a):(a)
#define minOfThree(a, b, c) (min(a, b)) < (c)?(min(a, b)):(c)
int edgeDistance(char *a, char *b)
{
    int lenOfFirst = strlen(a);
    int lenOfSecond = strlen(b);

    int dp[lenOfFirst+1][lenOfSecond+1];
    for (int i = 0; i < lenOfFirst; i++)
    {
        dp[i][0] = i;//情况一
    }
    for (int i = 1; i < lenOfSecond; i++)
    {
        dp[0][i] = i;//情况二
    }
    for (int i = 2; i < lenOfFirst; i++)
    {
        for (int j = 0; j < lenOfSecond; j++)
        {
            int temp = a[i-1] == b[j-1]?0 : 1;
            dp[i][j] = minOfThree(  dp[i-1][j] + 1,
                                    dp[i][j-1] + 1,
                                    dp[i-1][j-1] + temp);
        }
        return dp[lenOfFirst][lenOfSecond];
    }
    
    
    
}

int main(int argc, char const *argv[])
{
    char a[10] = "abcefg\0";
    char b[10] = "bbdefg\0";
    printf("%d\n", 2);
    return 0;
}
