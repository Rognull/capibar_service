package cfg

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Cfg struct { 
	Port   string
	DbName string
	DbUser string
	DbPass string
	DbHost string
	DbPort string
}

func LoadAndStoreConfig() Cfg {
	v := viper.New()

	// v.SetEnvPrefix("SERV") 
	v.SetDefault("PORT", "8080") //ставим умолчальные настройки
	// v.SetDefault("DBUSER", "postgres")
	// v.SetDefault("DBPASS", "tmppas")
	// v.SetDefault("DBHOST", "")
	// v.SetDefault("DBPORT", "")
	// v.SetDefault("DBNAME", "capy")
	v.AutomaticEnv()

	var cfg Cfg

	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Panic(err)
	}
	return cfg
}

func (cfg *Cfg) GetDBString() string { 

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("SERV_DBUSER"), os.Getenv("SERV_DBPASS"), os.Getenv("SERV_DBHOST"), os.Getenv("SERV_DBPORT"), os.Getenv("SERV_DBNAME"))

}
