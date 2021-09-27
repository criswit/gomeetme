build:
	mkdir -p build
	go build -o build/gomeetme

windows:
	mkdir -p windows
	GOOS=windows GOARCHamd64 go build -o windows/gomeetme

test:
	go test -v ./factory
	go test -v ./timeconverters

