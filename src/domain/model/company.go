package model

type Company struct {
	BaseModel

	Name      string `gorm:"size:1000;type:string;not null,unique;"`
	ShortName string `gorm:"size:200;type:string;not null,unique;"`
	Status    int    `gorm:"not null;default:1"`
	Address   string `gorm:"size:2000;type:string;null;"`
	Phone     string `gorm:"size:400;type:string;null;"`
	Email     string `gorm:"size:100;type:string;null;"`
	Website   string `gorm:"size:200;type:string;null;"`
	Identity  string `gorm:"size:50;type:string;null;unique;"`
	Tz        string `gorm:"size:100;not null;default:'America/La_Paz'"`
	License   string `gorm:"size:1000;type:string;null;"`

	// Country Country `gorm:"foreignKey:CountryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}
