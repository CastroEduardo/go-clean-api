package middleware

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/CastroEduardo/go-clean-api/config"
	"github.com/CastroEduardo/go-clean-api/pkg/logging"
	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func DefaultStructuredLogger(cfg *config.Config) gin.HandlerFunc {
	logger := logging.NewLogger(cfg)
	return structuredLogger(logger)
}

func structuredLogger(logger logging.Logger) gin.HandlerFunc {
	const maxLogSize = 200 // máximo de bytes a loggear del ResponseBody

	return func(c *gin.Context) {
		if strings.Contains(c.FullPath(), "swagger") {
			c.Next()
			return
		}

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		start := time.Now()
		path := c.FullPath()
		raw := c.Request.URL.RawQuery

		// Leer Body sin consumirlo
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body.Close()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// Captura de respuesta
		c.Writer = blw
		c.Next()

		// Preparar parámetros
		param := gin.LogFormatterParams{}
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		// Construir el log
		keys := map[logging.ExtraKey]interface{}{}
		keys[logging.Path] = param.Path
		keys[logging.ClientIp] = param.ClientIP
		keys[logging.Method] = param.Method
		keys[logging.Latency] = param.Latency
		keys[logging.StatusCode] = param.StatusCode
		keys[logging.ErrorMessage] = param.ErrorMessage
		keys[logging.BodySize] = param.BodySize
		keys[logging.RequestBody] = string(bodyBytes)

		// Analizar el tipo de response y limitar tamaño
		responseBytes := blw.body.Bytes()
		contentType := c.Writer.Header().Get("Content-Type")
		if strings.Contains(contentType, "application/json") {
			if len(responseBytes) > maxLogSize {
				keys[logging.ResponseBody] = string(responseBytes[:maxLogSize]) + "...(truncated)"
			} else {
				keys[logging.ResponseBody] = string(responseBytes)
			}
		} else {
			keys[logging.ResponseBody] = fmt.Sprintf("[binary %d bytes]", len(responseBytes))
		}

		logger.Info(logging.RequestResponse, logging.Api, "", keys)
	}
}

// func structuredLogger(logger logging.Logger) gin.HandlerFunc {

// 	fmt.Println("REGISTER LOGGER__")
// 	return func(c *gin.Context) {
// 		if strings.Contains(c.FullPath(), "swagger") {
// 			c.Next()
// 		} else {
// 			blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
// 			start := time.Now() // start
// 			path := c.FullPath()
// 			raw := c.Request.URL.RawQuery

// 			bodyBytes, _ := io.ReadAll(c.Request.Body)
// 			c.Request.Body.Close()
// 			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

// 			c.Writer = blw
// 			c.Next()

// 			param := gin.LogFormatterParams{}
// 			param.TimeStamp = time.Now() // stop
// 			param.Latency = param.TimeStamp.Sub(start)
// 			param.ClientIP = c.ClientIP()
// 			param.Method = c.Request.Method
// 			param.StatusCode = c.Writer.Status()
// 			param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
// 			param.BodySize = c.Writer.Size()

// 			if raw != "" {
// 				path = path + "?" + raw
// 			}
// 			param.Path = path

// 			keys := map[logging.ExtraKey]interface{}{}
// 			keys[logging.Path] = param.Path
// 			keys[logging.ClientIp] = param.ClientIP
// 			keys[logging.Method] = param.Method
// 			keys[logging.Latency] = param.Latency
// 			keys[logging.StatusCode] = param.StatusCode
// 			keys[logging.ErrorMessage] = param.ErrorMessage
// 			keys[logging.BodySize] = param.BodySize
// 			keys[logging.RequestBody] = string(bodyBytes)
// 			keys[logging.ResponseBody] = blw.body.String()

// 			// jsonBytes, _ := json.MarshalIndent(keys, "", "  ")
// 			// fmt.Println("START LOGGER__👍\n", string(jsonBytes))

// 			fmt.Println("END LOGGER 👍 " + logging.Api)
// 			logger.Info(logging.RequestResponse, logging.Api, "", keys)
// 		}
// 	}
// }
