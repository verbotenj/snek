// Copyright 2023 Blink Labs, LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chainsync

import (
	"encoding/hex"

	ocommon "github.com/blinklabs-io/gouroboros/protocol/common"
)

type RollbackEvent struct {
	BlockHash  string `json:"blockHash"`
	SlotNumber uint64 `json:"slotNumber"`
}

func NewRollbackEvent(point ocommon.Point) RollbackEvent {
	blockHashHex := hex.EncodeToString(point.Hash)
	evt := RollbackEvent{
		BlockHash:  blockHashHex,
		SlotNumber: point.Slot,
	}
	return evt
}
