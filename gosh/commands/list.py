#!/usr/bin/python3
import os

files1 = [(f, True) for f in os.listdir('.')
          if os.path.isfile(f) and not os.access(f, os.X_OK)]
files2 = [(f, False) for f in os.listdir('.') if os.path.isdir(f)]
files3 = [(f, None) for f in os.listdir('.')
          if os.path.isfile(f) and os.access(f, os.X_OK)]
files = files1 + files2 + files3
print("----------------------------------")
for i in files:
    os.system("/workspace/gosh/src/commands/bin/resetColor")
    char = ' '
    z = 34 - len(i[0]) - 3
    spaces = z*char
    item = "| " + i[0] + spaces + "|"
    if i[1] == True:
        print(item[:1], end="")
        print(item[1:-2], end="")
    elif i[1] == None:
        print(item[:1], end="")
        os.system("/workspace/gosh/src/commands/bin/makeYellow")
        print(item[1:-2], end="")
    else:
        print(item[:1], end="")
        os.system("/workspace/gosh/src/commands/bin/makeCyan")
        print(item[1:-2], end="")
    print(item[-2:])
    os.system("/workspace/gosh/src/commands/bin/resetColor")
print("----------------------------------")
