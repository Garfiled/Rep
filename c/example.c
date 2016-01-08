#include <stdio.h>

void addGlobal();

int x = 1;

int main(int argc, char const *argv[])
{
	printf("%d\n", x);
	addGlobal();
	printf("%d\n", x);
	return 0;
}

void addGlobal()
{
	x++;
}