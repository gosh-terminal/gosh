'''Puts all items in ls in a box'''
#!/usr/bin/python3
import os
files = [f for f in os.listdir('.') if os.path.isfile(f) or os.path.isdir]
print("----------------------------------")
for i in files:
    char = ' '
    z = 34 - len(i) - 3
    spaces = z*char
    item = "| " + i + spaces + "|"
    print(item)
print("----------------------------------")
