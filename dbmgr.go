package gbatis

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/mallbook/commandline"
)

var (
	_dbmgr *dbMgr
	_once  sync.Once
)

func getDBMgrInstance() *dbMgr {
	_once.Do(func() {
		_dbmgr = &dbMgr{
			dbs: make(map[string]*sql.DB),
		}
	})
	return _dbmgr
}

type dbMgr struct {
	sync.RWMutex
	defaultID string
	dbs       map[string]*sql.DB
}

func (d *dbMgr) getDB(dbName string) (db *sql.DB, ok bool) {
	d.RLock()
	defer d.RUnlock()
	db, ok = d.dbs[dbName]
	return db, ok
}

func (d *dbMgr) setDB(dbName string, db *sql.DB) {
	d.Lock()
	defer d.Unlock()
	d.dbs[dbName] = db
}

func (d *dbMgr) setDefaultID(defaultID string) {
	d.defaultID = defaultID
}

func (d *dbMgr) getDefaultID() (defaultID string) {
	return d.defaultID
}

// openDB open database
func openDB(confFile string) (err error) {
	conf, err := loadConfig(confFile)
	if err != nil {
		return fmt.Errorf("Load config fail, confFile = %s", confFile)
	}

	dbmgr := getDBMgrInstance()
	dbmgr.setDefaultID(conf.DBs.Default)

	for _, dbParams := range conf.DBs.DBList {

		if _, ok := dbmgr.getDB(dbParams.ID); ok {
			// DB exist no create
			continue
		}

		di := newDBInfo()
		if err := di.parse(dbParams.Properties); err != nil {
			continue
		}

		if !di.check() {
			continue
		}

		db, err := sql.Open(di.Driver, di.DataSource)
		if err != nil {
			log.Printf("Open database fail, err=%v, driver=%s, dbID=%s", err, di.Driver, dbParams.ID)
			continue
		}

		log.Printf("Open database success, dbID=%s", dbParams.ID)

		err = db.Ping()
		if err != nil {
			log.Printf("Ping database fail, err=%v, driver=%s, dbID=%s", err, di.Driver, dbParams.ID)
			continue
		}
		db.SetMaxOpenConns(di.MaxOpenConns)
		db.SetMaxIdleConns(di.MaxIdleConns)

		dbmgr.setDB(dbParams.ID, db)
	}

	return
}

func init() {
	p := commandline.PrefixPath()
	err := openDB(p + "/etc/conf/gbatis.xml")
	if err != nil {
		log.Println(err)
		return
	}
	// log.Println("Open database success.")
}
