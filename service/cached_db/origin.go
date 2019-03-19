package data

import (
	"github.com/qinhan-shu/gp-server/service/cached_db/driver/mysql"
)

type originDB struct {
	conn *driver.Pool
}

func (p *originDB) Set(document string, data Data, where Data) error {
	if err := getCondition(document, data, where); err != nil {
		return err
	}
	return driver.Set(p.conn, document, data, where)
}

func (p *originDB) Get(document string, column []string, where Data) (Data, error) {
	if err := setCondition(document, column, where); err != nil {
		return nil, err
	}
	return driver.Get(p.conn, document, column, where)
}

func (p *originDB) Inc(document string, column []string, where Data) error {
	if err := setCondition(document, column, where); err != nil {
		return err
	}
	return driver.Inc(p.conn, document, column, where)
}

func (p *originDB) GetAll(document string, column []string, where Data) ([]Data, error) {
	return driver.GetAll(p.conn, document, column, where)
}