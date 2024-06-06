package cmd

import (
	"github.com/Tharolloo/go-playground/vin"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, service vin.Kmainstore) {
	// Register routes
	r.POST("/authenticate", service.GetAuthenticationToken())

	r.POST("/dynamic", service.CreateDynamicVirtualAccount())
	r.POST("/static", service.CreateStaticVirtualAccount())

	r.POST("/deactivate", service.DeactivateVirtualAccount())
	r.POST("/reactivate", service.ReactivateVirtualAccount())

	r.POST("/reallocate", service.ReallocateVirtualAccount())
	r.POST("/confirm", service.ConfirmVirtualPayment())

}
