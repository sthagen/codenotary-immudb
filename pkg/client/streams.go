/*
Copyright 2021 CodeNotary, Inc. All rights reserved.

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
	"github.com/codenotary/immudb/pkg/api/schema"
	"github.com/codenotary/immudb/pkg/stream"
	"io"
)

func (c *immuClient) streamSet(ctx context.Context) (schema.ImmuService_StreamSetClient, error) {
	if !c.IsConnected() {
		return nil, ErrNotConnected
	}
	return c.ServiceClient.StreamSet(ctx)
}

func (c *immuClient) streamGet(ctx context.Context, in *schema.KeyRequest) (schema.ImmuService_StreamGetClient, error) {
	if !c.IsConnected() {
		return nil, ErrNotConnected
	}
	return c.ServiceClient.StreamGet(ctx, in)
}

func (c *immuClient) StreamSet(ctx context.Context, kvs []*stream.KeyValue) (*schema.TxMetadata, error) {
	if !c.IsConnected() {
		return nil, ErrNotConnected
	}
	s, err := c.streamSet(ctx)
	if err != nil {
		return nil, err
	}

	kvss := stream.NewKvStreamSender(stream.NewMsgSender(s))

	for _, kv := range kvs {
		err = kvss.Send(kv)
		if err != nil {
			return nil, err
		}
	}

	return s.CloseAndRecv()
}

func (c *immuClient) StreamGet(ctx context.Context, k *schema.KeyRequest) (*schema.Entry, error) {
	gs, err := c.streamGet(ctx, k)

	kvr := stream.NewKvStreamReceiver(stream.NewMsgReceiver(gs))

	key, err := kvr.NextKey()
	if err != nil {
		return nil, err
	}

	vr, err := kvr.NextValueReader()
	if err != nil {
		return nil, err
	}

	b := bytes.NewBuffer([]byte{})
	vl := 0
	chunk := make([]byte, stream.ChunkSize)
	for {
		l, err := vr.Read(chunk)
		if err != nil && err != io.EOF {
			return nil, err
		}
		vl += l
		b.Write(chunk)
		if err == io.EOF || l == 0 {
			break
		}
	}
	value := make([]byte, vl)
	_, err = b.Read(value)
	if err != nil {
		return nil, err
	}

	return &schema.Entry{
		Key:   key,
		Value: value,
	}, nil
}
