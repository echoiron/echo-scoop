// 在 main.go 中
package main

import (
	"det/echo"
	"log"
	"sync"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	files, err := echo.BucketFiles()
	if err != nil {
		log.Panic(err)
		return
	}

	var wg sync.WaitGroup
	for _, name := range files {
		fileName := echo.NameNoExt(name)
		if echo.ShieldDet(fileName) {
			continue
		}
		wg.Add(1)
		go echo.ProcessApp(&wg, name)
	}
	wg.Wait()
}
