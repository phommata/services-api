package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/phommata/services-api/config"
	"github.com/phommata/services-api/responses"
	"github.com/phommata/services-api/models"
	"log"
	"strconv"
	"net/http"
)

const error = "error"

type ServiceRepo struct {
	Db *gorm.DB
}

func New(conf *config.Conf) *ServiceRepo {
	db := models.InitDb(conf)

	return &ServiceRepo{Db: db}
}

// GET /services
// Get all services
func (repository *ServiceRepo)  FindServices(c *gin.Context) {
	var services []responses.Service
	var count int64
	var searchWhere string

	limitStr := c.DefaultQuery("limit", "12")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errMsg := "limit is not an int"
		log.Println(errMsg, err)

		c.JSON(http.StatusBadRequest,
			gin.H{error: errMsg},
		)
		return
	}
	offsetStr := c.DefaultQuery("offset", "0")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		errMsg := "offset is not an int"
		log.Println(errMsg, err)
		c.JSON(http.StatusBadRequest,
			gin.H{error: errMsg},
		)
		return
	}
	search := c.Query("search")

	if search != "" {
		searchWhere = "s.name LIKE '%" + search + "%' OR s.description LIKE '%" + search + "%'"
		repository.Db.Table("services s").
		Select("count(*) count").Where("s.name LIKE '%" + search + "%' OR s.description LIKE '%" + search + "%'").
		Count(&count)
	} else {
		repository.Db.Table("services").Select("count(*)").Count(&count)
	}

	repository.Db.Debug().Limit(limit).Offset(offset).Preload("Versions").Table("services s").
	Select("*").
	Joins("left join " +
		"(" +
			"select s.id services_id, count(v.service_id) as version_count " +
			"from services s " +
			"join versions v on v.service_id = s.id " +
			"group by s.id" +
		") version_counts " +
		"ON version_counts.services_id = s.id").
	Find(&services, searchWhere).Order("s.created_at")

	c.JSON(http.StatusOK,
		gin.H{
			"data": services,
			"limit": limit,
			"offset": offset,
			"total": count,
		},
	)
}