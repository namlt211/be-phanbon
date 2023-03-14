package order

import (
	"encoding/json"
	"green/helpers"
	"green/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order *models.Order
	decoder := json.NewDecoder(r.Body)
	
	if err := decoder.Decode(&order); err != nil {
		response := map[string]string {"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()
	if err := models.MysqlConn.Create(&order).Error; err != nil {
		response := map[string]string {"message": "Lỗi khi thêm đơn hàng mới !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	for i, detail := range order.OrderDetails{
		detail.OrderID = order.Id
		models.MysqlConn.Create(&order.OrderDetails[i])
	}
	response := map[string]interface{}{"message": "Thêm đơn hàng thành công !", "status": true, "data": order}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func GetAllOrder(w http.ResponseWriter, r *http.Request) {
	var order []models.Order
	if err := models.MysqlConn.Find(&order).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	response := map[string]interface{}{"message": "Lấy danh sách đơn hàng thành công !", "status": true, "data": order}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func GetOneOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    idParam := vars["id"]
	if idParam == "" {
		response := map[string]string{"message": "Không có truyền lên mã đơn hàng !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := map[string]string{"message": "Không có truyền lên mã đơn hàng !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	var order models.Order
	result := models.MysqlConn.First(&order, id)
	if result.Error != nil {
		response := map[string]string{"message": "Không tìm thấy đơn hàng !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
    }
	response := map[string]interface{}{"message": "Đã thấy !", "status": true, "data": order}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {}


func CreateOrderDetail(w http.ResponseWriter, r *http.Request) {
	var orderDetail *models.OrderDetails

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&orderDetail); err != nil {
		response := map[string]string {"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()
	if err := models.MysqlConn.Create(&orderDetail).Error; err != nil {
		response := map[string]string {"message": "Lỗi khi thêm đơn hàng mới !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]interface{}{"message": "Thêm đơn hàng thành công !", "status": true, "data": orderDetail}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func UpdateOrderDetail(w http.ResponseWriter, r *http.Request) {}


func DeleteOrderDetail(w http.ResponseWriter, r *http.Request) {}


func GetAllOrderDetail(w http.ResponseWriter, r *http.Request) {
	var orderDetail []models.OrderDetails
	if err := models.MysqlConn.Preload("Order").Find(&orderDetail).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	response := map[string]interface{}{"message": "Lấy danh sách đơn hàng thành công !", "status": true, "data": orderDetail}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func GetOneOrderDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    idParam := vars["id"]
	if idParam == "" {
		response := map[string]string{"message": "Không có truyền lên mã chi đơn hàng !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response := map[string]string{"message": "Không có truyền lên mã chi đơn hàng !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	var orderDetail models.OrderDetails
	result := models.MysqlConn.First(&orderDetail, id)
	if result.Error != nil {
		response := map[string]string{"message": "Không tìm thấy chi đơn hàng !"}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
    }
	response := map[string]interface{}{"message": "Đã thấy !", "status": true, "data": orderDetail}
	helpers.ResponseJSON(w, http.StatusOK, response)
}



