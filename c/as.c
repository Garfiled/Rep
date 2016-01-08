#include <stdio.h>

void add(int flag)
{
	int i,j;
	i = 1;
	j = 2;
	if (flag>0) 
	{
		int k[2];
		k[0]=3;
		k[1]=4;
	} else {
		int l[5];
		l[0]=5;
		l[1]=6;
		l[2]=7;
		l[3]=8;
		l[4]=9;
	}
	// printf("k is %d\n", k[0],k[1],k[2]);
	// printf("l is %d\n", l[0],l[1],l[2]);  don't use undeclare
}

int main(void)
{
	add(1);
	return 0;
}