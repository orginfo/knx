# knx

## Установка для Ubuntu 16.04 
1. Устанавливаем golang-1.10 (ниже может не быть модуля context)
`a@am:~$ sudo apt-get install golang-1.10-go`
2. Создаем следующую структуру workspace-а gocode:
```
/home/a/gocode/
├── bin
│   └── knx
├── data
│   └── db.sqlite3
├── html
│   └── test.html
└── src
    ├── github.com
    │   └── mattn
    │       └── go-sqlite3
    └── knx
```
```
a@md:~$ mkdir -p /home/a/gocode/src
a@md:~$ mkdir -p /home/a/gocode/bin
a@md:~$ mkdir -p /home/a/gocode/html
a@md:~$ echo "test" > /home/a/gocode/html/test.html
a@md:~$ echo "projects" > /home/a/gocode/html/projects.html
```
3. Прописываем пути до go, gopath, goroot:
``` diff
--- .profile.bak	2018-04-06 09:53:00.312654571 +0700
+++ .profile	2018-04-06 09:16:21.221152165 +0700
@@ -20,3 +20,7 @@
 if [ -d "$HOME/bin" ] ; then
     PATH="$HOME/bin:$PATH"
 fi
+
+export GOPATH=$HOME/gocode
+export GOROOT=/usr/lib/go-1.10
+export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```
4. Применяем пути
`a@am:~$ source ~/.profile`
5. Клонируем проект:
`a@am:~/gocode/src$ git clone https://github.com/orginfo/knx.git`
6. Скачиваем зависимый пакет:
`a@am:~/gocode/src$ go get -v github.com/mattn/go-sqlite3`
7. Устанавливаем knx
`a@am:~/gocode/src$ go install knx`
8. Запускаем сервис:
`a@am:~/gocode/src$ knx`
