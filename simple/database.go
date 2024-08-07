package simple

type Database struct {
	Name string
}

type DatabasePostgres Database
type DatabaseMysql Database

func NewDatabaseMysql() *DatabaseMysql {
	database := &Database{Name: "Mysql"}

	return (*DatabaseMysql)(database)
}

func NewDatabasePostgres() *DatabasePostgres {
	database := &Database{Name: "Postgresql"}

	return (*DatabasePostgres)(database)
}

type DatabaseRepository struct {
	DatabasePostgresql *DatabasePostgres
	DatabaseMysql      *DatabaseMysql
}

func NewDatabaseRepository(databasePostgresql *DatabasePostgres, databaseMysql *DatabaseMysql) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgresql: databasePostgresql, DatabaseMysql: databaseMysql}
}
