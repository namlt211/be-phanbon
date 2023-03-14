package product

import (
	"encoding/json"
	"green/helpers"
	"green/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddSupplier(w http.ResponseWriter, r *http.Request) {
	var supplier models.Supplier
	decode := json.NewDecoder(r.Body)
	if err := decode.Decode(&supplier); err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()
	if err := models.MysqlConn.Create(&supplier).Error; err != nil {
		response := map[string]string{"message": "Lỗi khi thêm nhà cung cấp"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	response := map[string]interface{}{"message": "Thêm nhà cung cấp thành công !", "status": true, "data": supplier }
	helpers.ResponseJSON(w, http.StatusOK, response)
}


func GetAllSupplier(w http.ResponseWriter, r *http.Request){
	var supplier []models.Supplier
	if err := models.MysqlConn.Find(&supplier).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	response := map[string]interface{}{"message": "Lấy nhà cung cấp thành công !", "status": true, "data": supplier}
	helpers.ResponseJSON(w, http.StatusOK, response)
}


func UpdateSupplier(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    idParam := vars["id"]
	if idParam == "" {
		response := map[string]string{"message": "Không có truyền lên mã nhà cung cấp !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := map[string]string{"message": "Không có truyền lên mã nhà cung cấp !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var supplierUpdate models.Supplier

	err = json.NewDecoder(r.Body).Decode(&supplierUpdate)
	if err != nil {
		response := map[string]string{"message": "Dữ liệu không hợp lệ !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	supplier :=&models.Supplier{}
	result := models.MysqlConn.First(supplier, id)

	if result.Error != nil {
		response := map[string]string{"message": "Không tìm thấy nhà cung cấp hợp lệ !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	supplier.Name = supplierUpdate.Name
	supplier.Description = supplierUpdate.Description
	supplier.UpdatedAt = supplierUpdate.UpdatedAt
	result = models.MysqlConn.Save(supplier)
	if result.Error != nil {
		response := map[string]string{"message": "Không thể cập nhật nhà cung cấp !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]interface{}{"message": "Cập nhật nhà cung cấp thành công !", "status": true, "decode": supplierUpdate}
	helpers.ResponseJSON(w, http.StatusOK, response)
}


func GetSupplierById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    idParam := vars["id"]
	if idParam == "" {
		response := map[string]string{"message": "Không có truyền lên mã nhà cung cấp !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := map[string]string{"message": "Không có truyền lên mã nhà cung cấp !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	var supplier models.Supplier
	result := models.MysqlConn.First(&supplier, id)
	if result.Error != nil {
		response := map[string]string{"message": "Không tìm thấy nhà cung cấp !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
    }
	response := map[string]interface{}{"message": "Đã thấy !", "status": true, "data": supplier}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func DeleteSupplier(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
    idParam := vars["id"]
	if idParam == "" {
		response := map[string]string{"message": "Không có truyền lên mã nhà cung cấp !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := map[string]string{"message": "Không có truyền lên mã nhà cung cấp !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	var supplier models.Supplier
	if models.MysqlConn.Delete(&supplier, id).RowsAffected == 0 {
		response := map[string]string{"message": "Lỗi khi xóa !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]interface{}{"message": "Đã xóa !", "status": true}
	helpers.ResponseJSON(w, http.StatusOK, response)
}