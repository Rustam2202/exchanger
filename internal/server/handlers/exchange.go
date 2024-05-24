package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/llcmediatel/recruiting/golang-junior-dev/internal/service"
)

type ExchangeRequest struct {
	Amount    int   `json:"amount"`
	Banknotes []int `json:"banknotes"`
}

type ExchangeResponse struct {
	Exchanges [][]int `json:"exchanges"`
}

// @Summary	Calculate event data by Id
// @Description
// @Tags		Calculate
// @Accept		json
// @Produce	json
// @Param		request	body		ExchangeRequest	true	"Exchange Request"
// @Success	200		{object}	ExchangeResponse
// @Failure	400		{object}	handlers.ErrorResponce
// @Failure	500		{object}	handlers.ErrorResponce
// @Router		/exchange [post]
func ExchangePost(c *gin.Context) {
	var req ExchangeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exchanges, err := service.FindCombinations(req.Amount, req.Banknotes)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if exchanges==nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := ExchangeResponse{Exchanges: exchanges}
	c.JSON(http.StatusOK, resp)
}
