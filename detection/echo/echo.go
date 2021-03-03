package echo

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func init() {
	viper.SetConfigFile("conf/app.toml") //指定配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("read config failed: %v", err)
	}
}

// RegxValue 正则值
func RegxValue(name string) string {
	return REGX[name]
}

// Replace 替换应用链接
func Replace(name string) string {
	return ReplaceURL[name]
}

// ShieldDet 屏蔽应用检测
func ShieldDet(name string) bool {
	return ShieldApp[name]
}

func ProxyState(stat bool) bool {
	return stat
}

func Glob(dir string) string {
	sub := "bucket"
	ext := "*.json"
	separator := "\\"
	if !strings.Contains(dir, "\\") {
		separator = "/"
	}
	pattern := fmt.Sprintf("%s%s%s%s%s", dir, separator, sub, separator, ext)
	return pattern
}

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func GetParentPath(path string) string {
	parentPath := filepath.Dir(path)
	return strings.Replace(parentPath, "\\", "/", -1)
}
