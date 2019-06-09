package resource

import (
    "crypto/md5"
    "errors"
    "fmt"
    "github.com/rs/zerolog"
    "net/http"
    "strings"
)

var authPassword string

// TODO: Copied from PHP, insecure...
func init() {
    insecureAuthorizationPassword := []byte("aaron")
    authPassword = fmt.Sprintf("%x", md5.Sum(insecureAuthorizationPassword))
}

func authenticateRequest(header http.Header, logger *zerolog.Logger) (int, error) {

    authHeader := header.Get("Authorization")

    logger.Info().Msgf("authenticateRequest. Authorization=%s", authHeader)

    if len(authHeader) <= 0 {
        return http.StatusBadRequest, errors.New("Please provide authorize key.")
    }

    if strings.ToLower(authHeader) != authPassword {
        return http.StatusUnauthorized, errors.New("Unauthorized access.")
    }

    return 0, nil
}
