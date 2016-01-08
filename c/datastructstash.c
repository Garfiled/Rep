#include <stdio.h>
#include <stdlib.h>

typedef struct {
	int* addr;
	int num;
	int length;
} Stash;

void InitStash(Stash*);
void Push(Stash*,int);
int Pop(Stash*);
void Del(Stash*);
void Print(Stash*);

void Del(Stash* s)
{
	free(s->addr);
}

void Print(Stash* s)
{
	int i;
	for (i=0;i<s->num;i++)
	{
		printf("%d ", s->addr[i]);
	}
	printf("\n");
}

int Pop(Stash* s)
{
	if (s->num>0) {
		s->num--;
		return s->addr[s->num-1];
	}
}
void Push(Stash* s,int e)
{
	void* p;
	if (s->num<s->length) {
		s->addr[s->num]=e;
		s->num++;
	} else {
		p=realloc(s->addr, (s->length)*2*sizeof(int));
		if (p==NULL)
			exit(1);
		s->addr=(int*)p;
		s->addr[s->num]=e;
		s->num++;
		s->length *= 2;
	}
}

void InitStash(Stash* s)
{
	void* p;
	p=malloc(4*sizeof(int));
	if (p!=NULL)
	{
		s->addr=(int*)p;
		s->num=0;
		s->length=4;		
	} else {
		exit(1);
	}

}

int main(int argc, char const *argv[])
{
	Stash s;
	InitStash(&s);
	Push(&s,20);
	Push(&s,5);
	Push(&s,10);
	Print(&s);
	return 0;
}