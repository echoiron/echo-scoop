
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
scoop install echo/notepad4
scoop install 7zip git --global
scoop install aria2 curl grep sed less touch
scoop install python@3.12.4
scoop install wechat -f
scoop uninstall discord -p
scoop cache rm *
scoop cleanup *
```

###  Commonly used buckets
- powershell
```powershell
scoop bucket add main ;`
scoop bucket add extras ;`
scoop bucket add echo https://github.com/echoiron/echo-scoop ;`
scoop bucket add dodorz https://github.com/dodorz/scoop ;`
scoop bucket add dorado https://github.com/chawyehsu/dorado
```

- dos
```bash
scoop bucket add main & ^
scoop bucket add extras & ^
scoop bucket add echo https://github.com/echoiron/echo-scoop & ^
scoop bucket add dodorz https://github.com/dodorz/scoop & ^
scoop bucket add dorado https://github.com/chawyehsu/dorado
```

###  Other app buckets
<https://scoop.sh/#/buckets>

<https://rasa.github.io/scoop-directory/by-score>
