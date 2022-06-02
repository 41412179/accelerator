package table

type Version struct {
	Id      int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Version string `gorm:"column:version;type:varchar(45);NOT NULL" json:"version"`
	Url     string `gorm:"column:url;type:varchar(45);NOT NULL" json:"url"`
	Desc    string `gorm:"column:desc;type:varchar(45);NOT NULL" json:"desc"`
	CodeID  string `gorm:"column:code_id;type:varchar(45);NOT NULL" json:"code_id"`
}

func (m *Version) TableName() string {
	return "version"
}
