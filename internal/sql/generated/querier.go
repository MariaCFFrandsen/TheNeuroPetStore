// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package generated

import (
	"context"
)

type Querier interface {
	GetPets(ctx context.Context) ([]Pet, error)
}

var _ Querier = (*Queries)(nil)