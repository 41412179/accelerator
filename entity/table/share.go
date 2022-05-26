package table

type Share struct {
	Id  int    `gorm:"column:id;type:int(11);primary_key" json:"id"`
	Url string `gorm:"column:url;type:varchar(512);NOT NULL" json:"url"` // 分享地址
}

func (m *Share) TableName() string {
	return "share"
}
