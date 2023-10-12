default: build

build:
	mkdir -p dist
	go build -v -o dist/active-reception

run: build
	./dist/active-reception

install:
	go install

lint:
	golangci-lint run ./...
	cd terraform/ && golangci-lint run ./...

clean:
	rm -rf dist/
