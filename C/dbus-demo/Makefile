#gcc send_signal.c `pkg-config --cflags --libs glib-2.0 gthread-2.0 dbus-1 dbus-glib-1`
CC = gcc
CC_FLAGS = $(shell pkg-config --cflags glib-2.0 gthread-2.0 dbus-1 dbus-glib-1)
CC_FLAGS += -std=c99 -g
LD_FLAGS = $(shell pkg-config --libs glib-2.0 gthread-2.0 dbus-1 dbus-glib-1)
 
all:send receive
send:send.o
	$(CC) -o $@ $+ $(LD_FLAGS)
send.o:send.c
	$(CC) $(CC_FLAGS) -c -o $@ $<
 
receive:receive.o
	$(CC) -o $@ $+ $(LD_FLAGS)
receive.o:receive.c
	$(CC) $(CC_FLAGS) -c -o $@ $<
 
clean:
	rm send.o send receive.o receive