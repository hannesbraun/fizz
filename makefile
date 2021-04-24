all: fizz

fizz: main.go key.go xor.go
	go build -v -o fizz .

clean:
	rm fizz

install:
	mkdir -p /usr/local/bin
	cp fizz /usr/local/bin/fizz

uninstall:
	rm -f /usr/local/bin/fizz
