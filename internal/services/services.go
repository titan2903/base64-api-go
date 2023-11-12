package services

type services struct {
}

type Services interface {
	EncodeBase64(encodeData string) (string, error)
	DecodeBase64(encodedData string) (string, error)
}

func NewServices() Services {
	return &services{}
}
