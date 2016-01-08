#include <stdio.h>
#include <ctype.h>

void convertToUppercase(char *);

int main(int argc, char const *argv[])
{
	char string[] = "characters and $32.98";

	printf("%s\n", string);
	convertToUppercase(string);
	printf("%s\n", string);
	return 0;
}

void convertToUppercase(char *sPtr)
{
	while (*sPtr!='\0') {
		if (islower(*sPtr)) {
			*sPtr = toupper(*sPtr);
		}
		++sPtr;
	}
}