package main

import (
	"github.com/cpk/currency-rounding/c"
	"testing"
)

func TestRoundFloat64(t *testing.T) {
	tables := []struct {
		Input float64
		Output float64
	}{
		{0, 0},
		{0.5, 1},
		{0.4999, 0},
		{-0.5, -1},
		{-0.4999, 0},
	}
	for _, table := range tables {
		output := c.RoundFloat64(table.Input)
		if output != table.Output {
			t.Errorf("RoundFloat64 of %f was incorrect, got: %f, want: %f.", table.Input, output, table.Output)
		}
	}
}

func TestRoundCurrency(t *testing.T) {
	tables := []struct {
		Currency string
		Input float64
		Output float64
	}{
		{"VND", 0, 0},
		{"VND", 0.5, 1},
		{"VND", 0.4999, 0},
		{"VND", -0.5, -1},
		{"VND", -0.4999, 0},


		{"VNC", 0, 0},
		{"VNC", 0.5, 0.5},
		{"VNC", 0.4999, 0.5},
		{"VNC", -0.5, -0.5},
		{"VNC", -0.4999, -0.5},


		{"USD", 0, 0},
		{"USD", 0.5, 0.5},
		{"USD", 0.4999, 0.5},
		{"USD", -0.5, -0.5},
		{"USD", -0.4999, -0.5},


		{"", 0, 0},
		{"", 0.5, 0.5},
		{"", 0.4999, 0.5},
		{"", -0.5, -0.5},
		{"", -0.4999, -0.5},
	}

	for _, table := range tables {
		output := c.RoundCurrency(table.Input, table.Currency)
		if output != table.Output {
			t.Errorf("RoundCurrency of %f for %s was incorrect, got: %f, want: %f.", table.Input, table.Currency, output, table.Output)
		}
	}
}



func TestRoundCurrencyStr(t *testing.T) {
	tables := []struct {
		Currency string
		Input float64
		Output string
	}{
		{"VND", 0, "0"},
		{"VND", 0.5, "1"},
		{"VND", 0.4999, "0"},
		{"VND", -0.5, "-1"},
		{"VND", -0.4999, "0"},


		{"VNC", 0, "0.00"},
		{"VNC", 0.5, "0.50"},
		{"VNC", 0.4999, "0.50"},
		{"VNC", -0.5, "-0.50"},
		{"VNC", -0.4999, "-0.50"},


		{"USD", 0, "0.00"},
		{"USD", 0.5, "0.50"},
		{"USD", 0.4999, "0.50"},
		{"USD", -0.5, "-0.50"},
		{"USD", -0.4999, "-0.50"},


		{"", 0, "0.00"},
		{"", 0.5, "0.50"},
		{"", 0.4999, "0.50"},
		{"", -0.5, "-0.50"},
		{"", -0.4999, "-0.50"},
	}

	for _, table := range tables {
		output := c.RoundCurrencyStr(table.Input, table.Currency)
		if output != table.Output {
			t.Errorf("RoundCurrency of %f for %s was incorrect, got: %s, want: %s.", table.Input, table.Currency, output, table.Output)
		}
	}
}