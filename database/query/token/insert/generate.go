package insert

import (
	"encoding/hex"
	"errors"
	"strconv"
	"time"
)

// Generate token
func Generate(email string) (string, error) {
	time := strconv.FormatInt(time.Now().UnixNano(), 32)
	str := shuffle(email) + time
	token := hex.EncodeToString([]byte(str))
	saved := save(email, token)
	if saved {
		return token, nil
	}

	return "", errors.New("NotSaved")
}
