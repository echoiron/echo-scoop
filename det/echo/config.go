package echo

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 终端背景颜色
const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

func init() {
	viper.SetConfigFile("conf/app.toml") //指定配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("read config failed: %v", err)
	}
}

// colorPrint 配置终端打印颜色
// def: 终端默认颜色
// fg: 前景色
// bg: 背景色
func colorPrint(text string, def, fg, bg int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, def, bg, fg, text, 0x1B)
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

// bucketPattern bucket文件通配表达式
func bucketPattern(dir string) string {
	sub := "bucket"
	ext := "*.json"
	separator := "\\"
	if !strings.Contains(dir, "\\") {
		separator = "/"
	}
	return fmt.Sprintf("%s%s%s%s%s", dir, separator, sub, separator, ext)
}
