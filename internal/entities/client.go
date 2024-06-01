package entities

import "time"

type Client struct {
	ID                 int64      `json:"id"`
	CPF                string     `json:"cpf"`
	Private            bool       `json:"private"`
	Incomplete         bool       `json:"incomplete"`
	DateLastPurchase   *time.Time `json:"date_last_purchase"`
	AverageTicket      int64      `json:"average_ticket"`
	TicketLastPurchase int64      `json:"ticket_last_purchase"`
	FrequentStore      *string    `json:"frequent_store"`
	LastStore          *string    `json:"last_store"`
}
