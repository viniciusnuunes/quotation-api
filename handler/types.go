package handler

type Handlers struct {
	Currency *CurrencyHandler
}

func MustInit() *Handlers {
	return &Handlers{
		Currency: &CurrencyHandler{},
	}
}
