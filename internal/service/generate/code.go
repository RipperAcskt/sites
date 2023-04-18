package generate

var codeStr = `package generate

import (
	"danila/sites"
	"sync"
	"os"
)

func Check(email string) {
	var wg sync.WaitGroup
	file, _ := os.Create("result")

`
