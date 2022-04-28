package cache

var Token *tokenCache

type tokenCache struct {
}

func (c *tokenCache) SetToken(token string, uid int64) error {
	return nil
}
