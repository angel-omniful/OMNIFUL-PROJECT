package main

import (
	
	"log"
	"time"
	"github.com/omniful/go_commons/http"
	"github.com/angel-omniful/ims/routes"
	//"github.com/angel-omniful/ims/myContext"
    _"github.com/angel-omniful/ims/myDb"
	//"github.com/omniful/go_commons/db/sql/migration"
	//"github.com/omniful/go_commons/config"
	

	
)



func main() {

	log.Println("Connecting to the server...")
	server:= http.InitializeServer(
	":8080",                // listen address
	10 * time.Second,       // read timeout
	10 * time.Second,       // write timeout
	70 * time.Second,       // idle timeout
	false,
	// Add custom middleware here...
	)
	log.Println("Registering routes...")
	routes.RegisterAllRoutes(server)
	log.Println("Routes registered successfully!")
	// Start the server
	


	//1.context setup and config loading happening in myContext package
	//2.server setup and initialization happening in server package
	//3.database connection and redis connection happening in myDb package
	// Initialize the database connection
	// log.Println("Connecting to the database...")
	// db:=myDb.ConnectDB()
	// if db==nil {
	// 	log.Panic("failed to connect")
	// }
	// log.Println("Database connected successfully")
	// ctx:=myContext.GetContext()
    // //  // Build the database URL
	// log.Println("Starting procedure for migrations...")
	// myHost := config.GetString(ctx, "postgres.master.host")
	// myPort := config.GetString(ctx, "postgres.master.port")
	// myUsername := config.GetString(ctx, "postgres.master.username")
	// myPassword := config.GetString(ctx, "postgres.master.password")
	// myDbname := config.GetString(ctx, "postgres.master.dbname")
    
	// dbURL := migration.BuildSQLDBURL(
    //     myHost,
    //     myPort,
    //     myDbname,
    //     myUsername,
    //     myPassword,
    // )
    // log.Println("Database URL:", dbURL)
    // // Path to migration files
    // migrationPath :="file://C:/Users/angel/OneDrive/Desktop/OMNIFUL-PROJECT/ims/migrations"
    
    // // Initialize the migrator
    // migrator, err := migration.InitializeMigrate(migrationPath, dbURL)
    // if err != nil {
    //     log.Fatalf("Failed to initialize migrator: %v", err)
    // }
    
    // // Apply all pending migrations
    // err = migrator.Up()
    // if err != nil {
    //     log.Fatalf("Failed to apply migrations: %v", err)
    // }
	// log.Println("Migrations applied successfully")

	// redis_db := myDb.ConnectRedis()
	// if redis_db == nil {
	// 	log.Panic("failed to connect to redis")
	// }
	// log.Println("Redis connected successfully")
	// fmt.Println("Starting the IMS service...")
        
	// cache:= myDb.ConnectRedis()
	// if cache == nil{
	// 	log.Panic("failed to connect to redis")
	// }
	// log.Println("Redis connected successfully")

	// _,err:=cache.Set(ctx,"service","ims",0)
	// if err != nil {
	// 	log.Println("Error setting value in Redis:", err)
	// } else {
	// 	log.Println("Value set in Redis successfully")
	// 	val,_:= cache.Get(ctx,"service")
	// 	log.Println("Value from Redis:", val)
	// }
	
	//routes.InitializeRoutes() // Initialize the routes for the IMS service

	
   if err := server.StartServer("my-service"); err != nil {
		log.Println("server not connected")
     }else{
		log.Println("server connected successfully")
	}

	
}

