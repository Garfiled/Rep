import socket,time

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
print sock.connect(('192.168.1.220',8000))
data = r"GET / HTTP/1.1"
print sock.send(data)
while True:
	ret = sock.recv(1024)
	print "debug"
	time.sleep(1)
	print ret