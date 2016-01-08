#!/usr/bin/env python
# -*- coding: utf-8 -*-

'a server example which send hello to client.'

import time, socket, threading

def tcplink(sock, addr):
    global user_dict
    print 'Accept new connection from %s:%s...' % addr
    user = sock.recv(512)
    if user:
        user_dict[user] = sock
        sock.send("ok")
    while True:
        # time.sleep(1)
        data = sock.recv(1024)
        if data == 'q' or not data:
            break
        else:
            if data.startswith("@"):
                index = data.find(" ")
                if index != -1:
                    dst = data[1:index]
                    msg = data[index+1:]
                    msg = "From " + user +":" + msg
                    if user_dict.get(dst):
                        user_dict[dst].send(msg)
                    else:
                        print "not the user",dst
                        sock.send("not the user")
                else:
                    print "消息格式不对"
            else:
                print data
    sock.close()

user_dict = {}
s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
addr = ('', 9999)
# 监听端口:
s.bind(addr)
s.listen(10)
print '>Waiting for connection...'
while True:
    # 接受一个新连接:
    sock, addr = s.accept()
    # 创建新线程来处理TCP连接:
    t = threading.Thread(target=tcplink, args=(sock, addr))
    t.start()