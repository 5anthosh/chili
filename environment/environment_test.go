package environment

import (
	"testing"

	"github.com/5anthosh/chili/function"
)

type fields struct {
	symbolTable map[string]uint
	variables   map[string]interface{}
	functions   map[string]function.Function
}

func TestEnvironment_IsFunction(t *testing.T) {
	type args struct {
		name string
	}
	test1 := fields{}
	test1.symbolTable = map[string]uint{
		"testfunc": functionType,
	}

	test2 := fields{}
	test2.symbolTable = map[string]uint{
		"testfunc-1": functionType,
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "Is function is true",
			fields: test1,
			args: args{
				name: "testfunc",
			},
			want: true,
		},
		{
			name:   "Is function is false",
			fields: test2,
			args: args{
				name: "testfunc",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Environment{
				symbolTable: tt.fields.symbolTable,
				variables:   tt.fields.variables,
				functions:   tt.fields.functions,
			}
			if got := e.IsFunction(tt.args.name); got != tt.want {
				t.Errorf("Environment.IsFunction() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvironment_SetFunction(t *testing.T) {
	type args struct {
		function function.Function
	}

	func1 := function.Function{
		Name: "test-func",
	}

	func2 := function.Function{
		Name: "test-func2",
	}

	fieldsTest := fields{}
	fieldsTest.functions = make(map[string]function.Function)
	fieldsTest.symbolTable = make(map[string]uint)
	fieldsTest.variables = make(map[string]interface{})
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "SetFunction success",
			fields: fieldsTest,
			args: args{
				function: func1,
			},
			wantErr: false,
		},
		{
			name:   "SetFunction success (2)",
			fields: fieldsTest,
			args: args{
				function: func2,
			},
			wantErr: false,
		},
		{
			name:   "SetFunction failure",
			fields: fieldsTest,
			args: args{
				function: func1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Environment{
				symbolTable: tt.fields.symbolTable,
				variables:   tt.fields.variables,
				functions:   tt.fields.functions,
			}
			if err := e.SetFunction(tt.args.function); (err != nil) != tt.wantErr {
				t.Errorf("Environment.SetFunction() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
