build:
	GOOS=linux GOARCH=amd64 go build -gcflags='-N -l' .
as:
	objdump -d example >> example.s
clean:
	rm example
	rm example.s

