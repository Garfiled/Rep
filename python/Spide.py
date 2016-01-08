import socket
addr = ('192.168.1.220',9009)
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(addr)

while True:
	p = re.compile(r'[^ ]+')
	l = p.findall(msg)
	addr = (l[1],int(l[2]))
    sock.connect(addr)
    data = sock.recv(512)
    print data
elif addr:
 	sock.send(msg)
 	data = sock.recv(1024)
    print data
else:
	print "not connect server, not define local god!"