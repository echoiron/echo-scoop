package main

import (
	"det/echo"
	"log"
	"sync"
)

// conf/app.toml 配置

func main() {

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
		app := echo.ReadApp(name)
		go func() {
			defer wg.Done()
			echo.CheckApp(&app)
			echo.NewVersionDetected(&app)
		}()
	}
	wg.Wait()
}
