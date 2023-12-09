package services

import (
	"net/http"
	"net/url"
	"os"
	"os/exec"

	"github.com/kletskovg/accounting/packages/config"
	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
)

func UploadReport(report string) {
	reportFile, createError := os.CreateTemp("/tmp/", "report.*.csv")

	if createError != nil {
		message := "Cant create tmp file with report " + createError.Error()
		logger.Info(message)
		go http.Get(
			common.TelegramApiUrl + "/done/" + url.PathEscape(message),
		)
		return
	}

	reportFile.WriteString(report)

	defer reportFile.Close()
	bucketName := config.GetEnvVariable(config.ACC_AWS_BUCKET)
	uploadResult := exec.Command("aws", "s3", "--endpoint-url=https://storage.yandexcloud.net", "cp", reportFile.Name(), "s3://", bucketName+reportFile.Name())

	if uploadError := uploadResult.Run(); uploadError != nil {
		logger.Info("Cant upload file to S3 ", uploadError)
		return
	}

	http.Get(
		common.TelegramApiUrl + "/done/" + url.PathEscape("https://storage.yandexcloud.net/"+bucketName+reportFile.Name()),
	)
}
