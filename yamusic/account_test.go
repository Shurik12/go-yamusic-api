package yamusic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountService_GetStatus(t *testing.T) {
	setup()
	defer teardown()

	want := &AccountStatusResp{}
	want.InvocationInfo.ReqID = "Account.GetStatus"

	mux.HandleFunc("/account/status", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
		b, err := json.Marshal(want)
		assert.NoError(t, err)
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Account().GetStatus(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}
