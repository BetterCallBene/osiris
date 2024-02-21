import signal
import sys
import time

def interrupt_signal_handler(_signal_number, _frame):
    """SIGINT/SIGTSTP handler for gracefully stopping an application."""
    print("Caught interrupt signal. Stop application!")
    sys.exit(0)


def run():
    
    while True:
        print("I am still running...")
        time.sleep(5)


if __name__ == "__main__":
    signal.signal(signal.SIGINT, interrupt_signal_handler)
    signal.signal(signal.SIGTSTP, interrupt_signal_handler)
    signal.signal(signal.SIGTERM, interrupt_signal_handler)
    run()