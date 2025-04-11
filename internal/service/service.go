package service

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"go-run-reports/internal/model"
	"go-run-reports/internal/notifier"
	"go-run-reports/internal/repository"
	"go-run-reports/pkg/util"
)

type UsageService struct {
	repo repository.Repository
}

func NewUsageService() *UsageService {
	return &UsageService{repo: repository.NewRepository()}
}

func (s *UsageService) SetPlanHandler(c *gin.Context) {
	type req struct {
		OrgID string       `json:"org_id"`
		Plan  model.PlanType `json:"plan"`
	}
	var r req
	if err := c.BindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	plan := &model.OrgPlan{
		OrgID:     r.OrgID,
		Plan:      r.Plan,
		Month:     util.CurrentMonth(),
		UserUsage: map[string]int{},
	}
	s.repo.SaveOrgPlan(plan)
	c.JSON(http.StatusOK, gin.H{"status": "plan set"})
}

func (s *UsageService) SetUserLimitHandler(c *gin.Context) {
	type req struct {
		OrgID string `json:"org_id"`
		Limit int    `json:"limit"`
	}
	var r req
	if err := c.BindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	org, found := s.repo.GetOrgPlan(r.OrgID)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "org not found"})
		return
	}
	org.PerUser = r.Limit
	s.repo.SaveOrgPlan(org)
	c.JSON(http.StatusOK, gin.H{"status": "user limit set"})
}

func (s *UsageService) UseReportCreditHandler(c *gin.Context) {
	type req struct {
		OrgID    string `json:"org_id"`
		UserID   string `json:"user_id"`
		ReportID string `json:"report_id"`
	}
	var r req
	if err := c.BindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	org, found := s.repo.GetOrgPlan(r.OrgID)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "org not found"})
		return
	}

	if org.Month != util.CurrentMonth() {
		org.TotalUsed = 0
		org.UserUsage = map[string]int{}
		org.Month = util.CurrentMonth()
		org.TopUps = 0
	}

	totalLimit := util.PlanToLimit(org.Plan) + org.TopUps
	if org.TotalUsed >= totalLimit {
		notifier.SendLimitReached(r.OrgID)
		c.JSON(http.StatusForbidden, gin.H{"error": "limit reached"})
		return
	}

	if org.UserUsage[r.UserID] >= org.PerUser {
		c.JSON(http.StatusForbidden, gin.H{"error": "user limit reached"})
		return
	}

	org.TotalUsed++
	org.UserUsage[r.UserID]++
	s.repo.SaveOrgPlan(org)
	c.JSON(http.StatusOK, gin.H{"status": "report counted"})
}

func (s *UsageService) TopUpCreditsHandler(c *gin.Context) {
	type req struct {
		OrgID string `json:"org_id"`
		Count int    `json:"count"`
	}
	var r req
	if err := c.BindJSON(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	org, found := s.repo.GetOrgPlan(r.OrgID)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "org not found"})
		return
	}
	org.TopUps += r.Count
	s.repo.SaveOrgPlan(org)
	c.JSON(http.StatusOK, gin.H{"status": "credits topped up"})
}