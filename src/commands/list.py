#!/usr/bin/python3
import os
files = [f for f in os.listdir('.') if os.path.isfile(f)]
print("----------------------------------")
for i in files:
    char = ' '
    z = 34 - len(i) - 3
    spaces = z*char
    item = "| " + i + spaces + "|"
    print(item)
print("----------------------------------")
