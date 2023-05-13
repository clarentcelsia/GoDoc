package worker

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"simple-go/database/config"
	"simple-go/utils"
	"sync"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// QUEUE, CHANNELS, BATCH
var wg sync.WaitGroup
var mtx sync.Mutex
var counters = 0

func Basic() {
	start := time.Now()

	// SETTING RUNTIME PROCESSOR
	// runtime.GOMAXPROCS(2)

	// SETTING DB CONFIG
	config.SetConfig()

	db, err := config.ConnectDetail()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxIdleConns(DBMaxIdleConns)
	db.SetMaxOpenConns(DBMaxConns)
	db.SetConnMaxLifetime(2 * time.Second)

	// FILE
	file, errf := os.Open("file/data.csv")
	if errf != nil {
		panic(errf)
	}
	defer file.Close()

	var (
		headers = make([]string, 0)
		jobs    = make(chan []interface{}, 0)
		// mtx     sync.Mutex
	)

	// PREPARE BACKGROUND WORKER
	for i := 0; i <= 1*Workers; i++ { // 100 * 1000 = 10rb
		wg.Add(1)
		go works(jobs, db)
	}

	// PREPARE FILE
	reader := csv.NewReader(file)
	for {
		record, errread := reader.Read()
		if errread != nil {
			if errread == io.EOF {
				errread = nil
				// log.Fatal(errread)
			}
			break
		}

		if len(headers) == 0 {
			headers = record // >>> [NO, NAME]
		}

		var rows []interface{}
		for _, each := range record {
			rows = append(rows, each)
		}

		jobs <- rows // >>> [1 NAMA1] [2 NAMA2]
	}

	//================= 10000 ==========================
	// c := 0
	// for {
	// 	if c == 10000 {
	// 		break
	// 	}
	// 	var rows []interface{}

	// 	for k := 0; k < 7; k++ {
	// 		rows = append(rows, k)
	// 	}
	// 	c++
	// 	jobs <- rows
	// }

	close(jobs)

	wg.Wait()

	duration := utils.Tracker(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
	fmt.Println("COUNTER : ", counters)
}

func works(jobs chan []interface{}, db *sql.DB) {
	defer wg.Done()

	stmt, err := db.PrepareContext(context.Background(), "INSERT INTO [DETAIL].dbo.[WORKERS] (NO,NAME, NAME2, NAME3, NAME4, NAME5, NAME6) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	tx, errtx := db.Begin()
	if errtx != nil {
		panic(errtx)
	}

	for job := range jobs { //1000
		mtx.Lock()
		// multiple queries/batchin
		_, errst := tx.Stmt(stmt).ExecContext(context.Background(), job...)
		if errst != nil {
			tx.Rollback()
			// return
			panic(errst)
		}
		// tx.Stmt(stmt).Exec(job...)

		counters++
		mtx.Unlock()
	}

	if errcmt := tx.Commit(); errcmt != nil {
		panic(errcmt)
	}
}
