#include <stdio.h>


int main(int argc, char const *argv[])
{
	printf("%d %s\n", argc, *argv);

	for (int i=0;i<argc;i++)
	{
		printf("%d %s\n", i,argv[i]);
	}
	return 0;
}