// @APIVersion 1.0.0
// @APITitle Swagger Example API
// @APIDescription Swagger Example API
// @Contact varyous@gmail.com
// @TermsOfServiceUrl http://yvasiyarov.com/
// @License BSD
// @LicenseUrl http://yvasiyarov.com/
package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	SvcHost    string
	DbUser     string
	DbPassword string
	DbHost     string
	DbName     string
	Token      string
	Url        string
}

type SlackService struct {
}

func (s *SlackService) getDb(cfg Config) (gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "/tmp/"+cfg.DbName)
	//db.LogMode(true)
	return db, err
}

func (s *SlackService) Migrate(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		return err
	}
	db.SingularTable(true)

	db.AutoMigrate(&Slack{})
	return nil
}
func (s *SlackService) Run(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		return err
	}
	db.SingularTable(true)

	slackResource := &SlackResource{db: db}

	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	r.GET("/slack_message", slackResource.GetAllSlacks)
	r.GET("/slack_message/:id", slackResource.GetSlack)
	r.POST("/slack_message", slackResource.CreateSlack)
	r.PUT("/slack_message/:id", slackResource.UpdateSlack)
	r.PATCH("/slack_message/:id", slackResource.PatchSlack)
	r.DELETE("/slack_message/:id", slackResource.DeleteSlack)

	r.Run(cfg.SvcHost)

	return nil
}
