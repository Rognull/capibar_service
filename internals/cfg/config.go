package cfg

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Cfg struct { //наша структура для хранения конфигов, я полагаюсь на Viper для матчинга имен
	Port   string
	DbName string
	DbUser string
	DbPass string
	DbHost string
	DbPort string
}

func LoadAndStoreConfig() Cfg {
	v := viper.New() //создаем экземпляр нашего ридера для Env

	// v.SetEnvPrefix("SERV") //префикс, все переменные нашего сервера должны теперь стартовать с SERV_ для того, чтобы не смешиваться с системными
	v.SetDefault("PORT", "8080") //ставим умолчальные настройки
	// v.SetDefault("DBUSER", "postgres")
	// v.SetDefault("DBPASS", "tmppass")
	// v.SetDefault("DBHOST", "195.133.197.62")
	// v.SetDefault("DBPORT", "3030")
	// v.SetDefault("DBNAME", "capybaras")
	v.AutomaticEnv() //собираем наши переменные с системных

	var cfg Cfg

	err := v.Unmarshal(&cfg) //закидываем переменные в cfg после анмаршалинга
	if err != nil {
		log.Panic(err)
	}
	return cfg
}

func (cfg *Cfg) GetDBString() string { //маленький метод для сборки строки соединения с БД
	// return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("SERV_DBUSER"), os.Getenv("SERV_DBPASS"), os.Getenv("SERV_DBHOST"), os.Getenv("SERV_DBPORT"), os.Getenv("SERV_DBNAME"))

	// cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
}
