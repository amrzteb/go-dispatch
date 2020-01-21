package dao

type Worker struct {
	ID         int `json:"id" orm:"column(id);auto"`
	appkey     string
	channel    string
	host       string
	siteId     int
	capacity   float64
	active     int
	baseWeight float64
	weight     float64
}