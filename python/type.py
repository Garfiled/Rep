d = {'a':1,'b':2,'c':3}
for key in d:
	print key,d[key]
print '-----------------'
for i ,value in enumerate(['A','B','C']):
	print i,value
print '-----------------'
for x,y in [(1,1),(2,4),(3,9)]:
	print x,y
print '-----------------'
print [x*x for x in range(1,11)]