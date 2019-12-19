package daily

import (
"context"
)

type Service interface {
	Status(context.Context) (bool, error)
}