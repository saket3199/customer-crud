package controller

import (
	"customerCrud/customer/service"
	"customerCrud/model"

	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type Controller struct {
	CustService service.Service
}

func NewController(s *service.Service) *Controller {
	return &Controller{
		CustService: *s,
	}
}
func (h *Controller) HandleRequests(r *mux.Router) {
	// creates a new instance of a mux router
	// myRouter := mux.NewRouter().StrictSlash(true)

	// replace http.HandleFunc with myRouter.HandleFunc
	r.HandleFunc("/", h.CustService.HomePage)
	r.HandleFunc("/customers", h.GetAll)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	r.HandleFunc("/customer", h.CreateNewCustomer).Methods("POST")
	r.HandleFunc("/customer/{id}", h.ReturnSingleCustomer).Methods("GET")
	r.HandleFunc("/customer/{id}", h.DeleteCustomer).Methods("DELETE")
	r.HandleFunc("/customer/{id}", h.UpdateCustomer).Methods("PUT")

}

func (h *Controller) GetAll(w http.ResponseWriter, r *http.Request) {

	cust := []model.Customer{}
	h.CustService.GetAll(&cust)
	data, err := json.Marshal(&cust)
	if err != nil {
		fmt.Fprint(w, errors.New("internal error"))
		return
	}

	fmt.Fprint(w, string(data))
}

func (h *Controller) CreateNewCustomer(w http.ResponseWriter, r *http.Request) {
	cust := model.Customer{}
	// Unmarshal json.
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &cust)
	fmt.Fprint(w, &cust)

	// (r, &cust)
	if err != nil {
		// log.NewLogger().Error(err.Error())
		fmt.Fprint(w, "Error in adding ", err)
		// web.RespondError(w, errors.NewHTTPError("unable to parse requested data", http.StatusBadRequest))
		return
	}
	h.CustService.CreateNewCustomer(&cust)
	fmt.Fprint(w, "Record Added Successfully")

}
func (h *Controller) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	cust := model.Customer{}
	// Unmarshal json.
	input := mux.Vars(r)["id"]
	if len(input) == 0 {
		fmt.Fprint(w, errors.New("empty Id"))
	}
	id, er := uuid.FromString(input)

	if er != nil {
		fmt.Fprint(w, errors.New("cant Parse"))
		return
	}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &cust)
	// (r, &cust)
	if err != nil {
		// log.NewLogger().Error(err.Error())
		// web.RespondError(w, errors.NewHTTPError("unable to parse requested data", http.StatusBadRequest))
		return
	}
	cust.ID = id
	errs := h.CustService.UpdateCustomer(&cust)
	if errs != nil {
		fmt.Fprint(w, "Errors updating values ", errs)
		return
	}
	fmt.Fprint(w, "Record Updated Successfully")

}
func (h *Controller) ReturnSingleCustomer(w http.ResponseWriter, r *http.Request) {
	cust := model.Customer{}

	input := mux.Vars(r)["id"]
	if len(input) == 0 {
		fmt.Fprint(w, errors.New("empty Id"))
		return
	}
	id, err := uuid.FromString(input)

	if err != nil {
		fmt.Fprint(w, errors.New("cant Parse"))
		return
	}

	errs := h.CustService.ReturnSingleCustomer(&cust, id)
	data, err := json.Marshal(&cust)
	if err != nil {
		fmt.Fprint(w, errors.New("internal error"))
		return
	}
	if errs != nil {
		fmt.Fprint(w, "error: ", errs)
		return
	}

	fmt.Fprint(w, string(data))

}
func (h *Controller) DeleteCustomer(w http.ResponseWriter, r *http.Request) {

	cust := model.Customer{}
	input := mux.Vars(r)["id"]
	if len(input) == 0 {
		fmt.Fprint(w, errors.New("empty Id"))
	}
	id, err := uuid.FromString(input)

	if err != nil {
		fmt.Fprint(w, errors.New("cant Parse"))
		return
	}
	cust.ID = id
	errs := h.CustService.DeleteCustomer(&cust)
	if errs != nil {
		fmt.Fprint(w, "error: ", errs)
		return
	}
	fmt.Fprint(w, "Deleted successfully ")

}
