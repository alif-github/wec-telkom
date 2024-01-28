package main

import (
	"content-management/helper"
	http_x "content-management/internal/delivery/http"
	http_listpoin "content-management/internal/delivery/http/listpoin"
	http_shopsbanner "content-management/internal/delivery/http/shopspage/banner"
	http_shopsbundlingcard "content-management/internal/delivery/http/shopspage/bundlingcard"
	http_shopsdevicebrand "content-management/internal/delivery/http/shopspage/devicebrand"
	http_shopsexpressinput "content-management/internal/delivery/http/shopspage/expressinput"
	http_shopsinternetpackage "content-management/internal/delivery/http/shopspage/internetpackage"
	http_shopspaymentmethod "content-management/internal/delivery/http/shopspage/paymentmethod"
	service_subscriber "content-management/internal/infrastructure/service/subscriber"
	"strconv"

	"github.com/robfig/cron/v3"

	// http_dp "content-management/internal/delivery/http/digital-product"
	repository_atl "content-management/internal/repository/atl/mysql"
	repository_categorypoin "content-management/internal/repository/categorypoin/mysql"
	repository_dp "content-management/internal/repository/digital-product/mysql"
	repository_listpoin "content-management/internal/repository/listpoin/mysql"
	repository_promo "content-management/internal/repository/promo/mysql"
	repository_shopsbanner "content-management/internal/repository/shopspage/banner"
	repository_shopsbundlingcard "content-management/internal/repository/shopspage/bundlingcard"
	repository_shopsdevicebrand "content-management/internal/repository/shopspage/devicebrand"
	repository_shopsexpressinput "content-management/internal/repository/shopspage/expressinput"
	repository_shopsinternetpackage "content-management/internal/repository/shopspage/internetpackage"
	repository_shopspaymentmethod "content-management/internal/repository/shopspage/paymentmethod"

	usecase_atl "content-management/internal/usecase/atl"
	usecase_dp "content-management/internal/usecase/digital-product"
	usecase_listpoin "content-management/internal/usecase/listpoin"
	usecase_promo "content-management/internal/usecase/promo"
	usecase_shopsbanner "content-management/internal/usecase/shopspage/banner"
	usecase_shopsbundlingcard "content-management/internal/usecase/shopspage/bundlingcard"
	usecase_shopsdevicebrand "content-management/internal/usecase/shopspage/devicebrand"
	usecase_shopsexpressinput "content-management/internal/usecase/shopspage/expressinput"
	usecase_shopsinternetpackage "content-management/internal/usecase/shopspage/internetpackage"
	usecase_shopspaymentmethod "content-management/internal/usecase/shopspage/paymentmethod"
	"crypto/tls"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/newrelic/go-agent/v3/integrations/nrlogrus"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

// var logg = logrus.New()

var newtrans http.RoundTripper = &http.Transport{
	Proxy: http.ProxyFromEnvironment,
}

const (
	MaxIdleConns       int  = 100
	MaxIdleConnections int  = 100
	RequestTimeout     int  = 30
	SSL                bool = true
)

func createHTTPClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: SSL},
		MaxIdleConns:        MaxIdleConns,
		MaxIdleConnsPerHost: MaxIdleConnections,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(RequestTimeout) * time.Second,
	}
	return client
}

