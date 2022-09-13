package dbmodel

type Post struct {
	ID           uint64 `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Title        string `gorm:"not null"`
	Content      string `gorm:"not null"`
	Author       string `gorm:"not null; unique"`
	Hero         string `json:"Hero"`
	Published_At string `json:"PublishedAt"`
	Updated_At   string `json:"UpdateAt"`
}
