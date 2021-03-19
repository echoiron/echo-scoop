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
	"sort"
	"strconv"
	"strings"
)

// App echo-scoop/bucket/xx.json
type App struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	NewVersion  string `json:"newVersion"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	//License     string `json:"license"`
	//URL         string `json:"url"`
	//Bin string `json:"bin"`
}

func ReadApp(name string) App {
	fileName := NameNoExt(name)
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

func CheckApp(app *App) {
	homePage := Replace(app.Name)
	if homePage == "" {
		homePage = app.Homepage
	}
	content := requestLink(homePage)
	// fmt.Println(content)
	re := regexp2.MustCompile(RegxValue(app.Name), 0)
	m, err := re.FindStringMatch(content)
	if err != nil || m == nil {
		log.Println(app.Name, "match data:", m, "| err:", err)
		app.NewVersion = app.Version
		return
	}
	app.NewVersion = m.String()
}

// NewVersionDetected 检测到App新版本
func NewVersionDetected(app *App) {
	if compareVersion(removerChar(app.Version), removerChar(app.NewVersion)) == -1 {
		text := fmt.Sprintf("%s %s %s %s", app.Name, app.Version, app.NewVersion, app.Homepage)
		fmt.Println(colorPrint(text, 0, TextRed, 0))
		return
	}
	fmt.Printf("%s %s %s %s\n", app.Name, app.Version, app.NewVersion, app.Homepage)
}

// BucketFiles bucket文件
func BucketFiles() ([]string, error) {
	wd := currDir()
	pd := parentDir(wd)
	pattern := bucketPattern(pd)

	bucketFiles, err := filepath.Glob(pattern)
	if err != nil {
		return []string{}, err
	}
	sort.Strings(bucketFiles)
	return bucketFiles, nil
}

// NameNoExt 不包含扩展名称的文件名
func NameNoExt(filePath string) string {
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

// removerChar 移除版本号中的特殊符号(eg:3.12rc3325)
func removerChar(str string) string {
	re := regexp2.MustCompile("[a-zA-Z]", regexp2.None)
	newStr, err := re.Replace(str, "0", -1, -1)
	if err != nil {
		return str
	}
	return newStr
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

func requestLink(url string) string {
	proxyState := viper.GetBool(fmt.Sprintf("network.%s", "enable_proxy")) //是否开启代理
	c := &fasthttp.Client{}
	c.Name = viper.GetString(fmt.Sprintf("network.%s", "user_agent"))
	if proxyState {
		c.Dial = fasthttpproxy.FasthttpHTTPDialer(viper.GetString(fmt.Sprintf("network.%s", "proxy_host")))
	}
	code, body, err := c.Get(nil, url)
	if err != nil {
		log.Printf("Request error: %s", err)
	}
	if code != fasthttp.StatusOK {
		log.Printf("Unexpected status code: %d. Expecting %d", code, fasthttp.StatusOK)
	}
	return string(body)
}
