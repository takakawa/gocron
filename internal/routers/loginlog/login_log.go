package loginlog

import (
	"github.com/takakawa/gocron/internal/models"
	"github.com/takakawa/gocron/internal/modules/logger"
	"github.com/takakawa/gocron/internal/modules/utils"
	"github.com/takakawa/gocron/internal/routers/base"
	"gopkg.in/macaron.v1"
)

func Index(ctx *macaron.Context) string {
	loginLogModel := new(models.LoginLog)
	params := models.CommonMap{}
	base.ParsePageAndPageSize(ctx, params)
	total, err := loginLogModel.Total()
	loginLogs, err := loginLogModel.List(params)
	if err != nil {
		logger.Error(err)
	}

	jsonResp := utils.JsonResponse{}

	return jsonResp.Success(utils.SuccessContent, map[string]interface{}{
		"total": total,
		"data":  loginLogs,
	})
}
