package utils

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func Pointer[T any](v T) *T {
	return &v
}

func NullableString(s *string) types.String {
	if s != nil {
		return types.StringValue(*s)
	}
	return types.StringNull()
}

func NullableBool(b *bool) types.Bool {
	if b != nil {
		return types.BoolValue(*b)
	}
	return types.BoolNull()
}

func NullableInt64(i *int64) types.Int64 {
	if i != nil {
		return types.Int64Value(*i)
	}
	return types.Int64Null()
}

func NullableObject[T any, R any](source *T, value R) *R {
	if source != nil {
		return &value
	}

	return nil
}

func NullableFloat64(f *float64) types.Float64 {
	if f != nil {
		return types.Float64Value(*f)
	}
	return types.Float64Null()
}

func MapList[T, R any](from *[]T, f func(T) R) *[]R {
	to := make([]R, len(*from))
	for i, v := range *from {
		to[i] = f(v)
	}
	return &to
}

func ToList(ctx context.Context, from any, toType attr.Type, diagnostics *diag.Diagnostics) types.List {
	result, err := types.ListValueFrom(ctx, toType, from)
	if err != nil {
		diagnostics.Append(err.Warnings()...)
		diagnostics.Append(err.Errors()...)
		return types.ListUnknown(toType)
	}

	return result
}

func FromListToPrimitiveSlice[T any](ctx context.Context, from types.List, toType attr.Type, diagnostics *diag.Diagnostics) *[]T {
	elements := from.Elements()
	result := make([]T, len(elements))
	for i, elem := range elements {
		conversionMethod, err := getConversionMethodName(toType)
		if err != nil {
			diagnostics.Append(diag.NewErrorDiagnostic("conversion error", err.Error()))
			return nil
		}

		res := reflect.ValueOf(elem).MethodByName(conversionMethod).Call([]reflect.Value{})
		result[i] = res[0].Interface().(T)

	}
	return &result
}

func getConversionMethodName(t attr.Type) (string, error) {
	if t.Equal(types.StringType) {
		return "ValueString", nil
	} else if t.Equal(types.BoolType) {
		return "ValueBool", nil
	} else if t.Equal(types.Float64Type) {
		return "ValueFloat64", nil
	} else if t.Equal(types.Int64Type) {
		return "ValueInt64", nil
	} else {
		return "", fmt.Errorf("unsupported type %s", t.String())
	}
}
