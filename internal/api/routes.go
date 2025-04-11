package api

import (
	"github.com/gin-gonic/gin"
	"go-run-reports/internal/service"
)

func RegisterRoutes(r *gin.Engine) {
	svc := service.NewUsageService()

	r.POST("/plan", svc.SetPlanHandler)
	r.POST("/limit", svc.SetUserLimitHandler)
	r.POST("/use", svc.UseReportCreditHandler)
	r.POST("/top-up", svc.TopUpCreditsHandler)
}