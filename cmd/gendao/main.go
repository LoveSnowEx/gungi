package main

import (
	"github.com/LoveSnowEx/gungi/internal/bootstrap"
	"github.com/LoveSnowEx/gungi/internal/infra/database"
	"github.com/LoveSnowEx/gungi/internal/infra/po"
	"gorm.io/gen"
)

func main() {
	bootstrap.SetupSlog()
	bootstrap.SetupDB()
	// Initialize the generator with configuration
	g := gen.NewGenerator(gen.Config{
		OutPath:       "../../internal/infra/dal", // output directory, default value is ./query
		ModelPkgPath:  "po",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	// Initialize a *gorm.DB instance
	db := database.Default()

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	// Generate default DAO interface for those generated structs from database
	g.ApplyBasic(
		po.Game{},
		po.User{},
	)
	// Execute the generator
	g.Execute()
}
