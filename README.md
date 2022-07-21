
# echo-scoop

### Quick Start

<https://github.com/lukesampson/scoop/wiki/Quick-Start>

### Installing Scoop

Make sure you have allowed PowerShell to execute local scripts

```powershell
set-executionpolicy remotesigned -scope currentuser
```

Installing Scoop to Custom Directory

```powershell
$env:SCOOP='D:\scoop'
[environment]::setEnvironmentVariable('SCOOP',$env:SCOOP,'User')
Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://get.scoop.sh')
```

### How to install bucket

```bash
scoop bucket add echo https://github.com/echoiron/echo-scoop
```

### App of installing bucket

```bash
scoop install echo/app_name

# e.g.
scoop install echo/fscapture
```

### Delete all buckets
```bash
scoop bucket rm *
```

### Commonly used commands
```bash
scoop update
scoop update *
scoop install extras/sumatrapdf
scoop install python@3.7.14
scoop install wechat -f
scoop uninstall discord -p
scoop cache rm notepad2
scoop cleanup *
```

###  Commonly used buckets
- powershell
```powershell
scoop bucket add main ;`
scoop bucket add extras ;`
scoop bucket add jetbrains https://github.com/Ash258/Scoop-JetBrains ;`
scoop bucket add echo https://github.com/echoiron/echo-scoop ;`
scoop bucket add dodorz https://github.com/dodorz/scoop ;`
scoop bucket add dorado https://github.com/chawyehsu/dorado ;`
scoop bucket add 42wim https://github.com/42wim/scoop-bucket ;`
scoop bucket add mochi https://github.com/Qv2ray/mochi
```

- dos
```bash
scoop bucket add main & ^
scoop bucket add extras & ^
scoop bucket add jetbrains https://github.com/Ash258/Scoop-JetBrains & ^
scoop bucket add echo https://github.com/echoiron/echo-scoop & ^
scoop bucket add dodorz https://github.com/dodorz/scoop & ^
scoop bucket add dorado https://github.com/chawyehsu/dorado & ^
scoop bucket add 42wim https://github.com/42wim/scoop-bucket & ^
scoop bucket add mochi https://github.com/Qv2ray/mochi
```

###  Other app buckets
<https://rasa.github.io/scoop-directory/by-score>

