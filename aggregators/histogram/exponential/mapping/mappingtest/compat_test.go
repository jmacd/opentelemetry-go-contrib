// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package histogram_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/contrib/aggregators/histogram/exponential/mapping"
	"go.opentelemetry.io/contrib/aggregators/histogram/exponential/mapping/logarithm"
	"go.opentelemetry.io/contrib/aggregators/histogram/exponential/mapping/lookuptable"
)

func testCompatibility(t *testing.T, histoScale, testScale int32) {
	src := rand.New(rand.NewSource(54979))
	t.Run(fmt.Sprintf("compat_%d_%d", histoScale, testScale), func(t *testing.T) {
		const trials = 1e5

		ltm := lookuptable.NewLookupTableMapping(histoScale)
		lgm := logarithm.NewLogarithmMapping(histoScale)

		for i := 0; i < trials; i++ {
			v := mapping.Scalb(1+src.Float64(), testScale)

			lti := ltm.MapToIndex(v)
			lgi := lgm.MapToIndex(v)

			assert.Equal(t, lti, lgi)
		}

		size := int64(1) << histoScale
		additional := int64(testScale) * size

		for i := int64(0); i < size; i++ {
			ltb, _ := ltm.LowerBoundary(i + additional)
			lgb, _ := lgm.LowerBoundary(i + additional)

			assert.InEpsilon(t, ltb, lgb, 0.000001, "hs %v ts %v sz %v add %v index %v ltb %v lgb %v", histoScale, testScale, size, additional, i+additional, ltb, lgb)
		}
	})
}

func TestCompat0(t *testing.T) {
	for scale := int32(3); scale <= 10; scale++ {
		for tscale := int32(-1); tscale <= 1; tscale++ {
			testCompatibility(t, scale, tscale)
		}
	}
}
