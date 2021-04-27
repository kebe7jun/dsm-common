package utils

import "testing"

func TestAnalysePemCA(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				value: `
-----BEGIN CERTIFICATE-----
MIIC3jCCAcYCAQAwDQYJKoZIhvcNAQELBQAwLTEVMBMGA1UECgwMZXhhbXBsZSBJ
bmMuMRQwEgYDVQQDDAtleGFtcGxlLmNvbTAeFw0yMTA0MTMxNzExMDNaFw0yMjA0
MTMxNzExMDNaMD0xHDAaBgNVBAMME2h0dHBiaW4uZXhhbXBsZS5jb20xHTAbBgNV
BAoMFGh0dHBiaW4gb3JnYW5pemF0aW9uMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8A
MIIBCgKCAQEA1og+0lvFskl8KeEDRQFe0l6ZvpxBz+JyXJbwMzKsMBUcsxGOsdL/
kxfg8qlEYRyZ7FVLuiTCkvsoIhkJ1OIeZ3T8xH6j4DLVenStXRjqVMhXoyMKBhHe
pkkTu6qeJfEKv18jN/Y2vGFnynsEbV4LJ670WhwB81QZCfF4Vn6+0ely/ypfKnBI
jnXEscxQxrPi0+fcdbuztAHQ/IkIA9uMrSkHKGFovTY/+WwkOzKgPO+cjfBWhMXK
9IswRfuN93ii3tBtXIqpY8SUX8IxAywyopu/UXWdScYSmaW+FV9GcrQoTQsPiw0A
4QoyQdQDkMRK4m75qcc1z/xisBQvYLC/WQIDAQABMA0GCSqGSIb3DQEBCwUAA4IB
AQB4zRtxQ9bamM6NoQa6g3dU0Zh7EJ0LKhwpzW1rQOrvBseSv6bZGPQYN8xY2L9T
6uuqWHzcTHpApH0abHO6nyM7sQSLfcDIR/NHwOqEgtRJj8FjtbdSFvZBnOD4IcGt
qB8BGXEHW4AQqFV8ck7gsDV7En3M/TrTGBPgDjkx1npVVvdIxMtBSvot427fWx5c
vZrtn2PeseZ9UyY3CcfUeFVowhAABsM5weKsvHEseojBb3yETW2dTI0iIeJwLiJc
WOp4TY0yNlWSYvJcvwDPjuAo7swz1Bv9XoEuQilLxhoe3eVAQ2UOrUD0dm/UEwMA
MZ9nAmMhW1c/peJ4Zm7DLj6d
-----END CERTIFICATE-----`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseToCertificate([]byte(tt.args.value))
		})
	}
}
