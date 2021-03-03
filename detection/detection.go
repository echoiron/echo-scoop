package main

import (
	"detection/echo"
	"log"
	"path/filepath"
	"sort"
)

// conf/app.toml 配置是否开启代理

func main() {

	files, err := filepath.Glob(echo.Glob(echo.GetParentPath(echo.GetCurrentPath())))
	if err != nil {
		log.Panic(err)
	}
	sort.Strings(files)
	for _, name := range files {
		fileName := echo.FileName(name)
		if echo.ShieldDet(fileName) {
			continue
		}
		app := echo.ReadApp(name)
		echo.CheckApp(&app)
		echo.NewVersionNotice(&app)
	}
}
