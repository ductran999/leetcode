from threading import Semaphore, Barrier

class H2O:
    def __init__(self):
        self.h_sem = Semaphore(2)   # allow 2 hydrogens
        self.o_sem = Semaphore(1)   # allow 1 oxygen
        self.barrier = Barrier(3)   # wait for 3 threads (2H + 1O)

    def hydrogen(self, releaseHydrogen: 'Callable[[], None]') -> None:
        self.h_sem.acquire()
        releaseHydrogen()           # print "H"
        self.barrier.wait()
        self.h_sem.release()        # ready for next round

    def oxygen(self, releaseOxygen: 'Callable[[], None]') -> None:
        self.o_sem.acquire()
        releaseOxygen()             # print "O"
        self.barrier.wait()
        self.o_sem.release()        # ready for next round
