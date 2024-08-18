package refreshtokenservice

type Service struct {
	cfg tokenConfig
}

func New(cfg tokenConfig) *Service {
	return &Service{
		cfg: cfg,
	}
}

type tokenConfig interface {
	SecretKey() string
	TTLSecond() int
}
