package main

import (
	"context"
	"rental/app/router"
	"rental/connection"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.Default()
	ctx := context.Background()
	log := logrus.NewEntry(logrus.StandardLogger())
	db := connection.ConnectionDB(ctx, log)

	rh := &router.Handlers{
		Ctx: ctx,
		DB:  db,
		R:   r,
		Log: log,
	}

	rh.Routes()

	r.Run()

}
