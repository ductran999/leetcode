from threading import Event

class ZeroEvenOdd:
    def __init__(self, n: int):
        self.n = n
        self.zero_turn = Event()
        self.odd_turn = Event()
        self.even_turn = Event()
        self.zero_turn.set()  # start with zero()

    def zero(self, printNumber: 'Callable[[int], None]') -> None:
        for i in range(1, self.n + 1):
            self.zero_turn.wait()         # wait until it’s zero’s turn
            self.zero_turn.clear()
            printNumber(0)                # always print zero
            if i % 2 == 1:
                self.odd_turn.set()       # signal odd next
            else:
                self.even_turn.set()      # signal even next

    def even(self, printNumber: 'Callable[[int], None]') -> None:
        for i in range(2, self.n + 1, 2):
            self.even_turn.wait()         # wait until even’s turn
            self.even_turn.clear()
            printNumber(i)                # print the even number
            self.zero_turn.set()          # give control back to zero

    def odd(self, printNumber: 'Callable[[int], None]') -> None:
        for i in range(1, self.n + 1, 2):
            self.odd_turn.wait()          # wait until odd’s turn
            self.odd_turn.clear()
            printNumber(i)                # print the odd number
            self.zero_turn.set()          # give control back to zero
