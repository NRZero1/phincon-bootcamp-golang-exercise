package domain

type Package struct {
	PackageID          int     `json:"package_id"`
	Name               string  `json:"package_name"`
	Price              float64 `json:"price"`
	BalanceRequirement float64 `json:"requirement"`
}
