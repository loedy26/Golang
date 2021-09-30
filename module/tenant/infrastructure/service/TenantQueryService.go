package service

import (
	"context"

	"golang-api/module/tenant/domain/entity"
	"golang-api/module/tenant/domain/repository"
)

// TenantQueryService handles the tenant query service logic
type TenantQueryService struct {
	repository.TenantQueryRepositoryInterface
}

// GetTenantByID retrieves the tenant record provided by its id
func (service *TenantQueryService) GetTenantByID(ctx context.Context, tenantID string) (entity.Tenant, error) {
	res, err := service.TenantQueryRepositoryInterface.SelectTenantByID(tenantID)
	if err != nil {
		return res, err
	}

	return res, nil
}
