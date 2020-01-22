package dao

type Worker struct {
	ID          int `json:"id" orm:"column(id);auto"`
	Appkey      string
	Channel     string
	Host        string
	SiteId      int
	Capacity    float64 `gorm:"default:1.5"`
	CurrentLoad float64 `gorm:"default:0"`
	Active      int `gorm:"default:1"`
	BaseWeight  float64  `gorm:"default:1.5"`
	Weight      float64	`gorm:"default:1.5"`
}