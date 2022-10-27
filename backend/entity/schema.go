package entity

import (
	"time"

	"gorm.io/gorm"
)

type Authoritie struct {
	gorm.Model
	Name      string
	Email     string `gorm:"uniqeIndex"`
	//Password  string
	Schedules []Schedule `gorm:"foreignKey:AuthoritieID"`
}

type User struct {
	gorm.Model
	Name      string
	Email     string `gorm:"uniqeIndex"`
	Password  string
	Role      string
	Schedules []Schedule `gorm:"foreignKey:AuthoritieID"`
}

type Doctor struct {
	gorm.Model
	Name string

	Schedules []Schedule `gorm:"foreignKey:DoctorID" `
}

type Location struct {
	gorm.Model
	Name string

	Schedules []Schedule `gorm:"foreignKey:LocationID"`
}

type Room struct {
	gorm.Model
	Name string

	Schedules []Schedule `gorm:"foreignKey:RoomID"`
}

type Department struct {
	gorm.Model
	Name string

	Schedules []Schedule `gorm:"foreignKey:DepartmentID"`
}

type Schedule struct {
	gorm.Model
	ScheduleTime time.Time


	// DoctorID เป็น FK
	UserID *uint
	// ข้อมูลของ Doctor เมื่อ join ตาราง
	User User `gorm:"references:id"`



	// DoctorID เป็น FK
	DoctorID *uint
	// ข้อมูลของ Doctor เมื่อ join ตาราง
	Doctor Doctor `gorm:"references:id"`

	// LocationID  เป็น FK
	LocationID *uint
	// ข้อมูลของ Location เมื่อ join ตาราง
	Location Location `gorm:"references:id"`

	// AuthoritiesID เป็น FK
	AuthoritieID *uint
	// ข้อมูลของ Authorities เมื่อ join ตาราง
	Authoritie Authoritie `gorm:"references:id"`

	// AuthoritiesID เป็น FK
	RoomID *uint
	// ข้อมูลของ Authorities เมื่อ join ตาราง
	Room Room `gorm:"references:id"`

	// AuthoritiesID เป็น FK
	DepartmentID *uint
	// ข้อมูลของ Authorities เมื่อ join ตาราง
	Department Department `gorm:"references:id"`
}
