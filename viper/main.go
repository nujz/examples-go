package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("test1")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(viper.Get("mysql.addr"))
	fmt.Println(viper.GetInt("redis.db"))
	fmt.Println(viper.Get("aaa"))
}
