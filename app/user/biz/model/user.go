package model

import "gorm.io/gorm"

// 首字母小写的字段是私有的（即不可导出）,不能gorm
type User struct {
	gorm.Model
	Email          string `gorm:"uniqueIndex;type:varchar(255) "`
	PasswordHashed string `gorm:"varchar(255) not null"`
}

func (receiver User) TableName() string {
	return "user"
}

// 创建用户
func Create(db *gorm.DB, user *User) error {
	return db.Create(&user).Error
}

func GetUserByEmail(db *gorm.DB, email string) (user *User, err error) {
	user = &User{}
	err = db.Where("email=?", email).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
