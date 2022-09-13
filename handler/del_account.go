package handler

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iamsrikants/rest-curd-api/startup/db"
)

func DeleteAccount(c *gin.Context) {
	if err := DeleteSingleAccount(c); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "DELETE record SUCCESS"})
	}
}

func DeleteSingleAccount(c *gin.Context) error {
	// var accounts []AccountLoyaltyMemberRef
	// var account *AccountLoyaltyMemberRef
	conn := db.Postgres

	account_id := c.Query("account_id")
	// contact_id := c.Query("contact_id")
	// loyalty_id := c.Query("loyalty_id")

	// fmt.Printf("account_id: %s; contact_id: %s; loyalty_id: %s \n", account_id, contact_id, loyalty_id)
	fmt.Printf("account_id: %s\n", account_id)

	if account_id == "" {
		return errors.New("all mandatory values NOT Passed")
	}

	// _, err := conn.Exec(context.Background(), "DELETE FROM catch.account_loyalty_member_reference WHERE cl_account_id=$1", account_id, contact_id, loyalty_id)
	_, err := conn.Exec(context.Background(), "DELETE FROM catch.account_loyalty_member_reference WHERE cl_account_id=$1", account_id)
	if err != nil {
		log.Println("error while executing query: ", err.Error())
		return errors.New("error while executing DELETE query")
	}

	// iterate through the rows
	// for rows.Next() {
	// 	values, err := rows.Values()
	// 	if err != nil {
	// 		log.Fatal("error while iterating dataset")
	// 	}

	return err
}
