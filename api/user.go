package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/tester/db/sqlc"
	"github.com/tester/util"
)

type UserCreateRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserCreateResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type GetUserResponse struct {
	ID int64 `uri:"id" binding:"required"`
}

type GetUserByNameResponse struct {
	Username string `json:"username"`
}

func (s *Server) createUser(ctx *gin.Context) {

	var req UserCreateRequest

	fmt.Println("1", req)

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("1")
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	HashedPassword, err := util.HashPassword(req.Password)

	if err != nil {
		fmt.Println("2")
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	user, err := s.store.CreateUser(ctx, db.CreateUserParams{
		Username:     req.Username,
		HashPassword: HashedPassword,
	})

	if err != nil {
		fmt.Println("3")
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, UserCreateResponse{ID: user.ID, Username: user.Username})

}

func (s *Server) getUser(ctx *gin.Context) {

	var req GetUserResponse

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	user, err := s.store.GetUser(ctx, req.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, UserCreateResponse{ID: user.ID, Username: user.Username})

}

func (s *Server) getUserByname(ctx *gin.Context) {

	var req GetUserByNameResponse

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	user, err := s.store.GetUserByName(ctx, req.Username)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, UserCreateResponse{ID: user.ID, Username: user.Username})

}

func (s *Server) getUsers(ctx *gin.Context) {

	users, err := s.store.GetUsersList(ctx)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)

}

func (s *Server) deleteUser(ctx *gin.Context) {

	var req GetUserResponse

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	err := s.store.DeleteUser(ctx, req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nil)

}

func (s *Server) updateUser(ctx *gin.Context) {

	var req UserCreateResponse

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	user, err := s.store.UpdateUser(ctx, db.UpdateUserParams{
		ID:       req.ID,
		Username: req.Username,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, UserCreateResponse{ID: user.ID, Username: user.Username})

}
