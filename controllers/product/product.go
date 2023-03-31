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
		response := map[string]string{"message": "Dữ liệu không hợp lệ 1 !"}
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
	product.Unit = productUpdate.Unit
	product.Discount = productUpdate.Discount
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
	models.MysqlConn.Table("products").Where("id = ?", id).Update("is_delete", 0)
	response := map[string]interface{}{"message": "Đã xóa !", "status": true}
	helpers.ResponseJSON(w, http.StatusOK, response)
}
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []struct{
		Id uint `json:"id"`
		Name string `json:"name" gorm:"column:product_name;"`
		Image string `json:"image" gorm:"column:product_image;"`
		Price float64 `json:"price" gorm:"not null;type:float(10,2);column:product_price;"`
		Unit string `json:"unit"`
		Discount float64 `json:"discount" gorm:"type:float(10,2);column:discount;"`
		Supplier string `json:"supplier"`
	}
	models.MysqlConn.Table("products").
	Select(`products.id, products.product_name, products.product_image, products.product_price, products.unit,products.discount, suppliers.name AS supplier`).
	Where(`products.status = 1`).
	Joins(`JOIN suppliers ON products.supplier_id = suppliers.id`).
	Scan(&products)
	response := map[string]interface{}{"message": "Lấy sản phẩm thành công !", "status": true, "data": products}
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
	var product struct{
		Id uint `json:"id"`
		Name string `json:"name" gorm:"column:product_name;"`
		Image string `json:"image" gorm:"column:product_image;"`
		Price float64 `json:"price" gorm:"type:float(10,2);column:product_price;"`
		Unit string `json:"unit"`
		Description string `json:"description" gorm:"column:product_description;"`
		Supplier string `json:"supplier"`
		SupplierID uint `json:"supplier_id"`
	}
	models.MysqlConn.Table("products").
	Select(`products.id, products.product_name, products.product_image, products.product_price, products.unit, products.product_description, products.supplier_id, suppliers.name AS supplier`).
	Where(`products.status = 1 && products.id = ` + strconv.Itoa(int(id))).
	Joins(`JOIN suppliers ON products.supplier_id = suppliers.id`).
	Scan(&product)
	if product.Id == 0 {
		response := map[string]interface{}{"message": "Không có sản phẩm !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	response := map[string]interface{}{"message": "Lấy sản phẩm thành công !", "status": true, "data": product}
	helpers.ResponseJSON(w, http.StatusOK, response)
}