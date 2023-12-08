package vendors

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dbModel "github.com/synctv-org/synctv/internal/model"
	"github.com/synctv-org/synctv/internal/vendor"
	"github.com/synctv-org/synctv/server/model"
	"golang.org/x/exp/maps"
)

func Backends(ctx *gin.Context) {
	var backends []string
	switch dbModel.StreamingVendor(ctx.Param("vendor")) {
	case dbModel.StreamingVendorBilibili:
		backends = maps.Keys(vendor.BilibiliClients())
	case dbModel.StreamingVendorAlist:
		backends = maps.Keys(vendor.AlistClients())
	default:
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.NewApiErrorStringResp("invalid vendor"))
		return
	}
	ctx.JSON(http.StatusOK, model.NewApiDataResp(backends))
}