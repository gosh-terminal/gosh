import os
import platform


def install():
    if "Linux" in platform.system():
        os.system("go build -o gosh *.go")
    elif "macos" in platform.system():
        os.system("GOOS=darwin GOARCH=amd64 go build *.go")
    elif "windows" in platform.system():
        os.system("GOOS=windows GOARCH=amd64 go build -o gosh.exe *.go")
    else:
        print("Unknown OS")


install()
