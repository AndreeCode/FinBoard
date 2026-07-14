package domains

type CreateInvestmentRequest struct {
	TransactionId string  `json:"transaction_id"`
	ExpectedGain  float64 `json:"expected_gain"`
	RiskLevel     string  `json:"risk_level"`
	Status        string  `json:"status"`
}
type UpdateInvestmentRequest struct {
	ExpectedGain *float64 `json:"expected_gain"`
	RiskLevel   *string  `json:"risk_level"`
	Status      *string  `json:"status"`
}
