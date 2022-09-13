package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iamsrikants/rest-curd-api/startup/db"
)

// ListAccounts returns List of Accounts.
func ListAccounts(c *gin.Context) {
	if res, err := getAllAccounts(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func getAllAccounts(c *gin.Context) ([]AccountLoyaltyMemberRef, error) {
	var accounts []AccountLoyaltyMemberRef
	conn := db.Postgres

	query := "SELECT * FROM catch.account_loyalty_member_reference"
	fmt.Println("Query: ", query)
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	count := 0
	// iterate through the rows
	for rows.Next() {
		count++
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}

		// convert DB types to Go types
		account := AccountLoyaltyMemberRef{
			Account_id:                values[0].(int64),
			Contact_id:                values[1].(int64),
			Loyalty_id:                values[2].(int64),
			Primary_loyalty_member_id: values[3].(bool),
			Loyalty_type_code:         values[4].(string),
			Loyalty_matched_code:      values[5].(string),
			Brand:                     values[6].(string),
			Create_rcd_timestamp:      values[7].(time.Time),
			Create_rcd_by_who:         values[8].(string),
			Create_rcd_by_app:         values[9].(string),
			Update_rcd_timestamp:      values[10].(time.Time),
			Update_rcd_by_who:         values[11].(string),
			Update_rcd_by_app:         values[12].(string),
			Original_data_source:      values[13].(string),
		}

		accounts = append(accounts, account)
	}

	if count > 0 {
		return accounts, nil
	} else {
		return nil, errors.New("record Not Found")
	}
}
