package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	application "github.com/yunusemre12500/poll-api/internal/application/poll/v1"
	v1 "github.com/yunusemre12500/poll-api/internal/application/poll/v1"
	domain "github.com/yunusemre12500/poll-api/internal/domain/poll/v1"
)

type HTTPPollController struct {
	repository domain.PollRepository
}

func NewHTTPPollController(repository domain.PollRepository) *HTTPPollController {
	return &HTTPPollController{
		repository: repository,
	}
}

func (controller *HTTPPollController) AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/v1/polls", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controller.List(w, r)
		case http.MethodPost:
			controller.Create(w, r)
		default:
			http.Error(w, fmt.Sprintf("Request method (%s) not supported for this route.", r.Method), http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/v1/polls/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			controller.GetByID(w, r)
		default:
			http.Error(w, fmt.Sprintf("Request method (%s) not supported for this route.", r.Method), http.StatusMethodNotAllowed)
		}
	})
}

func (controller *HTTPPollController) Create(w http.ResponseWriter, r *http.Request) {
	acceptHeaderValue := r.Header.Get("Accept")

	if acceptHeaderValue == "" {
		http.Error(w, "Accept header is missing.", http.StatusBadRequest)

		return
	} else if acceptHeaderValue != "application/json" {
		http.Error(w, fmt.Sprintf("Requested response type (%s) not supported.", acceptHeaderValue), http.StatusBadRequest)

		return
	}

	contentTypeHeaderValue := r.Header.Get("Content-Type")

	if contentTypeHeaderValue == "" {
		http.Error(w, "Content-Type header is missing.", http.StatusBadRequest)

		return
	} else if contentTypeHeaderValue != "application/json" {
		http.Error(w, fmt.Sprintf("Requested Content-Type (%s) not supported.", contentTypeHeaderValue), http.StatusBadRequest)

		return
	}

	var createPollRequestBody *application.CreatePollRequestBody

	switch contentTypeHeaderValue {
	case "application/json":
		if err := json.NewDecoder(r.Body).Decode(&createPollRequestBody); err != nil {
			http.Error(w, "Failed to decode request body.", http.StatusBadRequest)

			return
		}
	}

	newPoll := domain.NewPollFromCreatePollRequestBody(createPollRequestBody)

	err := controller.repository.Create(newPoll)

	if err != nil {
		http.Error(w, "Failed to save created poll. Please try again later.", http.StatusInternalServerError)

		return
	}

	switch acceptHeaderValue {
	case "application/json":
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(newPoll.IntoCreatePollResponseBody()); err != nil {
			http.Error(w, "Failed to encode response body.", http.StatusInternalServerError)

			return
		}
	}
}

func (controller *HTTPPollController) GetByID(w http.ResponseWriter, r *http.Request) {
	acceptHeaderValue := r.Header.Get("Accept")

	if acceptHeaderValue == "" {
		http.Error(w, "Accept header is missing.", http.StatusBadRequest)

		return
	} else if acceptHeaderValue != "application/json" {
		http.Error(w, fmt.Sprintf("Requested response type (%s) not supported.", acceptHeaderValue), http.StatusBadRequest)

		return
	}

	id, err := uuid.Parse(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Failed to parse 'id' path parameter.", http.StatusBadRequest)

		return
	}

	poll, err := controller.repository.GetByID(id)

	if err != nil {
		if err == domain.ErrNotFound {
			http.Error(w, "Poll not found.", http.StatusNotFound)

			return
		}

		http.Error(w, "Failed to get poll.", http.StatusNotFound)

		return
	}

	switch acceptHeaderValue {
	case "application/json":
		if err = json.NewEncoder(w).Encode(poll.IntoGetPollByIdResponseBody()); err != nil {
			http.Error(w, "Failed to encode response body.", http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
	}
}

func (controller *HTTPPollController) List(w http.ResponseWriter, r *http.Request) {
	acceptHeaderValue := r.Header.Get("Accept")

	if acceptHeaderValue == "" {
		http.Error(w, "Accept header is missing.", http.StatusBadRequest)

		return
	} else if acceptHeaderValue != "application/json" {
		http.Error(w, fmt.Sprintf("Requested response type (%s) not supported.", acceptHeaderValue), http.StatusBadRequest)

		return
	}

	limit, err := strconv.ParseUint(r.URL.Query().Get("limit"), 10, 8)

	if err != nil {
		http.Error(w, "Failed to parse 'limit' query parameter.", http.StatusBadRequest)

		return
	}

	if limit < 2 {
		http.Error(w, "Query parameter value of 'limit' lower than 2.", http.StatusBadRequest)

		return
	} else if limit > 100 {
		http.Error(w, "Query parameter value of 'limit' greater than 100.", http.StatusBadRequest)

		return
	}

	offset, err := strconv.ParseInt(r.URL.Query().Get("offset"), 10, 64)

	if err != nil {
		http.Error(w, "Failed to parse 'offset' query parameter.", http.StatusBadRequest)

		return
	}

	if offset < 0 {
		http.Error(w, "Query parameter value of 'offset' must be positive integer.", http.StatusBadRequest)

		return
	}

	polls, err := controller.repository.List(uint(limit), uint(offset))

	if err != nil {
		http.Error(w, "Failed to list polls.", http.StatusInternalServerError)

		return
	}

	if polls == nil {
		http.Error(w, "No polls found.", http.StatusNoContent)

		return
	}

	switch acceptHeaderValue {
	case "application/json":
		var listedPolls []*v1.ListPollsResponseBody

		for _, poll := range polls {
			listedPolls = append(listedPolls, poll.IntoListPollsResponseBody())
		}

		if err = json.NewEncoder(w).Encode(&listedPolls); err != nil {
			http.Error(w, "Failed to encode polls.", http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
	}
}
