package models

import (
	"github.com/phommata/services-api/config"
	dbConn "github.com/phommata/services-api/adapter/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func InitDb(conf *config.Conf) *gorm.DB {
	var DB, err = dbConn.New(conf)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connection Established")
	DB.DropTableIfExists(&Service{}, &Version{})
	DB.AutoMigrate(&Service{}, &Version{})

	services := []Service{
		{
			Name: "Security",
			Description: "Lorem ipsum dolor",
			Versions: []Version{
				{
					Version: "1.0.0",
				},
				{
					Version: "1.0.1",
				},
				{
					Version: "1.0.2",
				},
			},
		},
		{
			Name: "Security",
			Description: "Lorem ipsum dolor",
			Versions: []Version{
				{
					Version: "1.0.0",
				},
				{
					Version: "1.0.1",
				},
				{
					Version: "1.0.2",
				},
			},
		},
		{
			Name: "Reporting",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor...",
			Versions: []Version{
				{
					Version: "1.0.0",
				},
				{
					Version: "1.0.1",
				},
				{
					Version: "1.0.2",
				},
			},
		},
		{
			Name: "Priority Services",
			Versions: []Version{
				{
					Version: "1.0.0",
				},
				{
					Version: "1.0.1",
				},
				{
					Version: "1.0.2",
				},
			},
		},
		{
			Name: "Notifications",
			Versions: []Version{
				{
					Version: "1.0.0",
				},
				{
					Version: "1.0.1",
				},
				{
					Version: "1.0.2",
				},
			},
		},
		{
			Name: "Notifications",
			Versions: []Version{
				{
					Version: "1.0.0",
				},
				{
					Version: "1.0.1",
				},
				{
					Version: "1.0.2",
				},
			},
		},
		{
			Name: "FX Rates International...",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor",
			Versions: []Version{
				{
					Version: "1.0.0",
				},
				{
					Version: "1.0.1",
				},
				{
					Version: "1.0.2",
				},
			},
		},
		{
			Name: "FX Rates International",
			Description: "Lorem ipsum dolor",
			Versions: []Version{
				{
					Version: "1.0.0",
				},
				{
					Version: "1.0.1",
				},
				{
					Version: "1.0.2",
				},
			},
		},
		{
			Name: "Contact Us",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing",
			Versions: []Version{
				{
					Version: "1.0.0",
				},
				{
					Version: "1.0.1",
				},
				{
					Version: "1.0.2",
				},
			},
		},
		{
			Name: "Contact Us",
			Description:"Lorem ipsum dolor sit amet, consectetur adipiscing",
			Versions: []Version{
				{
					Version: "1.0.0",
				},
				{
					Version: "1.0.1",
				},
				{
					Version: "1.0.2",
				},
			},
		},
		{
			Name: "Collect Monday",
			Versions: []Version{
				{
					Version: "1.0.0",
				},
				{
					Version: "1.0.1",
				},
				{
					Version: "1.0.2",
				},
			},
		},
		{
			Name: "Locate Us",
			Description: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor...",
			Versions: []Version{
				{
					Version: "1.0.0",
				},
				{
					Version: "1.0.1",
				},
				{
					Version: "1.0.2",
				},
			},
		},
	}

	for _, service := range services {
		DB.Create(&service)
	}

	return DB
}