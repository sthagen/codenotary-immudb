// +build integration

/*
Copyright 2019-2020 vChain, Inc.

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
package client

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"github.com/codenotary/immudb/pkg/api/schema"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"math"
	rand2 "math/rand"
	"strings"
	"testing"
	"time"
)

func TestImmuClient_ExecAllMultiSortedSetsIntegration(t *testing.T) {
	ts := NewTokenService().WithTokenFileName("testTokenFile").WithHds(DefaultHomedirServiceMock())
	client, err := NewImmuClient(DefaultOptions().WithDialOptions(&[]grpc.DialOption{grpc.WithInsecure()}).WithTokenService(ts))
	if err != nil {
		log.Fatal(err)
	}
	lr, err := client.Login(context.TODO(), []byte(`immudb`), []byte(`immudb`))
	if err != nil {
		log.Fatal(err)
	}
	md := metadata.Pairs("authorization", lr.Token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	ur, err := client.UseDatabase(ctx, &schema.Database{Databasename: "defaultdb"})
	if err != nil {
		log.Fatal(err)
	}
	md = metadata.Pairs("authorization", ur.Token)
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	setName1 := []byte(`set13`)
	setName2 := []byte(`set14`)

	zaddOpts1 := &schema.ZAddRequest{
		Set:      setName1,
		Score:    float64(1),
		Key:      []byte(`key1.0`),
		BoundRef: true,
	}
	zaddOpts2 := &schema.ZAddRequest{
		Set:      setName1,
		Score:    float64(1),
		Key:      []byte(`key1.1`),
		BoundRef: true,
	}
	zaddOpts3 := &schema.ZAddRequest{
		Set:      setName1,
		Score:    float64(2),
		Key:      []byte(`key1.2`),
		BoundRef: true,
	}
	zaddOpts4 := &schema.ZAddRequest{
		Set:      setName2,
		Score:    float64(2),
		Key:      []byte(`key1.0`),
		BoundRef: true,
	}
	zaddOpts5 := &schema.ZAddRequest{
		Set:      setName2,
		Score:    float64(2),
		Key:      []byte(`key1.1`),
		BoundRef: true,
	}
	zaddOpts6 := &schema.ZAddRequest{
		Set:      setName2,
		Score:    float64(3),
		Key:      []byte(`key1.2`),
		BoundRef: true,
	}

	bOps1 := &schema.ExecAllRequest{
		Operations: []*schema.Op{
			{
				Operation: &schema.Op_Kv{
					Kv: &schema.KeyValue{
						Key:   []byte(`key1.0`),
						Value: []byte(`key1.0`),
					},
				},
			},
			{
				Operation: &schema.Op_ZAdd{
					ZAdd: zaddOpts1,
				},
			},
		},
	}
	client.ExecAll(ctx, bOps1)

	bOps2 := &schema.ExecAllRequest{
		Operations: []*schema.Op{
			{
				Operation: &schema.Op_Kv{
					Kv: &schema.KeyValue{
						Key:   []byte(`key1.1`),
						Value: []byte(`key1.1`),
					},
				},
			},
			{
				Operation: &schema.Op_ZAdd{
					ZAdd: zaddOpts2,
				},
			},
		},
	}
	client.ExecAll(ctx, bOps2)

	bOps3 := &schema.ExecAllRequest{
		Operations: []*schema.Op{
			{
				Operation: &schema.Op_Kv{
					Kv: &schema.KeyValue{
						Key:   []byte(`key1.2`),
						Value: []byte(`key1.2`),
					},
				},
			},
			{
				Operation: &schema.Op_ZAdd{
					ZAdd: zaddOpts3,
				},
			},
		},
	}
	client.ExecAll(ctx, bOps3)

	bOps4 := &schema.ExecAllRequest{
		Operations: []*schema.Op{
			{
				Operation: &schema.Op_Kv{
					Kv: &schema.KeyValue{
						Key:   []byte(`key1.0`),
						Value: []byte(`key1.0`),
					},
				},
			},
			{
				Operation: &schema.Op_ZAdd{
					ZAdd: zaddOpts4,
				},
			},
		},
	}
	client.ExecAll(ctx, bOps4)

	bOps5 := &schema.ExecAllRequest{
		Operations: []*schema.Op{
			{
				Operation: &schema.Op_Kv{
					Kv: &schema.KeyValue{
						Key:   []byte(`key1.1`),
						Value: []byte(`key1.1`),
					},
				},
			},
			{
				Operation: &schema.Op_ZAdd{
					ZAdd: zaddOpts5,
				},
			},
		},
	}
	client.ExecAll(ctx, bOps5)

	bOps6 := &schema.ExecAllRequest{
		Operations: []*schema.Op{
			{
				Operation: &schema.Op_Kv{
					Kv: &schema.KeyValue{
						Key:   []byte(`key1.2`),
						Value: []byte(`key1.2`),
					},
				},
			},
			{
				Operation: &schema.Op_ZAdd{
					ZAdd: zaddOpts6,
				},
			},
		},
	}
	meta, err := client.ExecAll(ctx, bOps6)
	require.NoError(t, err)
	require.NotNil(t, meta)

	zScanOption1 := &schema.ZScanRequest{
		Set:       setName1,
		SeekKey:   []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		SeekScore: math.MaxFloat64,
		SeekAtTx:  math.MaxUint64,
		SinceTx:   math.MaxUint64,
		Desc:      true,
		NoWait:    true,
	}

	list1, err := client.ZScan(ctx, zScanOption1)
	require.NoError(t, err)
	require.Len(t, list1.Entries, 3)

	zScanOption2 := &schema.ZScanRequest{
		Set:       setName2,
		SeekKey:   []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		SeekScore: math.MaxFloat64,
		SeekAtTx:  math.MaxUint64,
		SinceTx:   math.MaxUint64,
		Desc:      true,
		NoWait:    true,
	}

	list2, err := client.ZScan(ctx, zScanOption2)
	require.NoError(t, err)
	require.Len(t, list2.Entries, 3)

	client.Disconnect()
}

func TestImmuClient_ExecAllConcurrentIntegration(t *testing.T) {

	const numExecAll = 10000
	jobs := make(chan *schema.ExecAllRequest, numExecAll)
	res := make(chan *results, numExecAll)
	errors := make(chan error, numExecAll)

	for w := 1; w <= 50; w++ {
		go execAll(w, jobs, res, errors)
	}

	for j := 1; j <= numExecAll; j++ {
		jobs <- getRandomExecOps()
	}
	close(jobs)

	go func() {
		for a := 1; a <= numExecAll; a++ {
			err := <-errors
			e := fmt.Sprintf("error %s", err.Error())
			println(e)
		}
	}()
	for a := 1; a <= numExecAll; a++ {
		r := <-res
		s := fmt.Sprintf("worker %d res %d", r.workerId, r.txMeta.Id)
		println(s)
	}
}

func execAll(w int, jobs <-chan *schema.ExecAllRequest, res chan<- *results, errors chan<- error) {
	ts := NewTokenService().WithTokenFileName("testTokenFile").WithHds(DefaultHomedirServiceMock())
	client, err := NewImmuClient(DefaultOptions().WithDialOptions(&[]grpc.DialOption{grpc.WithInsecure()}).WithTokenService(ts))
	if err != nil {
		log.Fatal(err)
	}
	lr, err := client.Login(context.TODO(), []byte(`immudb`), []byte(`immudb`))
	if err != nil {
		log.Fatal(err)
	}
	md := metadata.Pairs("authorization", lr.Token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	ur, err := client.UseDatabase(ctx, &schema.Database{Databasename: "defaultdb"})
	if err != nil {
		log.Fatal(err)
	}
	md = metadata.Pairs("authorization", ur.Token)
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	for j := range jobs {
		metaTx, err := client.ExecAll(ctx, j)
		if err != nil {
			errors <- err
			return
		}
		r := &results{
			txMeta:   metaTx,
			workerId: w,
		}
		res <- r
	}
	client.Disconnect()

}

type results struct {
	txMeta   *schema.TxMetadata
	workerId int
}

func getRandomExecOps() *schema.ExecAllRequest {
	tn1 := time.Now()

	keyItemDate := make([]byte, 8)
	binary.BigEndian.PutUint64(keyItemDate, uint64(tn1.UnixNano()))
	keyItemDate = bytes.Join([][]byte{[]byte("_ITEM.INSERTION-DATE."), keyItemDate}, nil)

	tb, _ := tn1.MarshalBinary()

	rand2.Seed(tn1.UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand2.Intn(len(chars))])
	}

	sha256Token := sha256.Sum256([]byte(b.String()))

	prefix := []byte(`vcn.SGHn32iPIu87WQbNf8sEiTIq6V0_LztEzQdb4VmZImw=.`)
	key := append(prefix, sha256Token[:]...)
	aOps := &schema.ExecAllRequest{
		Operations: []*schema.Op{
			{
				Operation: &schema.Op_Kv{
					Kv: &schema.KeyValue{
						Key:   key,
						Value: []byte(`{"kind":"file","name":".gitignore","hash":"87b7515a98f78ed4ce0c6c7bb272e9ceb73e93770ac0ac98f98e1d1a085f7ba7","size":371,"timestamp":"0001-01-01T00:00:00Z","contentType":"application/octet-stream","metadata":{},"signer":"SGHn32iPIu87WQbNf8sEiTIq6V0_LztEzQdb4VmZImw=","status":0}`),
					},
				},
			},
			{
				Operation: &schema.Op_Kv{
					Kv: &schema.KeyValue{
						Key:   keyItemDate,
						Value: tb,
					},
				},
			},
			{
				Operation: &schema.Op_ZAdd{
					ZAdd: &schema.ZAddRequest{
						Set:      sha256Token[:],
						Key:      key,
						Score:    float64(tn1.UnixNano()),
						BoundRef: true,
					},
				},
			},
		},
	}

	return aOps
}

func TestImmuClient_MultiExecAll(t *testing.T) {
	ts := NewTokenService().WithTokenFileName("testTokenFile").WithHds(DefaultHomedirServiceMock())
	client, err := NewImmuClient(DefaultOptions().WithDialOptions(&[]grpc.DialOption{grpc.WithInsecure()}).WithTokenService(ts))
	if err != nil {
		log.Fatal(err)
	}
	lr, err := client.Login(context.TODO(), []byte(`immudb`), []byte(`immudb`))
	if err != nil {
		log.Fatal(err)
	}
	md := metadata.Pairs("authorization", lr.Token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	ur, err := client.UseDatabase(ctx, &schema.Database{Databasename: "defaultdb"})
	if err != nil {
		log.Fatal(err)
	}
	md = metadata.Pairs("authorization", ur.Token)
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	bOps1 := getExecAllRequest()
	_, err = client.ExecAll(ctx, bOps1)
	require.NoError(t, err)
	time.Sleep(200 * time.Millisecond)
	bOps2 := getExecAllRequest()
	_, err = client.ExecAll(ctx, bOps2)
	require.NoError(t, err)
	time.Sleep(200 * time.Millisecond)
	bOps3 := getExecAllRequest()
	_, err = client.ExecAll(ctx, bOps3)
	require.NoError(t, err)
	time.Sleep(200 * time.Millisecond)
	bOps4 := getExecAllRequest()
	_, err = client.ExecAll(ctx, bOps4)
	require.NoError(t, err)
	time.Sleep(200 * time.Millisecond)
	bOps5 := getExecAllRequest()
	_, err = client.ExecAll(ctx, bOps5)
	require.NoError(t, err)
	time.Sleep(200 * time.Millisecond)
	zScanOption1 := &schema.ZScanRequest{
		Set:       []byte(`vcn.signerId.hash`),
		SeekKey:   []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		SeekScore: math.MaxFloat64,
		SeekAtTx:  math.MaxUint64,
		SinceTx:   math.MaxUint64,
		Desc:      true,
		NoWait:    true,
	}

	time.Sleep(1 * time.Second)
	list1, err := client.ZScan(ctx, zScanOption1)
	require.NoError(t, err)
	require.True(t, len(list1.Entries) > 0)

	client.Disconnect()
}

const IndexDateRangePrefix = "_INDEX.ITEM.INSERTION-DATE."

func getExecAllRequest() *schema.ExecAllRequest {
	kv := &schema.KeyValue{
		Key:   []byte(`vcn.signerId.hash`),
		Value: []byte(`blablabla`),
	}

	prefix := "vcn"

	key := kv.Key
	val := kv.Value

	inspectSigIDIndex := stripPrefix(prefix, key)

	t := time.Now()
	tb, _ := t.MarshalBinary()

	setIdxDR := []byte(IndexDateRangePrefix)
	setIdxDR = append(setIdxDR, stripPrefix(prefix, key)...)

	keyItemDate := make([]byte, 8)
	binary.BigEndian.PutUint64(keyItemDate, uint64(t.UnixNano()))
	keyItemDate = bytes.Join([][]byte{[]byte(prefix), keyItemDate}, nil)

	bOps := &schema.ExecAllRequest{
		Operations: []*schema.Op{
			{
				Operation: &schema.Op_Kv{
					Kv: &schema.KeyValue{
						Key:   key,
						Value: val,
					},
				},
			},
			{
				Operation: &schema.Op_Kv{
					Kv: &schema.KeyValue{
						Key:   keyItemDate,
						Value: tb,
					},
				},
			},
			{
				Operation: &schema.Op_ZAdd{
					ZAdd: &schema.ZAddRequest{
						Set:      inspectSigIDIndex,
						Key:      key,
						Score:    float64(t.UnixNano()),
						BoundRef: true,
					},
				},
			},
			{
				Operation: &schema.Op_ZAdd{
					ZAdd: &schema.ZAddRequest{
						Set:      setIdxDR,
						Score:    float64(t.UnixNano()),
						Key:      key,
						BoundRef: true,
					},
				},
			},
			{
				Operation: &schema.Op_ZAdd{
					ZAdd: &schema.ZAddRequest{
						Set:      key,
						Score:    float64(t.UnixNano()),
						Key:      key,
						BoundRef: true,
					},
				},
			},
		},
	}
	return bOps
}

func stripPrefix(prefix string, k []byte) []byte {
	kLen := len(k)
	prefixLen := len(prefix) + 1
	unwrappedLen := kLen - prefixLen
	var unwrapped = make([]byte, unwrappedLen)
	copy(unwrapped[0:], k[prefixLen:])
	return unwrapped
}

func TestImmuClient_ExecAllSimple(t *testing.T) {
	client, err := NewImmuClient(DefaultOptions())
	if err != nil {
		log.Fatal(err)
	}
	lr, err := client.Login(context.TODO(), []byte(`immudb`), []byte(`immudb`))
	if err != nil {
		log.Fatal(err)
	}
	md := metadata.Pairs("authorization", lr.Token)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	ur, err := client.UseDatabase(ctx, &schema.Database{Databasename: "defaultdb"})
	if err != nil {
		log.Fatal(err)
	}
	md = metadata.Pairs("authorization", ur.Token)
	ctx = metadata.NewOutgoingContext(context.Background(), md)

	setName1 := []byte(`set134`)

	zaddOpts1 := &schema.ZAddRequest{
		Set:      setName1,
		Score:    float64(1),
		Key:      []byte(`key1.0`),
		BoundRef: true,
	}

	bOps1 := &schema.ExecAllRequest{
		Operations: []*schema.Op{
			{
				Operation: &schema.Op_Kv{
					Kv: &schema.KeyValue{
						Key:   []byte(`key1.0`),
						Value: []byte(`key1.0`),
					},
				},
			},
			{
				Operation: &schema.Op_ZAdd{
					ZAdd: zaddOpts1,
				},
			},
		},
	}
	_, err = client.ExecAll(ctx, bOps1)
	require.NoError(t, err)

	zScanOption1 := &schema.ZScanRequest{
		Set:       setName1,
		SeekKey:   []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
		SeekScore: math.MaxFloat64,
		SeekAtTx:  math.MaxUint64,
		SinceTx:   math.MaxUint64,
		Desc:      true,
		NoWait:    true,
	}

	list1, err := client.ZScan(ctx, zScanOption1)
	require.NoError(t, err)
	require.Len(t, list1.Entries, 3)

	client.Disconnect()
}
