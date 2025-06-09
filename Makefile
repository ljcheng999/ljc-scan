BIN="bin/*"



help:
	go run main.go jfrog-xray -h

run:
	go run main.go jfrog-xray -d -u username -p test123 --url test.com
# go run cmd/ljc-scan/ljc-scan.go

b:
# chmod +x bin/ljc-scan
	./bin/ljc-scan -h

build:
	go build -o bin/ljc-scan cmd/ljc-scan/ljc-scan.go
# go build -o bin/ljc-scan cmd/ljc-scan/ljc-scan.go && chmod +x bin/ljc-scan


clean:
	rm -rf $(BIN)