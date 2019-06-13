package cdek

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

//ClientConf SDK client configuration
type ClientConf struct {
	Auth      Auth
	XmlApiUrl string
}

//Auth account credentials
type Auth struct {
	Account string
	Secure  string
}

//EncodedSecure encode secure according to CDEK api
func (a Auth) EncodedSecure() (date string, encodedSecure string) {
	date = time.Now().Format("2006-01-02 15:04:05")
	encoder := md5.New()
	_, _ = encoder.Write([]byte(date + "&" + a.Secure))

	return date, hex.EncodeToString(encoder.Sum(nil))
}