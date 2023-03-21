package internals

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

type Migration struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

var log = logrus.New()

func New(pool *pgxpool.Pool, ctx context.Context) *Migration {
	return &Migration{
		pool: pool,
		ctx:  ctx,
	}
}

// CreateTables Создает таблицы, которые указаны в файле "sql/create-table.sql"
func (m *Migration) CreateTables() {
	sql := readSqlFile("sql/create-tables.sql")

	_, err := m.pool.Exec(m.ctx, sql)
	if err != nil {
		log.Errorln("Error execute. Error: ", err)
		return
	}

	log.Infoln("Successfully create tables")
}

// UploadDevData Загрузка тестовых данных в таблицы
func (m *Migration) UploadDevData() {
	sql := readSqlFile("sql/dev/create-dev-data.sql")

	_, err := m.pool.Exec(m.ctx, sql)
	if err != nil {
		log.Errorln("Error execute. Error: ", err)
		return
	}

	log.Infoln("Successfully upload dev data")
}

// DropTables Удаляет таблицы
func (m *Migration) DropTables() {
	sql := readSqlFile("sql/dev/drop-dev-data.sql")

	_, err := m.pool.Exec(m.ctx, sql)
	if err != nil {
		log.Errorln("Error execute. Error: ", err)
		return
	}

	log.Infoln("Successfully drop tables")
}

// Метод для чтения файла с sql запросами
func readSqlFile(sqlFilePath string) string {
	file, err := os.Open(sqlFilePath)
	if err != nil {
		log.Fatalln("Error open file. Error: ", err)
	}

	sql, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln("Error read file. Error: ", err)
	}

	defer file.Close()

	return string(sql)
}
