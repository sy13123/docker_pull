export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go mod download
go mod tidy
go build -x -o ./build/gopull
