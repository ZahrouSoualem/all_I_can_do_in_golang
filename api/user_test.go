package api

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mock_sqlc "github.com/tester/db/mock"
	db "github.com/tester/db/sqlc"
	"github.com/tester/util"
)

func TestGetUser(t *testing.T) {

	config, err := util.LoadConfig("../")

	if err != nil {
		log.Fatal("an erros has detected when loading configuration")
	}

	user := User()

	// gomock thhe framework
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// gomock the package
	store := mock_sqlc.NewMockStore(ctrl)

	//building the stubs

	store.EXPECT().
		GetUser(gomock.Any(), gomock.Eq(user.ID)).
		Times(1).
		Return(user, nil)

	//to start test server
	server, err := NewServer(store, "", config)
	require.NoError(t, err)

	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/user/%d", user.ID)

	req, err := http.NewRequest(http.MethodGet, url, nil)

	require.NoError(t, err)

	addAutorization(t, req, server.tokenMaker, AuthenticationKey, user.Username, time.Minute)

	server.router.ServeHTTP(recorder, req)

	require.Equal(t, http.StatusOK, recorder.Code)

}

func User() db.User {
	return db.User{
		ID:           int64(util.RandomInteger(2, 10)),
		Username:     util.RandomString(7),
		HashPassword: util.RandomString(7),
	}
}
