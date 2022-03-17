package data

import (
	"context"
	"github.com/henriSedjame/webee/utils"
)

//LoadDB loads datasource stored into the context
func LoadDB[T](ctx context.Context) T {
	return ctx.Value(utils.DBCtxKey).(T)
}
