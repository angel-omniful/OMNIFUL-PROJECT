package myDb

import(
	"log"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/db/sql/postgres"

	"github.com/angel-omniful/ims/myContext"

)


func ConnectDB() *postgres.DbCluster {
	ctx := myContext.GetContext()
	myHost := config.GetString(ctx, "postgres.master.host")
	myPort := config.GetString(ctx, "postgres.master.port")
	myUsername := config.GetString(ctx, "postgres.master.username")
	myPassword := config.GetString(ctx, "postgres.master.password")
	myDbname := config.GetString(ctx, "postgres.master.dbname")
	maxOpenConns := config.GetInt(ctx, "postgres.master.max_open_conns")
	maxIdleConns := config.GetInt(ctx, "postgres.master.max_idle_conns")
	connMaxLifetime := config.GetDuration(ctx, "postgres.master.conn_max_lifetime")
	debugMode := config.GetBool(ctx, "postgres.master.debug_mode")

	masterConfig := postgres.DBConfig{
		Host:               myHost,
		Port:               myPort,
		Username:           myUsername,
		Password:           myPassword,
		Dbname:             myDbname,
		MaxOpenConnections: maxOpenConns,
		MaxIdleConnections: maxIdleConns,
		ConnMaxLifetime:  connMaxLifetime,
		DebugMode:          debugMode,
	}
	// Initialize slavesConfig as an empty slice
	// they can be added later if needed
	slavesConfig := make([]postgres.DBConfig, 0) // read replicas

	db := postgres.InitializeDBInstance(masterConfig, &slavesConfig)

	//db is a cluster here
	log.Println("Connecting to the database...")
	return db
}

// func ConnectRedis() *redis.Client {
// 	config := &redis.Config{
// 		Hosts:       []string{"localhost:6379"},
// 		PoolSize:    50,
// 		MinIdleConn: 10,
// 	}
// 	client := redis.NewClient(config)
// 	return client
// }

	
// func GetRedis() *redis.Client {
// 	rds := ConnectRedis()
// 	return rds
// }