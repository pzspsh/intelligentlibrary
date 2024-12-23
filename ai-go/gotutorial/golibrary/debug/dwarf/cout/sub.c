/*
@File   : sub.c
@Author : pan
@Time   : 2023-11-30 14:04:34
*/
#include <stdio.h>
#include <string.h>
typedef struct Obj
{
    int a;
    int b;
    int c;
} Obj_t;

Obj_t g_a;
Obj_t g_ar[12];

int sub(int a, int b)
{
    memset(&g_a, 0, sizeof(g_a));
    memset(&g_ar, 0, sizeof(g_ar));
    return a - b + sizeof(g_a) + sizeof(g_ar);
}