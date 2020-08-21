all: fizz

fizz: main.go key.go xor.go
	go build -o fizz main.go key.go xor.go

clean:
	rm fizz

install:
	mkdir -p /usr/local/bin
	cp fizz /usr/local/bin/fizz

uninstall:
	rm -f /usr/local/bin/fizz
