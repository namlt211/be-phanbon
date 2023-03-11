package product

import (
	"encoding/json"
	"green/helpers"
	"green/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		response := map[string]string {"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()
	if err := models.MysqlConn.Create(&product).Error; err != nil {
		response := map[string]string {"message": "Lỗi khi thêm sản phẩm mới !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]interface{}{"message": "Thêm sản phẩm thành công !", "status": true, "data": product}
	helpers.ResponseJSON(w, http.StatusOK, response)
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    idParam := vars["id"]
	if idParam == "" {
		response := map[string]string{"message": "Không có truyền lên mã sản phẩm !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := map[string]string{"message": "Không có truyền lên mã sản phẩm !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	var productUpdate models.Product

	err = json.NewDecoder(r.Body).Decode(&productUpdate)
	if err != nil {
		response := map[string]string{"message": "Dữ liệu không hợp lệ !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	product :=&models.Product{}
	result := models.MysqlConn.First(product, id)

	if result.Error != nil {
		response := map[string]string{"message": "Không tìm thấy sản phẩm hợp lệ !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	product.ProductName = productUpdate.ProductName
	product.ProductDescription = productUpdate.ProductDescription
	product.ProductPrice = productUpdate.ProductPrice
	product.SupplierID = productUpdate.SupplierID
	product.ProductImage = productUpdate.ProductImage
	product.UpdatedAt = productUpdate.UpdatedAt
	result = models.MysqlConn.Save(product)
	if result.Error != nil {
		response := map[string]string{"message": "Không thể cập nhật sản phẩm !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]interface{}{"message": "Cập nhật sản phẩm thành công !", "status": true}
	helpers.ResponseJSON(w, http.StatusOK, response)
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    idParam := vars["id"]
	if idParam == "" {
		response := map[string]string{"message": "Không có truyền lên mã sản phẩm !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := map[string]string{"message": "Không có truyền lên mã sản phẩm !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	var product models.Product
	if models.MysqlConn.Delete(&product, id).RowsAffected == 0 {
		response := map[string]string{"message": "Lỗi khi xóa !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]interface{}{"message": "Đã xóa !", "status": true}
	helpers.ResponseJSON(w, http.StatusOK, response)
}
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var product []models.Product
	if err := models.MysqlConn.Preload("Suppliers").Find(&product).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	response := map[string]interface{}{"message": "Lấy sản phẩm thành công !", "status": true, "data": product}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func GetProductById(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
    idParam := vars["id"]
	if idParam == "" {
		response := map[string]string{"message": "Không có truyền lên mã sản phẩm !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := map[string]string{"message": "Không có truyền lên mã sản phẩm !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	var product models.Product
	result := models.MysqlConn.Preload("Suppliers").First(&product, id)
	if result.Error != nil {
		response := map[string]string{"message": "Không tìm thấy sản phẩm !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
    }
	response := map[string]interface{}{"message": "Đã thấy !", "status": true, "data": product}
	helpers.ResponseJSON(w, http.StatusOK, response)
}