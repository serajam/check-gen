package password

import (
	"reflect"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	type fields struct {
		Password string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "should success",
			fields: fields{
				Password: "wqerqwrwqrwqwe1R",
			},
			want: nil,
		},
		{
			name: "should no uppercase",
			fields: fields{
				Password: "wqerqwrwqrwqwe1",
			},
			want: []string{"password::invalid_password_format"},
		},
		{
			name: "should no number",
			fields: fields{
				Password: "wqerqwrwqrwqweR",
			},
			want: []string{"password::invalid_password_format"},
		},
		{
			name: "should len",
			fields: fields{
				Password: "wqewr1R",
			},
			want: []string{"password::min_length_is::8"},
		},
		{
			name: "should len err",
			fields: fields{
				Password: "",
			},
			want: []string{"password::min_length_is::8"},
		},
		{
			name: "should success with spec symbols",
			fields: fields{
				Password: "$qrqwrerwqrewqrerq1R",
			},
			want: nil,
		},
		{
			name: "should err with white space",
			fields: fields{
				Password: "$qrqwrerwq rewqrerq1R",
			},
			want: []string{"password::invalid_password_format"},
		},

		{
			name: "should err with russian chars",
			fields: fields{
				Password: "ыыыыыыыыыыыыыЫ2",
			},
			want: []string{"password::invalid_password_format"},
		},
		{
			name: "should success all supported alphabet",
			fields: fields{
				Password: "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				Password: tt.fields.Password,
			}
			if got := u.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func str(s string) *string {
	return &s
}

func TestUser2_Validate(t *testing.T) {
	type fields struct {
		Password *string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "should success",
			fields: fields{
				Password: str("wqerqwrwqrwqwe1R"),
			},
			want: nil,
		},
		{
			name: "should no uppercase",
			fields: fields{
				Password: str("wqerqwrwqrwqwe1"),
			},
			want: []string{"password_ref::invalid_password_format"},
		},
		{
			name: "should no number",
			fields: fields{
				Password: str("wqerqwrwqrwqweR"),
			},
			want: []string{"password_ref::invalid_password_format"},
		},
		{
			name: "should len",
			fields: fields{
				Password: str("wqewr1R"),
			},
			want: []string{"password_ref::min_length_is::8"},
		},
		{
			name: "should len err",
			fields: fields{
				Password: str(""),
			},
			want: []string{"password_ref::min_length_is::8"},
		},
		{
			name: "should len err",
			fields: fields{
				Password: str("$qrqwrerwqrewqrerq1R"),
			},
			want: nil,
		},
		{
			name: "should err with white space",
			fields: fields{
				Password: str("$qrqwrerwq rewqrerq1R"),
			},
			want: []string{"password_ref::invalid_password_format"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User2{
				PasswordRef: tt.fields.Password,
			}
			if got := u.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
