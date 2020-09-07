package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/vipin23/vmart-api/api_server"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseModel
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Mobile   string `gorm:"size:13;unique" json:"mobile"`
	Password string `gorm:"size:100;not null;" json:"password"`
	Nickname string `gorm:"size:100" json:"nickname"`
	RoleID   uint64 `gorm:"null" json:"role_id"`
	Role     Role
	Profile  *Profile
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil

	default:
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

func (u *User) Login() error {
	var err error
	validUser := User{}
	err = api_server.DefaultServer.DB.Debug().Where("email = ? OR mobile = ?", u.Email, u.Mobile).First(&validUser).Error
	if err != nil {
		return err
	}
	err = VerifyPassword(validUser.Password, u.Password)
	if err != nil {
		return err
	}
	u = &validUser
	return err
}

func (u *User) SaveUser() (*User, error) {

	var err error
	err = api_server.DefaultServer.DB.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers() (*[]User, error) {
	var err error
	users := []User{}
	err = api_server.DefaultServer.DB.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(uid uint32) (*User, error) {
	var err error
	err = api_server.DefaultServer.DB.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db := api_server.DefaultServer.DB.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":  u.Password,
			"email":     u.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = api_server.DefaultServer.DB.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(uid uint32) (int64, error) {

	db := api_server.DefaultServer.DB.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
