package routers

import (
	"green/controllers/auth"
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

	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/register", auth.Register).Methods("POST")
	r.HandleFunc("/logout", auth.Logout).Methods("GET")


	r.HandleFunc("/product/add", product.AddProduct).Methods("POST")
	r.HandleFunc("/product/getall", product.GetAllProducts).Methods("GET")
	r.HandleFunc("/product/getproductbyid/{id}",product.GetProductById).Methods("GET")
	r.HandleFunc("/product/update/{id}", product.UpdateProduct).Methods("PUT")
	r.HandleFunc("/product/delete/{id}", product.DeleteProduct).Methods("GET")


	r.HandleFunc("/supplier/add", product.AddSupplier).Methods("POST")
	r.HandleFunc("/supplier/getall", product.GetAllSupplier).Methods("GET")
	r.HandleFunc("/supplier/update/{id}",product.UpdateSupplier).Methods("PUT")
	r.HandleFunc("/supplier/getsupplierbyid/{id}", product.GetSupplierById).Methods("GET")
	r.HandleFunc("/supplier/delete/{id}", product.DeleteSupplier).Methods("GET")

	log.Fatal(http.ListenAndServe(post,handlers.CORS(originsOk, headersOk, methodsOk, credentialsTrue)(r)))

}