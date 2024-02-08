package api

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tester/token"
)

func addAutorization(
	t *testing.T,
	request *http.Request,
	tokenmaker token.Meker,
	authorizationType string,
	username string,
	duration time.Duration,
) {

	token, _, err := tokenmaker.CreateToken(username, duration)

	require.NoError(t, err)
	/* require.NotEmpty(t,payload)
	require.NotEmpty(t,token) */

	authorizationHeader := fmt.Sprintf("%s %s", AuthenticationBearer, token)
	request.Header.Set(AuthenticationKey, authorizationHeader)

}
