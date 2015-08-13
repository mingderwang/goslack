package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"strconv"
	"time"
)

type SlackResource struct {
	db gorm.DB
}

// @Title CreateSlack
// @Description get string by ID
// @Accept  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 201 {object} string
// @Failure 400 {object} APIError "problem decoding body"
// @Router /slack/ [post]
func (tr *SlackResource) CreateSlack(c *gin.Context) {
	var slack Slack

	if c.Bind(&slack) != nil {
		c.JSON(400, NewError("problem decoding body"))
		return
	}
	slack.Status = SlackStatus
	slack.Created = int32(time.Now().Unix())

	tr.db.Save(&slack)

	c.JSON(201, slack)
}

func (tr *SlackResource) GetAllSlacks(c *gin.Context) {
	var slacks []Slack

	tr.db.Order("created desc").Find(&slacks)

	c.JSON(200, slacks)
}

func (tr *SlackResource) GetSlack(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(400, NewError("problem decoding id sent"))
		return
	}

	var slack Slack

	if tr.db.First(&slack, id).RecordNotFound() {
		c.JSON(404, gin.H{"error": "not found"})
	} else {
		c.JSON(200, slack)
	}
}

func (tr *SlackResource) UpdateSlack(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(400, NewError("problem decoding id sent"))
		return
	}

	var slack Slack

	if c.Bind(&slack) != nil {
		c.JSON(400, NewError("problem decoding body"))
		return
	}
	slack.Id = int32(id)

	var existing Slack

	if tr.db.First(&existing, id).RecordNotFound() {
		c.JSON(404, NewError("not found"))
	} else {
		tr.db.Save(&slack)
		c.JSON(200, slack)
	}

}

func (tr *SlackResource) PatchSlack(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(400, NewError("problem decoding id sent"))
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
			c.JSON(400, NewError("PATCH support is limited and can only replace the /status path"))
			return
		}
		var slack Slack

		if tr.db.First(&slack, id).RecordNotFound() {
			c.JSON(404, NewError("not found"))
		} else {
			slack.Status = json[0].Value

			tr.db.Save(&slack)
			c.JSON(200, slack)
		}
	}
}

func (tr *SlackResource) DeleteSlack(c *gin.Context) {
	id, err := tr.getId(c)
	if err != nil {
		c.JSON(400, NewError("problem decoding id sent"))
		return
	}

	var slack Slack

	if tr.db.First(&slack, id).RecordNotFound() {
		c.JSON(404, NewError("not found"))
	} else {
		tr.db.Delete(&slack)
		c.Data(204, "application/json", make([]byte, 0))
	}
}

func (tr *SlackResource) getId(c *gin.Context) (int32, error) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return int32(id), nil
}
