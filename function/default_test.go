package function

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

func Test_absImpl(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "Run abs for negative values",
			args: args{
				args: []interface{}{decimal.NewFromInt(-234)},
			},
			want:    decimal.NewFromInt(234),
			wantErr: false,
		},
		{
			name: "Run abs for positive values",
			args: args{
				args: []interface{}{decimal.NewFromFloat32(34.345)},
			},
			want:    decimal.NewFromFloat32(34.345),
			wantErr: false,
		},

		{
			name: "Run abs for string",
			args: args{
				args: []interface{}{"hello"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := absImpl(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("absImpl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("absImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}
