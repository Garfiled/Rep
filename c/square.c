#include <stdio.h>

int square(int);

int main(int argc, char const *argv[])
{
	int x;
	scanf("%d", &x);
	printf("%d\n", square(x));
	return 0;
}

int square(int x)
{
	return x * x;
}