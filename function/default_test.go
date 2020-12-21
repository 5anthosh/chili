package function

import (
	"reflect"
	"testing"

	"github.com/shopspring/decimal"
)

type args struct {
	args []interface{}
}

type testArgs struct {
	name    string
	args    args
	want    interface{}
	wantErr bool
}

func Test_absImpl(t *testing.T) {
	tests := []testArgs{
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

func Test_cbrtImpl(t *testing.T) {
	tests := []testArgs{
		{
			name: "cbrt for positive integer",
			args: args{
				args: []interface{}{decimal.NewFromInt(27)},
			},
			want:    three,
			wantErr: false,
		},
		{
			name: "cbrt for negative integer",
			args: args{
				args: []interface{}{decimal.NewFromInt(-27)},
			},
			want:    three.Neg(),
			wantErr: false,
		},
		{
			name: "cbrt for decimal number",
			args: args{
				args: []interface{}{
					decimal.NewFromFloat(0.45),
				},
			},
			want:    decimal.NewFromFloat(0.7663094323935531),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cbrtImpl(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("cbrtImpl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cbrtImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ceilImpl(t *testing.T) {
	tests := []testArgs{
		{
			name: "ceil for positive decimal",
			args: args{
				args: []interface{}{decimal.NewFromFloat(2.2)},
			},
			want:    three,
			wantErr: false,
		},
		{
			name: "ceil for negative decimal",
			args: args{
				args: []interface{}{decimal.NewFromFloat(2.4).Neg()},
			},
			want:    two.Neg(),
			wantErr: false,
		},
		{
			name: "ceil for positive integer",
			args: args{
				args: []interface{}{three},
			},
			want:    three,
			wantErr: false,
		},
		{
			name: "ceil for negative integer",
			args: args{
				args: []interface{}{three.Neg()},
			},
			want:    three.Neg(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ceilImpl(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ceilImpl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ceilImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}
