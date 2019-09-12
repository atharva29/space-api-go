package db

import (
	"context"

	"github.com/spaceuptech/space-api-go/api/config"
	"github.com/spaceuptech/space-api-go/api/model"
	"github.com/spaceuptech/space-api-go/api/proto"
	"github.com/spaceuptech/space-api-go/api/transport"
)

// Aggr contains the methods for the aggregation operation
type Aggr struct {
	ctx      context.Context
	meta     *proto.Meta
	op       string
	pipeline []interface{}
	config   *config.Config
	httpMeta *model.Meta
}

func initAggr(ctx context.Context, db, col, op string, config *config.Config) *Aggr {
	m := &proto.Meta{Col: col, DbType: db, Project: config.Project, Token: config.Token}
	meta := &model.Meta{Col: col, DbType: db, Project: config.Project, Token: config.Token}
	p := []interface{}{}
	return &Aggr{ctx, m, op, p, config, meta}
}

// Pipe sets the pipeline to run on the backend
func (a *Aggr) Pipe(pipeline []interface{}) *Aggr {
	a.pipeline = pipeline
	return a
}

// Apply executes the operation and returns the result
func (a *Aggr) Apply() (*model.Response, error) {
	transport.Send("aggr", a.createAggrReq(), a.httpMeta)
	return a.config.Transport.Aggr(a.ctx, a.meta, a.op, a.pipeline)
}

func (a *Aggr) createAggrReq() *model.AggregateRequest {
	return &model.AggregateRequest{Pipeline: a.pipeline, Operation: a.op}
}
