from threading import Event

class Foo:
    def __init__(self):
        self.first_done = Event()
        self.second_done = Event()

    def first(self, printFirst: 'Callable[[], None]') -> None:
        # printFirst() outputs "first". Do not change or remove this line.
        printFirst()
        self.first_done.set()  # Signal that first() is done

    def second(self, printSecond: 'Callable[[], None]') -> None:
        self.first_done.wait()  # Wait until first() is done
        printSecond()
        self.second_done.set()  # Signal that second() is done

    def third(self, printThird: 'Callable[[], None]') -> None:
        self.second_done.wait()  # Wait until second() is done
        printThird()
