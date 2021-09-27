build:
	mkdir -p build
	go build -o build/gomeetme

test:
	go test -v ./factory
	go test -v ./timeconverters

