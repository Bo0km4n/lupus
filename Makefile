MAKE = make
linux:
	GOOS=linux GOARCH=amd64 go build .
	$(MAKE) -C example build
pb: 
	$(MAKE) -C protobuf pb