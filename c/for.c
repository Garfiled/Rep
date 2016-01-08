#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#define SIZE 6

int main(int argc, char const *argv[])
{
	int face;
	int roll;
	int frequency[SIZE]={0};

	srand(time(NULL));

	for (roll=1;roll<=6000;roll++) {
		face = 1 + rand()%6;
		++frequency[face-1];
	}

	printf("%s%17s\n", "Face", "Frequency");
	for (face=0;face<SIZE;face++) {
		printf("%4d%17d\n", face+1,frequency[face]);
	}
	return 0;
}