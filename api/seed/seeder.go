package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kkmonlee/gemini-backend/gemini-backend/api/models"
)

var users = []models.User{
	models.User{
		Username: "kkmonlee",
		Email:    "kkmonlee@gmail.com",
		Password: "password",
	},
	models.User{
		Username: "Isaac Newton",
		Email:    "i.newton@science.com",
		Password: "password",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "My first blog post",
		Content: "Lorem ipsum",
	},
	models.Post{
		Title:   "My second blog post",
		Content: "Doret salet",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
