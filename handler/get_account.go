package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iamsrikants/rest-curd-api/startup/db"
)

func GetAccount(c *gin.Context) {
	if eval := GetSingleAccount(c); eval == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
	} else {
		c.JSON(http.StatusOK, eval)
	}
}

func GetSingleAccount(c *gin.Context) *AccountLoyaltyMemberRef {
	var account *AccountLoyaltyMemberRef
	conn := db.Postgres

	account_id := c.Query("account_id")
	// contact_id := c.Query("contact_id")
	// loyalty_id := c.Query("loyalty_id")

	// fmt.Printf("account_id: %s; contact_id: %s; loyalty_id: %s \n", account_id, contact_id, loyalty_id)
	fmt.Printf("account_id: %s\n", account_id)
	// if account_id == "" {
	// 	return errors.New("all mandatory values NOT Passed")
	// }

	// rows, err := conn.Query(context.Background(), "SELECT * FROM catch.account_loyalty_member_reference WHERE cl_account_id=$1 OR cl_contact_id=$2 OR cl_loyalty_id=$3 LIMIT 1", account_id, contact_id, loyalty_id)
	rows, err := conn.Query(context.Background(), "SELECT * FROM catch.account_loyalty_member_reference WHERE cl_account_id=$1 LIMIT 1", account_id)
	if err != nil {
		log.Fatal("error while executing query")
	}

	// iterate through the rows
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Fatal("error while iterating dataset")
		}

		// convert DB types to Go types
		account = &AccountLoyaltyMemberRef{
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
		log.Println("[Account_id:", account.Account_id, "]")
	}

	return account
}
