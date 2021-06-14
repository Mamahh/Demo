#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

//定义一个链表
typedef int datatype;
typedef struct  List
{
	datatype data;
	struct List *next;
}node,*list;

struct node *tail; //预定尾节点 

//创建一个单链表
list Create_list(datatype n)
{
	list p = (list)malloc(sizeof(list));
	if(p != NULL)
		p->next = NULL; //创建头结点
	
	for(int i=0;i<n;i++)
	{
		list new = (list)malloc(sizeof(list));
		new->next = NULL;

		new->data = n - i - 1; 
		new->next = p; //头插
		p = new;
	}
	return p;
}

//删除链表的一个值为-d的结点
bool Del_list(datatype d,list head)
{	
	list p = head;
	datatype data = -d;

	if(p == NULL)
		return false;
	while(p->next != NULL) //找到要删除节点的前一个节点
	{
		if(data == p->next->data)
			break;
		p = p->next;
	}
	list del = p->next;
	p->next = del->next;
	free(del);

	return true;
}

//摧毁一个链表
void Destory_list(list head)
{
	list p = head;
	list tmp = NULL;
	if(p == NULL)
		return;
	while(p != NULL)
	{
		tmp = p->next;
		free(p);
		p = tmp;
	}
}

//清理一个链表(保留头结点)
void Clean_list(list head)
{
	list p = head->next;
	list tmp = NULL;
	if(p == NULL)
		return;
	while(p != NULL)
	{
		tmp = p->next;
		free(p);
		p = tmp;
	}
}

//打印链表
void Show_list(list head)
{
	list p = head;
	if(p == NULL)
		return;
	while(p->next != NULL)
	{
		printf("%4d",p->data);
		p = p->next;
	}
	printf("\n");
}


int main(){
	list p = (list)malloc(sizeof(list));
	
	p = Create_list(6);//创建链表
	printf("初始链表：\n");
	Show_list(p);//打印链表

	printf("删除后链表：\n");
	Del_list(-3,p);
	Show_list(p);

	// printf("清空后链表：\n");
	// Clean_list(p);
	// Show_list(p);

	// printf("摧毁后链表：\n");
	// Destory_list(p);
	// Show_list(p);
	return 0;
}