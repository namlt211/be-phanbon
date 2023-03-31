package order

import (
	"encoding/json"
	"green/helpers"
	"green/models"
	"net/http"
)

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	var payment *models.Payment
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payment); err != nil {
		response := map[string]string {"message": err.Error()}
		helpers.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()
	if err := models.MysqlConn.Create(&payment).Error; err != nil {
		response := map[string]string {"message": "Lỗi khi thêm thanh toán mới !"}
		helpers.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]interface{}{"message": "Thêm thanh toán thành công !", "status": true, "data": payment}
	helpers.ResponseJSON(w, http.StatusOK, response)
}