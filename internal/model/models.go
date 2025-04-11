package model

type PlanType string

const (
	PlanUltimate  PlanType = "Ultimate"
	PlanEnterprise PlanType = "Enterprise"
	PlanBasic     PlanType = "Basic"
	PlanLite      PlanType = "Lite"
	PlanTrial     PlanType = "Trial"
)

type OrgPlan struct {
	OrgID     string
	Plan      PlanType
	PerUser   int
	Month     string
	TotalUsed int
	UserUsage map[string]int
	TopUps    int
}