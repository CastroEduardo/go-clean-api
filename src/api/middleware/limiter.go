// package middleware

// import (
// 	"net/http"

// 	"github.com/CastroEduardo/go-clean-api/api/helper"
// 	"github.com/didip/tollbooth"
// 	"github.com/gin-gonic/gin"
// )

// func LimitByRequest() gin.HandlerFunc {
// 	lmt := tollbooth.NewLimiter(1, nil)
// 	return func(c *gin.Context) {
// 		err := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusTooManyRequests,
// 				helper.GenerateBaseResponseWithError(nil, false, helper.LimiterError, err))
// 			return
// 		} else {
// 			c.Next()
// 		}
// 	}
// }

package middleware

import (
	"net"
	"net/http"
	"strings"

	"github.com/CastroEduardo/go-clean-api/api/helper"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func LimitByRequest() gin.HandlerFunc {

	lmt := tollbooth.NewLimiter(1, nil) // 5 req/s
	lmt.SetBurst(5)

	return func(c *gin.Context) {

		// 1. Intentar por X-API-Key
		apiKey := c.GetHeader("X-API-Key")
		if apiKey != "" {
			key := "apiKey|" + apiKey

			if httpErr := tollbooth.LimitByKeys(lmt, []string{key}); httpErr != nil {
				c.AbortWithStatusJSON(http.StatusTooManyRequests,
					helper.GenerateBaseResponseWithError(nil, false, helper.LimiterError, httpErr))
				return
			}

			c.Next()
			return
		}

		// 2. Intentar por Authorization
		auth := c.GetHeader("Authorization")
		if auth != "" {
			key := "auth|" + auth

			if httpErr := tollbooth.LimitByKeys(lmt, []string{key}); httpErr != nil {
				c.AbortWithStatusJSON(http.StatusTooManyRequests,
					helper.GenerateBaseResponseWithError(nil, false, helper.LimiterError, httpErr))
				return
			}

			c.Next()
			return
		}

		// 3. Fallback → IP real usando Gin (la forma correcta)
		clientIP := c.ClientIP()
		clientIP = cleanIP(clientIP)

		key := "ip|" + clientIP

		if httpErr := tollbooth.LimitByKeys(lmt, []string{key}); httpErr != nil {
			c.AbortWithStatusJSON(http.StatusTooManyRequests,
				helper.GenerateBaseResponseWithError(nil, false, helper.LimiterError, httpErr))
			return
		}

		c.Next()
	}
}

// Limpia el puerto y normaliza IPv6/IPv4
func cleanIP(raw string) string {
	// Si viene en formato "IP:PORT"
	if strings.Contains(raw, ":") {
		ip, _, err := net.SplitHostPort(raw)
		if err == nil {
			return ip
		}
	}

	return raw
}

// func LimitByRequest() gin.HandlerFunc {
// 	// Tasa de Mantenimiento: 5 peticiones por segundo
// 	// Burst (Estallido): Permite hasta 10 peticiones consecutivas rápidamente.
// 	// TTL (Time To Live): El contador se reinicia después de 1 segundo (por defecto, pero buena práctica establecerlo)
// 	lmt := tollbooth.NewLimiter(5, nil) // 5 peticiones por segundo
// 	// Configurar el límite de ráfaga (burst) para 10
// 	lmt.SetBurst(10)

// 	// Opcional: Puedes configurar un límite por usuario/IP para peticiones específicas
// 	// lmt.SetMessage("Demasiadas peticiones. Por favor, inténtalo más tarde.")

// 	return func(c *gin.Context) {
// 		// En este ejemplo, el límite se aplica por la IP del cliente (por defecto en tollbooth)
// 		err := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusTooManyRequests,
// 				helper.GenerateBaseResponseWithError(nil, false, helper.LimiterError, err))
// 			return
// 		} else {
// 			c.Next()
// 		}
// 	}
// }
