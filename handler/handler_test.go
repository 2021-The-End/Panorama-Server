package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPort(t *testing.T) {
	a := os.Getenv("PORT")
	fmt.Println(a)
	assert := assert.New(t)

	ah := MakeHandler()
	ts := httptest.NewServer(ah)

	_, err := http.Get(ts.URL + "/swagger/index.html")

	assert.NoError(err)
}
