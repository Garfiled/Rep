#include <stdio.h>

void add(int flag)
{
	int i[5];
	i[0] = 1;
	printf("%p %p\n",i,&i[1]);
	// printf("k is %d\n", k[0],k[1],k[2]);
	// printf("l is %d\n", l[0],l[1],l[2]);  don't use undeclare
}

int main(void)
{
	add(1);
	return 0;
}
