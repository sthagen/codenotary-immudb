/*
Copyright 2022 CodeNotary, Inc. All rights reserved.

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
package runner

import (
	"github.com/codenotary/immudb/test/performance-test-suite/pkg/benchmarks"
	"github.com/codenotary/immudb/test/performance-test-suite/pkg/benchmarks/writetxs"
)

func getBenchmarksToRun() []benchmarks.Benchmark {
	return []benchmarks.Benchmark{
		writetxs.NewBenchmark(writetxs.Config{
			Name:       "Write TX/s",
			Workers:    30,
			BatchSize:  1,
			KeySize:    32,
			ValueSize:  128,
			AsyncWrite: true,
		}),

		writetxs.NewBenchmark(writetxs.Config{
			Name:       "Write KV/s",
			Workers:    30,
			BatchSize:  1000,
			KeySize:    32,
			ValueSize:  128,
			AsyncWrite: true,
		}),
	}
}
