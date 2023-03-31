package customer

import (
	"encoding/json"
	"green/helpers"
	"green/helpers/crypt"
	"green/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	var customer []struct{
		Id uint `json:"id"`
		Name string `json:"name" gorm: "column:name;"`
		Phone string `json:"phone" gorm: "column:phone;"`
		Address string `json:"address" gorm: "column:address;"`
	}
	models.MysqlConn.Table(`customers`).
	Select(`customers.id, customers.name, customers.phone, customers.address`).
	Where(`customers.is_delete = 1`).
	Scan(&customer)
	response := map[string]interface{}{"message": "Lấy danh sách khách hàng thành công !", "status": true, "data": customer}
	helpers.ResponseJSON(w, http.StatusOK, response)
}


func AddCustomer(w http.ResponseWriter, r *http.Request) {
	var customer models.Customer
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&customer); err != nil {
		response := map[string]string {"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()
	//check phone
	if !crypt.IsValidPhoneNumber(customer.Phone){
		response := map[string]string {"message": "Số điện thoại không đúng định dạng !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	//check  exist user
	if !crypt.PhoneExists(models.MysqlConn.Table(`customers`), customer.Phone){
		response := map[string]string {"message": "Số điện thoại đã tồn tại !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	if err := models.MysqlConn.Create(&customer).Error; err != nil {
		response := map[string]string {"message": "Lỗi khi thêm khách hàng mới !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]interface{}{"message": "Thêm khách hàng thành công !", "status": true, "data": customer}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    idParam := vars["id"]
	if idParam == "" {
		response := map[string]string{"message": "Không có truyền lên mã khách hàng !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := map[string]string{"message": "Không có truyền lên mã khách hàng !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var customerUpdate models.Customer

	err = json.NewDecoder(r.Body).Decode(&customerUpdate)
	if err != nil {
		response := map[string]string{"message": "Dữ liệu không hợp lệ !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	//check phone
	if !crypt.IsValidPhoneNumber(customerUpdate.Phone){
		response := map[string]string {"message": "Số điện thoại không đúng định dạng !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	
	customer :=&models.Customer{}
	result := models.MysqlConn.First(customer, id)
	if result.Error != nil {
		response := map[string]string{"message": "Không tìm thấy khách hàng hợp lệ !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	customer.Name = customerUpdate.Name
	customer.Phone = customerUpdate.Phone
	customer.Address = customerUpdate.Address
	customer.UpdatedAt = customerUpdate.UpdatedAt
	result = models.MysqlConn.Save(customer)
	if result.Error != nil {
		response := map[string]string{"message": "Không thể cập nhật khách hàng !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	
	response := map[string]interface{}{"message": "Cập nhật thông tin khách hàng thành công !", "status": true}
	helpers.ResponseJSON(w, http.StatusOK, response)
}


func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	
	vars := mux.Vars(r)
    idParam := vars["id"]
	if idParam == "" {
		response := map[string]string{"message": "Không có truyền lên mã khách hàng !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := map[string]string{"message": "Không có truyền lên mã khách hàng !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	var customer struct{
		Id uint `json:"id"`
		Name string `json:"name" gorm: "column:name;"`
		Phone string `json:"phone" gorm: "column:phone;"`
		Address string `json:"address" gorm: "column:address;"`
	}
	models.MysqlConn.Table(`customers`).
	Select(`customers.id, customers.name, customers.phone, customers.address`).
	Where(`customers.is_delete = 1 && customers.id = ?`, id).
	Scan(&customer)
	response := map[string]interface{}{"message": "Đã thấy !", "status": true, "data": customer}
	helpers.ResponseJSON(w, http.StatusOK, response)
}


func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    idParam := vars["id"]
	if idParam == "" {
		response := map[string]string{"message": "Không có truyền lên mã khách hàng !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := map[string]string{"message": "Không có truyền lên mã khách hàng !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	models.MysqlConn.Table("customers").Where("id = ?", id).Update("is_delete", 0)
	response := map[string]interface{}{"message": "Đã xóa !", "status": true}
	helpers.ResponseJSON(w, http.StatusOK, response)
}
