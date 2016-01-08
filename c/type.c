#include <stdio.h>

typedef struct fraction {
	int num;
	int denum;
} fraction;

/*C program file*/
int main()
{
	printf("welcome to C!\n");
	char ch;
	int i;
	short sh;
	long l;
	float f;
	double d;
	printf("%lu %lu %lu %lu %lu %lu\n", (ch),sizeof(i),sizeof(sh),sizeof(l),sizeof(f),sizeof d);

	fraction pi;	
	pi.num = 22;
	return 0;
}