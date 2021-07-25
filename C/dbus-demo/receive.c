#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <dbus/dbus.h>
#include <dbus/dbus-glib.h>

void listen_signal()
{
	DBusMessage * msg;
	DBusMessageIter arg;
	DBusConnection * connection;
	DBusError err;
	int ret;
	char * sigvalue;

	//步骤1：建立与D-Bus后台的连接
	dbus_error_init(&err);
	connection = dbus_bus_get(DBUS_BUS_SESSION,&err);
	if(dbus_error_is_set(&err)){
		fprintf(stderr,"Connection Error %s\n",err.message);
		dbus_error_free(&err);
	}
	if(connection == NULL)
		return;
	
	//步骤2：给连接名分配一个可记忆名test.signal.dest作为Bus name,非必须步骤，推荐处理
	ret = dbus_bus_request_name(connection,"test.signal.dest",DBUS_NAME_FLAG_REPLACE_EXISTING,&err);
	if(dbus_error_is_set(&err)){
		fprintf(stderr,"Name Error %s\n",err.message);
		dbus_error_free(&err);
	}
	if(ret != DBUS_REQUEST_NAME_REPLY_PRIMARY_OWNER)
		return;

	//步骤3：通知D-Bus daemon，希望监听来自接口test.signal.Type的信号
	dbus_bus_add_match(connection,"type='signal',interface='test.signal.Type'",&err);
	//实际需要发送东西给daemon来通知希望监听的内容，所以需要flush
	dbus_connection_flush(connection);
	if(dbus_error_is_set(&err)){
		fprintf(stderr,"Match Error %s\n",err.message);
		dbus_error_free(&err);
	}

	//步骤4：在循环中监听，每隔1秒，就去试图获取此信号。这里给出的是连接中获取所有消息的方式，所以获取后去检查下这个消息是否为期望的信号，并获取内容。也可以通过这个方式来获取method call消息。
	while(1){
		dbus_connection_read_write(connection,0);
		msg = dbus_connection_pop_message(connection);
		if(msg == NULL){
			sleep(1);
			continue;
		}

		printf("listen signal\n");
		if(dbus_message_is_signal(msg,"test.signal.Type","Test")) {
			if(!dbus_message_iter_init(msg,&arg)){
				fprintf(stderr,"Message Has no Param");
				printf("Message Has no Parma\n");
			} 
			else if(dbus_message_iter_get_arg_type(&arg) != DBUS_TYPE_STRING)
				printf("Param is not string\n");
			else
				dbus_message_iter_get_basic(&arg,&sigvalue);
			printf("====Got Signal with value : %s\n",sigvalue);
		}
		//步骤5：释放相关的内存	
		dbus_message_unref(msg);
	}
}

int main(int argc,char **argv){
	listen_signal();
	return 0;
}