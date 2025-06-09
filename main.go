package main

import (
	_ "github.com/ljcheng999/ljc-scan/cmd/jfrog-xray"
	ljcscan "github.com/ljcheng999/ljc-scan/cmd/ljc-scan"
	// _ "github.com/ljcheng999/ljc-scan/pkg/cmd/jfrog-xray"
)

func main() {
	ljcscan.Execute()
}
