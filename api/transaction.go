// TODO: there should be 2 separate svcs - user and transaction

package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/web3dev6/token_transaction/db/sqlc"
	"github.com/web3dev6/token_transaction/token"
)

type createTransactionRequest struct {
	// Username string                `json:"username" binding:"required,alphanum"` // will come from middleware
	Context string          `json:"context" binding:"required,context"` // TODO: could be an enum of supported operations
	Payload json.RawMessage `json:"payload" binding:"required"`         // TODO: need to put validations on payload - defined structs only
}

type transactionResponse struct {
	Username    string          `json:"username"`
	Context     string          `json:"context"`
	Payload     json.RawMessage `json:"payload"`
	IsConfirmed bool            `json:"is_confirmed"`
	CreatedAt   time.Time       `json:"created_at"`
}

func newTransactionResponse(transaction db.Transaction) transactionResponse {
	return transactionResponse{
		Username:    transaction.Username,
		Context:     transaction.Context,
		Payload:     transaction.Payload,
		IsConfirmed: transaction.IsConfirmed,
		CreatedAt:   transaction.CreatedAt,
	}
}

func newTransactionResponses(transactions []db.Transaction) []transactionResponse {
	var responses []transactionResponse
	for _, transaction := range transactions {
		response := newTransactionResponse(transaction)
		responses = append(responses, response)
	}
	return responses
}

func (server *Server) createTransaction(ctx *gin.Context) {
	var req createTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// fmt.Println("req.Payload: ", string(req.Payload))

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateTransactionParams{
		Username: authPayload.Username,
		Context:  req.Context,
		Payload:  req.Payload,
	}

	transaction, err := server.store.CreateTransaction(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := newTransactionResponse(transaction)
	ctx.JSON(http.StatusOK, resp)
}

type getTransactionRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getTransactionDetails(ctx *gin.Context) {
	var req getTransactionRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	transaction, err := server.store.GetTransaction(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if transaction.Username != authPayload.Username {
		ctx.JSON(http.StatusUnauthorized, errorResponse(ErrFetchingUnauthorizedTransaction))
		return
	}

	resp := newTransactionResponse(transaction)
	ctx.JSON(http.StatusOK, resp)
}

type listTransactionsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listTransactions(ctx *gin.Context) {
	var req listTransactionsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListTransactionsParams{
		Username: authPayload.Username,
		Limit:    req.PageSize,
		Offset:   (req.PageID - 1) * req.PageSize,
	}
	transactions, err := server.store.ListTransactions(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := newTransactionResponses(transactions)
	ctx.JSON(http.StatusOK, resp)
}
