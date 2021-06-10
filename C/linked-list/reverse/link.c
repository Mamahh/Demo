#include <stdio.h>
#include <stdlib.h>

typedef struct Link {
    int data;
    struct Link *next; 
}Linklist;

//头插法翻转
Linklist *revert_link(Linklist *head)
{
    if(head == NULL || head->next == NULL)
        return head;
    Linklist *real_head = head; //为了不改变原有链表数据
    Linklist *temp_head = NULL; //搬运工角色，real_head的头结点。
    Linklist *newhead = NULL; //新建链表
    while(real_head != NULL)
    {
        temp_head = real_head; //temp_head保持为real_head的第一个结点
        real_head = real_head->next; //结点偏移

        temp_head->next = newhead; //指向新链表 1->NULL 2->1 3->2 ...
        newhead = temp_head; //NULL->1 1->2 ...  
    }
    return newhead;
}

//创建
Linklist *Creat_link(int n){
    int len = n;
    Linklist *Phead = NULL;
    for(int i=0;i<len;i++)
    {
        Linklist *p = (Linklist *)malloc(sizeof(Linklist));
        if(p == NULL)
            return p;
        p->data = len - 1 - i;
        p->next = Phead;
        Phead = p;
    }
    return Phead;
}

//打印
void Show_link(Linklist *head)
{
    Linklist *h = head;
    while(h != NULL)
    {
        printf("%4d",h->data);
        h = h->next;
    }
    printf("\n");
}

int main(int argc,char *argv[]){
    Linklist *real,*new;
    int len;
reset:
    printf("请输入您想创建的链表长度:\n");
    scanf("%d",&len);
    getchar();
    if(len < 1)
    {
        printf("输入长度无效，请重新输入\n");
        goto reset;
    }
    real = Creat_link(len);
    printf("初始链表:\n");
    Show_link(real);

    new = revert_link(real);
    printf("翻转后:\n");
    Show_link(new);

    return 0;
}
