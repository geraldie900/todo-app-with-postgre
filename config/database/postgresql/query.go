package postgresql

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type TransactionDB struct {
	Tx         *gorm.DB
	WhereField []string
	WhereValue []interface{}
	Limit      int
	Offset     int
	OrderBy    []string // "format= `field desc` if order not specified, asc will be used"
}

func StartTransaction() *TransactionDB {
	return &TransactionDB{
		Tx: PostgreDB.Connection.Begin(),
	}
}

func Start() *TransactionDB {
	return &TransactionDB{
		Tx: PostgreDB.Connection,
	}
}

func (tx *TransactionDB) CommitTransaction() error {
	queryResult := tx.Tx.Commit()
	return queryResult.Error
}

func (tx *TransactionDB) RollbackTransaction() error {
	queryResult := tx.Tx.Rollback()
	return queryResult.Error
}

func (tx *TransactionDB) CreateData(data interface{}) error {
	queryResult := tx.Tx.Create(data)
	return queryResult.Error
}

func (tx *TransactionDB) UpdateData(data interface{}) error {
	queryResult := tx.Tx.Clauses(clause.Returning{}).Updates(data)
	return queryResult.Error
}

func (tx *TransactionDB) DeleteData(data interface{}) error {
	if len(tx.WhereField) > 0 {
		tx.Tx.Where(tx.WhereField, tx.WhereValue)
	}

	queryResult := tx.Tx.Clauses(clause.Returning{}).Delete(data)
	return queryResult.Error
}

func (tx *TransactionDB) GetList(data interface{}) *gorm.DB {
	query := PostgreDB.Connection

	for idx, v := range tx.WhereField {
		if idx == 0 {
			query = query.Where(fmt.Sprintf("%s = ?", v), tx.WhereValue[idx])
		} else {
			query = query.Or(fmt.Sprintf("%s = ?", v), tx.WhereValue[idx])
		}
	}

	if tx.Limit > 0 {
		query = query.Limit(tx.Limit)
	}

	if tx.Offset > 0 {
		query = query.Offset(tx.Offset)
	}

	if len(tx.OrderBy) > 0 {
		query = query.Order(strings.Join(tx.OrderBy, ", "))
	}

	queryResult := query.Find(data)
	return queryResult
}

func (tx *TransactionDB) Get(data interface{}) *gorm.DB {
	query := PostgreDB.Connection

	queryResult := query.First(data)
	return queryResult
}

func (tx *TransactionDB) RawQuery(queryString string) *gorm.DB {
	query := PostgreDB.Connection.Raw(queryString)
	return query
}
