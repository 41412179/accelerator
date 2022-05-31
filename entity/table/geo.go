package table

type Geo struct {
	Id   int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"column:name;type:varchar(45);NOT NULL" json:"name"`
	Link string `gorm:"column:link;type:varchar(45);NOT NULL" json:"link"`
}

func (m *Geo) TableName() string {
	return "geo"
}
