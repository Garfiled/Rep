#include <stdio.h>

#define SIZE 20

typedef struct 
{
	int data[SIZE];
	int num;
} S_type;

void InitList(S_type*);
bool ListEmpty(S_type*);
void ListInsert(S_type*,int,int);
int ListLength(S_type*);
int GetElem(S_type*,int);
int LocateElem(S_type*,int);
int LocateElem(S_type* s,int e);
void ListDelElem(S_type*,int);
void ListPrint(S_type*);
void ClearList(S_type*);


// 时间复杂度

// o(1)
void ClearList(S_type* sp)
{
	sp->num=0;
}

// o(n)
void ListPrint(S_type* sp)
{
	int i;
	for (i=0;i<sp->num;i++)
	{
		printf("%d ", sp->data[i]);
	}
	printf("\n");
}

// o(n)
void ListDelElem(S_type* sp,int i)
{
	int j;
	if (sp->num>=i) 
	{
		for (j=i;j<sp->num;j++)
			sp->data[j-1] = sp->data[j];
		sp->data[sp->num]=0;
		(sp->num)--;
	}
}

// o(n)
int LocateElem(S_type* sp,int e)
{
	int i;
	for (i=0;i<sp->num;i++) {
		if ((sp->data)[i]==e) {
			return i+1;
		}
	}
	return -1;
}

// o(1)
int GetElem(S_type* sp,int i)
{
	if (i>sp->num) {
		return 0;
	}
	return sp->data[i-1];
}

// o(1)
int ListLength(S_type* sp)
{
	return sp->num;
}

// o(1)
void InitList(S_type* sp)
{
	sp->num = 0;
}

// o(1)
bool ListEmpty(S_type* sp)
{
	if (sp->num==0)
		return true;
	else
		return false;
}

// o(n)
void ListInsert(S_type* s,int addr,int e)
{
	int i;
	if (s->num==SIZE)
	{
		printf("enough elem for this list\n");
		return;
	}
	if (addr<=SIZE && addr>0) {
		if (addr>s->num)
		{
			s->data[s->num]=e;
			s->num++;
		} else {
			for (i=s->num;i>=addr;i--)
			{
				s->data[i]=s->data[i-1];
			}
			s->data[addr-1]=e;
			s->num++;
		}
	} else {
		printf("index out of range%d %d\n", addr,SIZE);
	} 
}

int main(int argc, char const *argv[])
{
	S_type s;
	printf("%d\n", sizeof s);
	InitList(&s);
	printf("%d %i %d\n",s.num,ListEmpty(&s),ListLength(&s)); 
	ListInsert(&s,1,10);
	ListInsert(&s,2,20);
	printf("%d %i %d\n",s.num,ListEmpty(&s),ListLength(&s)); 
	printf("%d\n", GetElem(&s,2));
	printf("%d\n", LocateElem(&s,20));
	ListPrint(&s);
	ListInsert(&s,3,35);
	ListPrint(&s);
	ListDelElem(&s,2);
	printf("%d %d\n", ListLength(&s),ListEmpty(&s));
	ListPrint(&s);
	// ClearList(&s);
	printf("%d %d\n", ListLength(&s),ListEmpty(&s));
	ListInsert(&s,1,40);
	ListPrint(&s);
	return 0;
}

/* Lanne

从时间复杂度上看，线性表的顺序存储结构方便读取，不易插入和删除
随机读比较快

*/