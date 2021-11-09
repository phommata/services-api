package responses

type Service struct {
	ID          uint   		`json:"id" gorm:"primary_key"`
	Name        string 		`json:"name"`
	Description string 		`json:"description"`
	Versions 	[]Version 	`json:"versions"`
	VersionCount int    	`json:"version_count"`
	CreatedAt 	string 		`json:"created_at"`
}