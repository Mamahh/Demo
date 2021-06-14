#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

typedef int datatype;
//定义一个双向链表
typedef struct  List
{
	datatype data;
	struct List *pre,*next;
}*list;

/*
创建新链表
 1.初始化
 2.创建新结点
 3.插入新结点
*/
list Create_list(datatype n)
{
	list p = (list)malloc(sizeof(list));
	if(p == NULL)
		return p;
	p->pre = p; //初始化
	p->next = p;

	for(int i=0;i<n;i++)
	{
		list new = (list)malloc(sizeof(list));//创建新结点
		new->data = i; 
		new->pre = p->pre; //尾插
		new->next = p;
		p->pre->next = new;
		p->pre = new;		
	}
	return p;
}

//打印链表数据
void Show_list(list head)
{
	list p = head->next;
	while(p != head)//遍历
	{
		printf("%4d",p->data);
		p = p->next;
	}
	printf("\n");
}

//根据指定'-n'删除d对应结点
void Del_link(datatype n, list head)
{	
	bool flag = false;
	datatype data = -n;
	list p = head->next;
	while(p->next != head)
	{
		if(p->data == data)//返回要删除的结点
		{
			flag = true;
			break;
		}			
		p = p->next;
	}

	/*
	处理好被删除结点的前后结点的连接关系
	*/
	if(flag)
	{
		p->pre->next = p->next;
		p->next->pre = p->pre;
		p->pre = NULL;
		p->next = NULL;
		free(p);
	}
	
}

//销毁链表
void Destory_link(list head)
{
	list p = head->next;
	list tmp = NULL;
	while(p != head)
	{
		tmp = p->next;
		free(p);
		p = tmp;
	}
	free(head);//链表头干掉
}

//清空链表，保留头
void Clear_link(list head)
{
	list p = head->next;
	list tmp = NULL;
	while(p != head)
	{
		tmp = p->next;
		free(p);
		p = tmp;
	}
}

int main(){
	int n = 8;//定义链表的长度
	int del_n = -3;//定义要删除的结点data值，在此负数代表删除。
	list p = (list)malloc(sizeof(list));
	printf("初始链表：\n");
	p = Create_list(n);//创建链表
	Show_list(p);//打印链表

	printf("删除后链表：\n");
	Del_link(del_n,p);
	Show_list(p);

	// printf("清空后的链表：\n");
	// Clear_link(p);
	// Show_list(p);

	// printf("销毁后的链表：\n");
	// Destory_link(p);
	// Show_list(p);
	return 0;
}