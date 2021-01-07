package day6_transaction

import (
	"database/sql"
	"go_basic/7days/gee-orm/day6-transaction/dialect"
	"go_basic/7days/gee-orm/day6-transaction/log"
	"go_basic/7days/gee-orm/day6-transaction/session"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	dialect, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s NOT FOUND", driver)
		return
	}

	e = &Engine{
		db:      db,
		dialect: dialect,
	}
	log.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("failed to close database")
	}
	log.Info("Close database success")
}

func (engine *Engine) NewSession() *session.Session {
	return session.New(engine.db, engine.dialect)
}

type TxFunc func(*session.Session) (interface{}, error)

func (engine *Engine) Transaction(f TxFunc) (result interface{}, err error) {
	s := engine.NewSession()
	if err := s.Begin(); err != nil {
		return nil, err
	}

	//todo 这个地方的事务处理很巧妙，要吸收要吸收
	defer func() {
		if p := recover(); p != nil {
			_ = s.Rollback()
			panic(p)
		} else if err != nil {
			_ = s.Rollback()
		} else {
			err = s.Commit()
		}
	}()
	return f(s)
}
