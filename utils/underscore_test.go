package utils

import "testing"

func TestUnderscore(t *testing.T) {
	cases := []struct {
		name string
		in   string
		out  string
	}{
		{
			"empty",
			"",
			"",
		},
		{
			"test",
			`{"aB": 1}`,
			`{"a_b": 1}`,
		},
		{
			"test array",
			`[{"aB": 1}]`,
			`[{"a_b": 1}]`,
		},
		{
			"test multiple _",
			`[{"aBC": 1}]`,
			`[{"a_bc": 1}]`,
		},
		{
			"test value",
			`[{"aB": "a_c"}]`,
			`[{"a_b": "a_c"}]`,
		},
		{
			"test id",
			`[{"aID": "a_c"}]`,
			`[{"a_id": "a_c"}]`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if res := UnderscoreJSONStr(c.in); res != c.out {
				t.Fatalf("underscore error, wanted %s, real: %s", c.out, res)
			}
		})
	}
}
