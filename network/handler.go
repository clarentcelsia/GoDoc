package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	Conf "simple-go/database/config"
	"simple-go/model"
	"simple-go/utils"

	"time"

	"github.com/gin-gonic/gin"
)

func I_Account(c *gin.Context) {
	// SETTING DB CONFIG
	Conf.SetConfig()

	db, err := Conf.ConnectDetail()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var param model.Param
	c.BindJSON(&param)

	trx, err := db.Begin()
	if err != nil {
		c.JSON(500, map[string]interface{}{
			"message": "connection failed!",
			"result":  nil,
		})
		return
	}

	var validation int
	var errormsg string
	errs := trx.QueryRow("EXEC SP_IU_CARD ?,?,?", param.UserId, param.CardNo, param.AccountNo).
		Scan(&validation, &errormsg)
	if errs != nil {
		c.JSON(500, map[string]interface{}{
			"message": "insert failed!",
			"result":  errs.Error(),
		})
		trx.Rollback()
		return
	}

	if validation != 200 {
		c.JSON(500, map[string]interface{}{
			"message": "insert failed with status != 200!",
			"result":  nil,
		})
		trx.Rollback()

		return
	}

	trx.Commit()
	c.JSON(500, map[string]interface{}{
		"message": "insert succeed!",
		"result":  nil,
	})

}

func GetAccount(c *gin.Context) {
	start := time.Now()

	res, errs := make(chan *http.Response), make(chan error)
	go Hit(res, errs)

	select {
	case result := <-res:
		c.JSON(200, result.Body)
	case err := <-errs:
		fmt.Println(err.Error())
		return
	}

	duration := utils.Tracker(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
}

func GetAccountManual(c *gin.Context) {
	start := time.Now()

	resp, err := http.Post("http://localhost:9090/test", "application/json", nil)
	if err != nil {
		return
	}

	c.JSON(200, resp.Body)

	duration := utils.Tracker(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
}

func Hit(res chan *http.Response, errs chan error) {

	body := map[string]interface{}{
		"UserId":    "6",
		"CardNo":    "12093093023",
		"AccountNo": "1234567890",
	}

	byt, _ := json.Marshal(body)
	resp, err := http.Post("http://localhost:9090/test", "application/json", bytes.NewBuffer(byt))
	if err != nil {
		errs <- err
	}

	fmt.Println(resp.Body)
	res <- resp
	// errs <- nil
}
