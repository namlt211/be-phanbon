package crypt

import (
	"fmt"
	"green/models"
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//kiểm tra định dạng của email
func ValidateEmail(email string) bool {
	patter := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(patter)
	return regex.MatchString(email) 
}
//kiểm tra tính toàn vẹn của mật khẩu
func VerifyPassword(s string) error {
	var (
		MinimunSeven, number, upper, special bool
		message string
	)
	letters := 0
	for _, s := range s {
		switch {
		case unicode.IsNumber(s):
			number = true
		case unicode.IsUpper(s):
			upper = true
		case unicode.IsPunct(s) || unicode.IsSymbol(s):
			special = true
		case unicode.IsLetter(s) || s == ' ':
		}
		letters ++
	} 
	MinimunSeven = letters >= 7
	message = ""
	if !MinimunSeven{
		message = message + "\nMật khẩu của bạn quá ngắn (ít hơn 7 ký tự) !"
	}
	if !number {
		message = message + "\nMật khẩu của bạn phải chứa ít nhất 1 số !"
	}
	if !upper{
		message = message + "\nMật khẩu của bạn phải chứa ít nhất 1 chữ hoa"
	}
	if !special{
		message = message + "\nMật khẩu của bạn phải chứa ít nhất 1 ký tự đặc biệt"
	}
	if message != ""{
		return fmt.Errorf(message)
	}
	return nil
}
// mã hóa mật khẩu
func HashPassword(password string) string {
	 hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	 return string(hashPassword)
}
//kiểm tra tính hợp lệ của số điện thoại
func IsValidPhoneNumber(number string) bool {
	regex := regexp.MustCompile(`^(\(\d{3}\)|\d{3})[- ]?\d{3}[- ]?\d{4}$`)
	return regex.MatchString(number)
}
// kiểm tra tên đăng nhập, email, phone có tồn tại trong data 
func UserExists(db * gorm.DB, username string, email string, phone string) bool {
	var user models.User
	db.Where("user_name = ?", username).Or("email = ?", email).Or("phone = ?", phone).First(&user)
	return user.Id == 0
}
//kiểm tra số điện thoại đã tồn tại
func PhoneExists(db * gorm.DB, phone string) bool {
	var user models.User
	db.Where("phone = ?", phone).First(&user)
	return user.Id == 0
}


func SupplierNameExists(db * gorm.DB, name string) bool {
	var supplier models.Supplier
	db.Where("name = ?", name).First(&supplier)
	return supplier.Id == 0
}