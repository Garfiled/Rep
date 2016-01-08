#include <stdio.h>


long factorial(long);

int main(int argc, char const *argv[])
{
	int n;
	n=10;
	printf("%d\n", factorial(n));
	return 0;
}

long factorial(long number)
{
	if (number<=1) {
		return 1;
	}
	else {
		return (number * factorial(number-1));
	}
}