func main() {
	log.Println("Telkomsel WEC API CMS CONSUMER")
	err := godotenv.Load(os.Getenv("BASE_DIR") + ".env")

	if err != nil {
		log.Fatalf("Error load .env file")
	}

	dbUrl := os.Getenv("DB_URL")
	db, err := sql.Open("mysql", dbUrl)

	if err != nil {
		fmt.Println("Failed to connect database")
		helper.StringLog("error", err.Error())
	}
	defer db.Close()

	handlerClient := createHTTPClient()
	fmt.Println(handlerClient)

	subscriberService := service_subscriber.NewSubscriberTypeService(handlerClient)

	atlRepository := repository_atl.NewATLPackageRepository(db)
	atlUseCase := usecase_atl.NewATLPackageUseCase(atlRepository)
	atlHttpDelivery := http_x.NewATLPackageHttpDelivery(atlUseCase)

	dpRepository := repository_dp.NewDigitalProductRepository(db)
	dpUsecase := usecase_dp.NewDigitalProductUseCase(dpRepository)
	dpHttpDelivery := http_x.NewDigitalProductDelivery(dpUsecase)

	promoRepository := repository_promo.NewPromoRepository(db)
	listpoinRepository := repository_listpoin.NewListPoinRepository(db)
	categoryRepository := repository_categorypoin.NewCategoryRepository(db)
	shopspageBannerRepository := repository_shopsbanner.NewShopspageBannerRepository(db)
	shopspageBundlingCardRepository := repository_shopsbundlingcard.NewShopspageBundlingCardRepository(db)
	shopspageDeviceBrandRepository := repository_shopsdevicebrand.NewShopspageDeviceBrandRepository(db)
	shopspageExpressInputRepository := repository_shopsexpressinput.NewShopspageExpressInputRepository(db)
	shopspageInternetPackageRepository := repository_shopsinternetpackage.NewShopspageInternetPackageRepository(db)
	shopspagePaymentMethodRepository := repository_shopspaymentmethod.NewShopspagePaymentMethodRepository(db)

	promoUseCase := usecase_promo.NewPromoUseCase(promoRepository)
	listpoinUsecase := usecase_listpoin.NewListPoinUsecase(listpoinRepository, categoryRepository, subscriberService)
	shopspageBannerUsecase := usecase_shopsbanner.NewShopspageBannerUsecase(shopspageBannerRepository)
	shopspageBundlingCardUsecase := usecase_shopsbundlingcard.NewShopspageBundlingCardUsecase(shopspageBundlingCardRepository)
	shopspageDeviceBrandUsecase := usecase_shopsdevicebrand.NewShopspageDeviceBrandUsecase(shopspageDeviceBrandRepository)
	shopspageExpressInputUsecase := usecase_shopsexpressinput.NewShopspageDeviceBrandUsecase(shopspageExpressInputRepository)
	shopspageInternetPackageUsecase := usecase_shopsinternetpackage.NewShopspageInternetPackageUsecase(shopspageInternetPackageRepository)
	shopspagePaymentMethodUsecase := usecase_shopspaymentmethod.NewShopspagePaymentMethodUsecase(shopspagePaymentMethodRepository)

	promoHttpDelivery := http_x.NewPromoHttpDelivery(promoUseCase)
	ListPoinHttpDelivery := http_listpoin.NewListPoinHttpDelivery(listpoinUsecase)
	shopspageBannerHttpDelivery := http_shopsbanner.NewShopspageBannerHttpDelivery(shopspageBannerUsecase)
	shopspageBundlingCardHttpDelivery := http_shopsbundlingcard.NewShopspageBundlingCardHttpDelivery(shopspageBundlingCardUsecase)
	shopspageDeviceBrandHttpDelivery := http_shopsdevicebrand.NewShopspageDeviceBrandHttpDelivery(shopspageDeviceBrandUsecase)
	shopspageExpressInputHttpDelivery := http_shopsexpressinput.NewShopspageExpressInputHttpDelivery(shopspageExpressInputUsecase)
	shopspageInternetPackageHttpDelivery := http_shopsinternetpackage.NewShopspageInternetPackageHttpDelivery(shopspageInternetPackageUsecase)
	shopspagePaymentMethodHttpDelivery := http_shopspaymentmethod.NewShopspagePaymentMethodHttpDelivery(shopspagePaymentMethodUsecase)

	if os.Getenv("ENABLE_CRON") == "true" {
		cronTime := os.Getenv("CRON_TIME")
		Cron(cronTime, listpoinUsecase)
	}

	router := mux.NewRouter()

	logrusLevel := LogLevel(os.Getenv("LOGRUS_LEVEL"))

	app, _ := newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("NEWRELIC_CONF_NAME")),
		newrelic.ConfigLicense(os.Getenv("NEWRELIC_CONF_LICENSE")),
		newrelic.ConfigDistributedTracerEnabled(true),
		func(config *newrelic.Config) {
			config.Enabled = false
			config.HostDisplayName = os.Getenv("NEWRELIC_HOSTNAME")
			config.Transport = newtrans
			logrus.SetLevel(logrus.Level(logrusLevel))
			config.Logger = nrlogrus.StandardLogger()
		},
	)

	router.HandleFunc(newrelic.WrapHandleFunc(app, "/manage-content/promo", promoHttpDelivery.GetPromo)).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/manage-content/promo/list-poin", ListPoinHttpDelivery.GetListPoin)).Methods("POST")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/manage-content/shops-page/banner", shopspageBannerHttpDelivery.GetShopspageBanners)).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/manage-content/shops-page/bundling-card", shopspageBundlingCardHttpDelivery.GetShopspageBundlingCard)).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/manage-content/shops-page/device-brand", shopspageDeviceBrandHttpDelivery.GetShopspageDeviceBrands)).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/manage-content/shops-page/express-input", shopspageExpressInputHttpDelivery.GetShopspageExpressInput)).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/manage-content/shops-page/payment-method", shopspagePaymentMethodHttpDelivery.GetShopspagePaymentMethods)).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/manage-content/shops-page/postpaid", shopspageInternetPackageHttpDelivery.GetShopspagePostpaid)).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/manage-content/shops-page/prepaid", shopspageInternetPackageHttpDelivery.GetShopspagePrepaid)).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/package/atl", atlHttpDelivery.GetPackage)).Methods("POST")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/package/atl/{id}", atlHttpDelivery.FindPackage)).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/region", atlHttpDelivery.GetRegion)).Methods("GET")

	router.HandleFunc(newrelic.WrapHandleFunc(app, "/digital-product/credit/{creditid}", dpHttpDelivery.FindCredit)).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/digital-product/credit-offer", dpHttpDelivery.GetCredit)).Methods("GET")
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/digital-product/stock/{id}", dpHttpDelivery.ReduceStock)).Methods("PATCH")

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	ssl, _ := strconv.ParseBool(os.Getenv("USE_SSL"))
	log.Printf("Service is running on port: %s \n", port)
	if os.Getenv("ENV") == "production" || os.Getenv("ENV") == "preproduction" {
		if ssl {
			http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: ssl}
			log.Println("Running on TLS server...")
			err := http.ListenAndServeTLS(port, os.Getenv("BASE_DIR")+os.Getenv("PATH_SSL_CERT"), os.Getenv("BASE_DIR")+os.Getenv("PATH_SSL_KEY"), router)
			if err != nil {
				log.Printf("Error ListenAndServeTLS: %s\n", err)
			}
		} else {
			log.Println("Running on non-TLS server...")
			err := http.ListenAndServe(port, router)
			if err != nil {
				log.Printf("Error ListenAndServe: %s\n", err)
			}
		}
	}
	select {}
}

func LogLevel(lvl string) logrus.Level {

	switch lvl {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	default:
		panic(fmt.Sprintf("the specified %s log level is not supported", lvl))
	}
}

func Cron(cronTime string, listPointUsecase *usecase_listpoin.ListPoinUsecase) {
	localTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := helper.CronNew(cron.WithLocation(localTime))

	defer scheduler.Stop()

	scheduler.AddFunc(cronTime, AutomateCheckKeyword(listPointUsecase))

	go scheduler.Start()
}

func AutomateCheckKeyword(listPointUsecase *usecase_listpoin.ListPoinUsecase) func() {
	return func() {
		err := listPointUsecase.DeleteKeywordAutomate()

		if err != nil {
			fmt.Println("failed check delete keyword")
		} else {
			fmt.Print("success check delete keyword")

		}
	}
}
