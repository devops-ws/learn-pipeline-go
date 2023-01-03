package oauth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetUserInfo sends the useinfo request, then returns the userInfo
func GetUserInfo(authServerURL, token string) (userInfo *UserInfo, err error) {
	var resp *http.Response
	if resp, err = http.Get(fmt.Sprintf("%s/oauth/userinfo?access_token=%s", authServerURL, token)); err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var data []byte
		if data, err = io.ReadAll(resp.Body); err != nil {
			return
		}

		userInfo = &UserInfo{}
		err = json.Unmarshal(data, userInfo)
	}
	return
}

type UserInfo struct {
	Sub               string   `json:"sub"`
	Name              string   `json:"name"`
	PreferredUsername string   `json:"preferred_username"`
	Email             string   `json:"email"`
	Picture           string   `json:"picture"`
	Groups            []string `json:"groups"`
}
