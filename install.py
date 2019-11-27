import os
import platform


def install():
    if "Linux" in platform.system():
        print("Installing for Linux")
        os.system("go install")
        print("Test gosh")
        os.system("$GOPATH/bin/gosh")
    elif "Darwin" in platform.system():
        print("Installing for OSX")
        os.system("go install")
        print("Test gosh")
        os.system("$GOPATH/bin/gosh")
    elif "Windows" in platform.system():
        print("Installing for Windows")
        os.system("go.exe install")
        print("Test gosh")
        os.system("%GOPATH/bin/gosh.exe")
    else:
        print("Unknown OS")


install()
