package xmongo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

func SetUpdateTime(update bson.M) bson.M {
	set, ok := update["$set"]
	if ok {
		set.(bson.M)["update_time"] = time.Now()
	} else {
		update["$set"] = bson.M{"update_time": time.Now()}
	}

	return update
}

type QueryOptions struct {
	Skip int
	Limit int
	Sort []string
}


func ApplyQueryOpts(query *mgo.Query, opts ...QueryOpt) *mgo.Query {
	qo := &QueryOptions{}
	for _, opt := range opts {
		opt(qo)
	}
	if len(qo.Sort) != 0 {
		query = query.Sort(qo.Sort...)
	}
	if qo.Skip != 0 {
		query = query.Skip(qo.Skip)
	}
	if qo.Limit != 0 {
		query = query.Limit(qo.Limit)
	}
	return query
}

type QueryOpt func(*QueryOptions)

func Skip(skip int) QueryOpt {
	return func(opts *QueryOptions) {
		opts.Skip = skip
	}
}

func Limit(limit int) QueryOpt {
	return func(opts *QueryOptions) {
		opts.Limit = limit
	}
}

func Sort(fields ...string) QueryOpt {
	return func(opts *QueryOptions) {
		opts.Sort = append(opts.Sort, fields...)
	}
}