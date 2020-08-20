all: fizz

fizz: main.go encrypt.go decrypt.go tools.go
	go build -o fizz main.go encrypt.go decrypt.go tools.go

clean:
	rm fizz

install:
	mkdir -p /usr/local/bin
	cp fizz /usr/local/bin/fizz

uninstall:
	rm -f /usr/local/bin/fizz
