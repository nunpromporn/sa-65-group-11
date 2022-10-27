package entity

import (

	// "fmt"
	// "time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema
	database.AutoMigrate(
		&User{}, 
		&Doctor{}, &Location{}, &Authoritie{}, &Department{}, &Room{}, &Schedule{},
	)

	db = database
	/////////////////////////////////////////
	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password1, err := bcrypt.GenerateFromPassword([]byte("654321"), 14)

	db.Model(&User{}).Create(&User{
			Name:     "Promporn Phinitphong",
			Email:    "promporn@gmail.com",
			Password: string(password),
			Role: 		"admin",
		})

		db.Model(&User{}).Create(&User{
			Name:     "abc",
			Email:    "abc@gmail.com",
			Password: string(password1),
			Role: 		"user",
		})


		var promporn User
	db.Raw("SELECT * FROM users WHERE email = ?", "promporn@gmail.com").Scan(&promporn)
	

	var abc User
	db.Raw("SELECT * FROM users WHERE email = ?", "abc@gmail.com").Scan(&abc)
	
	
	
	// db.Model(&Authoritie{}).Create(&Authoritie{
	// 	Name:     "Promporn Phinitphong",
	// 	Email:    "promporn@gmail.com",
	// 	Password: string(password),
	// })

	// var promporn Authoritie
	// db.Raw("SELECT * FROM authorities WHERE email = ?", "promporn@gmail.com").Scan(&promporn)

	//---Location Data
	A := Location{
		Name: "Building A",
	}
	db.Model(&Location{}).Create(&A)

	B := Location{
		Name: "Building B",
	}
	db.Model(&Location{}).Create(&B)

	//---Room A Data
	A1 := Room{
		Name: "A101",
	}
	db.Model(&Room{}).Create(&A1)

	A2 := Room{
		Name: "A102",
	}
	db.Model(&Room{}).Create(&A2)

	A3 := Room{
		Name: "A103",
	}
	db.Model(&Room{}).Create(&A3)

	B1 := Room{
		Name: "B101",
	}
	db.Model(&Room{}).Create(&B1)

	B2 := Room{
		Name: "B102",
	}
	db.Model(&Room{}).Create(&B2)

	B3 := Room{
		Name: "B103",
	}
	db.Model(&Room{}).Create(&B3)

	//---Doctor
	db.Model(&Doctor{}).Create(&Doctor{
		Name: "Thanawat Nitikarun",
	})

	DoctorNo1 := Doctor{
		Name: "Panadda Srisawat",
	}
	db.Model(&Doctor{}).Create(&DoctorNo1)

	DoctorNo2 := Doctor{
		Name: "Jakkrit   Chaiwan",
	}
	db.Model(&Doctor{}).Create(&DoctorNo2)

	DoctorNo3 := Doctor{
		Name: "Wallaya  Patisang",
	}
	db.Model(&Doctor{}).Create(&DoctorNo3)
	DoctorNo4 := Doctor{
		Name: "Ratchapol Piyaman",
	}
	db.Model(&Doctor{}).Create(&DoctorNo4)

	DoctorNo5 := Doctor{
		Name: "Thanawat Nitikarun",
	}
	db.Model(&Doctor{}).Create(&DoctorNo5)

	//---Department Data
	General := Department{
		Name: "แพทย์ทั่วไป",
	}
	db.Model(&Department{}).Create(&General)

	Orthopedics := Department{
		Name: "แพทย์กระดูก",
	}
	db.Model(&Department{}).Create(&Orthopedics)

	Cardiac := Department{
		Name: "แพทย์หัวใจ",
	}
	db.Model(&Department{}).Create(&Cardiac)

	Gynecologist := Department{
		Name: "สูตินารีแพทย์(ตรวจภายใน)",
	}
	db.Model(&Department{}).Create(&Gynecologist)

	Otolaryngology := Department{
		Name: "แพทย์เฉพาะทางด้าน ตา หู คอ จมูก",
	}
	db.Model(&Department{}).Create(&Otolaryngology)

	Psychology := Department{
		Name: "จิตเวช",
	}
	db.Model(&Department{}).Create(&Psychology)

	Phamarceutical := Department{
		Name: "เภสัชกรรม",
	}
	db.Model(&Department{}).Create(&Phamarceutical)

	Skin := Department{
		Name: "แพทย์ผิวหนัง",
	}
	db.Model(&Department{}).Create(&Skin)

}
