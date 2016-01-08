#include <stdio.h>
#include <stdlib.h>

int main(int argc, char const *argv[])
{
	int arr[2]={0,0};
	int i;
	for (i=0;i<1000;i++)
	{
		arr[rand()%2] += 1;
	}
	printf("%d %d", arr[0], arr[1]);
	return 0;
}