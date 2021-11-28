package distributedtracing

import (
	"context"
	"net/http"
	"strings"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func StartSegment(ctx context.Context, segmentName string) (context.Context, *newrelic.Segment) {
	txn := newrelic.FromContext(ctx)
	return ctx, txn.StartSegment(segmentName)
}

func StartPostgresSegment(ctx context.Context, segmentName string, query string) (context.Context, newrelic.DatastoreSegment) {
	txn := newrelic.FromContext(ctx)

	var operation string
	querySlice := strings.Split(query, " ")
	if len(querySlice) > 0 {
		operation = querySlice[0]
	}

	segment := newrelic.DatastoreSegment{
		StartTime:  txn.StartSegmentNow(),
		Product:    newrelic.DatastorePostgres,
		Collection: segmentName,
		Operation:  operation,
	}

	return ctx, segment
}

func StartExternalHTTPSegment(ctx context.Context, segmentName string, req *http.Request) (context.Context, *newrelic.Segment, *http.Request) {
	txn := newrelic.FromContext(ctx)
	requestWithCtx := newrelic.RequestWithTransactionContext(req, txn)

	return ctx, txn.StartSegment(segmentName), requestWithCtx
}
