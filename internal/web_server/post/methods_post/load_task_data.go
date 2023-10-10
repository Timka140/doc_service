package methods_post

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"projects/doc/doc_service/internal/db"
	"projects/doc/doc_service/internal/task_local"
	"projects/doc/doc_service/internal/web_server/sessions"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TLoadTaskData struct {
	// ctx *gin.Context
}

func newLoadTaskData(in *TInPostPage) IPostPage {
	t := &TLoadTaskData{}

	return t
}

func (t *TLoadTaskData) GetPath() string {
	return "/load_task_data"
}

func (t *TLoadTaskData) GetContext(c *gin.Context) {
	ses := sessions.Ses.GetSes(c)
	if ses == nil {
		return
	}

	err := c.Request.ParseForm()
	if err != nil {
		log.Println("TLoadTaskData.GetContext(): чтение формы, err=%w", err)
		return
	}

	form, _ := c.MultipartForm()

	task_id := form.Value["task_id"][0]
	file := form.File["file"]

	if task_id == "" {
		return
	}

	if len(file) != 1 {
		return
	}

	var name string
	var data bytes.Buffer

	for _, file := range file {
		f, err := file.Open()
		if err != nil {
			log.Println("TLoadTaskData.GetContext(): чтение файла, err=%w", err)
			return
		}
		name = file.Filename
		io.Copy(&data, f)
		f.Close()
	}

	fmt.Println(name)

	task := db.Task{}
	err = db.DB.Table("task").Where("task_id = ?", task_id).First(&task).Error
	if err != nil {
		log.Println("TLoadTaskData.GetContext(): чтение записи, err=%w", err)
	}

	db, err := gorm.Open(sqlite.Open(task.PathBase+".db3"), &gorm.Config{})
	if err != nil {
		log.Printf("NewDB(): открытие базы данных, err=%w", err)
	}

	// go func(db *gorm.DB) {
	tables, err := task_local.XlsxToBase(&task_local.TInXlsxToBase{File: data, List: "ad"})
	if err != nil {
		log.Println("TLoadTaskData.GetContext(): чтение таблицы, err=%w", err)
	}
	for name, table := range tables {
		table_sql := fmt.Sprintf("CREATE TABLE \"%v\" (\n", name)
		for ind, col := range table.Headers {
			if ind != len(table.Headers)-1 {
				table_sql += fmt.Sprintf("\"%v\"	TEXT,\n", col)
			} else {
				table_sql += fmt.Sprintf("\"%v\"	TEXT\n", col)
			}
		}
		table_sql += ");"
		err := db.Exec(table_sql).Error
		if err != nil {
			log.Println(err)
		}

		for _, row := range table.Data {
			err = db.Table(name).Create(&row).Error
			if err != nil {
				log.Println(err)
			}
		}
	}
	// }(db)

	c.Status(http.StatusOK)
}

func init() {
	err := constructors.Add("LoadTaskData", newLoadTaskData)
	if err != nil {
		log.Printf("LoadTaskData(): не удалось добавить в конструктор")
	}
}
