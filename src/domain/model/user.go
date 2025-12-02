package model

import (
	"gorm.io/datatypes"
)

type IdentityCompany struct {
	Company   Company `gorm:"foreignKey:CompanyID;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CompanyID int
}
type User struct {
	BaseModel

	Username     string `gorm:"size:50;not null;unique"`
	FirstName    string `gorm:"size:20"`
	LastName     string `gorm:"size:20"`
	MobileNumber string `gorm:"size:20;unique"`
	Email        string `gorm:"size:64;unique"`
	Password     string `gorm:"size:64;not null"`
	Enabled      bool   `gorm:"default:true"`

	IdentityCompany
	// UserRoles []UserRole  // opcional según tu diseño
}

type Role struct {
	BaseModel

	Name           string                      `gorm:"size:20;not null;unique"`
	AllowedURLs    datatypes.JSONSlice[string] `gorm:"type:jsonb"`
	AllowedActions datatypes.JSONSlice[string] `gorm:"type:jsonb"`

	UserRoles []UserRole
	IdentityCompany
}

type UserRole struct {
	BaseModel

	UserID int
	RoleID int

	User User `gorm:"foreignKey:UserID;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	Role Role `gorm:"foreignKey:RoleID;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`

	IdentityCompany
}
