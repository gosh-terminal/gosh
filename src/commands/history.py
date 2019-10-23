#!/usr/bin/python3
f = open("/workspace/gosh/src/commands/history.txt",'r')
count = 1
contents = f.readlines()
for line in contents:
    if line == '\n' or line == '\r':
        continue
    print(str(count) + " " + line,end='')
    count+=1
print()