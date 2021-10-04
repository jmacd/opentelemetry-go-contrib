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

package mapping

import (
	"math"
	"math/bits"
)

const (
	// MantissaWidth is the size of an IEEE 754 double-precision
	// floating-point mantissa.
	MantissaWidth = 52
	// ExponentWidth is the size of an IEEE 754 double-precision
	// floating-point exponent.
	ExponentWidth = 11

	// MantissaMask is the mask for the mantissa of an IEEE 754
	// double-precision floating-point value.
	MantissaMask = 1<<MantissaWidth - 1

	// ExponentBias is the exponent bias specified for encoding
	// the IEEE 754 double-precision floating point exponent.
	ExponentBias = 1<<(ExponentWidth-1) - 1

	// ExponentMask are set to 1 for the bits of an IEEE 754
	// floating point exponent (as distinct from the mantissa and
	// sign).
	ExponentMask = ((1 << ExponentWidth) - 1) << MantissaWidth

	// SignMask selects the sign bit of an IEEE 754 floating point
	// number.
	SignMask = (1 << 63)

	// MinNormalExponent is the minimum exponent of a normalized
	// floating point.
	MinNormalExponent int32 = -ExponentBias + 1

	// MaxNormalExponent is the maximum exponent of a normalized
	// floating point.
	MaxNormalExponent int32 = ExponentBias

	// SignedZeroSubnormalExponent is the exponent value after
	// removing bias for signed zero and subnormal values.
	SignedZeroSubnormalExponent int32 = -ExponentBias

	// InfAndNaNExponent is the exponent value after removing bias
	// for Inf and NaN values.
	InfAndNaNExponent int32 = ExponentBias + 1

	// Smallest positive subnormal exponent:
	MinSubnormalExponent int32 = MinNormalExponent - MantissaWidth
)

// java.lang.Math.scalb(float f, int scaleFactor) returns f x
// 2**scaleFactor, rounded as if performed by a single correctly
// rounded floating-point multiply to a member of the double value set.
func Scalb(f float64, sf int32) float64 {
	if f == 0 {
		return 0
	}
	valueBits := math.Float64bits(f)

	signBit := valueBits & SignMask
	mantissa := MantissaMask & valueBits

	exponent := (int64(valueBits) & ExponentMask) >> MantissaWidth
	exponent += int64(sf)

	return math.Float64frombits(signBit | uint64(exponent<<MantissaWidth) | mantissa)
}

// GetExponent extracts the 11-bit IEEE 754 exponent.
func GetExponent(value float64) int32 {
	valueBits := math.Float64bits(value)
	rawExponent := (int64(valueBits) & ExponentMask) >> MantissaWidth
	rawMantissa := valueBits & MantissaMask
	if rawExponent == 0 && rawMantissa != 0 {
		// Handle subnormals
		rawExponent -= int64(bits.LeadingZeros64(rawMantissa) - 12)
	}
	return int32(rawExponent - ExponentBias)
}
