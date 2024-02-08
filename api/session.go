package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/tester/db/sqlc"
	"github.com/tester/util"
)

type SessionRespense struct {
	SessionID    uuid.UUID `json:"id"`
	Userid       int64     `json:"userid"`
	RefreshToken string    `json:"refresh_token"`
	RefExpiresAt time.Time `json:"refexpires_at"`
	AccessToken  string    `json:"access_token"`
	AccExpiresAt time.Time `json:"accexpires_at"`
}

func (s *Server) loginuser(ctx *gin.Context) {

	var req UserCreateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	user, err := s.store.GetUserByName(ctx, req.Username)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	err = util.CheckPassword(user.HashPassword, req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	// Access Token
	acc_Token_dur, err := time.ParseDuration(s.config.AccTokDuration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	acc_Token, acc_payload, err := s.tokenMaker.CreateToken(req.Username, acc_Token_dur)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	// Refresh Token
	ref_Token_dur, err := time.ParseDuration(s.config.RefTokDuration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	Ref_Token, ref_payload, err := s.tokenMaker.CreateToken(req.Username, ref_Token_dur)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	session, err := s.store.CreateSessions(ctx, db.CreateSessionsParams{
		ID:           ref_payload.Id,
		Userid:       user.ID,
		RefreshToken: Ref_Token,
		ExpiresAt:    ref_payload.Expire_at,
		CreatedAt:    ref_payload.Create_at,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, SessionRespense{
		SessionID:    session.ID,
		Userid:       user.ID,
		RefreshToken: Ref_Token,
		RefExpiresAt: ref_payload.Expire_at,
		AccessToken:  acc_Token,
		AccExpiresAt: acc_payload.Expire_at,
	})

}

type RenewTokenRequestPram struct {
	RefToken string `json:"ref" binding:"required"`
}

type RenewTokenRespensePram struct {
	AccessToken  string    `json:"access_token"`
	AccExpiresAt time.Time `json:"accexpires_at"`
}

func (s *Server) renewtoken(ctx *gin.Context) {
	var req RenewTokenRequestPram

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	payload, err := s.tokenMaker.VerifyToken(req.RefToken)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "the refresh toen is invalid you ghave to login again"})
		return
	}

	session, err := s.store.GetSessions(ctx, payload.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	if session.Isbloced {
		err := errors.New("the session is alreader bloced try to login agains please")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if session.RefreshToken != req.RefToken {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "The is no such session stored in the server"})
		return
	}

	if time.Now().After(session.ExpiresAt) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "The refresh token Token has expired"})
		return
	}

	// Refresh Token
	acc_Token_dur, err := time.ParseDuration(s.config.AccTokDuration)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	token, payload, err := s.tokenMaker.CreateToken(payload.Username, acc_Token_dur)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, RenewTokenRespensePram{
		AccessToken:  token,
		AccExpiresAt: payload.Expire_at,
	})

}
