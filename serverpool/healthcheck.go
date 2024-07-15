package serverpool

import (
	"context"
	"time"

	"github.com/zakirkun/yae-miko/utils"
)

func LauchHealthCheck(ctx context.Context, sp ServerPool) {
	t := time.NewTicker(time.Second * 20)
	utils.Logger.Info("Starting health check...")
	for {
		select {
		case <-t.C:
			go HealthCheck(ctx, sp)
		case <-ctx.Done():
			utils.Logger.Info("Closing Health Check")
			return
		}
	}
}
