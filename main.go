package RbxOAuth

import (
	"encoding/json"
	"io"
	"net/http"
)

func UserINFO(JWT string) (Info, error) {
	resp, err := http.NewRequest("GET", "https://apis.roblox.com/oauth/v1/userinfo", nil)
	if err != nil {
		return Info{}, err
	}
	resp.Header.Add("Authorization", "Bearer "+JWT)
	r, err := http.DefaultClient.Do(resp)
	if err != nil {
		return Info{}, err
	}
	bb, _ := io.ReadAll(r.Body)
	var Data Info
	json.Unmarshal(bb, &Data)
	return Data, nil
}
