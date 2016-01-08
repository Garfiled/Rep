#include <stdio.h>

/* Lanne  

*/
void ArrPrint(int arr[],int n)
{
	int i;
	for (i=0;i<n;i++)
	{
		printf("%d ", arr[i]);
	}
	printf("\n");
}

/*
直接插入排序算法
时间复杂度 o(n^2 )
*/
void InsertSort(int arr[],int n)
{
	int i,j,x;
	for (i=1;i<n;i++)
	{
		j=i-1;
		x=arr[i];
		for (;j>=0;j--)
		{
			if (arr[j]>x)
				arr[j+1]=arr[j];
			else
				break;
		}
		arr[j+1]=x;
	}

}

/*
简单选择排序 o(n^2)
*/
void SelectSort(int arr[],int n)
{
	int i,x,j,temp;
	for (i=0;i<n;i++)
	{
		x=arr[i];
		for (j=i+1;j<n;j++)
		{
			if (x>arr[j])
			{
				temp=x;
				x=arr[j];
				arr[j]=temp;
			}
		}
		arr[i]=x;
		// ArrPrint(arr,n);
	}
}



int main(int argc, char const *argv[])
{
	int i;
	int arr[10]={32,1,4,7,3,28,50,2,99,9};
	// InsertSort(arr,10);
	SelectSort(arr,10);
	ArrPrint(arr,10);
	return 0;
}