package user

import "testing"

func Test_checkEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "email valid",
			args: args{
				email: "1234@qq.com",
			},
			want: true,
		},
		{
			name: "email  invalid",
			args: args{
				email: "test.com",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkEmail(tt.args.email); got != tt.want {
				t.Errorf("checkEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
