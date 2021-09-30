package repository

import (
	"golang-api/module/tenant/domain/entity"
)

type TenantQueryRepositoryInterface interface {
	SelectTenantByID(tenantID string) (entity.Tenant, error)
}
