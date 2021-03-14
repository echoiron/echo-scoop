package echo

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"sort"
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
	return viper.GetString(fmt.Sprintf("regx.%s", name))
}

// Replace 替换应用链接
func Replace(name string) string {
	return viper.GetString(fmt.Sprintf("replaceurl.%s", name))
}

// ShieldDet 屏蔽应用检测
func ShieldDet(name string) bool {
	return viper.GetBool(fmt.Sprintf("shieldapp.%s", name))
}

// currDir 当前目录
func currDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// parentDir 父级目录
func parentDir(path string) string {
	parentPath := filepath.Dir(path)
	return strings.Replace(parentPath, "\\", "/", -1)
}

func bucketPattern(dir string) string {
	sub := "bucket"
	ext := "*.json"
	separator := "\\"
	if !strings.Contains(dir, "\\") {
		separator = "/"
	}
	return fmt.Sprintf("%s%s%s%s%s", dir, separator, sub, separator, ext)
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
