all: 
	go build -o main main.go func.go test*.go

clean:
	rm -f main

