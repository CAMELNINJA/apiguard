package routes

import (
	"log"
	"net/http"

	"github.com/CAMELNINJA/apiguard/config"
	"github.com/CAMELNINJA/apiguard/middleware"
	zap_helper "github.com/CAMELNINJA/apiguard/pkg/zap_once"
	"github.com/CAMELNINJA/apiguard/proxy"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func SetupRouter(cfg *config.Config) http.Handler {
	r := chi.NewRouter()

	logger := zap_helper.GetLogger()

	r.Use(middleware.RequestId)
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())
	r.Use(middleware.AuthOptional(cfg))

	for _, route := range cfg.Routes {

		proxyHandler, err := proxy.NewReverseProxy(route.Upstream)
		if err != nil {
			log.Fatalf("error creating proxy for %s: %v", route.Name, err)
		}

		handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

			logger.Info("proxying request",
				zap.String("route", route.Name),
				zap.String("path", req.URL.Path),
				zap.String("method", req.Method),
			)
			proxyHandler.ServeHTTP(w, req)
		})

		r.Mount(route.MatchPrefix, handler)

	}

	return r
}
