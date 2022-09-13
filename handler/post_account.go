package handler

import (
	"context"
	"errors"
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

func InsertAccount(c *gin.Context) {
	if err := InsertSingleAccount(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "INSERT one Record SUCCESS"})
	}
}

// router.POST("/loginJSON", func(c *gin.Context) {
//     var json Login
//     if err := c.ShouldBindJSON(&json); err != nil {
//       c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//       return
//     }

//     if json.User != "manu" || json.Password != "123" {
//       c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
//       return
//     }

//     c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
//   })

func InsertSingleAccount(c *gin.Context) error {
	var account AccountLoyaltyMemberRefBind
	conn := db.Postgres

	if err := c.ShouldBindJSON(&account); err != nil {
		log.Printf("\n InPut Details: ERROR - shouldbindjson() : %v\n\n", err.Error())
		return err
	}

	log.Printf("\n InPut Details: %+v \n\n", account)

	_, err := conn.Exec(context.Background(),
		"INSERT INTO catch.account_loyalty_member_reference (\"cl_account_id\", \"cl_contact_id\", \"cl_loyalty_id\", \"cl_primary_loyalty_member_id\", \"cl_loyalty_type_code\", \"cl_loyalty_matched_code\", \"cl_brand\", \"cl_create_rcd_timestamp\", \"cl_create_rcd_by_who\", \"cl_create_rcd_by_app\", \"cl_update_rcd_timestamp\", \"cl_update_rcd_by_who\", \"cl_update_rcd_by_app\", \"cl_original_data_source\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)",
		account.Account_id, account.Contact_id, account.Loyalty_id, account.Primary_loyalty_member_id, account.Loyalty_type_code, account.Loyalty_matched_code, account.Brand, account.Create_rcd_timestamp, account.Create_rcd_by_who, account.Create_rcd_by_app, account.Update_rcd_timestamp, account.Update_rcd_by_who, account.Update_rcd_by_app, account.Original_data_source)

	if err != nil {
		log.Println("error while executing query: ", err.Error())
		return errors.New("error while executing INSERT query")
	}

	log.Printf("\n INSERT SUCCESS \n\n")

	return nil
}
