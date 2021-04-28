package pkg

import (
	"encoding/json"
	"net/http"
)

type GenericResponse struct {
	Success bool
	Message string
}

// HandlerHealthCheck function handles all the request calls coming
// on route /health-check and response with a static json message
func HandlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	resp := GenericResponse{
		Success: true,
		Message: "Go server is working !!",
	}
	jsResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsResp)
}

// HandleGetInvalidDeliveries function handles all the request calls
// for route /invalid-deliveries, this call makes a SQL query to
// fetch all the invalid deliveries and returns a json response
func HandleGetInvalidDeliveries(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("resource not found"))
		return
	}
	query := `SELECT * FROM delivery WHERE delivery.driver_id NOT IN (
    			SELECT driver.id FROM delivery
    			INNER JOIN supplier_bean_type ON delivery.supplier_id=supplier_bean_type.supplier_id
    			INNER JOIN carrier_bean_type ON supplier_bean_type.bean_type_id=carrier_bean_type.bean_type_id
    			INNER JOIN driver ON driver.carrier_id=carrier_bean_type.id
			);`
	res, err := db.Select(query)
	if err != nil {
		resp := GenericResponse{
			Success: false,
			Message: err.Error(),
		}
		byts, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(byts)
		return
	}
	resp, _ := json.Marshal(res)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}