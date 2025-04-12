package sqlite_impl

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (db SqliteDB) sqliteCommit(functionName string, request string, paramsInOrder ...any) error {
	req, err := db.DB.Prepare(request)
	if err != nil {
		return sqliteError(fmt.Errorf("%s sqlite req error. %s", functionName, err.Error()))
	}
	defer req.Close();

	_, err = req.Exec(paramsInOrder...)
	if err != nil {
		return sqliteError(fmt.Errorf("%s sqlite req exec error. %s", functionName, err.Error()))
	}

	log.Debugf("%s sqlite: OK", functionName)

	return nil
}


func (db SqliteDB) sqliteTransaction(functionName string, request string, paramsInOrder ...any) error {
	tx, err := db.Begin()
	if err != nil {
		return sqliteError(fmt.Errorf("transaction error"))
	}

	req, err := tx.Prepare(request)
	if err != nil {
		return sqliteError(fmt.Errorf("[Transaction] %s sqlite req error. %s", functionName, err.Error()))
	}
	defer req.Close();

	_, err = req.Exec(paramsInOrder...)
	if err != nil {
		tx.Rollback()
		return sqliteError(fmt.Errorf("[Transaction] %s sqlite req exec error. %s", functionName, err.Error()))
	}

	err = tx.Commit()
	if err != nil {
		return sqliteError(fmt.Errorf("[Transaction] %s sqlite req commit error. %s", functionName, err.Error()))
	}

	log.Debug("Transaction completed.")
	return nil
}

func sqliteError(err error) error {
	log.Error(err.Error())
	return err
}

func (db SqliteDB) createQuery(format string, args ...any) (*sql.Rows, error) {
	query := fmt.Sprintf(format, args...)
	log.Debug(query)
	return db.Query(query)
}