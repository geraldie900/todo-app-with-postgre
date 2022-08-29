package model

type HTTPQueryParameter struct {
	Limit   int    `query:"limit"`
	Offset  int    `query:"offset"`
	OrderBy string `query:"order_by"`
	Desc    string `query:"desc"`
}
