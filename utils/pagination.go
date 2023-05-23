package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	Pagination struct {
		Limit int `json:"limit"`
		Page  int `json:"page"`
	}
)

func JSONFILE() map[string]interface{} {
	//CALL API.JSON
	file, err := os.Open("file/api.json")
	if err != nil {
		log.Print("File can't be opened,", err)
		return nil
	}
	defer file.Close()

	fileBytes, _ := ioutil.ReadAll(file)

	var result map[string]interface{}
	json.Unmarshal(fileBytes, &result)
	return result
}

func Paginating(c *gin.Context) Pagination {
	page := 1
	limit := 2

	q := c.Request.URL.Query()
	for k, v := range q {
		qVals := v[len(v)-1]
		switch k {
		case "limit":
			limit, _ = strconv.Atoi(qVals)
			break
		case "page":
			page, _ = strconv.Atoi(qVals)
			break
		}
	}

	return Pagination{
		Limit: limit,
		Page:  page,
	}
}

func GetFiles(c *gin.Context) {
	file := JSONFILE()
	req := Paginating(c)

	list := file["data"].([]interface{})
	// max_per_page := req.Page * len(list)

	var list_lim []interface{}
	start := (req.Page - 1) * req.Limit
	end := req.Limit + ((req.Page - 1) * req.Limit)
	lastpage := int(math.Ceil(float64(len(list)) / float64(req.Limit)))
	if req.Page > lastpage || end > len(list) {
		if req.Page > lastpage {
			start = 0
			end = 0
		} else if end > len(list) {
			end = len(list)
		}
	}

	list_lim = list[start:end]

	// if (req.Limit * req.Page) <= len(list) {
	// 	fmt.Println("MASIH DALAM RENTANG LIST")
	// 	list_lim = list[((req.Page * req.Limit) - req.Limit):(req.Limit * req.Page):len(list)]
	// } else if !((len(list) / req.Limit) < 1) && !((req.Page*req.Limit)-req.Limit >= len(list)) {
	// 	fmt.Println("DALAM RENTANG TAPI KELEBIHAN")
	// 	list_lim = list[((req.Page * req.Limit) - req.Limit):len(list)]
	// } else if ((req.Limit*req.Page)%len(list)) != 0 && (req.Limit <= max_per_page) && !((req.Page*req.Limit)-req.Limit >= len(list)) {
	// 	fmt.Println("LIMIT DILUAR RENTANG LIST TAPI MASIH ADA SISA LIST DI PAGE")
	// 	list_lim = list[((req.Page * req.Limit) - req.Limit):len(list)]
	// } else {
	// 	fmt.Println("LIMIT DAN PAGE DILUAR RENTANG LIST")
	// 	list_lim = nil
	// }

	c.JSON(200, list_lim)

}
