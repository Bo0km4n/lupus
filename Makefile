MAKE = make
GO_S_B = go tool compile
GO_S_L = go tool link
SRC_PATH=/home/vagrant/go/src
PKG_PATH=/home/vagrant/go/pkg/linux_amd64
linux:
	GOOS=linux GOARCH=amd64 go build -gcflags='-N -l' .
	$(MAKE) -C example build
pb: 
	$(MAKE) -C protobuf pb

clean:
	rm *.a
	rm *.o

lupus_build: internal_build
	$(GO_S_B) -o lupus.a -I $SRC_PATH -I $PKG_PATH -pack ./*.go
