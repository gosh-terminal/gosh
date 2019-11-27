import os
import platform


def install():
    if "Linux" in platform.system():
        print("Installing for Linux")
        os.system("go install")
    elif "Darwin" in platform.system():
        print("Installing for OSX")
        os.system("go install")
    elif "Windows" in platform.system():
        print("Installing for Windows")
        os.system("go.exe install")
    else:
        print("Unknown OS")


install()
