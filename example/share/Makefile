SRC_PATH=/home/vagrant/go/src
PKG_PATH=/home/vagrant/go/pkg/linux_amd64
unpatched.bin:
	go tool compile -o unpatched.a -N -l -p main -I $(SRC_PATH) -I $(PKG_PATH) -pack main.go
	go tool link -v -n -o unpatched.bin -L $(SRC_PATH) -L $(PKG_PATH) unpatched.a
patch.bin:
	go tool compile -o patched.a -N -l -p main -I $(SRC_PATH) -I $(PKG_PATH) -pack main.go patch.go
	go tool link -v -n -o patched.bin -L $(SRC_PATH) -L $(PKG_PATH) patched.a 

dump_as:
	objdump -d patched.bin >> patch.s
	objdump -d unpatched.bin >> unpatch.s
clean:
	-rm *.bin
