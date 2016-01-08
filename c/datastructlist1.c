#include <stdio.h>
#include <stdlib.h>

struct Node {
	int data;
	Node *next;
};

void InitList(Node*);
bool ListEmpty(Node*);
void ListInsert(Node*,int,int);
int ListLength(Node*);
int GetElem(Node*,int);
void ListDelElem(Node*,int);
void ListPrint(Node*);
void ClearList(Node*);

void ClearList(Node* n)
{
	Node* temp;
	Node* na;
	na=n->next;
	while (na!=NULL) {
		temp=na->next;
		free(na);
		na=temp;
	}
	n->next=NULL;
}

void ListDelElem(Node* n,int addr)
{
	int i=0;
	Node* temp;
	while (n->next!=NULL && i<addr-1)
	{
		i++;
		n=n->next;
	}
	if (n->next!=NULL) {
		temp=n->next;
		n->next=temp->next;
		free(temp);
		
	}
}

void ListPrint(Node* n)
{
	while (n->next!=NULL)
	{
		n=n->next;
		printf("%d ", n->data);
	}
	printf("\n");
}

int GetElem(Node* n,int addr)
{
	int i=0;
	while (n->next!=NULL && i<addr)
	{
		i++;
		n=n->next;
	}
	if (i==addr)
		return n->data;
	else
		return -1;
}

int ListLength(Node* n)
{
	int i=0;
	while (n->next!=NULL)
	{
		i++;
		n=n->next;

	}
	return i;
}

void ListInsert(Node* n,int addr,int elem)
{
	int i=0;
	Node* temp;
	while (n->next!=NULL && i<addr-1)
	{
		i++;
		n=n->next;
	}
	printf("%d %d\n", i,n->data);
	temp=(Node*)malloc(sizeof(Node));
	temp->data=elem;
	temp->next=n->next;
	n->next=temp;
}

bool ListEmpty(Node* n)
{
	if (n->next==NULL)
		return true;
	return false;
}

void InitList(Node* n)
{
	n->data=0;
	n->next=NULL;
}

int main(int argc, char const *argv[])
{
	Node n,temp;
	InitList(&n);
	printf("%d\n", sizeof n); // 内存对齐
	printf("%i %d\n", ListEmpty(&n),ListLength(&n));
	ListInsert(&n,1,10);
	ListInsert(&n,2,20);
	ListInsert(&n,1,33);
	printf("%i %d\n", ListEmpty(&n),ListLength(&n));
	ListPrint(&n);
	ListDelElem(&n,2);
	ListPrint(&n);
	ClearList(&n);
	ListPrint(&n);
	return 0;
}