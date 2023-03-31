package order

import (
	"encoding/json"
	"green/helpers"
	"green/models"
	"net/http"
	"strconv"
	"time"

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
	response := map[string]interface{}{"message": "Thêm đơn hàng thành công !", "status": true, "data": order}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func GetAllOrder(w http.ResponseWriter, r *http.Request) {
	var order []struct{
		Id int64              `json:"id"`
		CustomerID int64      `json:"customer_id" gorm:"column:customer_id;"`
		CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;"`
    	TotalAmount float64 `json:"total_amount" gorm:"column:total_amount;"`
		CustomerName string `json:"customer_name"`
		Paid float64 `json:"paid" gorm:"column:paid;"`
		
	}
	models.MysqlConn.Table("orders").
	Select(`orders.id, orders.customer_id, orders.created_at, orders.total_amount, customers.name AS customer_name, SUM(payments.money) AS paid`).
	Joins(`JOIN customers ON orders.customer_id = customers.id`).
	Joins(`JOIN payments ON orders.id = payments.order_id`).
	Where(`orders.is_delete = 1`).
	Group(`orders.id`).
	Scan(&order)
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
	var order []struct{
		Id int64              `json:"id"`
		CustomerID int64      `json:"customer_id" gorm:"column:customer_id;"`
    	TotalAmount float64 `json:"total_amount" gorm:"column:total_amount;"`
    	OrderDate *time.Time `json:"order_date" gorm:"column:order_date;"`
		CustomerName string `json:"customer_name"`
	}
	models.MysqlConn.Table("orders").
	Select(`orders.id, orders.customer_id, orders.total_amount, customers.name AS customer_name`).
	Joins(`JOIN customers ON orders.customer_id = customers.id`).
	Where(`orders.is_delete = 1 && orders.id = ?`, id).
	Scan(&order)
	response := map[string]interface{}{"message": "Lấy danh sách đơn hàng thành công !", "status": true, "data": order}
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
		response := map[string]string {"message": "Lỗi khi thêm chi tiết đơn hàng mới !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]interface{}{"message": "Thêm chi tiết đơn hàng thành công !", "status": true, "data": orderDetail}
	helpers.ResponseJSON(w, http.StatusOK, response)
}

func UpdateOrderDetail(w http.ResponseWriter, r *http.Request) {}


func DeleteOrderDetail(w http.ResponseWriter, r *http.Request) {}


func GetAllOrderDetail(w http.ResponseWriter, r *http.Request) {
	var orderDetails []struct {
		Id uint `json:"id"`
		OrderId uint `json:"order_id"`
		ProductId uint `json:"product_id"`
		Quantity uint `json:"quantity"`
		ProductName string `json:"product_name"`
		OrderDetailDate time.Time `json:"created_at" gorm:"type:datetime;column:created_at;"`
		Unit string `json:"unit"`
		Discount           float64         `json:"discount" gorm:"type:float(10,2);default:0;column:discount;"`
		Price float64 `json:"price" gorm:"not null;type:float(10,2);column:product_price;"`
		Total     float64    `json:"total" gorm:"type:decimal(10,2);column:total;"`
		CustomerName string `json:"customer_name"`
	}
	models.MysqlConn.Table("order_details").
	Select(`order_details.id, order_details.order_id, order_details.product_id,
	order_details.quantity, products.product_name, order_details.created_at, products.unit, products.discount,
	products.product_price, orders.total_amount AS total,
	customers.name AS customer_name`).
	Joins("JOIN products ON  products.id = order_details.product_id").
	Joins("JOIN orders ON  orders.id = order_details.order_id").
	Joins("JOIN customers ON customers.id = orders.customer_id").
	Where(`order_details.is_delete = 1`).
	Scan(&orderDetails)
	response := map[string]interface{}{"message": "Lấy danh sách đơn hàng thành công !", "status": true, "data": orderDetails}
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
	var orderDetail struct {
		Id uint `json:"id"`
		OrderId uint `json:"order_id"`
		ProductId uint `json:"product_id"`
		Quantity uint `json:"quantity"`
		ProductName string `json:"product_name"`
		OrderDetailDate time.Time `json:"created_at" gorm:"type:datetime;column:created_at;"`
		Unit string `json:"unit"`
		Discount           float64         `json:"discount" gorm:"type:float(10,2);default:0;column:discount;"`
		Price float64 `json:"price" gorm:"not null;type:float(10,2);column:product_price;"`
		Total     float64    `json:"total" gorm:"type:decimal(10,2);column:total;"`
		CustomerName string `json:"customer_name"`
		
	}
	models.MysqlConn.Table("order_details").
	Select(`order_details.id, order_details.order_id, order_details.product_id,
	order_details.quantity, products.product_name, order_details.created_at, products.unit, products.discount,
	products.product_price, order_details.total,
	customers.name AS customer_name`).
	Joins("JOIN products ON  products.id = order_details.product_id").
	Joins("JOIN orders ON  orders.id = order_details.order_id").
	Joins("JOIN customers ON customers.id = orders.customer_id").
	Where(`order_details.is_delete = 1 && order_details.id = ?`, id).
	Scan(&orderDetail)
	response := map[string]interface{}{"message": "Lấy chi tiết đơn hàng thành công !", "status": true, "data": orderDetail}
	helpers.ResponseJSON(w, http.StatusOK, response)
}



