package config

import (
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
)

type ApplicationConfig struct {
	IDLen             int    `mapstructure:"idLen"`
	Domain            string `mapstructure:"domain"`
	DefaultExpireDays int    `mapstructure:"defaultExpireDays"`
}

func (o *ApplicationConfig) Show(log *logrus.Entry) {
	root := reflect.ValueOf(o).Elem()
	rootTypeName := root.Type().Name()
	for i := 0; i < root.NumField(); i++ {
		propTypeName := root.Type().Field(i).Name
		propValue := root.Field(i).Interface()
		st := fmt.Sprintf("%s.%s=%v", rootTypeName, propTypeName, propValue)
		log.Info(st)
	}
}
