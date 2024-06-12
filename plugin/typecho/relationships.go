package typecho

type Relationships struct {
	Cid uint `gorm:"primaryKey;not null;autoIncrement:false"`
	Mid uint `gorm:"primaryKey;not null;autoIncrement:false"`
}
