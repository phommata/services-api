package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/phommata/services-api/config"
	"github.com/phommata/services-api/models"
	"github.com/phommata/services-api/requests"
	"github.com/phommata/services-api/responses"
	"log"
	"net/http"
)

const Error = "error"

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
	var service  requests.Service
	var services []responses.Service
	var count int64
	var searchWhere string

	if err := c.Bind(&service); err != nil {
		errMsg := "cannot bind service"
		log.Println(errMsg, err)

		c.JSON(http.StatusBadRequest,
			gin.H{Error: err.Error()},
		)
		return
	}

	if service.Search != "" {
		searchWhere = "s.name LIKE '%" + service.Search + "%' OR s.description LIKE '%" + service.Search + "%'"
		repository.Db.Table("services s").
		Select("count(*) count").Where("s.name LIKE '%" + service.Search + "%' OR s.description LIKE '%" + service.Search + "%'").
		Count(&count)
	} else {
		repository.Db.Table("services").Select("count(*)").Count(&count)
	}

	repository.Db.Limit(service.Limit).Offset(service.Offset).Preload("Versions").Table("services s").
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
			"limit": service.Limit,
			"offset": service.Offset,
			"total": count,
		},
	)
}