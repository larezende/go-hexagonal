package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"go-hexagonal/adapters/dto"
	"go-hexagonal/application"
	"net/http"
)

const CONTENT_TYPE = "Content-Type"
const APPLICATION_JSON = "application/json"

func MakeProductWebHandler(r *mux.Router, n *negroni.Negroni, service application.ProductServiceInterface) {
	r.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	r.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product", n.With(
		negroni.Wrap(getAllProducts(service)),
	)).Methods("GET", "OPTIONS")

	r.Handle("/product/{id}/enable", n.With(
		negroni.Wrap(enableProduct(service)),
	)).Methods("PUT", "OPTIONS")

	r.Handle("/product/{id}/disable", n.With(
		negroni.Wrap(disableProduct(service)),
	)).Methods("PUT", "OPTIONS")
}

func createProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		var productDTO dto.ProductDTO
		err := json.NewDecoder(r.Body).Decode(&productDTO)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(jsonError(err.Error()))
			return
		}
		product, err := service.Create(productDTO.Name, productDTO.Price)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

func getProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

func getAllProducts(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		products, err := service.List()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(products)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

func enableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonError(err.Error()))
			return
		}
		result, err := service.Enable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}

func disableProduct(service application.ProductServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(CONTENT_TYPE, APPLICATION_JSON)
		vars := mux.Vars(r)
		id := vars["id"]
		product, err := service.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(jsonError(err.Error()))
			return
		}
		result, err := service.Disable(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonError(err.Error()))
			return
		}
	})
}
