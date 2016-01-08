import ssdb,threading

num = 200
thread_list = []
def worker(i, con):
	value = "liu" + str(i)
	con.multi_set(set_name=value)
for i in range(num):
	con = ssdb.SSDB("localhost",8888)
	t = threading.Thread(target=worker,args=(i,con))
	thread_list.append(t)

for t in thread_list:
	t.start()

for t in thread_list:
	t.join()

print "Done!"