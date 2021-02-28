package count

import (
	"reflect"
	"testing"
)

func TestChecker_Validate(t *testing.T) {
	type fields struct {
		CyrillicLen string
		PolishLen   string
		GermanLen   string

		CyrillicMin string
		PolishMin   string
		GermanMin   string

		CyrillicMax string
		PolishMax   string
		GermanMax   string

		Multiple     []string
		MultipleRefs []*string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "should success",
			fields: fields{
				CyrillicLen: `Язык общего назначения с широкими возможностями и понятным синтаксисом`,
				CyrillicMax: `Язык общего назначения с широкими возможностями и понятным синтаксисом`,
				CyrillicMin: `Язык общего назначения с широкими возможностями и понятным синтаксисом`,

				PolishLen: `Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`,
				PolishMax: `Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`,
				PolishMin: `Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`,

				GermanLen: `Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`,
				GermanMax: `Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`,
				GermanMin: `Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`,
			},
			want: nil,
		},

		{
			name: "should min errors on each language",
			fields: fields{
				CyrillicLen: `Язык общего назначения с широкими возможностями и понятным синтаксисом`,
				CyrillicMax: `Язык общего назначения с широкими возможностями и понятным синтаксисом`,
				CyrillicMin: `Язык общего назначения с широкими возможностями и понятным синтаксисо`,

				PolishLen: `Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`,
				PolishMax: `Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`,
				PolishMin: `Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzi`,

				GermanLen: `Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`,
				GermanMax: `Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`,
				GermanMin: `Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg`,
			},
			want: []string{
				"cyrillic_min::min_length_is::70", "polish_min::min_length_is::111", "german_min::min_length_is::120",
			},
		},

		{
			name: "should len errors on each language",
			fields: fields{
				CyrillicLen: `Язык общего назначения с широкими возможностями и понятным синтаксисо`,
				CyrillicMax: `Язык общего назначения с широкими возможностями и понятным синтаксисом`,
				CyrillicMin: `Язык общего назначения с широкими возможностями и понятным синтаксисом`,

				PolishLen: `Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzi`,
				PolishMax: `Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`,
				PolishMin: `Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`,

				GermanLen: `Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg`,
				GermanMax: `Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`,
				GermanMin: `Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`,
			},
			want: []string{
				"cyrillic_len::length_is_not::70", "polish_len::length_is_not::111", "german_len::length_is_not::120",
			},
		},

		{
			name: "should max errors on each language",
			fields: fields{
				CyrillicLen: `Язык общего назначения с широкими возможностями и понятным синтаксисом`,
				CyrillicMax: `Язык общего назначения с широкими возможностями и понятным синтаксисом.`,
				CyrillicMin: `Язык общего назначения с широкими возможностями и понятным синтаксисом`,

				PolishLen: `Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`,
				PolishMax: `Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić.`,
				PolishMin: `Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`,

				GermanLen: `Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`,
				GermanMax: `Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg..`,
				GermanMin: `Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`,
			},
			want: []string{
				"cyrillic_max::max_length_is::70", "polish_max::max_length_is::111", "german_max::max_length_is::120",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Checker{
				CyrillicLen: tt.fields.CyrillicLen,
				PolishLen:   tt.fields.PolishLen,
				GermanLen:   tt.fields.GermanLen,
				CyrillicMin: tt.fields.CyrillicMin,
				PolishMin:   tt.fields.PolishMin,
				GermanMin:   tt.fields.GermanMin,
				CyrillicMax: tt.fields.CyrillicMax,
				PolishMax:   tt.fields.PolishMax,
				GermanMax:   tt.fields.GermanMax,
			}
			if got := c.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckerRef_Validate(t *testing.T) {
	type fields struct {
		CyrillicLen *string
		PolishLen   *string
		GermanLen   *string

		CyrillicMin *string
		PolishMin   *string
		GermanMin   *string

		CyrillicMax *string
		PolishMax   *string
		GermanMax   *string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "should success",
			fields: fields{
				CyrillicLen: sToP(`Язык общего назначения с широкими возможностями и понятным синтаксисом`),
				CyrillicMax: sToP(`Язык общего назначения с широкими возможностями и понятным синтаксисом`),
				CyrillicMin: sToP(`Язык общего назначения с широкими возможностями и понятным синтаксисом`),

				PolishLen: sToP(`Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`),
				PolishMax: sToP(`Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`),
				PolishMin: sToP(`Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`),

				GermanLen: sToP(`Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`),
				GermanMax: sToP(`Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`),
				GermanMin: sToP(`Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`),
			},
			want: nil,
		},

		{
			name: "should min errors on each language",
			fields: fields{
				CyrillicLen: sToP(`Язык общего назначения с широкими возможностями и понятным синтаксисом`),
				CyrillicMax: sToP(`Язык общего назначения с широкими возможностями и понятным синтаксисом`),
				CyrillicMin: sToP(`Язык общего назначения с широкими возможностями и понятным синтаксисо`),

				PolishLen: sToP(`Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`),
				PolishMax: sToP(`Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`),
				PolishMin: sToP(`Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzi`),

				GermanLen: sToP(`Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`),
				GermanMax: sToP(`Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`),
				GermanMin: sToP(`Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg`),
			},
			want: []string{
				"cyrillic_min::min_length_is::70", "polish_min::min_length_is::111", "german_min::min_length_is::120",
			},
		},

		{
			name: "should len errors on each language",
			fields: fields{
				CyrillicLen: sToP(`Язык общего назначения с широкими возможностями и понятным синтаксисо`),
				CyrillicMax: sToP(`Язык общего назначения с широкими возможностями и понятным синтаксисом`),
				CyrillicMin: sToP(`Язык общего назначения с широкими возможностями и понятным синтаксисом`),

				PolishLen: sToP(`Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzi`),
				PolishMax: sToP(`Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`),
				PolishMin: sToP(`Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`),

				GermanLen: sToP(`Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg`),
				GermanMax: sToP(`Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`),
				GermanMin: sToP(`Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`),
			},
			want: []string{
				"cyrillic_len::length_is_not::70", "polish_len::length_is_not::111", "german_len::length_is_not::120",
			},
		},

		{
			name: "should max errors on each language",
			fields: fields{
				CyrillicLen: sToP(`Язык общего назначения с широкими возможностями и понятным синтаксисом`),
				CyrillicMax: sToP(`Язык общего назначения с широкими возможностями и понятным синтаксисом.`),
				CyrillicMin: sToP(`Язык общего назначения с широкими возможностями и понятным синтаксисом`),

				PolishLen: sToP(`Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`),
				PolishMax: sToP(`Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić.`),
				PolishMin: sToP(`Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`),

				GermanLen: sToP(`Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`),
				GermanMax: sToP(`Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg..`),
				GermanMin: sToP(`Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`),
			},
			want: []string{
				"cyrillic_max::max_length_is::70", "polish_max::max_length_is::111", "german_max::max_length_is::120",
			},
		},

		{
			name: "should max errors on each language",
			fields: fields{
				CyrillicLen: sToP(`Язык общего назначения с широкими возможностями и понятным синтаксисом`),
				CyrillicMax: sToP(`Язык общего назначения с широкими возможностями и понятным синтаксисом.`),
				CyrillicMin: sToP(`Язык общего назначения с широкими возможностями и понятным синтаксисом`),

				PolishLen: sToP(`Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`),
				PolishMax: sToP(`Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić.`),
				PolishMin: sToP(`Ola szykuje się do szkoły. Jest już w piątej klasie. Dawniej obawiała się szkoły, teraz bardzo lubi tam chodzić`),

				GermanLen: sToP(`Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`),
				GermanMax: sToP(`Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg..`),
				GermanMin: sToP(`Juliana kommt aus Paris. Das ist die Hauptstadt von Frankreich. In diesem Sommer macht sie einen Sprachkurs in Freiburg.`),
			},
			want: []string{
				"cyrillic_max::max_length_is::70", "polish_max::max_length_is::111", "german_max::max_length_is::120",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CheckerRef{
				CyrillicLen: tt.fields.CyrillicLen,
				PolishLen:   tt.fields.PolishLen,
				GermanLen:   tt.fields.GermanLen,
				CyrillicMin: tt.fields.CyrillicMin,
				PolishMin:   tt.fields.PolishMin,
				GermanMin:   tt.fields.GermanMin,
				CyrillicMax: tt.fields.CyrillicMax,
				PolishMax:   tt.fields.PolishMax,
				GermanMax:   tt.fields.GermanMax,
			}
			if got := c.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckerSlice_Validate(t *testing.T) {
	type fields struct {
		Multiple     []string
		MultipleRefs []*string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "should success on multiple strings in slice",
			fields: fields{
				Multiple: []string{
					"Most people think a normal body temperature is an oral temperature (by mouth) of 98.6°F (37°C). This is an average of normal body temperatures. Your normal temperature may actually be 1°F (0.6°C) or more above or below this. Also, your normal temperature changes by as much as 1°F (0.6°C) during the day, depending on how active you are and the time of day. Body temperature is very sensitive to hormone levels. So a woman's temperature may be higher or lower when she is ovulating or having her menstrual period.",
					"Most people think a normal body temperature is an oral temperature (by mouth) of 98.6°F (37°C). This is an average of normal body temperatures. Your normal temperature may actually be 1°F (0.6°C) or more above or below this. Also, your normal temperature changes by as much as 1°F (0.6°C) during the day, depending on how active you are and the time of day. Body temperature is very sensitive to hormone levels. So a woman's temperature may be higher or lower when she is ovulating or having her menstrual period.",
				},
				MultipleRefs: []*string{
					sToP("Most people think a normal body temperature is an oral temperature (by mouth) of 98.6°F (37°C). This is an average of normal body temperatures. Your normal temperature may actually be 1°F (0.6°C) or more above or below this. Also, your normal temperature changes by as much as 1°F (0.6°C) during the day, depending on how active you are and the time of day. Body temperature is very sensitive to hormone levels. So a woman's temperature may be higher or lower when she is ovulating or having her menstrual period."),
					sToP("Most people think a normal body temperature is an oral temperature (by mouth) of 98.6°F (37°C). This is an average of normal body temperatures. Your normal temperature may actually be 1°F (0.6°C) or more above or below this. Also, your normal temperature changes by as much as 1°F (0.6°C) during the day, depending on how active you are and the time of day. Body temperature is very sensitive to hormone levels. So a woman's temperature may be higher or lower when she is ovulating or having her menstrual period."),
				},
			},
			want: nil,
		},

		{
			name: "should error on second value multiple strings in slice",
			fields: fields{
				Multiple: []string{
					"Most people think a normal body temperature is an oral temperature (by mouth) of 98.6°F (37°C). This is an average of normal body temperatures. Your normal temperature may actually be 1°F (0.6°C) or more above or below this. Also, your normal temperature changes by as much as 1°F (0.6°C) during the day, depending on how active you are and the time of day. Body temperature is very sensitive to hormone levels. So a woman's temperature may be higher or lower when she is ovulating or having her menstrual period.",
					"ost people think a normal body temperature is an oral temperature (by mouth) of 98.6°F (37°C). This is an average of normal body temperatures. Your normal temperature may actually be 1°F (0.6°C) or more above or below this. Also, your normal temperature changes by as much as 1°F (0.6°C) during the day, depending on how active you are and the time of day. Body temperature is very sensitive to hormone levels. So a woman's temperature may be higher or lower when she is ovulating or having her menstrual period.",
				},

				MultipleRefs: []*string{
					sToP("Most people think a normal body temperature is an oral temperature (by mouth) of 98.6°F (37°C). This is an average of normal body temperatures. Your normal temperature may actually be 1°F (0.6°C) or more above or below this. Also, your normal temperature changes by as much as 1°F (0.6°C) during the day, depending on how active you are and the time of day. Body temperature is very sensitive to hormone levels. So a woman's temperature may be higher or lower when she is ovulating or having her menstrual period."),
					sToP("ost people think a normal body temperature is an oral temperature (by mouth) of 98.6°F (37°C). This is an average of normal body temperatures. Your normal temperature may actually be 1°F (0.6°C) or more above or below this. Also, your normal temperature changes by as much as 1°F (0.6°C) during the day, depending on how active you are and the time of day. Body temperature is very sensitive to hormone levels. So a woman's temperature may be higher or lower when she is ovulating or having her menstrual period."),
				},
			},
			want: []string{
				"multiple_1::length_is_not::513", "multiple_refs_1::length_is_not::513",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CheckerSlice{
				Multiple:     tt.fields.Multiple,
				MultipleRefs: tt.fields.MultipleRefs,
			}
			if got := c.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sToP(s string) *string {
	return &s
}
