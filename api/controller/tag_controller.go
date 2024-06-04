package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"

	model "golang-gin-api/model"
	services "golang-gin-api/service/interface"
)

type TagController struct {
	tagService services.TagService
}

type TagResponse struct {
	ID   uint   `copier:"must"`
	Name string `copier:"must"`
}

func NewTagController(service services.TagService) *TagController {
	return &TagController{
		tagService: service,
	}
}

func (p *TagController) FindAll(c *gin.Context) {
	tags, err := p.tagService.FindAll(c.Request.Context())

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := []TagResponse{}
		copier.Copy(&response, &tags)

		c.JSON(http.StatusOK, response)
	}
}

func (p *TagController) FindByID(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	tag, err := p.tagService.FindByID(c.Request.Context(), uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := TagResponse{}
		copier.Copy(&response, &tag)

		c.JSON(http.StatusOK, response)
	}
}

func (p *TagController) Save(c *gin.Context) {
	var tag model.Tags

	if err := c.BindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	tag, err := p.tagService.Save(c.Request.Context(), tag)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		response := TagResponse{}
		copier.Copy(&response, &tag)

		c.JSON(http.StatusOK, response)
	}
}

func (p *TagController) Delete(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "cannot parse id",
		})
		return
	}

	ctx := c.Request.Context()
	tag, err := p.tagService.FindByID(ctx, uint(id))

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	if tag == (model.Tags{}) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "tag doesnt exist",
		})
		return
	}

	p.tagService.Delete(ctx, tag)

	c.JSON(http.StatusOK, gin.H{"message": "tag is deleted successfully"})
}
