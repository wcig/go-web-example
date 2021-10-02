package db

import "github.com/go-xorm/xorm"

type DBConfig struct {
	Primary    string                   `yaml:"primary" json:"primary"`
	Datasource map[string]*DbItemConfig `yaml:"datasource" json:"datasource"`
}

type DbItemConfig struct {
	DriveName string `yaml:"drive_name" json:"drive_name"`
	Url       string `yaml:"url" json:"url"`
	MaxIdle   int    `yaml:"max_idle" json:"max_idle"`
	MaxOpen   int    `yaml:"max_open" json:"max_open"`
	ShowSql   bool   `json:"show_sql" json:"show_sql"`
}

var (
	dbMap   = make(map[string]*xorm.Engine)
	primary string
)

func Init(cfg *DBConfig) {
	// check primary
	primary = cfg.Primary
	if primary == "" {
		panic("db config primary empty")
	}
	if _, ok := cfg.Datasource[primary]; !ok {
		panic("db config datasource primary empty")
	}

	// init engine
	for k, v := range cfg.Datasource {
		engine, err := xorm.NewEngine(v.DriveName, v.Url)
		if err != nil {
			panic(err)
		}

		engine.SetMaxOpenConns(v.MaxOpen)
		engine.SetMaxIdleConns(v.MaxIdle)
		engine.ShowSQL(v.ShowSql)
		dbMap[k] = engine
	}

	// check connection
	for _, v := range dbMap {
		if err := v.Ping(); err != nil {
			panic(err)
		}
	}
}

func GetDB(key ...string) *xorm.Engine {
	if len(key) == 0 {
		return dbMap[primary]
	}
	return dbMap[key[0]]
}
