API to fetch new word and the meaning from GRE Words pool

How to use:

**Requirements:**
GIN,
Mysql

**Installation**

1.MySql

`$ go get -u github.com/go-sql-driver/mysql`

2.To install Gin package, you need to install Go and set your Go workspace first.

The first need Go installed (version 1.13+ is required), then you can use the below Go command to install Gin.

`$ go get -u github.com/gin-gonic/gin`

Import it in your code:

`import "github.com/gin-gonic/gin"`

(Optional) Import net/http. This is required for example if using constants such as http.StatusOK.

`import "net/http"`


```
$ go run main.go
```
