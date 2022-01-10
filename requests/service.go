package requests

type Service struct {
	Search string 	`form:"search"`
	Limit  int		`form:"limit"`
	Offset int		`form:"offset"`
}