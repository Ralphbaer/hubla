package handler

import (
	"net/http"

	"github.com/Ralphbaer/hubla/backend/common/hlog"
	commonHTTP "github.com/Ralphbaer/hubla/backend/common/net/http"
	uc "github.com/Ralphbaer/hubla/backend/transaction/usecase"
	"github.com/gorilla/mux"
)

// SellerHandler represents a handler which deal with Seller resource operations
type SellerHandler struct {
	UseCase *uc.SellerUseCase
}

// GetSellerBalanceByID returns an http.Handler that retrieves the seller balance by its ID.
func (handler *SellerHandler) GetSellerBalanceByID() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		logger := hlog.NewLoggerFromContext(ctx)
		logger.Debug("Get seller balance by ID")

		sellerID := mux.Vars(r)["id"]
		view, err := handler.UseCase.GetSellerBalanceByID(r.Context(), sellerID)
		if err != nil {
			logger.Error(err.Error())
			commonHTTP.WithError(w, err)
			return
		}

		commonHTTP.OK(w, view)
	})
}
