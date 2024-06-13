from threading import Thread
from threading import Condition

class Queue:
    def __init__(self, max_size):
        self.max_size = max_size
        self.curr_size = 0
        self.buffer = []
        self.lock = Condition()

    def enqueue(self, element):
        self.lock.acquire()
        while self.curr_size == self.max_size:
            # Wait while the queue is full
            self.lock.wait()

        # Increment buffer size, append the element
        self.buffer.append(element)
        self.curr_size += 1

        self.lock.release()
        # Notify waiting threads
        self.lock.notifyAll()

    def dequeue(self):
        self.lock.acquire()

        while self.curr_size == 0:
            # Wait till there is anything to deque
            self.lock.wait()

        # Pop the item, update size
        element = self.buffer.pop(0)
        self.curr_size -= 1

        self.lock.release()
        self.lock.notifyAll()

        return element


# Define test cases