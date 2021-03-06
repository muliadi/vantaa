package auth

import (
	lediscfg "github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
)

var cfg *lediscfg.Config = lediscfg.NewConfigDefault()
var l, _ = ledis.Open(cfg)
var db, _ = l.Select(0)

// BlacklistToken adds the token to a ledis database
func BlacklistToken(token []byte, expireAt int64) error {

	db.Set(token, token)
	db.ExpireAt(token, expireAt)

	return nil
}

// IsBlackListed checks if a token is in the blacklist list
func IsBlacklisted(token []byte) (bool, error) {
	res, err := db.Get(token)
	if err != nil {
		return false, err
	}

	if res == nil {
		return false, err
	}

	return true, nil
}

func ClearAllTokens() (int64, error) {
	l, err := ledis.Open(cfg)
	defer l.Close()
	if err != nil {
		return 0, err
	}

	db, err := l.Select(0)
	if err != nil {
		return 0, err
	}

	return db.FlushAll()
}
