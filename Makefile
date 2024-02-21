build:
	GOOS=darwin   GOARCH=amd64 go build -o bin/macos/orca       cmd/orca_cli/main.go
	#GOOS=linux    GOARCH=amd64 go build -o bin/linux/orca         cmd/orca_cli/main.go
clean:
	rm -fr bin/*
