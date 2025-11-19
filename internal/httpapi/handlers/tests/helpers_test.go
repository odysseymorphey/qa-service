package tests

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func withChiURLParam(r *http.Request, key, value string) *http.Request {
	routeCtx := chi.NewRouteContext()
	routeCtx.URLParams.Add(key, value)
	ctx := context.WithValue(r.Context(), chi.RouteCtxKey, routeCtx)
	return r.WithContext(ctx)
}
