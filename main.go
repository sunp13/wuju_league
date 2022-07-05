package main

import (
	"io/ioutil"
	"time"
	"wuju_league/entity"
	"wuju_league/initlize"
	"wuju_league/services"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

var (
	conf entity.Conf
)

func init() {
	if err := initlize.InitAll(); err != nil {
		panic(err)
	}

	confByte, err := ioutil.ReadFile("./conf/conf.yml")
	if err != nil {
		log.Fatal().Str("err", err.Error()).Send()
	}

	if err = yaml.Unmarshal(confByte, &conf); err != nil {
		log.Fatal().Str("err", err.Error()).Send()
	}
}

func main() {
	for {
		for i := 1; i < 50; i++ {
			end, err := services.UpComingService.GetData(&conf, i)
			if err != nil {
				log.Error().Str("err", err.Error()).Send()
				continue
			}
			if end {
				time.Sleep(time.Duration(conf.UpComingInterval) * time.Second)
				break
			}
		}
	}
}
