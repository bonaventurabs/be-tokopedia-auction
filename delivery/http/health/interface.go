package health

import "context"

type Usecases interface {
	HealthCheck(ctx context.Context) error
}
