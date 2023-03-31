package routers

import (
	"green/controllers/auth"
	"green/controllers/customer"
	"green/controllers/order"
	"green/controllers/product"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(engine *gin.Engine) {
	post := ":8080"
	r := mux.NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	credentialsTrue := handlers.AllowCredentials()

	//user router start
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/register", auth.Register).Methods("POST")
	r.HandleFunc("/logout", auth.Logout).Methods("GET")

	//user router end

	//product router start
	r.HandleFunc("/product/add", product.AddProduct).Methods("POST")
	r.HandleFunc("/product/getall", product.GetAllProducts).Methods("GET")
	r.HandleFunc("/product/getproductbyid/{id}",product.GetProductById).Methods("GET")
	r.HandleFunc("/product/update/{id}", product.UpdateProduct).Methods("PUT")
	r.HandleFunc("/product/delete/{id}", product.DeleteProduct).Methods("GET")

	//product router end

	//supplier router start
	r.HandleFunc("/supplier/add", product.AddSupplier).Methods("POST")
	r.HandleFunc("/supplier/getall", product.GetAllSupplier).Methods("GET")
	r.HandleFunc("/supplier/update/{id}",product.UpdateSupplier).Methods("PUT")
	r.HandleFunc("/supplier/getsupplierbyid/{id}", product.GetSupplierById).Methods("GET")
	r.HandleFunc("/supplier/delete/{id}", product.DeleteSupplier).Methods("GET")


	//supplier router end

	//customer router start
	r.HandleFunc("/customer/add", customer.AddCustomer).Methods("POST")
	r.HandleFunc("/customer/getall", customer.GetAllCustomer).Methods("GET")
	r.HandleFunc("/customer/getbyid/{id}", customer.GetCustomerById).Methods("GET")
	r.HandleFunc("/customer/update/{id}", customer.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customer/delete/{id}", customer.DeleteCustomer).Methods("GET")
	//customer router end

	//order router start
	r.HandleFunc("/order/getall", order.GetAllOrder).Methods("GET")
	r.HandleFunc("/order/add", order.CreateOrder).Methods("POST")
	r.HandleFunc("/order/getone", order.GetOneOrder).Methods("GET")
	//order router end

	//order details router start
	r.HandleFunc("/order-details/getall", order.GetAllOrderDetail).Methods("GET")
	r.HandleFunc("/order-details/add", order.CreateOrderDetail).Methods("POST")
	//order details router end



	//payment router start
	r.HandleFunc("/payment/add", order.CreatePayment).Methods("POST")
	//payment router end




	log.Fatal(http.ListenAndServe(post,handlers.CORS(originsOk, headersOk, methodsOk, credentialsTrue)(r)))

}