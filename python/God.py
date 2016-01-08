import socket
addr_server = ('',9009)
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.bind(addr_server)
sock.listen(10)
users = {}

def tcplink(sock, addr):
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


while True:
	s, addr = sock.accept()
    t = threading.Thread(target=tcplink, args=(s, addr))
    t.start()