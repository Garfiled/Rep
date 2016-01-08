# -*- coding: utf-8 -*-
 
import signal,time,socket,sys,re,thread
from Tkinter import *

def send_func():
    global addr
    # global sock,addr,text,
    msg = text.get()
    if msg == "q":
            return
    elif msg.startswith("conn"):
        p = re.compile(r'[^ ]+')
        l = p.findall(msg)
        addr = (l[1],int(l[2]))
        sock.connect(addr)
        sock.send(l[3])
        chat_his['text'] = chat_his['text'] + '\n' + msg
        text.set('')
    else:
        if addr == "":
            sys.stdout.write("not connect the server\n>")
        else:
            sock.send(msg)
            chat_his['text'] = chat_his['text'] + '\n' + msg
            text.set('')

def recv_func (sock):
    while True:
        if addr:
            data = sock.recv(1024)
            if data:
               chat_his['text'] = chat_his['text'] + '\n' + data
        else:
            time.sleep(2)
        
addr = ""
sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
root = Tk(className='ChatRoom')
chat_his = Label(root)
# chat_his['text'] = 'be on your own'
chat_his.pack()
text = StringVar()
text.set('input')
entry = Entry(root)
entry['textvariable'] = text
entry.pack()
button = Button(root)
button['text'] = 'send'
button['command'] = send_func
button.pack()
thread.start_new_thread(recv_func, (sock,)) 
root.mainloop()
 