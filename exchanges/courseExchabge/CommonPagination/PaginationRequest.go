package CommonPagination

import (
	"basicCrudoperations/response"
	"github.com/gin-gonic/gin"
)

type PaginationAndOrderRequest struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Order  string `json:"order"`
}

func ApplyPaginationAndOrderDefaults(p *PaginationAndOrderRequest) {
	if p.Limit <= 0 {
		p.Limit = 5
	}
	if p.Offset < 0 {
		p.Offset = 0
	}
	if p.Order == "" {
		p.Order = "created_At asc"
	}
}
func ShawAllPaginationWithOrder(c *gin.Context) (error, PaginationAndOrderRequest) {
	var req PaginationAndOrderRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.BadRequest(c, err.Error())
		return nil, PaginationAndOrderRequest{}
	}
	ApplyPaginationAndOrderDefaults(&req)
	return err, req
}
