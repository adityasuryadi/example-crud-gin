package controllers

import (
	"app/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReportMerchant struct {
	MrId          int64  `json:"mr_id" gorm:"column:mr_id"`
	MerchantName  string `json:"merchant_name" gorm:"column:merchant_name"`
	TotalOnline   int64  `gorm:"column:total_online" json:"total_online"`
	TotalOnsite   int64  `gorm:"column:total_onsite" json:"total_onsite"`
	TotalSuccess  int64  `gorm:"column:total_success" json:"total_success"`
	TotalAbsent   int64  `gorm:"column:total_absent" json:"total_absent"`
	TotalNoAction int64  `gorm:"column:total_no_action" json:"total_no_action"`
	TotalOther    int64  `gorm:"column:total_other" json:"total_other"`
}

type MerchantReportResponse struct {
	Id                 int64  `gorm:"column:mr_id" json:"merchant_id"`
	MerchantName       string `json:"merchant_name"`
	KdMerchant         string `gorm:"column:kd_merchant" json:"merchant_code"`
	TotalAllTicket     int64  `gorm:"column:all_ticket" json:"total_all_ticket"`
	TotalSuccesTicket  int64  `gorm:"column:success_ticket" json:"total_success_ticket"`
	TotalWaitingTicket int64  `gorm:"column:waiting_ticket" json:"total_waiting_ticket"`
}

func Test(ctx *gin.Context) {
	db := configs.InitDB()
	var result []ReportMerchant
	err := db.Raw(`select merchant.mr_id,total_online,total_onsite,total_success,total_absent,total_no_action,total_other,merchant_name from merchant
	RIGHT JOIN (
		select mr_id,count(qu_id) as total_online from queue where (fromApp = 5 OR 4) AND DATE(created) = DATE(NOW()) group By mr_id
	) as report_online on merchant.mr_id = report_online.mr_id
	RIGHT JOIN (
		select mr_id,count(qu_id) as total_onsite from queue where fromApp = 1 OR fromApp = 2 OR fromApp =3 group By mr_id
	) as report_onsite on merchant.mr_id = report_onsite.mr_id
	RIGHT JOIN (
		select mr_id,count(qu_id) as total_success from queue where status = 2 group By mr_id
	) as report_success on merchant.mr_id = report_success.mr_id
	RIGHT JOIN (
		select mr_id,count(qu_id) as total_absent from queue where status = 3 group By mr_id
	) as report_absent on merchant.mr_id = report_absent.mr_id
	RIGHT JOIN (
		select mr_id,count(qu_id) as total_no_action from queue where status = 0 OR status = 5 OR status = 8 group By mr_id
	) as report_no_action on merchant.mr_id = report_no_action.mr_id
	RIGHT JOIN (
		select mr_id,count(qu_id) as total_other from queue where status = 1 OR status = 4 OR status = 6 OR status = 7 group By mr_id
	) as report_other on merchant.mr_id = report_other.mr_id
	where isActive = 1
	`).Scan(&result).Error

	// var merchant []MerchantReportResponse
	// // centerId := ct
	// db.Raw(`
	// select merchant.mr_id,merchant.merchant_name,merchant.kd_merchant,
	// success_ticket,all_ticket,(all_ticket - success_ticket) as waiting_ticket
	// from merchant
	// LEFT JOIN (select mr_id,status,count(sr_id) as success_ticket from queue where status = 2 AND DATE(created) = DATE(NOW()) group by mr_id) as total_ticket_success
	// ON total_ticket_success.mr_id = merchant.mr_id
	// LEFT JOIN (select mr_id,status,count(sr_id) as all_ticket from queue WHERE DATE(created) = DATE(NOW()) group by mr_id) as total_ticket_all
	// ON total_ticket_all.mr_id = merchant.mr_id
	// WHERE centerId = ?
	// `, 1909).Scan(&merchant)

	// err := db.Raw("select merchant.mr_id,merchant_name from merchant").Scan(&result).Error

	if err != nil {
		panic(result)
	}

	var responses []ReportMerchant
	for _, merchant := range result {
		responses = append(responses, ReportMerchant{
			MrId:          merchant.MrId,
			MerchantName:  merchant.MerchantName,
			TotalOnline:   merchant.TotalOnline,
			TotalOnsite:   merchant.TotalOnsite,
			TotalSuccess:  merchant.TotalSuccess,
			TotalAbsent:   merchant.TotalAbsent,
			TotalNoAction: merchant.TotalNoAction,
			TotalOther:    merchant.TotalOther,
		})
	}

	// fmt.Println(responses)

	// ctx.JSON(http.StatusOK, gin.H{"data": true})
	ctx.JSON(http.StatusOK, gin.H{"data": responses})

}
