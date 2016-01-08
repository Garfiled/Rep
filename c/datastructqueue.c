#include <stdio.h>
#include <stdlib.h>

// 队列节点
typedef struct node {
	int data;
	struct node *next;
} Queue;

// 队列
typedef struct pointer {
	Queue *front;
	Queue *rear;
} Qpointer;

void QueueInit(Qpointer *qp)
{
	qp->front=NULL;
	qp->rear=NULL;
}

int QueueEmpty(Qpointer *qp)
{
	if (qp->front!=NULL)
		return 1;
	return -1;
}

void QueuePush(Qpointer *qp,int elem)
{
	Queue *que;
	que=(Queue*)malloc(sizeof(Queue));
	que->data=elem;
	que->next=NULL;
	if (qp->rear!=NULL) {
		qp->rear->next=que;
		qp->rear=que;
	} else {
		qp->front=que;
		qp->rear=que;
	}
}

int QueuePop(Qpointer *qp)
{
	int ret;
	if (qp->front!=NULL)
	{
		ret=qp->front->data;
		qp->front=qp->front->next;
		if (qp->front==NULL)
			qp->rear=NULL;
		return ret;
	}
	return -1;
}

void QueuePrint(Qpointer *qp)
{
	Queue *que;
	que=qp->front;
	while (que!=NULL)
	{
		printf("%d ", que->data);
		que=que->next;
	}
	printf("\n");
}

int main(int argc, char const *argv[])
{
	Qpointer qp;
	QueueInit(&qp);
	QueuePush(&qp,20);
	QueuePush(&qp,30);
	QueuePush(&qp,10);
	QueuePrint(&qp);
	printf("%d %d\n",QueuePop(&qp),QueuePop(&qp));
	QueuePrint(&qp);
	return 0;
} 