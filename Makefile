main:
	@go build -o ./bin/cm ./src/*.go
	@./bin/cm -file ./src.txt -debug

release:
	# android arm64
	GOOS=android GOARCH=arm64 go build -o ./bin/cm-android-arm64 ./src/*.go

	# linux amd64
	GOOS=linux GOARCH=amd64 go build -o ./bin/cm-linux-amd64 ./src/*.go
	# linux arm
	GOOS=linux GOARCH=arm go build -o ./bin/cm-linux-arm ./src/*.go

    # windows x64
	GOOS=windows GOARCH=amd64 go build -o ./bin/cm-windows-x64.exe ./src/*.go
	# windows x86
	GOOS=windows GOARCH=arm go build -o ./bin/cm-windows-x86.exe ./src/*.go
	
clean:
	rm ./bin/*