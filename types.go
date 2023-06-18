package RbxOAuth

type Resource struct {
	ResourceInfos []ResourceInfos `json:"resource_infos"`
}
type Owner struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
type Universe struct {
	Ids []string `json:"ids"`
}
type Resources struct {
	Universe Universe `json:"universe"`
}
type ResourceInfos struct {
	Owner     Owner     `json:"owner"`
	Resources Resources `json:"resources"`
}

type Config struct {
	ClientID     string
	ClientSecret string
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token"`
	Scope        string `json:"scope"`
}

type TokenError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type IntrospectResp struct {
	Active    bool   `json:"active"`
	Jti       string `json:"jti"`
	Iss       string `json:"iss"`
	TokenType string `json:"token_type"`
	ClientID  string `json:"client_id"`
	Aud       string `json:"aud"`
	Sub       string `json:"sub"`
	Scope     string `json:"scope"`
	Exp       int    `json:"exp"`
	Iat       int    `json:"iat"`
}

type Info struct {
	Sub               string `json:"sub"`
	Name              string `json:"name"`
	Nickname          string `json:"nickname"`
	PreferredUsername string `json:"preferred_username"`
	CreatedAt         int    `json:"created_at"`
	Profile           string `json:"profile"`
}
