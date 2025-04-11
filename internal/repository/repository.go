package repository

import (
	"sync"
	"go-run-reports/internal/model"
)

var store = map[string]*model.OrgPlan{}
var mu sync.Mutex

type Repository interface {
	GetOrgPlan(orgID string) (*model.OrgPlan, bool)
	SaveOrgPlan(org *model.OrgPlan)
}

type InMemoryRepo struct{}

func NewRepository() Repository {
	return &InMemoryRepo{}
}

func (r *InMemoryRepo) GetOrgPlan(orgID string) (*model.OrgPlan, bool) {
	mu.Lock()
	defer mu.Unlock()
	org, exists := store[orgID]
	return org, exists
}

func (r *InMemoryRepo) SaveOrgPlan(org *model.OrgPlan) {
	mu.Lock()
	defer mu.Unlock()
	store[org.OrgID] = org
}