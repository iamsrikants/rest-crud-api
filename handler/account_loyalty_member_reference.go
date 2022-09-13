package handler

import (
	"time"
)

type AccountLoyaltyMemberRef struct {
	Account_id                int64     `json:"cl_account_id"`
	Contact_id                int64     `json:"cl_contact_id"`
	Loyalty_id                int64     `json:"cl_loyalty_id"`
	Primary_loyalty_member_id bool      `json:"cl_primary_loyalty_member_id"`
	Loyalty_type_code         string    `json:"cl_loyalty_type_code"`
	Loyalty_matched_code      string    `json:"cl_loyalty_matched_code"`
	Brand                     string    `json:"cl_brand"`
	Create_rcd_timestamp      time.Time `json:"cl_create_rcd_timestamp"`
	Create_rcd_by_who         string    `json:"cl_create_rcd_by_who"`
	Create_rcd_by_app         string    `json:"cl_create_rcd_by_app"`
	Update_rcd_timestamp      time.Time `json:"cl_update_rcd_timestamp"`
	Update_rcd_by_who         string    `json:"cl_update_rcd_by_who"`
	Update_rcd_by_app         string    `json:"cl_update_rcd_by_app"`
	Original_data_source      string    `json:"cl_original_data_source"`
}

type AccountLoyaltyMemberRefBind struct {
	Account_id                int64     `json:"cl_account_id" binding:"required"`
	Contact_id                int64     `json:"cl_contact_id" binding:"required"`
	Loyalty_id                int64     `json:"cl_loyalty_id" binding:"required"`
	Primary_loyalty_member_id bool      `json:"cl_primary_loyalty_member_id" binding:"required"`
	Loyalty_type_code         string    `json:"cl_loyalty_type_code" binding:"required"`
	Loyalty_matched_code      string    `json:"cl_loyalty_matched_code" binding:"required"`
	Brand                     string    `json:"cl_brand" binding:"required"`
	Create_rcd_timestamp      time.Time `json:"cl_create_rcd_timestamp" binding:"required"`
	Create_rcd_by_who         string    `json:"cl_create_rcd_by_who" binding:"required"`
	Create_rcd_by_app         string    `json:"cl_create_rcd_by_app" `
	Update_rcd_timestamp      time.Time `json:"cl_update_rcd_timestamp" binding:"required"`
	Update_rcd_by_who         string    `json:"cl_update_rcd_by_who" binding:"required"`
	Update_rcd_by_app         string    `json:"cl_update_rcd_by_app" `
	Original_data_source      string    `json:"cl_original_data_source" binding:"required"`
}
