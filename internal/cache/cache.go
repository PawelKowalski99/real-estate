package cache

import (
	"context"
	"fmt"
)

var (
	NotExistError = fmt.Errorf("Key does not exist in cache")
)

type Cache interface {
    Get(context.Context, string) (string, error)
	Set(context.Context, string) error
}