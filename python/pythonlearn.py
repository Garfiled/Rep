def fact(n):
	if n == 0:
		return 1
	return n * fact(n-1)

print 'hello world!'
while True:
	msg = raw_input(">")
	if msg == 'q':
		break
	elif msg == '':
		continue
	elif msg.isdigit():
		print fact(int(msg))
	else:
		print "undefine"

