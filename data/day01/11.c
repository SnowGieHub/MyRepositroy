
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <pthread.h>


void swap_0(int* i,int * j)
{

    *i ^= *j;
    *j ^= *i;
    *i ^= *j;

}

void swap( int * arry,int len)
{

 int temp, i,j=0;
    for(i =0;i < len-1;i++)
    {
        for(j =0;j<len-i-1;j++)

            if(arry[j] < arry[j + 1])
            {
                temp =  arry[j];
                arry[j] = arry[j+1];
                arry[j+1] = temp;

            }            

    }

}


void swap_1(int * arry,int len)
{

    int i,j,temp = 0;

    for(i = 0; i< len ;i++)
    {
        for(j = i + 1;j < len;j++)
        {

            if(arry[j] > arry[i])
            {
                temp = arry[j];
                arry[j] = arry[i];
                arry[i] = temp;


            }
        }
    }

}

void swap_2(int * arry,int len)
{
    int i,j,temp = 0;

    for(i = 1;i < len;i++)
    {
        if(arry[i] > arry[i-1])
        {
            temp = arry[i];
            for(j = i -1;j >= 0 && temp > arry[j];j--)
            {
                arry[j + 1] = arry[j];
    
            }
            arry[j + 1] = temp;

        }
    }

}

void swap_3(int * arry ,int start,int end)
{
    int i= start;
    int j= end;
    int target =arry[start];

    if(i < j)
    {
        while(i < j){

            while(i < j && arry[j] < target){
                j--;
            }
            if(i < j){
            arry[i] = arry[j];
                i++;
            }


            while(i < j && arry[i] > target){
                i++;
            }
            if(i < j){
            arry[j] = arry[i];
            j--;
            }
        }

        arry[i] = target;
        swap_3(arry,start,i -1);
        swap_3(arry,i+1,end);
    }
}

    




int main(int argc, char *argv[])
{
    int arry[] = {10,5,67,7,8,3,23,54,65,76,8};
    int len = sizeof(arry) / sizeof(int);
    int i =0;
    swap_3(arry,0,len -1);
    for(i =0;i< len;i++)
    {

        printf("%d \n",arry[i]);

    }

    int a = 10;
    int b = 20;

    swap_0(&a,&b);

    printf("a = %d , b = %d \n",a,b);





    return 0;


}
