package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/jacky-htg/shonet-frontend/config"
	"github.com/jacky-htg/shonet-frontend/libraries"
	"github.com/jacky-htg/shonet-frontend/models"
	"github.com/jacky-htg/shonet-frontend/repositories"
	"github.com/yvasiyarov/php_session_decoder/php_serialize"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type Payload struct {
	IV    string
	Value string
	Mac   string
}

func GetSessionLaravel(cookie, key string) (string, error) {

	decodeBytes, err := base64.StdEncoding.DecodeString(cookie)
	if err != nil {
		return "", errors.New("cookie value must on base64 format")
	}

	payload := Payload{}

	err = json.Unmarshal(decodeBytes, &payload)
	if err != nil {
		return "", errors.New("cookie must be valid base64 format")
	}

	encryptedText, err := base64.StdEncoding.DecodeString(payload.Value)
	if err != nil {
		return "", errors.New("encrypting text must be valid base64 format")
	}

	iv, err := base64.StdEncoding.DecodeString(payload.IV)
	if err != nil {
		return "", errors.New("iv in payload must be valid base64 format")
	}

	var keyBytes []byte

	if (strings.HasPrefix(key, "base64:")) {
		keyBytes, err = base64.StdEncoding.DecodeString(string(key[7:]))
		if err != nil {
			return "", errors.New("seems like your provide a key in base64 format, but not valid")
		}
	} else {
		keyBytes = []byte(key)
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptedText, encryptedText)

	unserializer := php_serialize.NewUnSerializer(string(encryptedText))

	var sessionID php_serialize.PhpValue
	if sessionID, err = unserializer.Decode(); err != nil {
		return string(removePadding(encryptedText)), nil
	}

	return sessionID.(string), nil
}

func removePadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}

func parseSessionData(data string) (php_serialize.PhpArray, error) {

	decoder := php_serialize.NewUnSerializer(data)
	sessionDataDecoded, err := decoder.Decode()

	if err != nil {
		return nil, err
	}

	if reflect.TypeOf(sessionDataDecoded).Kind() == reflect.String {
		return parseSessionData(sessionDataDecoded.(string))
	}

	return sessionDataDecoded.(php_serialize.PhpArray), nil
}

func GetUserClient(r *http.Request) (models.User, error) {

	userID := 0
	user := models.User{}
	cook, err := r.Cookie(config.GetString("laravel.sessionName"))
	if err != nil {
		return models.User{}, err
	}

	cookie, err := url.QueryUnescape(cook.Value)
	if err != nil {
		return models.User{}, err
	}

	sessionID, err := GetSessionLaravel(cookie, config.GetString("laravel.appKey"))
	if err != nil {
		return models.User{}, err
	}

	keySession, err := libraries.RedisGet(sessionID)
	if err != nil {
		return models.User{}, err
	}

	sessionParse, err := parseSessionData(string(keySession))
	if err != nil {
		return models.User{}, err
	}

	if len(sessionParse) > 0 {
		for k, val := range sessionParse {
			if strings.HasPrefix(k.(string), "login_web_") {
				userID = val.(int)
			}
		}

		if userID <= 0 {
			return models.User{}, nil
		}
	}

	if userID > 0 {
		user, err = repositories.GetUserByID(uint(userID))
		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}