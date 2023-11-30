/*
@File   : main.c
@Author : pan
@Time   : 2023-11-30 14:03:58
*/
#include <stdio.h>

extern int add(int a, int b);
extern int sub(int a, int b);

int main()
{
    int a = 10;
    int b = 5;
    int c = 0;
    int d = 0;

    c = add(a, b);
    d = sub(a, b);
    printf("a=%d,b=%d\n", a, b);
    return 0;
}