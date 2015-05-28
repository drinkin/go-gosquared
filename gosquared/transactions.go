package gosquared

import "github.com/dghubble/sling"

type TransactionOpts struct {
	Revenue  int `json:"revenue,omitempty"`
	Quantity int `json:"quantity,omitempty"`
}

type TransactionItem struct {
	Name       string   `json:"name"`
	Categories []string `json:"categories"`
	Revenue    int      `json:"revenue,omitempty"`
	Price      int      `json:"price"`
	Quantity   int      `json:"quantity"`
}

type TransactionRequest struct {
	PersonID    string `json:"person_id,omitempty"`
	Transaction struct {
		ID    string            `json:"id"`
		Opts  *TransactionOpts  `json:"opts"`
		Items []TransactionItem `json:"items"`
	} `json:"transaction"`
}

type TransactionService struct {
	sling *sling.Sling
}

func NewTransactionServicee(sling *sling.Sling) *TransactionService {
	return &TransactionService{
		sling: sling.Path("transaction/"),
	}
}
