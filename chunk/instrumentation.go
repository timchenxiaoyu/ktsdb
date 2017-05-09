// Copyright 2014 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chunk


// Usually, a separate file for instrumentation is frowned upon. Metrics should
// be close to where they are used. However, the metrics below are set all over
// the place, so we go for a separate instrumentation file in this case.


const (
	namespace = "ktsdb"
	subsystem = "local_storage"

	// OpTypeLabel is the label name for chunk operation types.
	OpTypeLabel = "type"

	// Op-types for ChunkOps.

	// CreateAndPin is the label value for create-and-pin chunk ops.
	CreateAndPin = "create" // A Desc creation with refCount=1.
	// PersistAndUnpin is the label value for persist chunk ops.
	PersistAndUnpin = "persist"
	// Pin is the label value for pin chunk ops (excludes pin on creation).
	Pin = "pin"
	// Unpin is the label value for unpin chunk ops (excludes the unpin on persisting).
	Unpin = "unpin"
	// Clone is the label value for clone chunk ops.
	Clone = "clone"
	// Transcode is the label value for transcode chunk ops.
	Transcode = "transcode"
	// Drop is the label value for drop chunk ops.
	Drop = "drop"

	// Op-types for ChunkOps and ChunkDescOps.

	// Evict is the label value for evict chunk desc ops.
	Evict = "evict"
	// Load is the label value for load chunk and chunk desc ops.
	Load = "load"
)

func init() {
	//ktsdb.MustRegister(Ops)
	//ktsdb.MustRegister(DescOps)
	//ktsdb.MustRegister(NumMemDescs)
}

var (
	// NumMemChunks is the total number of chunks in memory. This is a
	// global counter, also used internally, so not implemented as
	// metrics. Collected in MemorySeriesStorage.Collect.
	// TODO(beorn7): As it is used internally, it is actually very bad style
	// to have it as a global variable.
	NumMemChunks int64

)
