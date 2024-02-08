package api

import (
	"errors"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tester/token"
)

const (
	AuthenticationKey       = "Authorization"
	AuthenticationBearer    = "bearer"
	AyutheticationPyloadkey = "payload_key"
)

func authMiddleware(tokenMaker token.Meker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader(AuthenticationKey)

		if len(header) == 0 {
			err := errors.New("authorizationHeader header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse(err))
			return
		}

		Fields := strings.Fields(header)

		if len(Fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "check the authrization header"})
			return
		}

		headerBearer := strings.ToLower(Fields[0])

		if AuthenticationBearer != headerBearer {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "The beaer value is missed"})
			return
		}

		payload, err := tokenMaker.VerifyToken(Fields[1])

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.Set(AyutheticationPyloadkey, payload)

		ctx.Next()
	}

}
