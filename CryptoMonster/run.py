import os
import sys
import subprocess


def run_daemon():
    pwd = os.getcwd()
    try:
        pid = os.fork()
        if pid > 0:
            sys.exit(0)
    except:
        sys.exit(0)

    os.setsid()
    os.umask(0)

    try:
        pid = os.fork()
        if pid > 0:
            sys.exit(0)
    except:
        sys.exit(0)

    subprocess.Popen(
    [f"{pwd}/CryptoMonster"],
    stdout=subprocess.DEVNULL,
    stderr=subprocess.DEVNULL,)
    sys.exit(0)



if __name__ == '__main__':
    run_daemon()
