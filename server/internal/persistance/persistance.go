package persistance

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kradnoel/teka/internal/types"
)

var db *gorm.DB

func init() {
	conn, err := gorm.Open("sqlite3", "./devlb.db")
	if err != nil {
		log.Panic(err)
	}
	log.Println("Connection Established")

	db = conn
	//db.Debug().DropTable(&types.Subtask{}, &types.Project{}, &types.Task{})
	db.Debug().AutoMigrate(&types.Project{}, &types.Task{}, &types.Subtask{})
}

func GetPersistance() *gorm.DB {
	return db
}
