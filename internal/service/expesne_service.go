package service

import (
	"encoding/json"
	"fms/internal/request"
	"fmt"
	"io"
	"net/http"
)

func (s *Service) AddExpense(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("AddExpense called")

	// debug headers and length
	fmt.Println("Content-Length:", r.ContentLength)
	fmt.Println("Content-Type:", r.Header.Get("Content-Type"))

	// read raw body (so we can log / check emptiness / get better error messages)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("Raw request body:", string(body))

	if len(body) == 0 {
		http.Error(w, "invalid request body: empty", http.StatusBadRequest)
		return
	}

	// decode from bytes
	var expReq request.ExpenseCreateRequest
	if err := json.Unmarshal(body, &expReq); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	expenseID, err := s.ExpenseHandler.AddExpense(ctx, &expReq)
	if err != nil {
		http.Error(w, "failed to add expense: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. Return expense ID as JSON
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(expenseID); err != nil {
		http.Error(w, "failed to encode expense ID", http.StatusInternalServerError)
		return
	}
}
