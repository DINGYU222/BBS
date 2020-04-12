package block

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// /block/add/view
func BlockAddView(c *gin.Context) {
	c.HTML(http.StatusOK, "admin-block-add-view.tmpl", "")
}
