package services

import (
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/kletskovg/accounting/packages/config"
	"github.com/kletskovg/accounting/packages/logger"
	"github.com/kletskovg/packages/common"
)

func UploadReport(report, hostname string) {
	reportFile, createError := os.CreateTemp("/tmp/", "report.*.csv")

	if createError != nil {
		message := "Cant create tmp file with report " + createError.Error()
		logger.Info(message)
		go http.Get(
			hostname + "/done/" + url.PathEscape(message),
		)
		return
	}

	reportFile.WriteString(report)

	defer reportFile.Close()
	bucketName := strings.TrimSpace(config.GetEnvVariable(config.ACC_AWS_BUCKET))
	uploadResult := exec.Command("aws", "s3", "--endpoint-url="+common.Hosts().StorageAPIURL, "cp", reportFile.Name(), "s3://"+bucketName+reportFile.Name())

	if uploadError := uploadResult.Run(); uploadError != nil {
		logger.Info("Cant upload file to S3 ", uploadError)
		return
	}

	requestURL := hostname + "/done/" + url.PathEscape(common.Hosts().StorageAPIURL+"/"+bucketName+reportFile.Name())
	_, requestError := http.Get(requestURL)

	if requestError != nil {
		logger.Info("Cant send report to", hostname, " ", requestError)
	}
}
