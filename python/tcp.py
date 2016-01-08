

import socket

def tcplink(conn, addr):
	
	while True:
		data = conn.recv(1024)
		print data
		conn.send(data)


sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)  
sock.bind(('',9004))
sock.listen(5)
print 'waiting for the client'
conn,addr = sock.accept()
print 'welcome ',addr
while True:	
	data = conn.recv(1024)
	print '>' + data
	conn.send(str(len(data)))