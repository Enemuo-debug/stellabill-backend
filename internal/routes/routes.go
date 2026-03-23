package routes

import (
	"stellarbill-backend/internal/config"
	"stellarbill-backend/internal/cors"
	"stellarbill-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	cfg := config.Load()
	corsProfile := cors.ProfileForEnv(cfg.Env, cfg.AllowedOrigins)

	r.Use(cors.Middleware(corsProfile))

	api := r.Group("/api")
	{
		api.GET("/health", handlers.Health)
		api.GET("/subscriptions", handlers.ListSubscriptions)
		api.GET("/subscriptions/:id", handlers.GetSubscription)
		api.GET("/plans", handlers.ListPlans)
	}
}
