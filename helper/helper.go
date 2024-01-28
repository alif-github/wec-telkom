package helper

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/forgoer/openssl"
	"github.com/sirupsen/logrus"
)

type Log struct {
	Event        string
	StatusCode   int
	ResponseTime time.Duration
	Method       string
	Request      interface{}
	URL          string
	Message      string
	Response     interface{}
}

var (
	log                  = logrus.New()
	Base64DecodeString   = base64.StdEncoding.DecodeString
	OpensslAesECBDecrypt = openssl.AesECBDecrypt
	CronNew              = cron.New
)

func CreateLog(data *Log, types string) error {
	baseDir := ""

	if os.Getenv("ENV") == "preprod" || os.Getenv("ENV") == "production" {
		baseDir = "/app/"
	}

	file, err := os.OpenFile(baseDir+"logs/service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	if types == "warning" {
		log.WithFields(logrus.Fields{"event": data.Event, "status_code": data.StatusCode, "response_time": data.ResponseTime, "method": data.Method, "request": data.Request, "url": data.URL, "response": data.Response}).Warn(data.Message)
	}

	if types == "info" {
		log.WithFields(logrus.Fields{"event": data.Event, "status_code": data.StatusCode, "response_time": data.ResponseTime, "method": data.Method, "request": data.Request, "url": data.URL, "response": data.Response}).Info(data.Message)
	}

	if types == "error" {
		log.WithFields(logrus.Fields{"event": data.Event, "status_code": data.StatusCode, "response_time": data.ResponseTime, "method": data.Method, "request": data.Request, "url": data.URL, "response": data.Response}).Error(data.Message)
	}

	log.Out = os.Stdout

	return nil
}

func StringLog(level string, message string) {
	log.Out = os.Stdout
	logDir := ""

	if os.Getenv("ENV") == "production" || os.Getenv("ENV") == "preprod" {
		logDir = "/app/"
	}

	logDir += "logs/service.log"
	file, err := os.OpenFile(logDir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	if level == "info" {
		log.Info(message)
	} else if level == "warn" {
		log.Warn(message)
	} else if level == "error" {
		log.Error(message)
	}

}

func HttpLog(level, event, method string, status int, message string, responseTime time.Duration, url, response string) {
	log.Out = os.Stdout
	log_dir := ""

	if os.Getenv("ENV") == "production" || os.Getenv("ENV") == "preprod" {
		log_dir = "/app/"
	}

	file, err := os.OpenFile(log_dir+"logs/service.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr (http log)")
	}

	payload := logrus.Fields{
		"method":        method,
		"status_code":   status,
		"resposne_time": responseTime,
		"url":           url,
		"response":      response,
	}

	if level == "info" {
		log.WithFields(payload).Info(message)
	} else if level == "warn" {
		log.WithFields(payload).Warn(message)
	} else if level == "error" {
		log.WithFields(payload).Error(message)
	}

}

func GetStatus(status string) string {
	if status == "1" {
		return "published"
	} else {
		return "unpublished"
	}
}

func DateFormatter(date string) string {
	parsedDate := ""
	fmt.Println("date helper", date)

	if date != "" {
		parser, err := time.Parse(time.RFC3339, date)
		if err != nil {
			log.Println("error helper", err)
		}
		parsedDate = parser.Format("01-02-2006")
	}

	return parsedDate
}

func ChangeDatePosition(date string) string {
	if date != "" {
		split := strings.Split(date, "-")
		date = fmt.Sprintf("%s-%s-%s", split[2], split[1], split[0])
	}

	return date
}
