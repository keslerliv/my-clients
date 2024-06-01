package models

import (
	"fmt"
	"strings"

	"github.com/keslerliv/my-clients/internal/entities"
	"github.com/keslerliv/my-clients/pkg/db"
	_ "github.com/lib/pq"
)

// Client POST model
func ClientInsert(client entities.Client) (id int64, err error) {

	// open connection
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	// insert client query
	sql := `INSERT INTO client (cpf, private, incomplete, date_last_purchase, average_ticket, ticket_last_purchase, frequent_store, last_store) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	_, err = conn.Exec(sql, client.CPF, client.Private, client.Incomplete, client.DateLastPurchase, client.AverageTicket, client.TicketLastPurchase, client.FrequentStore, client.LastStore)
	if err != nil {
		return
	}

	return
}

// Client GET model
func ClientGet(id int64) (client entities.Client, err error) {

	// open connection
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	// get client by id query
	row := conn.QueryRow("SELECT * FROM client WHERE id = $1", id)

	err = row.Scan(&client.ID, &client.CPF, &client.Private, &client.Incomplete, &client.DateLastPurchase, &client.AverageTicket, &client.TicketLastPurchase, &client.FrequentStore, &client.LastStore)

	return
}

// Client GET-ALL model
func ClientGetAll() (client []entities.Client, err error) {
	// open connection
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	// get all clients query
	rows, err := conn.Query("SELECT * FROM client")
	if err != nil {
		return
	}

	// parse all clients
	for rows.Next() {
		var c entities.Client

		err = rows.Scan(&c.ID, &c.CPF, &c.Private, &c.Incomplete, &c.DateLastPurchase, &c.AverageTicket, &c.TicketLastPurchase, &c.FrequentStore, &c.LastStore)
		if err != nil {
			continue
		}

		client = append(client, c)
	}

	return
}

// Client PUT model
func ClientUpdate(id int64, client entities.Client) (int64, error) {
	// open connection
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	// update client query
	res, err := conn.Exec(
		"UPDATE client SET cpf = $1, private = $2, incomplete = $3, date_last_purchase = $4, average_ticket = $5, ticket_last_purchase = $6, frequent_store = $7, last_store = $8 WHERE id = $9",
		client.CPF, client.Private, client.Incomplete, client.DateLastPurchase, client.AverageTicket, client.TicketLastPurchase, client.FrequentStore, client.LastStore, id,
	)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Client DELETE model
func ClientDelete(id int64) (int64, error) {
	// open connection
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	// delete client query
	res, err := conn.Exec("DELETE FROM client WHERE id = $1", id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

// Client BATCH model
func ClientListInsert(clients []entities.Client) (err error) {

	// open connection
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	// begin transaction
	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	// load in groups of batch
	batchSize := 6000
	for i := 0; i < len(clients); i += batchSize {
		end := i + batchSize
		if end > len(clients) {
			end = len(clients)
		}
		batch := clients[i:end]

		// build query and values to batch
		sql := `INSERT INTO client (cpf, private, incomplete, date_last_purchase, average_ticket, ticket_last_purchase, frequent_store, last_store) VALUES `
		values := []interface{}{}
		valueStrings := []string{}

		// insert values on query
		for i, client := range batch {
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*8+1, i*8+2, i*8+3, i*8+4, i*8+5, i*8+6, i*8+7, i*8+8))

			values = append(values, client.CPF, client.Private, client.Incomplete, client.DateLastPurchase, client.AverageTicket, client.TicketLastPurchase, client.FrequentStore, client.LastStore)
		}
		sql += strings.Join(valueStrings, ", ") + " RETURNING id"

		// prepare the query
		stmt, err := tx.Prepare(sql)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		// execute the query
		_, err = stmt.Exec(values...)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
