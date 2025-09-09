#!/bin/sh

echo "=== Debugging Go Module Issues ==="
echo

echo "1. Current working directory:"
pwd
echo

echo "2. Files in current directory:"
ls -la
echo

echo "3. Content of go.mod:"
if [ -f go.mod ]; then
    cat go.mod
    echo
    echo "4. Hex dump of go.mod (first 100 bytes):"
    hexdump -C go.mod | head -10
else
    echo "go.mod not found!"
fi
echo

echo "5. Checking if go.sum exists:"
if [ -f go.sum ]; then
    echo "go.sum exists"
    wc -l go.sum
else
    echo "go.sum does not exist"
fi
echo

echo "6. Trying to create a fresh go.mod:"
rm -f go.mod go.sum
go mod init golang-app
echo

echo "7. Adding Gin dependency:"
go get github.com/gin-gonic/gin@v1.9.1
echo

echo "8. Final go.mod content:"
cat go.mod
echo

echo "9. Final go.sum exists:"
ls -la go.sum 2>/dev/null || echo "go.sum still not found"
echo

echo "10. Testing go mod download:"
go mod download
echo

echo "11. Testing go build:"
echo 'package main
import "fmt"
func main() { fmt.Println("Hello") }' > test.go

go build test.go
if [ $? -eq 0 ]; then
    echo "Build successful!"
    ./test
else
    echo "Build failed!"
fi

echo
echo "=== Debug Complete ==="
