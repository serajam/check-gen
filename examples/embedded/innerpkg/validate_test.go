package innerpkg

import (
	"testing"
)

func TestBar_Validate(t *testing.T) {
	type fields struct {
		Surname string
		Foo     Foo
	}
	tests := []struct {
		name       string
		fields     fields
		wantErrNum int
	}{
		{
			"should validate",
			fields{Surname: "Foo", Foo: Foo{Name: "Bar"}},
			0,
		},

		{
			"should return errors",
			fields{Surname: "", Foo: Foo{Name: ""}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Bar{
				Surname: tt.fields.Surname,
				Foo:     tt.fields.Foo,
			}
			if got := b.Validate(); len(got) != tt.wantErrNum {
				t.Errorf("Validate() = %v, want %v", got, tt.wantErrNum)
			}
		})
	}
}

func TestBarRef_Validate(t *testing.T) {
	type fields struct {
		Surname string
		Foo     *Foo
	}
	tests := []struct {
		name       string
		fields     fields
		wantErrNum int
	}{
		{
			"should validate",
			fields{Surname: "Foo", Foo: &Foo{Name: "Bar"}},
			0,
		},

		{
			"should return errors",
			fields{Surname: "", Foo: &Foo{Name: ""}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := BarRef{
				Surname: tt.fields.Surname,
				Foo:     tt.fields.Foo,
			}
			if got := b.Validate(); len(got) != tt.wantErrNum {
				t.Errorf("Validate() = %v, want %v", got, tt.wantErrNum)
			}
		})
	}
}
