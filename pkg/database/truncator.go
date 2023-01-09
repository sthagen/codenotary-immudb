/*
Copyright 2022 Codenotary Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package database

import (
	"context"
	"time"

	"github.com/codenotary/immudb/embedded/store"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Truncator provides truncation against an underlying storage
// of appendable data.
type Truncator interface {
	// Plan returns the transaction upto which the log can be truncated.
	// When resulting transaction before specified time does not exists
	//  * No transaction header is returned.
	//  * Returns nil TxHeader, and an error.
	Plan(time.Time) (*store.TxHeader, error)

	// Truncate runs truncation against the relevant appendable logs. Must
	// be called after result of Plan().
	Truncate(*store.TxHeader) error
}

func newVlogTruncator(d *db) Truncator {
	return &vlogTruncator{
		db:      d,
		metrics: newTruncatorMetrics(d.name),
	}
}

// vlogTruncator implements Truncator for the value-log appendable
type vlogTruncator struct {
	db      *db
	metrics *truncatorMetrics
}

// Plan returns the transaction upto which the value log can be truncated.
// When resulting transaction before specified time does not exists
//  * No transaction header is returned.
//  * Returns nil TxHeader, and an error.
func (v *vlogTruncator) Plan(ts time.Time) (*store.TxHeader, error) {
	return v.db.st.FirstTxSince(ts)
}

// Truncate runs truncation against the relevant appendable logs upto the specified transaction offset.
func (v *vlogTruncator) Truncate(hdr *store.TxHeader) error {
	defer func(t time.Time) {
		v.metrics.ran.Inc()
		v.metrics.duration.Observe(time.Since(t).Seconds())
	}(time.Now())

	// copy sql catalogue
	tx, err := v.db.CopyCatalog(context.Background())
	if err != nil {
		v.db.Logger.Errorf("error during truncation for database '%s' {err = %v, id = %v, type=sql_catalogue_copy}", v.db.name, err, hdr.ID)
		return err
	}
	defer tx.Cancel()

	// setting the metadata to record the transaction upto which the log was truncated
	tx.WithMetadata(store.NewTxMetadata().WithTruncatedTxID(hdr.ID))

	// commit catalogue as a new transaction
	_, err = tx.Commit(context.Background())
	if err != nil {
		v.db.Logger.Errorf("error during truncation for database '%s' {err = %v, id = %v, type=sql_catalogue_commit}", v.db.name, err, hdr.ID)
		return err
	}

	// truncate upto hdr.ID
	err = v.db.st.TruncateUptoTx(hdr.ID)
	if err != nil {
		v.db.Logger.Errorf("error during truncation for database '%s' {err = %v, id = %v, type=truncate_upto}", v.db.name, err, hdr.ID)
	}

	return err
}

type truncatorMetrics struct {
	ran      prometheus.Counter
	duration prometheus.Observer
}

func newTruncatorMetrics(db string) *truncatorMetrics {
	reg := prometheus.NewRegistry()

	m := &truncatorMetrics{}
	m.ran = promauto.With(reg).NewCounterVec(
		prometheus.CounterOpts{
			Name: "immudb_truncation_total",
			Help: "Total number of truncation that were executed for the database.",
		},
		[]string{"db"},
	).WithLabelValues(db)

	m.duration = promauto.With(reg).NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "immudb_truncation_duration_seconds",
			Help:    "Duration of truncation runs",
			Buckets: prometheus.ExponentialBuckets(1, 10.0, 16),
		},
		[]string{"db"},
	).WithLabelValues(db)

	return m
}
