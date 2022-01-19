package config

import (
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
)

type RedisConfig struct {
	Address  string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func (o *RedisConfig) Show(log *logrus.Entry) {
	root := reflect.ValueOf(o).Elem()
	rootTypeName := root.Type().Name()
	for i := 0; i < root.NumField(); i++ {
		propTypeName := root.Type().Field(i).Name
		propValue := root.Field(i).Interface()
		st := fmt.Sprintf("%s.%s=%v", rootTypeName, propTypeName, propValue)
		log.Info(st)
	}
}
