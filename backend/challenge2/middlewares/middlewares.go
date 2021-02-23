package middlewares

import (
	"github.com/rs/zerolog/log"
	"github.com/savsgio/atreugo/v11"
)

func LogRequest(ctx *atreugo.RequestCtx) error {
	log.Info().Msg(ctx.String())

	return ctx.Next()
}
