package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

type _mingderwang_Resource struct {
	db gorm.DB
}

// @Title Create_mingderwang_
// @Description get string by ID
// @Accept  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 201 {object} string
// @Failure 400 {object} APIError "problem decoding body"
// @Router /_log4analytics_/ [post]
func (tr *_mingderwang_Resource) Create_mingderwang_(c *gin.Context) {
	var _log4analytics_ _mingderwang_

	if c.Bind(&_log4analytics_) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "problem decoding body"})
		return
	}
	_log4analytics_.Status = _mingderwang_Status
	_log4analytics_.Created = int32(time.Now().Unix())

	tr.db.Save(&_log4analytics_)

	c.JSON(http.StatusCreated, _log4analytics_)
}

func (tr *_mingderwang_Resource) GetAll_mingderwang_s(c *gin.Context) {
	var _log4analytics_s []_mingderwang_

	tr.db.Order("created desc").Find(&_log4analytics_s)

	c.JSON(http.StatusOK, _log4analytics_s)
}

func (tr *_mingderwang_Resource) Get_mingderwang_(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "problem decoding id sent"})
		return
	}

	var _log4analytics_ _mingderwang_

	if tr.db.First(&_log4analytics_, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		c.JSON(http.StatusOK, _log4analytics_)
	}
}

func (tr *_mingderwang_Resource) Update_mingderwang_(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "problem decoding id sent"})
		return
	}

	var _log4analytics_ _mingderwang_

	if c.Bind(&_log4analytics_) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "problem decoding body"})
		return
	}
	_log4analytics_.Id = int32(id)

	var existing _mingderwang_

	if tr.db.First(&existing, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		tr.db.Save(&_log4analytics_)
		c.JSON(http.StatusOK, _log4analytics_)
	}

}

func (tr *_mingderwang_Resource) Patch_mingderwang_(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "problem decoding id sent"})
		return
	}

	// this is a hack because Gin falsely claims my unmarshalled obj is invalid.
	// recovering from the panic and using my object that already has the json body bound to it.
	var json []Patch

	r := c.Bind(&json)
	if r != nil {
		fmt.Println(r)
	} else {
		if json[0].Op != "replace" && json[0].Path != "/status" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "PATCH support is limited and can only replace the /status path"})
			return
		}
		var _log4analytics_ _mingderwang_

		if tr.db.First(&_log4analytics_, id).RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
		} else {
			_log4analytics_.Status = json[0].Value

			tr.db.Save(&_log4analytics_)
			c.JSON(http.StatusOK, _log4analytics_)
		}
	}
}

func (tr *_mingderwang_Resource) Delete_mingderwang_(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "problem decoding id sent"})
		return
	}

	var _log4analytics_ _mingderwang_

	if tr.db.First(&_log4analytics_, id).RecordNotFound() {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	} else {
		tr.db.Delete(&_log4analytics_)
		c.Data(http.StatusNoContent, "application/json", make([]byte, 0))
	}
}

func (tr *_mingderwang_Resource) getId(c *gin.Context) (int32, error) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return int32(id), nil
}

/**
* on patching: http://williamdurand.fr/2014/02/14/please-do-not-patch-like-an-idiot/
 *
  * patch specification https://tools.ietf.org/html/rfc5789
   * json definition http://tools.ietf.org/html/rfc6902
*/

type Patch struct {
	Op    string `json:"op" binding:"required"`
	From  string `json:"from"`
	Path  string `json:"path"`
	Value string `json:"value"`
}
