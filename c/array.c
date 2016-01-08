#include <stdio.h>

#define SIZE 10

int main(int argc, char const *argv[])
{
	printf("%d %s\n", argc, *argv);
	int n[SIZE]={1, 3, 5, 4, 2, 99, 16};
	int i,sum=0;

	for (i=0;i<SIZE;i++) {
		// n[i]=0;
		sum += n[i];
	}

	printf("%s%13s\n", "Element", "Value");
	for (i=0;i<SIZE;i++)
		printf("%7d%13d\n", i, n[i]);

	printf("%7s%13d\n", "Sum:", sum);

	return 0;
}