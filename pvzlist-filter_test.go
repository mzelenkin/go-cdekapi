package cdekapi

import (
	"reflect"
	"testing"
)

func TestOfficesRequestParams(t *testing.T) {
	type fields struct {
		params map[string]string
	}

	type argument struct {
		name  OfficesFilterParamName
		value string
	}

	tests := []struct {
		name   string
		fields fields
		args   argument
		want   *OfficesFilter
	}{
		{
			name:   "add single argument",
			fields: fields{params: nil},
			args: argument{
				name:  OfficesFilterCityID,
				value: "33",
			},
			want: &OfficesFilter{params: map[string]string{string(OfficesFilterCityID): "33"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filterBuilder := &OfficesFilter{
				params: tt.fields.params,
			}
			if got := filterBuilder.AddParam(tt.args.name, tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}
