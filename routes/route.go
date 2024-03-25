package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes -> define endpoints
func Routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home Page for you custom api"))
	}).Methods("GET")

	// router.HandleFunc("/person/{id}", controllers.DeletePersonEndpoint).Methods("DELETE")
	// router.HandleFunc("/person/{id}", controllers.UpdatePersonEndpoint).Methods("PUT")
	// router.HandleFunc("/upload", controllers.UploadFileEndpoint).Methods("POST")

	return router
}
