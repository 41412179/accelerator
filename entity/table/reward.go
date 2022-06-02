package table

type Reward struct {
	Id    int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Qq    string `gorm:"column:qq;type:varchar(45)" json:"qq"`
	Count int    `gorm:"column:count;type:int(11);default:0;NOT NULL" json:"count"`
}

func (m *Reward) TableName() string {
	return "reward"
}
