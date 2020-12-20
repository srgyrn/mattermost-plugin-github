package graphql

import (
	"flag"
	"testing"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var userToken string
var username string

func init() {
	flag.StringVar(&username, "username", "", "Github username")
	flag.StringVar(&userToken, "token", "", "Github user access token")
}

func TestClient(t *testing.T) {
	if userToken == "" || username == "" {
		t.Skipf("empty username or access token, skipping test")
	}

	tok := oauth2.Token{AccessToken: userToken}
	client := NewClient(tok, username, "", "")

	var query struct {
		Viewer struct {
			Login githubv4.String
		}
	}

	err := client.executeQuery(&query, nil)
	if err != nil {
		t.Fatalf("executeQuery() err: %v", err)
	}

	if query.Viewer.Login == "" {
		t.Errorf("struct field Login is empty, expected value")
	}
}
