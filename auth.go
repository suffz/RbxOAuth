package RbxOAuth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (C *Config) Token(Code, Verifier string) (Token, error) {

	req, err := http.NewRequest("POST", "https://apis.roblox.com/oauth/v1/token", bytes.NewBuffer([]byte(fmt.Sprintf(`client_id=%v&code_verifier=%v&client_secret=%v&grant_type=authorization_code&code=%v`, C.ClientID, Verifier, C.ClientSecret, Code))))
	if err != nil {
		return Token{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Token{}, err
	}
	js, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		var T Token
		json.Unmarshal(js, &T)
		return T, nil
	} else {
		var T TokenError
		json.Unmarshal(js, &T)
		return Token{}, errors.New(fmt.Sprintf(`Error: %v - %v`, T.Error, T.ErrorDescription))
	}
}

func (C *Config) RefreshToken(Refresh string) (Token, error) {
	req, err := http.NewRequest("POST", "https://apis.roblox.com/oauth/v1/token", bytes.NewBuffer([]byte(fmt.Sprintf(`client_id=%v&client_secret=%v&grant_type=refresh_token&refresh_token=%v`, C.ClientID, C.ClientSecret, Refresh))))
	if err != nil {
		return Token{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Token{}, err
	}
	js, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		var T Token
		json.Unmarshal(js, &T)
		return T, nil
	} else {
		var T TokenError
		json.Unmarshal(js, &T)
		return Token{}, errors.New(fmt.Sprintf(`Error: %v - %v`, T.Error, T.ErrorDescription))
	}
}

func (C *Config) Introspect(Bearer string) (IntrospectResp, error) {
	req, err := http.NewRequest("POST", "https://apis.roblox.com/oauth/v1/token/introspect", bytes.NewBuffer([]byte(fmt.Sprintf(`client_id=%v&client_secret=%v&token=%v`, C.ClientID, C.ClientSecret, Bearer))))
	if err != nil {
		return IntrospectResp{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return IntrospectResp{}, err
	}
	js, _ := io.ReadAll(resp.Body)
	var T IntrospectResp
	json.Unmarshal(js, &T)
	return T, nil
}

func (C *Config) Resources(Bearer string) (Resource, error) {
	req, err := http.NewRequest("POST", "https://apis.roblox.com/oauth/v1/token/resources", bytes.NewBuffer([]byte(fmt.Sprintf(`client_id=%v&client_secret=%v&token=%v`, C.ClientID, C.ClientSecret, Bearer))))
	if err != nil {
		return Resource{}, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Resource{}, err
	}
	js, _ := io.ReadAll(resp.Body)
	var T Resource
	json.Unmarshal(js, &T)
	return T, nil
}

func (C *Config) Revoke(Refresh string) error {
	req, err := http.NewRequest("POST", "https://apis.roblox.com/oauth/v1/token/resources", bytes.NewBuffer([]byte(fmt.Sprintf(`client_id=%v&client_secret=%v&token=%v`, C.ClientID, C.ClientSecret, Refresh))))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("Error: Unknown status code %v | %v", resp.StatusCode, string(body)))
	}
	return nil
}
