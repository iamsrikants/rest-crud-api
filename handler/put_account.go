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

// type AccountLoyaltyMemberRefBind struct {
// 	Account_id                int64     `json:"cl_account_id" binding:"required"`
// 	Contact_id                int64     `json:"cl_contact_id" binding:"required"`
// 	Loyalty_id                int64     `json:"cl_loyalty_id" binding:"required"`
// 	Primary_loyalty_member_id bool      `json:"cl_primary_loyalty_member_id" binding:"required"`
// 	Loyalty_type_code         string    `json:"cl_loyalty_type_code" binding:"required"`
// 	Loyalty_matched_code      string    `json:"cl_loyalty_matched_code" binding:"required"`
// 	Brand                     string    `json:"cl_brand" binding:"required"`
// 	Create_rcd_timestamp      time.Time `json:"cl_create_rcd_timestamp" binding:"required"`
// 	Create_rcd_by_who         string    `json:"cl_create_rcd_by_who" binding:"required"`
// 	Create_rcd_by_app         string    `json:"cl_create_rcd_by_app" `
// 	Update_rcd_timestamp      time.Time `json:"cl_update_rcd_timestamp" binding:"required"`
// 	Update_rcd_by_who         string    `json:"cl_update_rcd_by_who" binding:"required"`
// 	Update_rcd_by_app         string    `json:"cl_update_rcd_by_app" `
// 	Original_data_source      string    `json:"cl_original_data_source" binding:"required"`
// }

func UpdateAccount(c *gin.Context) {
	if err := UpdateSingleAccount(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "UPDATE one Record SUCCESS"})
	}
}

func UpdateSingleAccount(c *gin.Context) error {
	eval := GetSingleAccount(c)
	if eval == nil {
		return errors.New("record Not Found")
	}
	var account AccountLoyaltyMemberRefBind
	conn := db.Postgres

	if err := c.BindJSON(&account); err != nil {
		log.Printf("\n InPut Details: ERROR - bindjson() : %v\n\n", err.Error())
		return err
	}

	log.Printf("\n InPut Details: %+v \n\n", account)

	account_id := c.Query("account_id")
	// contact_id := c.Query("contact_id")
	// loyalty_id := c.Query("loyalty_id")

	// fmt.Printf("account_id: %s; contact_id: %s; loyalty_id: %s \n", account_id, contact_id, loyalty_id)
	fmt.Printf("account_id: %s\n", account_id)
	if account_id == "" {
		return errors.New("all mandatory values NOT Passed")
	}

	_, err := conn.Exec(context.Background(),
		"UPDATE catch.account_loyalty_member_reference SET \"cl_account_id\"=$1, \"cl_contact_id\"=$2, \"cl_loyalty_id\"=$3, \"cl_primary_loyalty_member_id\"=$4, \"cl_loyalty_type_code\"=$5, \"cl_loyalty_matched_code\"=$6, \"cl_brand\"=$7, \"cl_create_rcd_timestamp\"=$8, \"cl_create_rcd_by_who\"=$9, \"cl_create_rcd_by_app\"=$10, \"cl_update_rcd_timestamp\"=$11, \"cl_update_rcd_by_who\"=$12, \"cl_update_rcd_by_app\"=$13, \"cl_original_data_source\"=$14 WHERE cl_account_id=$15",
		account.Account_id, account.Contact_id, account.Loyalty_id, account.Primary_loyalty_member_id, account.Loyalty_type_code, account.Loyalty_matched_code, account.Brand, account.Create_rcd_timestamp, account.Create_rcd_by_who, account.Create_rcd_by_app, account.Update_rcd_timestamp, account.Update_rcd_by_who, account.Update_rcd_by_app, account.Original_data_source, account_id)
	if err != nil {
		log.Println("error while executing query: ", err.Error())
		return errors.New("error while executing UPDATE query")
	}

	log.Printf("\n UPDATE SUCCESS \n\n")

	return nil
}
