package handler

import (
	"log"
	"net/http"

	commonHTTP "github.com/Ralphbaer/hubla/backend/common/net/http"
	uc "github.com/Ralphbaer/hubla/backend/transaction/usecase"
	"github.com/gorilla/mux"
)

// SellerHandler represents a handler which deal with Seller resource operations
type SellerHandler struct {
	UseCase *uc.SellerUseCase
}

func (handler *SellerHandler) GetSellerBalanceByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sellerID := mux.Vars(r)["id"]
		view, err := handler.UseCase.GetSellerBalanceByID(r.Context(), sellerID)
		if err != nil {
			log.Printf("ERRO %v", err)
			commonHTTP.WithError(w, err)
			return
		}

		commonHTTP.OK(w, view)
	})
}
