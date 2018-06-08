## DTransfer
##### Application for transfer os directories through network connections
---
#### Powered by 
##### 1. [Webview](https://github.com/zserge/webview)
##### 2. [Go-bindata](https://github.com/jteeuwen/go-bindata)
##### 3. [Gorilla Websocket](https://github.com/gorilla/websocket)

![example](https://github.com/AlexeyArno/DTransfer/blob/master/present/1st.png)

---
#### Setup

```
$ go get github.com/AlexeyArno/DTransfer
$ cd GOPATH/src/github.com/AlexeyArno/DTransfer
$ go get -d ./...
```
#### Build
```
# Linux
$ go build -o DTransfer && ./DTransfer

# MacOS uses app bundles for GUI apps
$ mkdir -p DTransfer.app/Contents/MacOS
$ go build -o DTransfer.app/Contents/MacOS/DTransfer
$ open DTransfer.app # Or click on the app in Finder

# Windows requires special linker flags for GUI apps.
# It's also recommended to use TDM-GCC-64 compiler for CGo.
# http://tdm-gcc.tdragon.net/download
$ go build -ldflags="-H windowsgui" -o DTransfer.exe
```


