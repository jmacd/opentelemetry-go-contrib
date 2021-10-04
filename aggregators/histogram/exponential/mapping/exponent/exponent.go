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

package exponent

import (
	"fmt"
	"math"

	"go.opentelemetry.io/contrib/aggregators/histogram/exponential/mapping"
)

// exponentMapping is used for negative scales, effectively a
// mapping of the base-2 logarithm of the exponent.
type exponentMapping struct {
	scale int32
}

func NewExponentMapping(scale int32) mapping.Mapping {
	if scale > 0 {
		panic("requires scale <= 0")
	}
	return exponentMapping{scale}
}

func (e exponentMapping) MapToIndex(value float64) int64 {
	// Note: we can assume not a 0, Inf, or NaN, ignore sign bit.

	// GetExponent compensates for subnormal values.
	exp := mapping.GetExponent(value)

	// Note: bit-shifting does the right thing for negative
	// exponents, e.g., -1 >> 1 == -1.
	return int64(exp >> -e.scale)
}

var ErrUnderflow = fmt.Errorf("underflow")
var ErrOverflow = fmt.Errorf("overflow")

func (e exponentMapping) LowerBoundary(index int64) (float64, error) {
	unbiased := int64(index << -e.scale)

	var bits uint64

	if unbiased < int64(mapping.MinNormalExponent) {
		diff := mapping.MinNormalExponent - int32(unbiased)
		if diff > mapping.SignificandWidth {
			return 0, ErrUnderflow
		}
		// the encoded exponent is zero
		bits = uint64(1) << (mapping.SignificandWidth - diff)
	} else if unbiased > int64(mapping.MaxNormalExponent) {
		return 0, ErrOverflow
	} else {
		exponent := unbiased + mapping.ExponentBias
		bits = uint64(exponent << mapping.SignificandWidth)
	}

	return math.Float64frombits(bits), nil
}

func (e exponentMapping) Scale() int32 {
	return e.scale
}
