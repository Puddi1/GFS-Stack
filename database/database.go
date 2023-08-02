package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Puddi1/GFS-Stack/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init_db(config *gorm.Config) {
	SUPABASE_DB_SSLMODE := env.ENVs["SUPABASE_DB_SSLMODE"]
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		env.ENVs["SUPABASE_DB_USER"],
		env.ENVs["SUPABASE_PROJECT_PASSWORD"],
		env.ENVs["SUPABASE_DB_HOST"],
		env.ENVs["SUPABASE_DB_PORT"],
		env.ENVs["SUPABASE_DB_NAME"],
		SUPABASE_DB_SSLMODE,
	)

	if SUPABASE_DB_SSLMODE == "verify-full" {
		cert, errRead := os.ReadFile(env.ENVs["SUPABASE_DB_SSLCERT_PATH"])
		if errRead != nil {
			log.Fatalf("Error reading cerificate: %s", errRead)
		}
		dsn = dsn + fmt.Sprintf(" sslcert=%s", cert)
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
}

// MigrateData creates a table of the data in the DB
func MigrateData(d ...any) {
	for _, d := range d {
		DB.AutoMigrate(d)
	}
}
