package main

import (
	"log"

	"github.com/labstack/echo"

	"github.com/bonaventurabs/be-tokopedia-auction/config"
	healthDeliveryHTTP "github.com/bonaventurabs/be-tokopedia-auction/delivery/http/health"
	itemDeliveryHTTP "github.com/bonaventurabs/be-tokopedia-auction/delivery/http/item"
	userDeliveryHTTP "github.com/bonaventurabs/be-tokopedia-auction/delivery/http/user"
	"github.com/bonaventurabs/be-tokopedia-auction/repository/postgre"
	"github.com/bonaventurabs/be-tokopedia-auction/usecase/health"
	"github.com/bonaventurabs/be-tokopedia-auction/usecase/item"
	"github.com/bonaventurabs/be-tokopedia-auction/usecase/user"
)

func main() {

	// read config
	cfg := &config.Configuration{}
	config.ReadModuleConfig(cfg, "./")

	e := echo.New()
	// timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	// repoHTTP := httpRepo.NewHTTP(cfg.URL)
	repoPostgre := postgre.NewDB(cfg.DB)

	// e.Use(middL.CORS)
	hUc := health.NewUsecase()
	uUc := user.NewUsecase(repoPostgre)
	iUc := item.NewUsecase(repoPostgre)

	healthDeliveryHTTP.NewHealthHTTP(e, hUc)
	userDeliveryHTTP.NewUserHTTP(e, uUc)
	itemDeliveryHTTP.NewItemHTTP(e, iUc)

	log.Fatal(e.Start(":8080"))

}
