# exploring-go
Exploring Go.

## 1. Installation
```bash
wget -O /tmp/go1.20.5.linux-amd64.tar.gz https://go.dev/dl/go1.20.5.linux-amd64.tar.gz
rm -rf /usr/local/go 
tar -C /usr/local -xzf /tmp/go1.20.5.linux-amd64.tar.gz

```

## Note(s)
```bash
mkdir hello
cd hello
go mod init example/hello
touch hello.go
go run .
go mod tidy
go run .

cd trash
go mod edit -replace=example.com/greetings=../greetings
    
```
```bash




```