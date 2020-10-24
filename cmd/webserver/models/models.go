package models

// Data Model
type Data_req struct {
	User string `gorm:"column:first_name" json:"user"`
}

type Data_resp struct {
	Status int `gorm:"column:status_id" json:"status"`
}

type Version struct {
	BuildTime string `gorm:"column:BuildTime" json:"buildTime"`
	Commit    string `gorm:"column:Commit" json:"commit"`
	Release   string `gorm:"column:first_name" json:"release"`
}
