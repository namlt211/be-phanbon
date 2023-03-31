package auth

import (
	"encoding/json"
	"green/helpers"
	"green/helpers/crypt"
	"green/models"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Person struct {
    Name string `json:"name"`
    Age int `json:"age"`
}

// Chức năng đăng ký
func Register(w http.ResponseWriter, r *http.Request) {
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string {"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()
	//check password
	if err := crypt.VerifyPassword(userInput.Password); err != nil {
		response := map[string]string {"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	userInput.Password = crypt.HashPassword(userInput.Password)
		//check email
	if !crypt.ValidateEmail(userInput.Email){
		response := map[string]string {"message": "Email không đúng định dạng !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	//check phone
	if !crypt.IsValidPhoneNumber(userInput.Phone){
		response := map[string]string {"message": "Số điện thoại không đúng định dạng !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	//check  exist user
	if !crypt.UserExists(models.MysqlConn, userInput.UserName, userInput.Email, userInput.Phone){
		response := map[string]string {"message": "Tên đăng nhập, Email, Phone đã tồn tại !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	// create new user
	if err := models.MysqlConn.Create(&userInput).Error; err != nil {
		response := map[string]string {"message": "Lỗi khi tạo tài khoản mới !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]string{"message": "Tạo tài khoản thành công !"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

//Chức năng đăng nhập
func Login(w http.ResponseWriter, r *http.Request) {
	var userInput models.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	defer r.Body.Close()

	var user models.User
	if err := models.MysqlConn.Where("user_name = ?", userInput.UserName).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]interface{}{
				"Message": "Tài khoản hoặc mật khẩu không đúng !",
				"Code": 101,
			}
			helpers.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": "Xuất hiện lỗi khi đăng nhập hệ thống !"}
			helpers.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Tài khoản hoặc mật khẩu không đúng 1 !"}
		helpers.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}
	
	refreshToken, err := crypt.RegisterAccessToken(&user)
	if err != nil{
		response := map[string]string {"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
		//set token cookie
		http.SetCookie(w, &http.Cookie{
		Name: "user_token",
		Path: "/",
		Value: refreshToken,
		HttpOnly: true,
	})

	response := map[string]interface{} {"message": "Đăng nhập thành công", "status": true, "user" : user}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

//Chức năng đăng xuất
func Logout(w http.ResponseWriter, r *http.Request) {
	//xóa toke cookie
	http.SetCookie(w, &http.Cookie{
		Name: "user_token",
		Path: "/",
		Value: "",
		HttpOnly: false,
		MaxAge: -1,
	})

	response := map[string]string{"message":"Tạm biệt !"}
	helpers.ResponseJSON(w, http.StatusOK, response)
}


