package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func InitMiddleware(e *echo.Echo, wrapper Wrapper) {
	e.Use(RequestLoggerMiddleware(wrapper))
}

func RequestLoggerMiddleware(wrapper Wrapper) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			fields := []zapcore.Field{
				zap.String("ClientIP", c.RealIP()),
				zap.String("Latency", time.Since(start).String()),
				zap.String("URL", fmt.Sprintf("%s %s", req.Method, req.RequestURI)),
				zap.Int("Status", res.Status),
				zap.String("UserAgent", req.UserAgent()),
			}

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
				fields = append(fields, zap.String("request_id", id))
			}

			n := res.Status
			switch {
			case n >= 500:
				wrapper.GetLogger().GetZapLogger().With(zap.Error(err)).Error("Server error", fields)
			case n >= 400:
				wrapper.GetLogger().GetZapLogger().With(zap.Error(err)).Warn("Client error", fields)
			case n >= 300:
				wrapper.GetLogger().GetZapLogger().Info("Redirection", fields)
			default:
				wrapper.GetLogger().GetZapLogger().Info("Success", fields)
			}

			return nil
		}
	}
}
