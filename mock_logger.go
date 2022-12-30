package main

import (
	"log"
	"math/rand"
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	sugar := logger.Sugar()

	msgPerMinute := 30

	timevar := time.Duration(60/msgPerMinute) * time.Second

	min := 1
	max := 3

	rand.Seed(time.Now().UnixNano())

	for {
		time.Sleep(timevar)

		randomInt := min + rand.Intn(max-min+1)
		switch randomInt {
		case 1:
			sugar.Error("This is a error")
		case 2:
			sugar.Warn("This is a warning")
		case 3:
			sugar.Info("This is an info")
		}

	}
}
