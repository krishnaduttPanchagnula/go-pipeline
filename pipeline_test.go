package pipeline

import (
	"errors"
	"reflect"
	"testing"
)

func TestExecute(t *testing.T) {
	tests := []struct {
		name        string
		functions   []func(interface{}) (interface{}, error)
		input       interface{}
		expected    interface{}
		expectedErr error
	}{
		{
			name: "Test with integers",
			functions: []func(interface{}) (interface{}, error){
				addTwo,
				multiplyByThree,
			},
			input:       5,
			expected:    21,
			expectedErr: nil,
		},
		{
			name: "Test with integer slice",
			functions: []func(interface{}) (interface{}, error){
				addTwo,
				multiplyByThree,
			},
			input:       []int{1, 2, 3},
			expected:    []int{9, 12, 15},
			expectedErr: nil,
		},
		{
			name: "Test with unsupported type",
			functions: []func(interface{}) (interface{}, error){
				addTwo,
				multiplyByThree,
			},
			input:       "unsupported",
			expected:    nil,
			expectedErr: errors.New("unsupported input type"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPipeline()
			for _, f := range tt.functions {
				p.AddFunction(f)
			}

			result, err := p.Execute(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected result %v, but got %v", tt.expected, result)
			}

			if !reflect.DeepEqual(err, tt.expectedErr) {
				t.Errorf("Expected error %v, but got %v", tt.expectedErr, err)
			}
		})
	}
}
func addTwo(x interface{}) (interface{}, error) {
	switch val := x.(type) {
	case int:
		return val + 2, nil
	case []int:
		// Add 2 to each element in the slice
		for i := range val {
			val[i] += 2
		}
		return val, nil
	default:
		return nil, errors.New("unsupported input type")
	}
}

func multiplyByThree(x interface{}) (interface{}, error) {
	switch val := x.(type) {
	case int:
		return val * 3, nil
	case []int:
		// Multiply each element in the slice by 3
		for i := range val {
			val[i] *= 3
		}
		return val, nil
	default:
		return nil, errors.New("unsupported input type")
	}
}
