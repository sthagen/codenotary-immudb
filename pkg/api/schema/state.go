/*
Copyright 2025 Codenotary Inc. All rights reserved.

SPDX-License-Identifier: BUSL-1.1
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://mariadb.com/bsl11/

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package schema

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/binary"
	"errors"

	"github.com/codenotary/immudb/pkg/signer"
)

func (state *ImmutableState) ToBytes() []byte {
	b := make([]byte, 4+len(state.Db)+8+sha256.Size)
	i := 0

	binary.BigEndian.PutUint32(b[i:], uint32(len(state.Db)))
	i += 4

	copy(b[i:], []byte(state.Db))
	i += len(state.Db)

	binary.BigEndian.PutUint64(b[i:], state.TxId)
	i += 8

	copy(b[i:], state.TxHash[:])

	return b
}

// CheckSignature
func (state *ImmutableState) CheckSignature(key *ecdsa.PublicKey) error {
	if state.Signature == nil {
		return errors.New("no signature provided")
	}

	return signer.Verify(state.ToBytes(), state.Signature.Signature, key)
}
