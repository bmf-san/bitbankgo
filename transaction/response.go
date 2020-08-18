package transaction

// TransactionResponse is a response for Transaction.
type TransactionResponse struct {
	Success int `json:"success"`
	Data    struct {
		Code         int `json:"code"`
		Transactions []struct {
			TransactionID int    `json:"transaction_id"`
			Side          string `json:"side"`
			Price         string `json:"price"`
			Amount        string `json:"amount"`
			ExecutedAt    int    `json:"executed_at"`
		} `json:"transactions"`
	} `json:"data"`
}
