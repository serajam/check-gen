// Copyright 2020 SK-Telemed GmbH. All rights reserved.

package phone

import (
	"reflect"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	ref := func(str string) *string {
		return &str
	}

	type fields struct {
		Phone    string
		PhoneRef *string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "should return nil - correct phone number",
			fields: fields{
				Phone:    "+380116246520",
				PhoneRef: ref("+380116246520"),
			},
			want: nil,
		},

		{
			name: "should return nil - 7 digits after '+",
			fields: fields{
				Phone:    "+3801162",
				PhoneRef: ref("+3801162"),
			},
			want: nil,
		},

		{
			name: "should return nil - 15 digits after '+",
			fields: fields{
				Phone:    "+380116238011621",
				PhoneRef: ref("+380116238011621"),
			},
			want: nil,
		},

		{
			name: "should return nil - no phone specified",
			fields: fields{
				Phone:    "",
				PhoneRef: ref(""),
			},
			want: nil,
		},

		{
			name: "should return nil - no phone specified, phoneRef nil",
			fields: fields{
				Phone:    "",
				PhoneRef: nil,
			},
			want: nil,
		},

		{
			name: "should return nil - no numbers specified",
			fields: fields{
				Phone:    "+",
				PhoneRef: ref("+"),
			},

			want: []string{"phone::invalid_phone_format::wrong_length", "phone_ref::invalid_phone_format::wrong_length"},
		},

		{
			name: "should return error - min length 6",
			fields: fields{
				Phone:    "+380116",
				PhoneRef: ref("+380116"),
			},
			want: []string{"phone::invalid_phone_format::wrong_length", "phone_ref::invalid_phone_format::wrong_length"},
		},

		{
			name: "should return error - min length 16",
			fields: fields{
				Phone:    "+3801162380116212",
				PhoneRef: ref("+3801162380116212"),
			},
			want: []string{"phone::invalid_phone_format::wrong_length", "phone_ref::invalid_phone_format::wrong_length"},
		},

		{
			name: "should return error - letter",
			fields: fields{
				Phone:    "+3801q380116212",
				PhoneRef: ref("+3801q380116212"),
			},
			want: []string{"phone::invalid_phone_format::not_a_digits", "phone_ref::invalid_phone_format::not_a_digits"},
		},

		{
			name: "should return error - spec symbol",
			fields: fields{
				Phone:    "+3801#380116212",
				PhoneRef: ref("+3801#380116212"),
			},
			want: []string{"phone::invalid_phone_format::not_a_digits", "phone_ref::invalid_phone_format::not_a_digits"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				Phone:    tt.fields.Phone,
				PhoneRef: tt.fields.PhoneRef,
			}
			if got := u.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
