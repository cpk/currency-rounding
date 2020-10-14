package c

import (
	"fmt"
	"github.com/cpk/currency-rounding/constant"
	"math"
)

func RoundFloat64(x float64) float64 {
	const (
		signMask = 1 << 63
		fracMask = (1 << constant.SHIFT) - 1
		halfMask = 1 << (constant.SHIFT - 1)
		one      = constant.BIAS << constant.SHIFT
	)
	bits := math.Float64bits(x)
	e := uint(bits>>constant.SHIFT) & constant.MASK

	switch {
	case e < constant.BIAS:
		// Round abs(x)<1 including denormals.
		bits &= signMask // +-0
		if e == constant.BIAS-1 {
			bits |= one // +-1
		}
	case e < constant.BIAS+constant.SHIFT:
		// Round any abs(x)>=1 containing a fractional component [0,1).
		e -= constant.BIAS
		bits += halfMask >> e
		bits &^= fracMask >> e
	}

	return math.Float64frombits(bits)
}

func FloorCurrency(price float64, currencyCode string) float64 {
	c, exist := constant.Currencies[currencyCode]
	decimal := constant.DEFAULT_DECIMAL
	if exist {
		decimal = c.Decimal
	}

	rate := math.Pow10(decimal)
	f := math.Floor(price*rate) / rate
	if f == -0 {
		f = 0
	}
	return f
}

func FloorCurrencyStr(price float64, currencyCode string) string {
	c, exist := constant.Currencies[currencyCode]
	decimal := constant.DEFAULT_DECIMAL
	if exist {
		decimal = c.Decimal
	}

	rate := math.Pow10(decimal)

	f := math.Floor(price*rate) / rate
	if f == -0 {
		f = 0
	}
	str := fmt.Sprintf("%."+fmt.Sprintf("%d", decimal)+"f", f)

	return str
}

func FloorStr(price float64, decimal int) string {
	rate := math.Pow10(decimal)

	f := math.Floor(price*rate) / rate
	if f == -0 {
		f = 0
	}
	str := fmt.Sprintf("%."+fmt.Sprintf("%d", decimal)+"f", f)

	return str
}



func CeilCurrency(price float64, currencyCode string) float64 {
	c, exist := constant.Currencies[currencyCode]
	decimal := constant.DEFAULT_DECIMAL
	if exist {
		decimal = c.Decimal
	}

	rate := math.Pow10(decimal)
	f := math.Floor(price*rate) / rate
	if f == -0 {
		f = 0
	}
	return f
}

func CeilCurrencyStr(price float64, currencyCode string) string {
	c, exist := constant.Currencies[currencyCode]
	decimal := constant.DEFAULT_DECIMAL
	if exist {
		decimal = c.Decimal
	}

	rate := math.Pow10(decimal)

	f := math.Floor(price*rate) / rate
	if f == -0 {
		f = 0
	}
	str := fmt.Sprintf("%."+fmt.Sprintf("%d", decimal)+"f", f)

	return str
}

func CeilStr(price float64, decimal int) string {
	rate := math.Pow10(decimal)

	f := math.Floor(price*rate) / rate
	if f == -0 {
		f = 0
	}
	str := fmt.Sprintf("%."+fmt.Sprintf("%d", decimal)+"f", f)

	return str
}

func RoundCurrency(price float64, currencyCode string) float64 {
	c, exist := constant.Currencies[currencyCode]
	decimal := constant.DEFAULT_DECIMAL
	if exist {
		decimal = c.Decimal
	}

	rate := math.Pow10(decimal)
	f := math.Round(price*rate) / rate
	if f == -0 {
		f = 0
	}
	return f
}

func RoundCurrencyStr(price float64, currencyCode string) string {
	c, exist := constant.Currencies[currencyCode]
	decimal := constant.DEFAULT_DECIMAL
	if exist {
		decimal = c.Decimal
	}

	rate := math.Pow10(decimal)

	f := math.Round(price*rate) / rate
	if f == -0 {
		f = 0
	}
	str := fmt.Sprintf("%."+fmt.Sprintf("%d", decimal)+"f", f)

	return str
}

func RoundStr(price float64, decimal int) string {
	rate := math.Pow10(decimal)

	f := math.Round(price*rate) / rate
	if f == -0 {
		f = 0
	}
	str := fmt.Sprintf("%."+fmt.Sprintf("%d", decimal)+"f", f)

	return str
}