API to fetch new word and the meaning from GRE Word pool



**Requirements:**
Go,
GIN,
Mysql

**Installation**

1.MySql

`$ go get -u github.com/go-sql-driver/mysql`

Import it in your code:

`"github.com/go-sql-driver/mysql"`

2.To install Gin package, you need to install Go and set your Go workspace first.

After installing Go (version 1.13+ is required), then you can use the below Go command to install Gin.

`$ go get -u github.com/gin-gonic/gin`

Import it in your code:

`import "github.com/gin-gonic/gin"`

(Optional) Import net/http. This is required for example if using constants such as http.StatusOK.

`import "net/http"`

**Check dbconfig file for database info**

**External API used for searching words and adding them to the database**

https://github.com/meetDeveloper/freeDictionaryAPI
```
$ go run main.go
```
