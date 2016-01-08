#include <stdio.h>
/*
交换排序
快速排序
*/
void swap(int *a, int *b)  
{  
    int tmp = *a;  
    *a = *b;  
    *b = tmp;  
}

int partition(int a[],int low,int high)
{
	int privoKey=a[low];
	while (low<high)
	{
		while (low<high && a[high]>=privoKey) --high;
		swap(&a[low],&a[high]);
		while (low<high && a[low]<=privoKey) ++low;
		swap(&a[low],&a[high]);
	}
}

void quickSort(int a[],int low,int high)
{
	if (low<high) {
		int privotLoc=partition(a,low,high);
		quickSort(a,low,privotLoc-1);
		quickSort(a,privotLoc+1,high);
	}
}

void PrintSort(int a[],int num)
{
	int i;
	for (i=0;i<num;i++) 
	{
		printf("%d ", a[i]);
	}
	printf("\n");
}

int main(int argc, char const *argv[])
{
	int a[]={90,1,8,6,23,15,11,220,16,10};
	PrintSort(a,10);
	quickSort(a,0,9);
	PrintSort(a,10);
	return 0;
}