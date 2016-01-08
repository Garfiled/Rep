#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef struct _node {
	char *name;
	char *desc;
	struct _node *next;
} node;

#define HASHSIZE 256
static node* hashtab[HASHSIZE];


void InitHashtab() {
	int i;
	for (i=0;i<HASHSIZE;i++)
		hashtab[i]=NULL;
}

unsigned hash(char *s)
{
	unsigned hashval;

	for (hashval=0;*s!='\0';s++)
		hashval=*s+31*hashval;
	return hashval % HASHSIZE;
}

node* lookup(char *n)
{
	unsigned int hi=hash(n);
	node* np=hashtab[hi];
	for (;np!=NULL;np=np->next)
	{
		if (!strcmp(np->name,n))
			return np;
	}
	return NULL;
}

char* m_strdup(char *o)
{
	int l=strlen(o)+1;
	char *ns=(char*)malloc(l*sizeof(char));
	strcpy(ns,o);
	if (ns==NULL)
		return NULL;
	else
		return ns;
}

void* get(char* name)
{
	node* n=lookup(name);
	if (n==NULL)
		return NULL;
	else
		return n->desc;
}

int install(char* name,char* desc)
{
	unsigned int hi;
	node* np;

	if ((np=lookup(name))==NULL) {
		hi=hash(name);
		np=(node*)malloc(sizeof(node));
		if (np==NULL)
			return 0;
		np->name=m_strdup(name);
		if (np->name==NULL) return 0;
		np->next=hashtab[hi];
		hashtab[hi]=np;
	} else {
		free(np->desc);
	}
	np->desc=m_strdup(desc);
	if (np->desc==NULL) return 0;
	return 1;
}

void displaytable() 
{
	int i;
	node* t;

	for (i=0;i<HASHSIZE;i++)
	{
		t=hashtab[i];
		if (t!=NULL) {
			printf("---------- %d -----------\n", i);
			while (t!=NULL) {
				printf("%s %s\n",t->name,t->desc);
				t=t->next;
			}
			printf("----------------------\n", i);
		}
	}
}

int main(int argc, char const *argv[])
{
	// char s[]="lanne";
	// printf("%d\n", hash(s));
	InitHashtab();
	printf("%d\n", sizeof hashtab);
	char name1[]="liu";
	char desc1[]="xiang";
	char name2[]="wu";
	char desc2[]="sheng";
	// char* names[]={"liu","wu"};
	// char* descs[]={"xiang","sheng"}
	install(name1,desc1);
	install(name2,desc2);
	displaytable();
	printf("%s\n", get(name1));
	return 0;
}