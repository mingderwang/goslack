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

type _mingderwang_Service struct {
}

func (s *_mingderwang_Service) getDb(cfg Config) (gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "/tmp/"+cfg.DbName)
	//db.LogMode(true)
	return db, err
}

func (s *_mingderwang_Service) Migrate(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		return err
	}
	db.SingularTable(true)

	db.AutoMigrate(&_mingderwang_{})
	return nil
}
func (s *_mingderwang_Service) Run(cfg Config) error {
	db, err := s.getDb(cfg)
	if err != nil {
		return err
	}
	db.SingularTable(true)

	slackResource := &_mingderwang_Resource{db: db}

	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)

	r.GET("/_mingderwang_", slackResource.GetAll_mingderwang_s)
	r.GET("/_mingderwang_/:id", slackResource.Get_mingderwang_)
	r.POST("/_mingderwang_", slackResource.Create_mingderwang_)
	r.PUT("/_mingderwang_/:id", slackResource.Update_mingderwang_)
	r.PATCH("/_mingderwang_/:id", slackResource.Patch_mingderwang_)
	r.DELETE("/_mingderwang_/:id", slackResource.Delete_mingderwang_)

	r.Run(cfg.SvcHost)

	return nil
}
