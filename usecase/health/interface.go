package health

import "context"

type HealthCheck interface {
	HealthCheck(ctx context.Context) error
}
