#include <stdio.h>

typedef struct  {
	long long a1;
	short a2;
} S;

int main(int argc, char const *argv[])
{
	S s;
	printf("%d %d %d\n", sizeof s, sizeof(long long), sizeof(short));
	return 0;
}