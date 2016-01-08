# -*- coding: utf-8 -*-
 
import signal,time,socket,sys,re
 
class AlarmException(Exception):
    pass
def alarm_handler(signum, frame):
    raise AlarmException
 
def my_input():
    signal.signal(signal.SIGALRM, alarm_handler)
    signal.alarm(3)
    try:
        msg = raw_input()
        signal.alarm(0)
        return msg
    except AlarmException:
        return -1
    signal.signal(signal.SIGALRM, signal.SIG_IGN)
    return


addr = ""
sys.stdout.write('>')
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
while True:
    msg = my_input()
    if msg == -1:
        pass
    elif msg :
        if msg == "q":
            break
        elif msg.startswith("conn"):
            p = re.compile(r'[^ ]+')
            l = p.findall(msg)
            addr = (l[1],int(l[2]))
            sock.connect(addr)
            sock.send(l[3])
            data = sock.recv(512)
            sys.stdout.write(data + '\n>')
            sock.setblocking(0)
        elif msg.startswith("@"):
            if addr == "":
                sys.stdout.write("not connect the server\n>")
            else:
                sock.send(msg)
                sys.stdout.write('>')
        else:
            sock.send(msg)
            sys.stdout.write('>')
    else:
        sys.stdout.write('>')

    if addr:
        try:
            data = sock.recv(1024)
            if data:
                sys.stdout.write(data+ '\n>')
        except:
            pass

        
if addr != "":
    sock.send("q")
    sock.close()