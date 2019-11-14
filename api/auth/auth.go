package auth

import (
	"fmt"

	"github.com/willyrgf/rerest/pkg/account"
	"github.com/willyrgf/rerest/pkg/repository"
	"github.com/garyburd/redigo/redis"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

const (
	uuidMinLength = 36
	// SetAPIKeysEnabled set of api_keys enabled
	SetAPIKeysEnabled = "api_keys:enabled"
)

func isValid(apiKey string) (valid bool) {
	_, err := uuid.Parse(apiKey)
	if err != nil {
		log.Debugf("auth.isValid() uuid.Parse() error=%w", err)
		return
	}

	valid = true
	log.Debugf("auth.isValid()=%+v", valid)
	return
}

// IsAuthorized check if the apikey is authorized
func IsAuthorized(apiKey string) (auth bool) {
	if !isValid(apiKey) {
		return
	}

	key := fmt.Sprintf("%s:%s", "api_key", apiKey)

	conn := repository.GetConnToControlDB()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("SISMEMBER", SetAPIKeysEnabled, key))
	if err != nil {
		log.Errorf("auth.IsAuthorized() check key with sismemeber=%w", err)
		return
	}

	log.Debugf("auth.IsAuthorized() the api key exists=%+v", exists)

	if exists {
		go account.Authorized(key)
	}

	auth = exists
	return
}
