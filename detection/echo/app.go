package echo

import (
	"encoding/json"
	"fmt"
	"github.com/dlclark/regexp2"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

type App struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	NewVersion  string `json:"newVersion"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	License     string `json:"license"`
	//URL         string `json:"url"`
	Bin string `json:"bin"`
}

func FileName(filePath string) string {
	base := filepath.Base(filePath)
	s := strings.Split(base, ".")
	var fileName string
	if len(s) == 0 {
		fileName = ""
	} else {
		fileName = s[0]
	}
	return fileName
}

func ReadApp(name string) App {
	fileName := FileName(name)
	file, err := ioutil.ReadFile(name)
	if err != nil {
		log.Println(err)
	}

	app := App{}
	err = json.Unmarshal(file, &app)
	if err != nil {
		log.Println(err)
	}
	app.Name = fileName
	return app
}

func request(url string) string {

	proxyState := viper.GetBool("enable_proxy") //是否开启代理
	c := &fasthttp.Client{}
	c.Name = Header
	if proxyState {
		c.Dial = fasthttpproxy.FasthttpHTTPDialer(ProxyHost)
	}
	code, body, err := c.Get(nil, url)
	if err != nil {
		log.Printf("Error when loading google page through local proxy: %s", err)
	}
	if code != fasthttp.StatusOK {
		log.Printf("Unexpected status code: %d. Expecting %d", code, fasthttp.StatusOK)
	}

	return string(body)
}

func CheckApp(app *App) {
	homePage := Replace(app.Name)
	if homePage == "" {
		homePage = app.Homepage
	}
	content := request(homePage)
	// fmt.Println(content)
	re := regexp2.MustCompile(RegxValue(app.Name), 0)
	m, err := re.FindStringMatch(content)
	if err != nil || m == nil {
		fmt.Println("match data:", m, "| err:", err)
		app.NewVersion = app.Version
		return
	}
	app.NewVersion = m.String()
}

// compareVersion 版本号比较
// 1:version1 > version2
//-1:version1 < version2
// 0:version1 = version2
func compareVersion(version1 string, version2 string) int {
	versionA := strings.Split(version1, ".")
	versionB := strings.Split(version2, ".")

	for i := len(versionA); i < 4; i++ {
		versionA = append(versionA, "0")
	}
	for i := len(versionB); i < 4; i++ {
		versionB = append(versionB, "0")
	}
	for i := 0; i < 4; i++ {
		version1, _ := strconv.Atoi(versionA[i])
		version2, _ := strconv.Atoi(versionB[i])
		if version1 == version2 {
			continue
		} else if version1 > version2 {
			return 1
		} else {
			return -1
		}
	}
	return 0
}

// NewVersionNotice 新版本高亮提示
func NewVersionNotice(app *App) {
	fmt.Println("app-->", app)
	if compareVersion(removerChar(app.Version), removerChar(app.NewVersion)) == -1 {
		fmt.Printf("\033[1;31m%s %s %s\033[0m\n", app.Name, app.Version, app.NewVersion)
		return
	}

	/*if app.Version < app.NewVersion {
		fmt.Printf("\033[1;31m%s %s %s\033[0m\n", app.Name, app.Version, app.NewVersion)
		return
	}
	if versionCompare2(app.Version, app.NewVersion) { //解析成浮点数有很大问题
		fmt.Printf("\033[1;31m%s %s %s\033[0m\n", app.Name, app.Version, app.NewVersion)
		return
	}*/
}

// removerChar 移除版本号中的特殊符号(eg:3.12rc3325)
func removerChar(str string) string {
	re := regexp2.MustCompile("[a-zA-Z]", regexp2.None)
	newStr, err := re.Replace(str, "0", -1, -1)
	if err != nil {
		return str
	}
	return newStr
}
