package main

import "testing"

func Test_bank_withdraw(t *testing.T) {
	type fields struct {
		firstName string
		lastName  string
		account   string
		balance   int
	}
	type args struct {
		cash int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &bank{
				firstName: tt.fields.firstName,
				lastName:  tt.fields.lastName,
				account:   tt.fields.account,
				balance:   tt.fields.balance,
			}
			b.withdraw(tt.args.cash)
		})
	}
}
