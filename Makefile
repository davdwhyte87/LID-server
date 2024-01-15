build:
	set GOOS=windows set GOARCH=amd64 
	go build -o bin/dau.exe app.go

run:
	go run main.go