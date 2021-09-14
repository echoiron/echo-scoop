
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

###  Other app buckets
<https://rasa.github.io/scoop-directory/by-score>

