package responses

type Version struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Version      string `json:"name"`
	ServiceID    string `json:"service_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}