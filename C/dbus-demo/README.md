## Dbus简易通信流程

### Send

1. 建立与 D-Bus 后台的连接

   ```c
   connection = dbus_bus_get(DBUS_BUS_SESSION, &err);
   ```

2. 发送信号

   1. 创建message，提供接口及信号名

      ```c
      msg = dbus_message_new_signal ("/test/signal/Object","test.signal.Type","Test")
      ```

   2. 给message赋值

      ```C
      dbus_message_iter_init_append(msg,&arg);
      dbus_message_iter_append_basic (&arg,DBUS_TYPE_STRING,&sigvalue)
      ```

   3. 将message从连接中发送出去

      ```C
      dbus_connection_send(connection,msg,&serial)
      ```

3. 释放内存

   ```C
   dbus_message_unref(msg)
   ```




### Receive

1. 建立连接

   ```C
   connection = dbus_bus_get(DBUS_BUS_SESSION,&err)
   ```

2. 通知dbus-daemon,让其监听xxx接口（test.signal.Type）的信号

   ```C
   dbus_bus_add_match(connection,"type='signal',interface='test.signal.Type'",&err);
   ```

3. 在循环中监听信号

   ```C
   msg = dbus_connection_pop_message(connection); //从总线获取消息
   dbus_message_iter_get_basic(&arg,&sigvalue); //提取参数的类型和参数
   ```

4. 释放内存

   ```C
   dbus_message_unref(msg)
   ```



