package controllers

import (
	"fmt"
	"net/http"

	"GREWordsv2/models"

	"github.com/gin-gonic/gin"
)

type showWords struct {
	SWord    string `json:"word"`
	SMeaning string `json:"meaning"`
}

func Testfunc(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})

}
func GetWord(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	var wr []models.Word
	var sw []showWords
	result, err := models.DB.Query("SELECT * FROM Words WHERE learnt = 0 ORDER BY RAND() LIMIT 1")
	if err != nil {
		panic(err)

	}

	defer result.Close()
	for result.Next() {
		var w1 models.Word
		var sw1 showWords
		err := result.Scan(&w1.Word, &w1.Meaning, &w1.Learnt, &w1.LearntAt)
		if err != nil {
			panic(err)
		}
		wr = append(wr, w1)
		_, err2 := models.DB.Query("update Words set learntAt = Now() where word = ?", w1.Word)
		if err2 != nil {
			panic(err2)
		}

		//fmt.Println(w1.Word, w1.Meaning)
		_, err1 := models.DB.Exec("UPDATE Words set learnt = 1 where word = ? AND meaning = ?", w1.Word, w1.Meaning)
		if err1 != nil {
			panic(err1)
		}
		sw1.SWord = w1.Word
		sw1.SMeaning = w1.Meaning
		sw = append(sw, sw1)
		fmt.Println("Successfully fetched a new word.")
	}

	c.JSON(http.StatusOK, sw)
}

func Learntwords(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	var lw []models.Word
	var sw []showWords
	result, err := models.DB.Query("SELECT * FROM Words WHERE learnt = 1 order by learntAt asc ")
	if err != nil {
		panic(err)
	}
	defer result.Close()
	for result.Next() {
		var l1 models.Word
		var sw1 showWords
		err := result.Scan(&l1.Word, &l1.Meaning, &l1.Learnt, &l1.LearntAt)
		if err != nil {
			panic(err)
		}
		lw = append(lw, l1)
		sw1.SWord = l1.Word
		sw1.SMeaning = l1.Meaning
		sw = append(sw, sw1)
		//fmt.Println(l1.Word, l1.Meaning)

	}
	fmt.Println("Successfully retrieved.")

	c.JSON(http.StatusOK, sw)
}

func CheckWord(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	var Resword []models.Word2
	var FinalRes []models.SearchWord
	param := c.Param("word")
	if param == "" {
		c.JSON(http.StatusOK, "Please enter a word ")
	}

	result, err := models.DB.Query("SELECT word FROM Words WHERE word = ? ", param)
	if err != nil {
		panic(err)
	}
	defer result.Close()
	//c.JSON(http.StatusOK, gin.H{"data": result})
	i := 0

	for result.Next() {
		i++
		var check models.Word2
		var scheck models.SearchWord
		err := result.Scan(&check.Word)
		if err != nil {
			panic(err)
		}
		Resword = append(Resword, check)
		scheck.Word = check.Word
		FinalRes = append(FinalRes, scheck)
		c.JSON(http.StatusOK, gin.H{"data": FinalRes})

	}
	if i == 0 {
		c.JSON(http.StatusOK, "empty")
	}

}

func GetMeaning(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	var Resword []models.Word
	var sw []showWords
	param := c.Param("word")
	if param == "" {
		c.JSON(http.StatusOK, "Please enter a word ")
	}
	result, err := models.DB.Query("SELECT * FROM Words WHERE word = ? ", param)
	if err != nil {
		panic(err)
	}
	defer result.Close()
	for result.Next() {
		var l1 models.Word
		var sw1 showWords
		err := result.Scan(&l1.Word, &l1.Meaning, &l1.Learnt, &l1.LearntAt)
		if err != nil {
			panic(err)
		}
		Resword = append(Resword, l1)
		sw1.SWord = l1.Word
		sw1.SMeaning = l1.Meaning
		sw = append(sw, sw1)
		//fmt.Println(l1.Word, l1.Meaning)

	}
	fmt.Println("Successfully retrieved.")

	c.JSON(http.StatusOK, sw)
}
