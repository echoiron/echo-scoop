
# echo-scoop

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

### Multi-line command for powershell
```bash
scoop bucket add extras ;`
scoop bucket add jetbrains ;`
scoop bucket add echo https://github.com/echoiron/echo-scoop ;`
scoop bucket add dodorz https://github.com/dodorz/scoop-bucket ;`
scoop bucket add dorado https://github.com/chawyehsu/dorado
```

---
### Multi-line command for dos/cmd
```bash
scoop bucket add extras & ^
scoop bucket add jetbrains & ^
scoop bucket add echo https://github.com/echoiron/echo-scoop & ^
scoop bucket add dodorz https://github.com/dodorz/scoop-bucket & ^
scoop bucket add dorado https://github.com/chawyehsu/dorado
```

### Delete bucket
```bash
scoop bucket rm extras
scoop bucket rm jetbrains
scoop bucket rm echo
scoop bucket rm dodorz
scoop bucket rm dorado
```

### Delete all buckets
```bash
scoop bucket rm *
```
