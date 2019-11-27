import os
import platform


def install():
    if "Linux" in platform.system():
        os.system("go install")
        os.system("gosh")
    elif "Darwin" in platform.system():
        os.system("go install")
        os.system("gosh")
    elif "Windows" in platform.system():
        os.system("go.exe install")
    else:
        print("Unknown OS")


install()
