#!/usr/bin/python3
import os
files = [f for f in os.listdir('.') if os.path.isfile(f)]
for i in files:
    char = ' '
    if len(i) < 9:
        z = 9 - len(i) - 2
    else:
        z = 1
    spaces = z*char
    item = "| " + i + spaces + "|"
    print(" --------")
    print(item)
print(" --------")
