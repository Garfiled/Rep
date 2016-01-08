  # -*- coding: utf-8 -*- 

import urllib2
import json
import gzip 
import StringIO

servers = [("21",'192.168.1.21'),("22","192.168.1.22"),("23","192.168.1.23"),("171","123.57.11.171"),("213","192.168.1.213")]
server= servers[0] #默认服务器
uids = {}
uid = "0" #默认uid
sid = "0" #默认sid
command = r"http://" + server[1] + "/api/"
case_path = r"case.txt"    #测试案例路径
casefile_object = open(case_path)
case_file = casefile_object.readlines()

def check_server (msg):
	global servers,server,command
	if msg == "server":
		print "当前连接服务器：".decode("utf-8"),server
		return
	msg0 = msg[7:]
	msg0 = msg0.strip()
	if not msg0.startswith("-"):
		print "参数选项请加 -".decode("utf-8"),msg0
		return
	r = msg0.find(" ")
	if r == -1:
		option = msg0[1:]
	else:
		option = msg0[1:r]

	if 'c' in option:
		if r == -1:
			print "缺少参数".decode("utf-8")
		else:
			msg1 = msg0[r+1:]
			msg1 = msg1.strip()
			if msg1 == "21":
				server= servers[0]
				command = r"http://" + server[1] + "/api/"
				print "ok"
			elif msg1 == "22":			
				server = servers[1]
				command = r"http://" + server[1] + "/api/"
				print "ok"
			elif msg1 == "23":			
				server = servers[2]
				command = r"http://" + server[1] + "/api/"
				print "ok"
			elif msg1 == "171":			
				server = servers[3]
				command = r"http://" + server[1] + "/api/"
				print "ok"
			elif msg1 == "213":			
				server = servers[4]
				command = r"http://" + server[1] + "/"
				print "ok"
			else:
				print "无效的服务器名称".decode("utf-8"),msg1

	if 'a' in option:
		print "服务器列表:".decode("utf-8")
		for i in servers:
			print i
def check_command (s):
	global command
	global sid
	s0 = s.strip(" ")

	s1 = s0.split(" ")
	r = command + s1[0] + "?"
	login_f = False
	if s1[0] == "login":
		login_f = True
		flag = 0
	else:
		r = r + "sid=" + sid
		flag = 1
	for i in s1[1:]:
		if flag == 1:
			r =r + "&" + i
			flag = 2
		elif flag == 2:
			r = r + "=" + i
			flag = 1
		elif flag == 0:
			r =r + i
			flag = 2
	return [r,login_f]

def check_uid (msg):
	global uids,uid,sid
	if msg == "uid":
		print "当前连接用户名:".decode("utf-8"),uid,uids.get(uid)
		return
	msg0 = msg[3:]
	msg0 = msg0.strip()
	if not msg0.startswith("-"):
		print "参数选项请加 -".decode("utf-8"),msg0
		return
	r = msg0.find(" ")
	if r == -1:
		option = msg0[1:]
	else:
		option = msg0[1:r]
	if 'c' in option:
		if r == -1:
			print "缺少参数".decode("utf-8")
		else:
			uid_in = msg0[r+1:].strip()
			v = uids.get(uid_in)
			if v == None:
				print "未记录在案".decode("utf-8"),uid_in
				return
			else:
				uid = uid_in
				sid = v
				print "ok"
				return 
	if 'a' in option:
		print "uid列表:".decode("utf-8")
		for i in uids:
			print i,uids[i]
def store_uid (r):
	global uids,uid,sid
	s = json.loads(r)
	uid = s["Uid"]
	sid = s["Sid"].replace("+","%2b")
	uids[uid] = sid
def check_case (msg):
	global case_file,case_path
	msg0 = msg[5:]
	msg0 = msg0.strip()
	option = ""
	if msg0.startswith("-"):
		index = msg0.find(" ")
		if index == -1:
			option = msg0[1:]
			msg1 = ""
		else:
			option = msg0[1:index]
			msg1 = msg0[index+1:]
		if 'u' in option:
			casefile_object = open(case_path)
			case_file = casefile_object.readlines()
			print "ok"
	else:
		msg1 = msg0
	# 测试案例行数信息
	# print msg1
	if msg1:
		row_l = []
		msg_l = msg1.split(",")
		for i in msg_l:
			i0 = i.strip()
			if "-" in i0:
				j0 = i0.index("-")
				j1 = int(i0[:j0])
				j2 = int(i0[j0+1:])
				while j1<=j2:
					row_l.append(j1)
					j1 = j1+1
			else:
				row_l.append(int(i0))
		# print row_l
		if row_l:
			for row in row_l:
				if row >=1 and row <= len(case_file):
					rowmsg = case_file[row-1]
					rowmsg = rowmsg.strip()
					if 'r' in option:
						print row,rowmsg
					else:
						if not (rowmsg.startswith("--") or rowmsg.startswith("//") or rowmsg == ""):
							# print "debug",row,rowmsg
							err = check_gm(rowmsg)
							if not err:
								print "错误行:".decode("utf-8"),row
				else:
					print "行数不在case文件内".decode("utf-8"),row
					return
	else:
		if 'r' in option:
			i = 1
			for row in case_file:
				row = row.strip()
				print i,row
				i = i+1
def check_gm (msg):
	err = True
	try:
		l = check_command(msg)
		response = urllib2.urlopen(l[0])
		r= response.read()
		if response.headers.has_key('content-encoding'):	
			r1 = StringIO.StringIO(r)
			gz = gzip.GzipFile(fileobj=r1)
			r = gz.read()
			gz.close()
		print r.decode('utf-8')
		if l[1]:
			store_uid(r)
	except:
		print "error command: ",l[0]
		err = False
	finally:
		return err

if __name__ == '__main__':
	print "当前连接服务器:".decode("utf-8"),  server
	while True:
		msg = raw_input(">")
		msg = msg.strip()
		if msg == "":
			continue
		elif msg == "q":
			break
		elif msg.startswith("server"):
			check_server(msg)
		elif msg.startswith("uid"):
			check_uid(msg)
		elif msg.startswith("case"):
			check_case(msg)
		else:
			check_gm(msg)


