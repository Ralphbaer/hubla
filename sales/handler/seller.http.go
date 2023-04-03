package handler

import (
	"net/http"

	commonHTTP "github.com/Ralphbaer/hubla/common/net/http"
	uc "github.com/Ralphbaer/hubla/sales/usecase"
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
		if view == nil {
			commonHTTP.WithError(w, err)
			return
		}
		if err != nil {
			commonHTTP.WithError(w, err)
			return
		}

		commonHTTP.OK(w, view)
	})
}
