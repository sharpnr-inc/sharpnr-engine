package cashfreegateway

import cashfree_pg "github.com/cashfree/cashfree-pg/v5"

type CashfreeProvider struct {
	ClientID     string
	ClientSecret string
	Environment  cashfree_pg.CFEnvironment
}
