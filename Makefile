main:
	@go build -o ./bin/cm ./src/*.go
	@./bin/cm -file ./src.txt -debug
clean:
	rm ./bin/*