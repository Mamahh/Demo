#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <unistd.h>
#include <dbus/dbus.h>
#include <dbus/dbus-glib.h>

int send_a_signal(char * sigvalue)
{
	DBusError err;
	DBusConnection * connection;
	DBusMessage * msg;
	DBusMessageIter arg;
	dbus_uint32_t serial = 0;
	int ret;

	//步骤1：建立与D-Bus后台的连接
	/* initialise the erroes */
	dbus_error_init(&err);
	/* Connect to Bus*/
	connection = dbus_bus_get(DBUS_BUS_SESSION, &err);
	if(dbus_error_is_set(&err)){
		fprintf(stderr,"Connection Err : %s\n",err.message);
		dbus_error_free(&err);
	}
	if(connection == NULL)
		return -1;
	
	//步骤2：给连接名分配一个well-known的名字作为Bus name,非必须步骤，可用if 0 注释，也可以用名字来检查，是否已经开启了这个应用另外的进程。
	#if 1
		ret = dbus_bus_request_name(connection,"test.signal.source1",DBUS_NAME_FLAG_REPLACE_EXISTING,&err);
		if(dbus_error_is_set(&err)){
			fprintf(stderr,"Name Err : %s\n",err.message);
			dbus_error_free(&err);
		}
	if(ret != DBUS_REQUEST_NAME_REPLY_PRIMARY_OWNER)
		return -1;
	#endif

	//步骤3：发送一个信号
	//给出的这个信号的路径(即可以指向对象)，接口，以及信号名，创建一个Message
	if((msg = dbus_message_new_signal ("/test/signal/Object","test.signal.Type","Test")) == NULL) {
		fprintf(stderr,"Out of Memory!\n");
		return -1;
	}
	
	//给这个信号（message）具体的内容
	dbus_message_iter_init_append(msg,&arg);
	if(!dbus_message_iter_append_basic (&arg,DBUS_TYPE_STRING,&sigvalue)) {
		fprintf(stderr,"Out of Memory!\n");
		return -1;
	}

	//步骤4：将信号从连接中发送
	if(!dbus_connection_send(connection,msg,&serial)){
		fprintf(stderr,"Out of Memory!\n");
		return -1;
	}
	dbus_connection_flush(connection);
	printf("Signal Send\n");

	while(1);

	//步骤5：释放相关的内存
	dbus_message_unref(msg);
	return 0;
}

int main(int argc, char **argv){
	send_a_signal("Hello World");
	return 0;
}