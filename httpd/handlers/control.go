package handlers

import (
	"net/http"

	"github.com/ledgerhq/satstack/config"
	"github.com/ledgerhq/satstack/httpd/svc"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func ImportAccounts(s svc.ControlService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Accounts []config.Account `json:"accounts" binding:"required"`
		}

		if err := ctx.BindJSON(&request); err != nil {
			log.Error("Failed to bind JSON request")
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		s.ImportAccounts(request.Accounts)

		ctx.JSON(http.StatusOK, gin.H{"Status": "OK"})
	}
}

func HasDescriptor(s svc.ControlService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request struct {
			Descriptor string `json:"descriptor" binding:"required"`
		}

		if err := ctx.BindJSON(&request); err != nil {
			log.Error("Failed to bind JSON request")
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		exists, err := s.HasDescriptor(request.Descriptor)
		if err != nil {
			log.WithField("error", err).Error("Failed to handle descriptor")
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"exists": exists,
		})
	}
}
