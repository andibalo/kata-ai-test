package middleware

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"io"
	"pokemon-be/pkg/logger"
	"time"
)

// RequestLogger logs a gin HTTP request in JSON format. Allows to set the
// logger for testing purposes.
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {

		l := logger.GetLogger()

		requestId := uuid.New().String()

		c.Set("request_id", requestId)

		ctx := context.WithValue(c, "request_id", requestId)
		c.Request.WithContext(ctx)

		start := time.Now() // Start timer
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Read the request body
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		bodyString := string(bodyBytes)

		// Process request
		c.Next()

		// Fill the params
		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now() // Stop timer
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		// Log using the params
		var logEvent *zerolog.Event
		if c.Writer.Status() >= 500 {
			logEvent = l.Error()
		} else {
			logEvent = l.Info()
		}

		logEvent.Str("client_id", param.ClientIP).
			Str("method", param.Method).
			Int("status_code", param.StatusCode).
			Int("body_size", param.BodySize).
			Str("path", param.Path).
			Str("latency", param.Latency.String()).
			Str("req", bodyString).
			Str("request_id", requestId).
			Msg(param.ErrorMessage)

	}
}
