package application

import (
	"context"

	"golang-api/module/tenant/domain/entity"
)

// TenantQueryServiceInterface holds the implementable methods forthe tenant query service
type TenantQueryServiceInterface interface {
	GetTenantByID(ctx context.Context, tenantID string) (entity.Tenant, error)
}
