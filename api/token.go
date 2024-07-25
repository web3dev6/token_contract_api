package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/web3dev6/token_transaction/util"
)

type getTokenDetailsRequest struct {
	TokenAddress string `uri:"tokenAddress" binding:"required,address"`
}
type TokenDetailsResponse struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	TotalSupply string `json:"totalSupply"`
	TokenOwner  string `json:"tokenOwner"`
}

func (server *Server) getTokenDetails(ctx *gin.Context) {
	var req getTokenDetailsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Construct the URL with the route parameter
	url := fmt.Sprintf("%s/token/%s", server.config.TokenSvcBaseUrl, req.TokenAddress)
	// Set up the HTTP request options
	options := util.HTTPRequestOptions{
		Method: "GET",
		URL:    url,
		Headers: map[string]string{
			"Authorization": ctx.GetHeader(authorizationHeaderKey),
		},
	}
	// Make the HTTP request
	responseBody, err := util.HttpRequest(options)
	if err != nil {
		log.Fatalf("Error making request to %s: %v", url, err)
	}
	fmt.Printf("Response body from %s: %s\n", url, responseBody)

	// Unmarshal the response body into TokenDetailsResponse
	var tokenDetails TokenDetailsResponse
	if err := json.Unmarshal(responseBody, &tokenDetails); err != nil {
		log.Printf("Error unmarshaling token details response body: %v", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// Send the response back to the client
	ctx.JSON(http.StatusOK, tokenDetails)
}

type getTokenBalanceRequest struct {
	TokenAddress  string `uri:"tokenAddress" binding:"required,address"`
	WalletAddress string `uri:"walletAddress" binding:"required,address"`
}
type TokenBalanceResponse struct {
	Balance string `json:"balance"`
}

func (server *Server) getTokenBalance(ctx *gin.Context) {
	var req getTokenBalanceRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Construct the URL with the route parameters
	url := fmt.Sprintf("%s/token/%s/balance/%s", server.config.TokenSvcBaseUrl, req.TokenAddress, req.WalletAddress)
	// Set up the HTTP request options
	options := util.HTTPRequestOptions{
		Method: "GET",
		URL:    url,
		Headers: map[string]string{
			"Authorization": ctx.GetHeader(authorizationHeaderKey),
		},
	}
	// Make the HTTP request
	responseBody, err := util.HttpRequest(options)
	if err != nil {
		log.Fatalf("Error making request to %s: %v", url, err)
	}
	fmt.Printf("Response body from %s: %s\n", url, responseBody)

	// Unmarshal the response body into TokenBalanceResponse
	var tokenBalance TokenBalanceResponse
	if err := json.Unmarshal(responseBody, &tokenBalance); err != nil {
		log.Printf("Error unmarshaling token balance response body: %v", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// Send the response back to the client
	ctx.JSON(http.StatusOK, tokenBalance)
}

type Token struct {
	Username  string `json:"username"`
	Address   string `json:"address"`
	Name      string `json:"name"`
	Symbol    string `json:"symbol"`
	Amount    string `json:"amount"`
	Owner     string `json:"owner"`
	Authority string `json:"authority"`
}

func (server *Server) listTokens(ctx *gin.Context) {
	// Construct the URL
	url := fmt.Sprintf("%s/token", server.config.TokenSvcBaseUrl)
	// Set up the HTTP request options
	options := util.HTTPRequestOptions{
		Method: "GET",
		URL:    url,
		Headers: map[string]string{
			"Authorization": ctx.GetHeader(authorizationHeaderKey),
		},
	}
	// Make the HTTP request
	responseBody, err := util.HttpRequest(options)
	if err != nil {
		log.Fatalf("Error making request to %s: %v", url, err)
	}
	fmt.Printf("Response body from %s: %s\n", url, responseBody)

	// Unmarshal the response body into a slice of Token
	var tokens []Token
	if err := json.Unmarshal(responseBody, &tokens); err != nil {
		log.Printf("Error unmarshaling tokens response body: %v", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Send the response back to the client
	ctx.JSON(http.StatusOK, tokens)
}
