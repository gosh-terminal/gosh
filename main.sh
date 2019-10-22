nim c src/commands/help.nim && mv src/commands/help src/commands/bin
cd src/commands/pymath/ && python3 setup.py install
cd ..
cd ..
cd ..
clear
go run src/gosh.go