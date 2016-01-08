#include <stdio.h>

int main(int argc, char const *argv[])
{
	char string1[] = "first";
	printf("%d %s %c", sizeof(string1), string1, *string1);
	return 0;
}