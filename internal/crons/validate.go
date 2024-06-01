package crons

import (
	"regexp"
	"strings"
	"time"

	"github.com/keslerliv/my-clients/pkg/db"
	"github.com/keslerliv/my-clients/pkg/utils"
)

func ValidateClient() {
	go func() {
		for {

			// open DB connection
			conn, err := db.OpenConnection()
			if err != nil {
				time.Sleep(time.Minute)
				continue
			}

			// get all clients query
			rows, err := conn.Query("SELECT id, cpf, incomplete FROM client")
			if err != nil {
				conn.Close()
				time.Sleep(time.Minute)
				continue
			}

			for rows.Next() {
				var id int
				var cpf string
				var incomplete bool

				err = rows.Scan(&id, &cpf, &incomplete)

				// remove non-numeric characters
				cpfNumerico := strings.Join(regexp.MustCompile("[0-9]+").FindAllString(cpf, -1), "")

				// check validate cpf
				if utils.ValidaCPF(cpfNumerico) {
					// CPF is valid, update without changing incomplete flag
					_, err = conn.Exec("UPDATE client SET cpf = $1 WHERE id = $2", cpfNumerico, id)
				} else {
					// CPF is invalid, update and set incomplete flag
					_, err = conn.Exec("UPDATE client SET cpf = $1, incomplete = true WHERE id = $2", cpfNumerico, id)
				}

			}

			rows.Close()
			conn.Close()

			time.Sleep(time.Hour)
		}
	}()
}
