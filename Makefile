MAKE = make
linux:
	GOOS=linux GOARCH=amd64 go build -gcflags='-N -l' .
	$(MAKE) -C example build
pb: 
	$(MAKE) -C protobuf pb