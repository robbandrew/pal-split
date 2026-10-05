build:
	go build -o bin/splitpal cmd/main.go
	
clean:
	rm -f bin/splitpal
	rm -rf dist/

run:
	go run cmd/main.go

serve: clean build
	./bin/splitpal