package account

import (
	"github.com/willyrgf/rerest/pkg/repository"
	"github.com/prometheus/common/log"
)

// Authorized accounts all requests authorized
func Authorized(key string) {
	conn := repository.GetConnToControlDB()
	defer conn.Close()

	if _, err := conn.Do("HINCRBY", key, "count_auth", 1); err != nil {
		log.Errorf("account.Authorized() error on hincrby count_auth=%w", err)
		return
	}

}

// Response accounts all answered requests
func Response(key string) {
	conn := repository.GetConnToControlDB()
	defer conn.Close()

	if _, err := conn.Do("HINCRBY", key, "count", 1); err != nil {
		log.Errorf("account.Response() error on hincrby count=%w", err)
		return
	}

}
