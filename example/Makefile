build:
	GOOS=linux GOARCH=amd64 go build -gcflags='-N -l' .
as:
	objdump -d example >> example.s
clean:
	-rm example
	-rm example.s
	-rm example_hoge example.a example_hoge.a example.o hoge.o

inc_build:
	go tool compile -o example.a -p main -I $$SRC_PATH -I $$PKG_PATH -pack example.go
	go tool link -v -n -o example -L $$SRC_PATH -L $$PKG_PATH example.a
	go tool compile -o example_hoge.a -p main -I $$SRC_PATH -I $$PKG_PATH -pack hoge.go example.go
	go tool link -v -n -o example_hoge -L $$SRC_PATH -L $$PKG_PATH example_hoge.a