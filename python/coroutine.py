def coroutine(func):
    def ret():
        f = func()
        f.next()
        return f
    return ret
 
 
 
@coroutine
def consumer():
    print "Wait to getting a task"
    while 1:
        n = (yield)
        print "Got ",n
 
 
 
import time
def producer():
    c = consumer()
    while 1:
        time.sleep(3)
        print "Send a task to consumer"
        c.send("task")
        print "send over"
 
if __name__ == "__main__":
    producer()
