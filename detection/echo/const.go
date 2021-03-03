package echo

const (
	Header    = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36 Edg/88.0.705.81"
	ProxyHost = "localhost:1080"
)

// REGX 应用的版本更新检测正则表达式集
var REGX = map[string]string{
	"2345pic":            `(?<=class="version">v)[0-9.]+(?=</div>)`,
	"acmekit":            `(?<=软件版本：)[0-9.]+(?=</FONT)`,
	"beyondcompare":      `(?<=BCompare-)[0-9.]+(?=.exe)`,
	"ccleaner":           `(?<=CCleaner v)[0-9.]+`,
	"centbrowser":        `(?<=Version: )[0-9.]+`,
	"clash":              `(?<=v )[0-9.]+`,
	"claunch":            `(?<=CLaunch 64-bit Unicode Ver. )[0-9.]+`,
	"contextmenumanager": `(?<=tag/)[0-9.]+`,
	"cosbrowser":         `(?<=</a>v)[0-9.]+`,
	"dismpp":             `(?<=What's new ‒ version )[0-9.]+`,
	"ditto":              `(?<=versionDots=")[0-9.]+`,
	"es":                 `(?<=href="/ES-)[0-9.]+(?=\.zip")`,
	"emeditor":           `(?<=Download latest version: v)[0-9.]+`,
	"fastcopy":           `(?<=FastCopy ver )[0-9.]+`,
	"foobar2000":         `(?<=foobar2000 v)[0-9.]+`,
	"fscapture":          `(?<=Version\s+)[0-9.]+`,
	"fsresizer":          `(?<=Version\s+)[0-9.]+`,
	"hbuilderx":          `(?<="version":\s+")[0-9.]+`,
	"idm":                `(?<=IDM_Retail_v)[0-9.]+`,
	"iobitdriverbooster": `(?<=<span>V\s+)[0-9.]+`,
	"iobituninstaller":   `(?<=<span>V\s+)[0-9.]+`,
	"kodobrowser":        `(?<=kodo-browser-Windows-x64-v)[0-9.]+(?=.zip)`,
	"kodoimport":         `(?<=kodoimport_windowsOS_64_v)[0-9.]+(?=.tar.gz)`,
	"lightproxy":         `(?<=tag/v)[0-9.]+`,
	"newfiletime":        `(?<=NewFileTime\s+)[0-9.]+`,
	"n_m3u8dl-cli":       `(?<=N_m3u8DL-CLI_v)[0-9r.]+`,
	"notepad2":           `(?<=tag/v)[0-9r.]+`,
	"notepad3":           `(?<=Notepad3 Release )[0-9.]+`,
	"oss-browser":        `(?<=oss-browser/)[0-9.]+(?=/oss-browser-win32-x64.zip)`,
	"ossutil":            `(?<=ossutil/)[0-9.]+(?=/ossutil64.zip)`,
	"picgo":              `(?<=tree/v)[0-9.]+(?=")`,
	"potplayer":          `(?<=class="tit_version" title="version">)1\.[0-9.]+`,
	"qq":                 `(?<="new-version">)[0-9.]+`,
	"qshell":             `(?<=qshell-v)[0-9.]+(?=-windows-amd64)`,
	"renamer":            `(?<=/download/renamer/installer/)[0-9.]+`,
	"shadowsocksrr":      `(?<=tag/)[0-9.]+`,
	"totalcommander":     `(?<=Version\s+)[0-9.]+(?=,)`,
	"ultraiso":           `(?<=UltraISO\s+)[0-9.]+(?=\s+released)`,
	"uninstalltool":      `(?<=release-notes">)[0-9.]+`,
	"v2rayn":             `(?<=releases/tag/)[0-9.]+`,
	"wisecare365":        `(?<=当前版本: )[0-9.]+`,
	"wechat":             `(?<=class="version-tag">)[0-9.]+`,
	"wesing":             `(?<=version_pc:')[0-9.]+`,
	"xyplorer":           `(?<=class="smaller">)[0-9.]+`,
}

// ReplaceURL 替换应用Homepage URL
var ReplaceURL = map[string]string{
	"acmekit":       "http://blog.sina.com.cn/s/blog_89a729a40102wjwh.html",
	"ccleaner":      "https://www.ccleaner.com/ccleaner/download",
	"cosbrowser":    "https://github.com/tencentyun/cosbrowser/blob/master/changelog.md",
	"hbuilderx":     "https://download1.dcloud.net.cn/hbuilderx/release.json",
	"idm":           "https://www.lanzoux.com/b0i6zkze",
	"kodobrowser":   "https://developer.qiniu.com/kodo/tools/5972/kodo-browser",
	"oss-browser":   "https://help.aliyun.com/document_detail/61872.html",
	"ossutil":       "https://help.aliyun.com/document_detail/120075.html",
	"picgo":         "https://github.com/Molunerfinn/PicGo/releases",
	"renamer":       "https://www.den4b.com/downloads/renamer",
	"uninstalltool": "https://crystalidea.com/uninstall-tool/download",
	"wechat":        "https://pc.weixin.qq.com/",
	"wesing":        "https://kg.qq.com/gtimg/mediastyle/musiccm/kg/kg_feature.js",
}

// ShieldApp 需要屏蔽无法或不需要检测更新的应用
var ShieldApp = map[string]bool{
	"2345pic":    true,
	"foobar2000": true,
	"qq":         true,
	"idm":        true,
	"potplayer":  true,
	//"cosbrowser":   true,
	"eudic":        true,
	"ocam":         true,
	"xiuxiu":       true,
	"ludashi":      true,
	"neteasemusic": true, //301
	"notepad2-mod": true,
	"exescope":     true,
	"kbootopt":     true,
	"kcleaner":     true,
	"ksoftmgr":     true,
	"wetool":       true,
	"xtools":       true,
}
