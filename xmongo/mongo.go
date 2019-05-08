package xmongo

import (
	"errors"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/nickxb/pkg/xsync"
	"sync"
	"time"
)

var (
	mongoSessionsLock = new(sync.RWMutex)
	mongoSessions     = make(map[string]*mgo.Session)
)

type MongoConfig struct {
	Alias   string `json:"alias"`
	Url     string `json:"url"`
	Timeout int    `json:"timeout"`
}

func InitMongoConfigs(configs []*MongoConfig) error {
	for _, c := range configs {
		if _, ok := mongoSessions[c.Alias]; ok {
			return errors.New("duplicate session: " + c.Alias)
		}

		s, err := CreateMongoSession(c.Alias, c.Url, time.Duration(c.Timeout)*time.Second)
		if err != nil {
			return errors.New(fmt.Sprintf("mongo alias %s url %s error %v", c.Alias, c.Url, err))
		}
		xsync.WithLock(mongoSessionsLock, func() {
			mongoSessions[c.Alias] = s
		})

	}
	return nil
}

func CreateMongoSession(alias string, url string, timeout time.Duration) (*mgo.Session, error) {
	s, err := mgo.DialWithTimeout(url, timeout)
	if err != nil {
		return nil, err
	}

	err = s.Ping()
	if err != nil {
		return nil, err
	}
	return s, err
}

func GetMongoSession(alias string) *mgo.Session {
	mongoSessionsLock.RLock()
	defer mongoSessionsLock.RUnlock()
	return mongoSessions[alias].Copy()
}

type MongoDB struct {
	*mgo.Database
}

func (m *MongoDB) Close() {
	m.Session.Close()
}

func NewMongoDB(alias string, dbName string) *MongoDB {
	return &MongoDB{GetMongoSession(alias).DB(dbName)}
}

type MongoColl struct {
	*mgo.Collection
}

func (m *MongoColl) Close() {
	m.Database.Session.Close()
}

func NewMongoColl(alias string, dbName string, collName string) *MongoColl {
	return &MongoColl{NewMongoDB(alias, dbName).C(collName)}
}

func DBCollNames(alias string, dbName string) (collNames []string, err error) {
	db := NewMongoDB(alias, dbName)
	defer db.Close()
	collNames, err = db.CollectionNames()
	if err != nil {
		return nil, err
	}
	return collNames, nil
}

func DBDatabaseNames(alias string) ([]string, error)  {
	sess := GetMongoSession(alias)
	defer sess.Close()
	return sess.DatabaseNames()
}

func WithColl(alias string, dbName string, collName string, fn func(coll *MongoColl) error) error {
	coll := NewMongoColl(alias, dbName, collName)
	defer coll.Close()
	return fn(coll)
}