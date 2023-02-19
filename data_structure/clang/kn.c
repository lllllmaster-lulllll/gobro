#include <stdio.h>

void main()
{
    int maxNum=max(10,20);
    printf("x与 y 的大小结果%d \n",maxNum);
}
//这里定义函数
//函数也可以定义到外面
int max (int x,int y)
{
    int z;
    z=x>y?x:y;
    return z;

}