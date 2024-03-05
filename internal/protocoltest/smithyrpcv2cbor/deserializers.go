// Code generated by smithy-go-codegen DO NOT EDIT.

package smithyrpcv2cbor

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/smithyrpcv2cbor/types"
	smithy "github.com/aws/smithy-go"
	smithycbor "github.com/aws/smithy-go/encoding/cbor"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
	"io/ioutil"
	"strings"
	"time"
)

type smithyRpcv2cbor_deserializeOpEmptyInputOutput struct {
}

func (*smithyRpcv2cbor_deserializeOpEmptyInputOutput) ID() string {
	return "OperationDeserializer"
}

func (m *smithyRpcv2cbor_deserializeOpEmptyInputOutput) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	resp, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", out.RawResponse)
	}

	if resp.Header.Get("smithy-protocol") != "rpc-v2-cbor" {
		return out, metadata, &smithy.DeserializationError{
			Err: fmt.Errorf(
				"unexpected smithy-protocol response header '%s' (HTTP status: %s)",
				resp.Header.Get("smithy-protocol"),
				resp.Status,
			),
		}
	}

	if resp.StatusCode != 200 {
		return out, metadata, rpc2_deserializeOpErrorEmptyInputOutput(resp)
	}

	if _, err = io.Copy(ioutil.Discard, resp.Body); err != nil {
		return out, metadata, fmt.Errorf("discard response body: %w", err)
	}

	out.Result = &EmptyInputOutputOutput{}

	return out, metadata, nil
}

type smithyRpcv2cbor_deserializeOpFractionalSeconds struct {
}

func (*smithyRpcv2cbor_deserializeOpFractionalSeconds) ID() string {
	return "OperationDeserializer"
}

func (m *smithyRpcv2cbor_deserializeOpFractionalSeconds) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	resp, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", out.RawResponse)
	}

	if resp.Header.Get("smithy-protocol") != "rpc-v2-cbor" {
		return out, metadata, &smithy.DeserializationError{
			Err: fmt.Errorf(
				"unexpected smithy-protocol response header '%s' (HTTP status: %s)",
				resp.Header.Get("smithy-protocol"),
				resp.Status,
			),
		}
	}

	if resp.StatusCode != 200 {
		return out, metadata, rpc2_deserializeOpErrorFractionalSeconds(resp)
	}

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return out, metadata, err
	}

	if len(payload) == 0 {
		out.Result = &FractionalSecondsOutput{}
		return out, metadata, nil
	}

	cv, err := smithycbor.Decode(payload)
	if err != nil {
		return out, metadata, err
	}

	output, err := deserializeCBOR_FractionalSecondsOutput(cv)
	if err != nil {
		return out, metadata, err
	}

	out.Result = output

	return out, metadata, nil
}

type smithyRpcv2cbor_deserializeOpGreetingWithErrors struct {
}

func (*smithyRpcv2cbor_deserializeOpGreetingWithErrors) ID() string {
	return "OperationDeserializer"
}

func (m *smithyRpcv2cbor_deserializeOpGreetingWithErrors) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	resp, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", out.RawResponse)
	}

	if resp.Header.Get("smithy-protocol") != "rpc-v2-cbor" {
		return out, metadata, &smithy.DeserializationError{
			Err: fmt.Errorf(
				"unexpected smithy-protocol response header '%s' (HTTP status: %s)",
				resp.Header.Get("smithy-protocol"),
				resp.Status,
			),
		}
	}

	if resp.StatusCode != 200 {
		return out, metadata, rpc2_deserializeOpErrorGreetingWithErrors(resp)
	}

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return out, metadata, err
	}

	if len(payload) == 0 {
		out.Result = &GreetingWithErrorsOutput{}
		return out, metadata, nil
	}

	cv, err := smithycbor.Decode(payload)
	if err != nil {
		return out, metadata, err
	}

	output, err := deserializeCBOR_GreetingWithErrorsOutput(cv)
	if err != nil {
		return out, metadata, err
	}

	out.Result = output

	return out, metadata, nil
}

type smithyRpcv2cbor_deserializeOpNoInputOutput struct {
}

func (*smithyRpcv2cbor_deserializeOpNoInputOutput) ID() string {
	return "OperationDeserializer"
}

func (m *smithyRpcv2cbor_deserializeOpNoInputOutput) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	resp, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", out.RawResponse)
	}

	if resp.Header.Get("smithy-protocol") != "rpc-v2-cbor" {
		return out, metadata, &smithy.DeserializationError{
			Err: fmt.Errorf(
				"unexpected smithy-protocol response header '%s' (HTTP status: %s)",
				resp.Header.Get("smithy-protocol"),
				resp.Status,
			),
		}
	}

	if resp.StatusCode != 200 {
		return out, metadata, rpc2_deserializeOpErrorNoInputOutput(resp)
	}

	if _, err = io.Copy(ioutil.Discard, resp.Body); err != nil {
		return out, metadata, fmt.Errorf("discard response body: %w", err)
	}

	out.Result = &NoInputOutputOutput{}

	return out, metadata, nil
}

type smithyRpcv2cbor_deserializeOpOptionalInputOutput struct {
}

func (*smithyRpcv2cbor_deserializeOpOptionalInputOutput) ID() string {
	return "OperationDeserializer"
}

func (m *smithyRpcv2cbor_deserializeOpOptionalInputOutput) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	resp, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", out.RawResponse)
	}

	if resp.Header.Get("smithy-protocol") != "rpc-v2-cbor" {
		return out, metadata, &smithy.DeserializationError{
			Err: fmt.Errorf(
				"unexpected smithy-protocol response header '%s' (HTTP status: %s)",
				resp.Header.Get("smithy-protocol"),
				resp.Status,
			),
		}
	}

	if resp.StatusCode != 200 {
		return out, metadata, rpc2_deserializeOpErrorOptionalInputOutput(resp)
	}

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return out, metadata, err
	}

	if len(payload) == 0 {
		out.Result = &OptionalInputOutputOutput{}
		return out, metadata, nil
	}

	cv, err := smithycbor.Decode(payload)
	if err != nil {
		return out, metadata, err
	}

	output, err := deserializeCBOR_OptionalInputOutputOutput(cv)
	if err != nil {
		return out, metadata, err
	}

	out.Result = output

	return out, metadata, nil
}

type smithyRpcv2cbor_deserializeOpRecursiveShapes struct {
}

func (*smithyRpcv2cbor_deserializeOpRecursiveShapes) ID() string {
	return "OperationDeserializer"
}

func (m *smithyRpcv2cbor_deserializeOpRecursiveShapes) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	resp, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", out.RawResponse)
	}

	if resp.Header.Get("smithy-protocol") != "rpc-v2-cbor" {
		return out, metadata, &smithy.DeserializationError{
			Err: fmt.Errorf(
				"unexpected smithy-protocol response header '%s' (HTTP status: %s)",
				resp.Header.Get("smithy-protocol"),
				resp.Status,
			),
		}
	}

	if resp.StatusCode != 200 {
		return out, metadata, rpc2_deserializeOpErrorRecursiveShapes(resp)
	}

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return out, metadata, err
	}

	if len(payload) == 0 {
		out.Result = &RecursiveShapesOutput{}
		return out, metadata, nil
	}

	cv, err := smithycbor.Decode(payload)
	if err != nil {
		return out, metadata, err
	}

	output, err := deserializeCBOR_RecursiveShapesOutput(cv)
	if err != nil {
		return out, metadata, err
	}

	out.Result = output

	return out, metadata, nil
}

type smithyRpcv2cbor_deserializeOpRpcV2CborLists struct {
}

func (*smithyRpcv2cbor_deserializeOpRpcV2CborLists) ID() string {
	return "OperationDeserializer"
}

func (m *smithyRpcv2cbor_deserializeOpRpcV2CborLists) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	resp, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", out.RawResponse)
	}

	if resp.Header.Get("smithy-protocol") != "rpc-v2-cbor" {
		return out, metadata, &smithy.DeserializationError{
			Err: fmt.Errorf(
				"unexpected smithy-protocol response header '%s' (HTTP status: %s)",
				resp.Header.Get("smithy-protocol"),
				resp.Status,
			),
		}
	}

	if resp.StatusCode != 200 {
		return out, metadata, rpc2_deserializeOpErrorRpcV2CborLists(resp)
	}

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return out, metadata, err
	}

	if len(payload) == 0 {
		out.Result = &RpcV2CborListsOutput{}
		return out, metadata, nil
	}

	cv, err := smithycbor.Decode(payload)
	if err != nil {
		return out, metadata, err
	}

	output, err := deserializeCBOR_RpcV2CborListsOutput(cv)
	if err != nil {
		return out, metadata, err
	}

	out.Result = output

	return out, metadata, nil
}

type smithyRpcv2cbor_deserializeOpRpcV2CborMaps struct {
}

func (*smithyRpcv2cbor_deserializeOpRpcV2CborMaps) ID() string {
	return "OperationDeserializer"
}

func (m *smithyRpcv2cbor_deserializeOpRpcV2CborMaps) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	resp, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", out.RawResponse)
	}

	if resp.Header.Get("smithy-protocol") != "rpc-v2-cbor" {
		return out, metadata, &smithy.DeserializationError{
			Err: fmt.Errorf(
				"unexpected smithy-protocol response header '%s' (HTTP status: %s)",
				resp.Header.Get("smithy-protocol"),
				resp.Status,
			),
		}
	}

	if resp.StatusCode != 200 {
		return out, metadata, rpc2_deserializeOpErrorRpcV2CborMaps(resp)
	}

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return out, metadata, err
	}

	if len(payload) == 0 {
		out.Result = &RpcV2CborMapsOutput{}
		return out, metadata, nil
	}

	cv, err := smithycbor.Decode(payload)
	if err != nil {
		return out, metadata, err
	}

	output, err := deserializeCBOR_RpcV2CborMapsOutput(cv)
	if err != nil {
		return out, metadata, err
	}

	out.Result = output

	return out, metadata, nil
}

type smithyRpcv2cbor_deserializeOpSimpleScalarProperties struct {
}

func (*smithyRpcv2cbor_deserializeOpSimpleScalarProperties) ID() string {
	return "OperationDeserializer"
}

func (m *smithyRpcv2cbor_deserializeOpSimpleScalarProperties) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	resp, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, fmt.Errorf("unexpected transport type %T", out.RawResponse)
	}

	if resp.Header.Get("smithy-protocol") != "rpc-v2-cbor" {
		return out, metadata, &smithy.DeserializationError{
			Err: fmt.Errorf(
				"unexpected smithy-protocol response header '%s' (HTTP status: %s)",
				resp.Header.Get("smithy-protocol"),
				resp.Status,
			),
		}
	}

	if resp.StatusCode != 200 {
		return out, metadata, rpc2_deserializeOpErrorSimpleScalarProperties(resp)
	}

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return out, metadata, err
	}

	if len(payload) == 0 {
		out.Result = &SimpleScalarPropertiesOutput{}
		return out, metadata, nil
	}

	cv, err := smithycbor.Decode(payload)
	if err != nil {
		return out, metadata, err
	}

	output, err := deserializeCBOR_SimpleScalarPropertiesOutput(cv)
	if err != nil {
		return out, metadata, err
	}

	out.Result = output

	return out, metadata, nil
}
func deserializeCBOR_RecursiveShapesInputOutputNested1(v smithycbor.Value) (*types.RecursiveShapesInputOutputNested1, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	ds := &types.RecursiveShapesInputOutputNested1{}
	for key, sv := range av {
		_, _ = key, sv
		if key == "foo" {
			dv, err := deserializeCBOR_String(sv)
			if err != nil {
				return nil, err
			}
			ds.Foo = ptr.String(dv)
		}

		if key == "nested" {
			dv, err := deserializeCBOR_RecursiveShapesInputOutputNested2(sv)
			if err != nil {
				return nil, err
			}
			ds.Nested = dv
		}
	}
	return ds, nil
}

func deserializeCBOR_DenseBooleanMap(v smithycbor.Value) (map[string]bool, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	dm := map[string]bool{}
	for key, sv := range av {

		dv, err := deserializeCBOR_Bool(sv)
		if err != nil {
			return nil, err
		}
		dm[key] = dv
	}
	return dm, nil
}

func deserializeCBOR_StructureList(v smithycbor.Value) ([]types.StructureListMember, error) {
	av, ok := v.(smithycbor.List)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	var dl []types.StructureListMember
	for _, si := range av {

		di, err := deserializeCBOR_StructureListMember(si)
		if err != nil {
			return nil, err
		}
		dl = append(dl, *di)
	}
	return dl, nil
}

func deserializeCBOR_ComplexError(v smithycbor.Value) (*types.ComplexError, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	ds := &types.ComplexError{}
	for key, sv := range av {
		_, _ = key, sv
		if key == "TopLevel" {
			dv, err := deserializeCBOR_String(sv)
			if err != nil {
				return nil, err
			}
			ds.TopLevel = ptr.String(dv)
		}

		if key == "Nested" {
			dv, err := deserializeCBOR_ComplexNestedErrorData(sv)
			if err != nil {
				return nil, err
			}
			ds.Nested = dv
		}
	}
	return ds, nil
}

func deserializeCBOR_BooleanList(v smithycbor.Value) ([]bool, error) {
	av, ok := v.(smithycbor.List)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	var dl []bool
	for _, si := range av {

		di, err := deserializeCBOR_Bool(si)
		if err != nil {
			return nil, err
		}
		dl = append(dl, di)
	}
	return dl, nil
}

func deserializeCBOR_SparseStructMap(v smithycbor.Value) (map[string]*types.GreetingStruct, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	dm := map[string]*types.GreetingStruct{}
	for key, sv := range av {
		if _, ok := sv.(*smithycbor.Nil); ok {
			dm[key] = nil
			continue
		}
		dv, err := deserializeCBOR_GreetingStruct(sv)
		if err != nil {
			return nil, err
		}
		dm[key] = dv
	}
	return dm, nil
}

func deserializeCBOR_GreetingStruct(v smithycbor.Value) (*types.GreetingStruct, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	ds := &types.GreetingStruct{}
	for key, sv := range av {
		_, _ = key, sv
		if key == "hi" {
			dv, err := deserializeCBOR_String(sv)
			if err != nil {
				return nil, err
			}
			ds.Hi = ptr.String(dv)
		}
	}
	return ds, nil
}

func deserializeCBOR_SparseStringMap(v smithycbor.Value) (map[string]*string, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	dm := map[string]*string{}
	for key, sv := range av {
		if _, ok := sv.(*smithycbor.Nil); ok {
			dm[key] = nil
			continue
		}
		dv, err := deserializeCBOR_String(sv)
		if err != nil {
			return nil, err
		}
		dm[key] = &dv
	}
	return dm, nil
}

func deserializeCBOR_SparseSetMap(v smithycbor.Value) (map[string][]string, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	dm := map[string][]string{}
	for key, sv := range av {
		if _, ok := sv.(*smithycbor.Nil); ok {
			dm[key] = nil
			continue
		}
		dv, err := deserializeCBOR_StringSet(sv)
		if err != nil {
			return nil, err
		}
		dm[key] = dv
	}
	return dm, nil
}

func deserializeCBOR_Float64(v smithycbor.Value) (float64, error) {
	return smithycbor.AsFloat64(v)
}

func deserializeCBOR_SparseStringList(v smithycbor.Value) ([]*string, error) {
	av, ok := v.(smithycbor.List)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	var dl []*string
	for _, si := range av {
		if _, ok := si.(*smithycbor.Nil); ok {
			dl = append(dl, nil)
			continue
		}
		di, err := deserializeCBOR_String(si)
		if err != nil {
			return nil, err
		}
		dl = append(dl, &di)
	}
	return dl, nil
}

func deserializeCBOR_InvalidGreeting(v smithycbor.Value) (*types.InvalidGreeting, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	ds := &types.InvalidGreeting{}
	for key, sv := range av {
		_, _ = key, sv
		if key == "Message" {
			dv, err := deserializeCBOR_String(sv)
			if err != nil {
				return nil, err
			}
			ds.Message = ptr.String(dv)
		}
	}
	return ds, nil
}

func deserializeCBOR_DenseStringMap(v smithycbor.Value) (map[string]string, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	dm := map[string]string{}
	for key, sv := range av {

		dv, err := deserializeCBOR_String(sv)
		if err != nil {
			return nil, err
		}
		dm[key] = dv
	}
	return dm, nil
}

func deserializeCBOR_StringSet(v smithycbor.Value) ([]string, error) {
	av, ok := v.(smithycbor.List)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	var dl []string
	for _, si := range av {

		di, err := deserializeCBOR_String(si)
		if err != nil {
			return nil, err
		}
		dl = append(dl, di)
	}
	return dl, nil
}

func deserializeCBOR_Int32(v smithycbor.Value) (int32, error) {
	return smithycbor.AsInt32(v)
}

func deserializeCBOR_RecursiveShapesOutput(v smithycbor.Value) (*RecursiveShapesOutput, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	ds := &RecursiveShapesOutput{}
	for key, sv := range av {
		_, _ = key, sv
		if key == "nested" {
			dv, err := deserializeCBOR_RecursiveShapesInputOutputNested1(sv)
			if err != nil {
				return nil, err
			}
			ds.Nested = dv
		}
	}
	return ds, nil
}

func deserializeCBOR_Int16(v smithycbor.Value) (int16, error) {
	return smithycbor.AsInt16(v)
}

func deserializeCBOR_String(v smithycbor.Value) (string, error) {
	av, ok := v.(smithycbor.String)
	if !ok {
		return "", fmt.Errorf("unexpected value type %T", v)
	}
	return string(av), nil
}

func deserializeCBOR_StringList(v smithycbor.Value) ([]string, error) {
	av, ok := v.(smithycbor.List)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	var dl []string
	for _, si := range av {

		di, err := deserializeCBOR_String(si)
		if err != nil {
			return nil, err
		}
		dl = append(dl, di)
	}
	return dl, nil
}

func deserializeCBOR_FractionalSecondsOutput(v smithycbor.Value) (*FractionalSecondsOutput, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	ds := &FractionalSecondsOutput{}
	for key, sv := range av {
		_, _ = key, sv
		if key == "datetime" {
			dv, err := deserializeCBOR_Time(sv)
			if err != nil {
				return nil, err
			}
			ds.Datetime = ptr.Time(dv)
		}
	}
	return ds, nil
}

func deserializeCBOR_DenseStructMap(v smithycbor.Value) (map[string]types.GreetingStruct, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	dm := map[string]types.GreetingStruct{}
	for key, sv := range av {

		dv, err := deserializeCBOR_GreetingStruct(sv)
		if err != nil {
			return nil, err
		}
		dm[key] = *dv
	}
	return dm, nil
}

func deserializeCBOR_SimpleScalarPropertiesOutput(v smithycbor.Value) (*SimpleScalarPropertiesOutput, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	ds := &SimpleScalarPropertiesOutput{}
	for key, sv := range av {
		_, _ = key, sv
		if key == "trueBooleanValue" {
			dv, err := deserializeCBOR_Bool(sv)
			if err != nil {
				return nil, err
			}
			ds.TrueBooleanValue = ptr.Bool(dv)
		}

		if key == "falseBooleanValue" {
			dv, err := deserializeCBOR_Bool(sv)
			if err != nil {
				return nil, err
			}
			ds.FalseBooleanValue = ptr.Bool(dv)
		}

		if key == "byteValue" {
			dv, err := deserializeCBOR_Int8(sv)
			if err != nil {
				return nil, err
			}
			ds.ByteValue = ptr.Int8(dv)
		}

		if key == "doubleValue" {
			dv, err := deserializeCBOR_Float64(sv)
			if err != nil {
				return nil, err
			}
			ds.DoubleValue = ptr.Float64(dv)
		}

		if key == "floatValue" {
			dv, err := deserializeCBOR_Float32(sv)
			if err != nil {
				return nil, err
			}
			ds.FloatValue = ptr.Float32(dv)
		}

		if key == "integerValue" {
			dv, err := deserializeCBOR_Int32(sv)
			if err != nil {
				return nil, err
			}
			ds.IntegerValue = ptr.Int32(dv)
		}

		if key == "longValue" {
			dv, err := deserializeCBOR_Int64(sv)
			if err != nil {
				return nil, err
			}
			ds.LongValue = ptr.Int64(dv)
		}

		if key == "shortValue" {
			dv, err := deserializeCBOR_Int16(sv)
			if err != nil {
				return nil, err
			}
			ds.ShortValue = ptr.Int16(dv)
		}

		if key == "stringValue" {
			dv, err := deserializeCBOR_String(sv)
			if err != nil {
				return nil, err
			}
			ds.StringValue = ptr.String(dv)
		}
	}
	return ds, nil
}

func deserializeCBOR_SparseBooleanMap(v smithycbor.Value) (map[string]*bool, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	dm := map[string]*bool{}
	for key, sv := range av {
		if _, ok := sv.(*smithycbor.Nil); ok {
			dm[key] = nil
			continue
		}
		dv, err := deserializeCBOR_Bool(sv)
		if err != nil {
			return nil, err
		}
		dm[key] = &dv
	}
	return dm, nil
}

func deserializeCBOR_StructureListMember(v smithycbor.Value) (*types.StructureListMember, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	ds := &types.StructureListMember{}
	for key, sv := range av {
		_, _ = key, sv
		if key == "a" {
			dv, err := deserializeCBOR_String(sv)
			if err != nil {
				return nil, err
			}
			ds.A = ptr.String(dv)
		}

		if key == "b" {
			dv, err := deserializeCBOR_String(sv)
			if err != nil {
				return nil, err
			}
			ds.B = ptr.String(dv)
		}
	}
	return ds, nil
}

func deserializeCBOR_Int8(v smithycbor.Value) (int8, error) {
	return smithycbor.AsInt8(v)
}

func deserializeCBOR_ComplexNestedErrorData(v smithycbor.Value) (*types.ComplexNestedErrorData, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	ds := &types.ComplexNestedErrorData{}
	for key, sv := range av {
		_, _ = key, sv
		if key == "Foo" {
			dv, err := deserializeCBOR_String(sv)
			if err != nil {
				return nil, err
			}
			ds.Foo = ptr.String(dv)
		}
	}
	return ds, nil
}

func deserializeCBOR_NestedStringList(v smithycbor.Value) ([][]string, error) {
	av, ok := v.(smithycbor.List)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	var dl [][]string
	for _, si := range av {
		if _, ok := si.(*smithycbor.Nil); ok {
			dl = append(dl, nil)
			continue
		}
		di, err := deserializeCBOR_StringList(si)
		if err != nil {
			return nil, err
		}
		dl = append(dl, di)
	}
	return dl, nil
}

func deserializeCBOR_TimestampList(v smithycbor.Value) ([]time.Time, error) {
	av, ok := v.(smithycbor.List)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	var dl []time.Time
	for _, si := range av {

		di, err := deserializeCBOR_Time(si)
		if err != nil {
			return nil, err
		}
		dl = append(dl, di)
	}
	return dl, nil
}

func deserializeCBOR_Bool(v smithycbor.Value) (bool, error) {
	av, ok := v.(smithycbor.Bool)
	if !ok {
		return false, fmt.Errorf("unexpected value type %T", v)
	}
	return bool(av), nil
}

func deserializeCBOR_DenseNumberMap(v smithycbor.Value) (map[string]int32, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	dm := map[string]int32{}
	for key, sv := range av {

		dv, err := deserializeCBOR_Int32(sv)
		if err != nil {
			return nil, err
		}
		dm[key] = dv
	}
	return dm, nil
}

func deserializeCBOR_Float32(v smithycbor.Value) (float32, error) {
	return smithycbor.AsFloat32(v)
}

func deserializeCBOR_FooEnum(v smithycbor.Value) (types.FooEnum, error) {
	av, ok := v.(smithycbor.String)
	if !ok {
		return types.FooEnum(""), fmt.Errorf("unexpected value type %T", v)
	}
	return types.FooEnum(av), nil
}

func deserializeCBOR_Time(v smithycbor.Value) (time.Time, error) {
	return smithycbor.AsTime(v)
}

func deserializeCBOR_IntegerList(v smithycbor.Value) ([]int32, error) {
	av, ok := v.(smithycbor.List)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	var dl []int32
	for _, si := range av {

		di, err := deserializeCBOR_Int32(si)
		if err != nil {
			return nil, err
		}
		dl = append(dl, di)
	}
	return dl, nil
}

func deserializeCBOR_IntegerEnum(v smithycbor.Value) (types.IntegerEnum, error) {
	av, err := smithycbor.AsInt32(v)
	if err != nil {
		return 0, err
	}
	return types.IntegerEnum(av), nil
}

func deserializeCBOR_SparseNumberMap(v smithycbor.Value) (map[string]*int32, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	dm := map[string]*int32{}
	for key, sv := range av {
		if _, ok := sv.(*smithycbor.Nil); ok {
			dm[key] = nil
			continue
		}
		dv, err := deserializeCBOR_Int32(sv)
		if err != nil {
			return nil, err
		}
		dm[key] = &dv
	}
	return dm, nil
}

func deserializeCBOR_RpcV2CborListsOutput(v smithycbor.Value) (*RpcV2CborListsOutput, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	ds := &RpcV2CborListsOutput{}
	for key, sv := range av {
		_, _ = key, sv
		if key == "stringList" {
			dv, err := deserializeCBOR_StringList(sv)
			if err != nil {
				return nil, err
			}
			ds.StringList = dv
		}

		if key == "sparseStringList" {
			dv, err := deserializeCBOR_SparseStringList(sv)
			if err != nil {
				return nil, err
			}
			ds.SparseStringList = dv
		}

		if key == "stringSet" {
			dv, err := deserializeCBOR_StringSet(sv)
			if err != nil {
				return nil, err
			}
			ds.StringSet = dv
		}

		if key == "integerList" {
			dv, err := deserializeCBOR_IntegerList(sv)
			if err != nil {
				return nil, err
			}
			ds.IntegerList = dv
		}

		if key == "booleanList" {
			dv, err := deserializeCBOR_BooleanList(sv)
			if err != nil {
				return nil, err
			}
			ds.BooleanList = dv
		}

		if key == "timestampList" {
			dv, err := deserializeCBOR_TimestampList(sv)
			if err != nil {
				return nil, err
			}
			ds.TimestampList = dv
		}

		if key == "enumList" {
			dv, err := deserializeCBOR_FooEnumList(sv)
			if err != nil {
				return nil, err
			}
			ds.EnumList = dv
		}

		if key == "intEnumList" {
			dv, err := deserializeCBOR_IntegerEnumList(sv)
			if err != nil {
				return nil, err
			}
			ds.IntEnumList = dv
		}

		if key == "nestedStringList" {
			dv, err := deserializeCBOR_NestedStringList(sv)
			if err != nil {
				return nil, err
			}
			ds.NestedStringList = dv
		}

		if key == "structureList" {
			dv, err := deserializeCBOR_StructureList(sv)
			if err != nil {
				return nil, err
			}
			ds.StructureList = dv
		}
	}
	return ds, nil
}

func deserializeCBOR_RecursiveShapesInputOutputNested2(v smithycbor.Value) (*types.RecursiveShapesInputOutputNested2, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	ds := &types.RecursiveShapesInputOutputNested2{}
	for key, sv := range av {
		_, _ = key, sv
		if key == "bar" {
			dv, err := deserializeCBOR_String(sv)
			if err != nil {
				return nil, err
			}
			ds.Bar = ptr.String(dv)
		}

		if key == "recursiveMember" {
			dv, err := deserializeCBOR_RecursiveShapesInputOutputNested1(sv)
			if err != nil {
				return nil, err
			}
			ds.RecursiveMember = dv
		}
	}
	return ds, nil
}

func deserializeCBOR_DenseSetMap(v smithycbor.Value) (map[string][]string, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	dm := map[string][]string{}
	for key, sv := range av {
		if _, ok := sv.(*smithycbor.Nil); ok {
			dm[key] = nil
			continue
		}
		dv, err := deserializeCBOR_StringSet(sv)
		if err != nil {
			return nil, err
		}
		dm[key] = dv
	}
	return dm, nil
}

func deserializeCBOR_RpcV2CborMapsOutput(v smithycbor.Value) (*RpcV2CborMapsOutput, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	ds := &RpcV2CborMapsOutput{}
	for key, sv := range av {
		_, _ = key, sv
		if key == "denseStructMap" {
			dv, err := deserializeCBOR_DenseStructMap(sv)
			if err != nil {
				return nil, err
			}
			ds.DenseStructMap = dv
		}

		if key == "sparseStructMap" {
			dv, err := deserializeCBOR_SparseStructMap(sv)
			if err != nil {
				return nil, err
			}
			ds.SparseStructMap = dv
		}

		if key == "denseNumberMap" {
			dv, err := deserializeCBOR_DenseNumberMap(sv)
			if err != nil {
				return nil, err
			}
			ds.DenseNumberMap = dv
		}

		if key == "denseBooleanMap" {
			dv, err := deserializeCBOR_DenseBooleanMap(sv)
			if err != nil {
				return nil, err
			}
			ds.DenseBooleanMap = dv
		}

		if key == "denseStringMap" {
			dv, err := deserializeCBOR_DenseStringMap(sv)
			if err != nil {
				return nil, err
			}
			ds.DenseStringMap = dv
		}

		if key == "sparseNumberMap" {
			dv, err := deserializeCBOR_SparseNumberMap(sv)
			if err != nil {
				return nil, err
			}
			ds.SparseNumberMap = dv
		}

		if key == "sparseBooleanMap" {
			dv, err := deserializeCBOR_SparseBooleanMap(sv)
			if err != nil {
				return nil, err
			}
			ds.SparseBooleanMap = dv
		}

		if key == "sparseStringMap" {
			dv, err := deserializeCBOR_SparseStringMap(sv)
			if err != nil {
				return nil, err
			}
			ds.SparseStringMap = dv
		}

		if key == "denseSetMap" {
			dv, err := deserializeCBOR_DenseSetMap(sv)
			if err != nil {
				return nil, err
			}
			ds.DenseSetMap = dv
		}

		if key == "sparseSetMap" {
			dv, err := deserializeCBOR_SparseSetMap(sv)
			if err != nil {
				return nil, err
			}
			ds.SparseSetMap = dv
		}
	}
	return ds, nil
}

func deserializeCBOR_IntegerEnumList(v smithycbor.Value) ([]types.IntegerEnum, error) {
	av, ok := v.(smithycbor.List)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	var dl []types.IntegerEnum
	for _, si := range av {

		di, err := deserializeCBOR_IntegerEnum(si)
		if err != nil {
			return nil, err
		}
		dl = append(dl, di)
	}
	return dl, nil
}

func deserializeCBOR_GreetingWithErrorsOutput(v smithycbor.Value) (*GreetingWithErrorsOutput, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	ds := &GreetingWithErrorsOutput{}
	for key, sv := range av {
		_, _ = key, sv
		if key == "greeting" {
			dv, err := deserializeCBOR_String(sv)
			if err != nil {
				return nil, err
			}
			ds.Greeting = ptr.String(dv)
		}
	}
	return ds, nil
}

func deserializeCBOR_OptionalInputOutputOutput(v smithycbor.Value) (*OptionalInputOutputOutput, error) {
	av, ok := v.(smithycbor.Map)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	ds := &OptionalInputOutputOutput{}
	for key, sv := range av {
		_, _ = key, sv
		if key == "value" {
			dv, err := deserializeCBOR_String(sv)
			if err != nil {
				return nil, err
			}
			ds.Value = ptr.String(dv)
		}
	}
	return ds, nil
}

func deserializeCBOR_Int64(v smithycbor.Value) (int64, error) {
	return smithycbor.AsInt64(v)
}

func deserializeCBOR_FooEnumList(v smithycbor.Value) ([]types.FooEnum, error) {
	av, ok := v.(smithycbor.List)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T", v)
	}
	var dl []types.FooEnum
	for _, si := range av {

		di, err := deserializeCBOR_FooEnum(si)
		if err != nil {
			return nil, err
		}
		dl = append(dl, di)
	}
	return dl, nil
}
func rpc2_deserializeOpErrorEmptyInputOutput(resp *smithyhttp.Response) error {
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("read response body: %w", err)}
	}

	typ, msg, v, err := getProtocolErrorInfo(payload)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("get error info: %w", err)}
	}

	if len(typ) == 0 {
		typ = "UnknownError"
	}
	if len(msg) == 0 {
		msg = "UnknownError"
	}

	_ = v
	switch string(typ) {

	default:

		return &smithy.GenericAPIError{Code: typ, Message: msg}
	}
}

func rpc2_deserializeOpErrorFractionalSeconds(resp *smithyhttp.Response) error {
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("read response body: %w", err)}
	}

	typ, msg, v, err := getProtocolErrorInfo(payload)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("get error info: %w", err)}
	}

	if len(typ) == 0 {
		typ = "UnknownError"
	}
	if len(msg) == 0 {
		msg = "UnknownError"
	}

	_ = v
	switch string(typ) {

	default:

		return &smithy.GenericAPIError{Code: typ, Message: msg}
	}
}

func rpc2_deserializeOpErrorGreetingWithErrors(resp *smithyhttp.Response) error {
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("read response body: %w", err)}
	}

	typ, msg, v, err := getProtocolErrorInfo(payload)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("get error info: %w", err)}
	}

	if len(typ) == 0 {
		typ = "UnknownError"
	}
	if len(msg) == 0 {
		msg = "UnknownError"
	}

	_ = v
	switch string(typ) {
	case "aws.protocoltests.rpcv2Cbor#ComplexError":
		verr, err := deserializeCBOR_ComplexError(v)
		if err != nil {
			return &smithy.DeserializationError{
				Err:      fmt.Errorf("deserialize aws.protocoltests.rpcv2Cbor#ComplexError: %w", err),
				Snapshot: payload,
			}
		}

		return verr
	case "aws.protocoltests.rpcv2Cbor#InvalidGreeting":
		verr, err := deserializeCBOR_InvalidGreeting(v)
		if err != nil {
			return &smithy.DeserializationError{
				Err:      fmt.Errorf("deserialize aws.protocoltests.rpcv2Cbor#InvalidGreeting: %w", err),
				Snapshot: payload,
			}
		}

		return verr
	default:

		return &smithy.GenericAPIError{Code: typ, Message: msg}
	}
}

func rpc2_deserializeOpErrorNoInputOutput(resp *smithyhttp.Response) error {
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("read response body: %w", err)}
	}

	typ, msg, v, err := getProtocolErrorInfo(payload)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("get error info: %w", err)}
	}

	if len(typ) == 0 {
		typ = "UnknownError"
	}
	if len(msg) == 0 {
		msg = "UnknownError"
	}

	_ = v
	switch string(typ) {

	default:

		return &smithy.GenericAPIError{Code: typ, Message: msg}
	}
}

func rpc2_deserializeOpErrorOptionalInputOutput(resp *smithyhttp.Response) error {
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("read response body: %w", err)}
	}

	typ, msg, v, err := getProtocolErrorInfo(payload)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("get error info: %w", err)}
	}

	if len(typ) == 0 {
		typ = "UnknownError"
	}
	if len(msg) == 0 {
		msg = "UnknownError"
	}

	_ = v
	switch string(typ) {

	default:

		return &smithy.GenericAPIError{Code: typ, Message: msg}
	}
}

func rpc2_deserializeOpErrorRecursiveShapes(resp *smithyhttp.Response) error {
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("read response body: %w", err)}
	}

	typ, msg, v, err := getProtocolErrorInfo(payload)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("get error info: %w", err)}
	}

	if len(typ) == 0 {
		typ = "UnknownError"
	}
	if len(msg) == 0 {
		msg = "UnknownError"
	}

	_ = v
	switch string(typ) {

	default:

		return &smithy.GenericAPIError{Code: typ, Message: msg}
	}
}

func rpc2_deserializeOpErrorRpcV2CborLists(resp *smithyhttp.Response) error {
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("read response body: %w", err)}
	}

	typ, msg, v, err := getProtocolErrorInfo(payload)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("get error info: %w", err)}
	}

	if len(typ) == 0 {
		typ = "UnknownError"
	}
	if len(msg) == 0 {
		msg = "UnknownError"
	}

	_ = v
	switch string(typ) {

	default:

		return &smithy.GenericAPIError{Code: typ, Message: msg}
	}
}

func rpc2_deserializeOpErrorRpcV2CborMaps(resp *smithyhttp.Response) error {
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("read response body: %w", err)}
	}

	typ, msg, v, err := getProtocolErrorInfo(payload)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("get error info: %w", err)}
	}

	if len(typ) == 0 {
		typ = "UnknownError"
	}
	if len(msg) == 0 {
		msg = "UnknownError"
	}

	_ = v
	switch string(typ) {

	default:

		return &smithy.GenericAPIError{Code: typ, Message: msg}
	}
}

func rpc2_deserializeOpErrorSimpleScalarProperties(resp *smithyhttp.Response) error {
	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("read response body: %w", err)}
	}

	typ, msg, v, err := getProtocolErrorInfo(payload)
	if err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("get error info: %w", err)}
	}

	if len(typ) == 0 {
		typ = "UnknownError"
	}
	if len(msg) == 0 {
		msg = "UnknownError"
	}

	_ = v
	switch string(typ) {

	default:

		return &smithy.GenericAPIError{Code: typ, Message: msg}
	}
}
func getProtocolErrorInfo(payload []byte) (typ, msg string, v smithycbor.Value, err error) {
	v, err = smithycbor.Decode(payload)
	if err != nil {
		return "", "", nil, fmt.Errorf("decode: %w", err)
	}

	mv, ok := v.(smithycbor.Map)
	if !ok {
		return "", "", nil, fmt.Errorf("unexpected payload type %T", v)
	}

	if ctyp, ok := mv["__type"]; ok {
		if ttyp, ok := ctyp.(smithycbor.String); ok {
			typ = string(ttyp)
		}
	}

	if cmsg, ok := mv["message"]; ok {
		if tmsg, ok := cmsg.(smithycbor.String); ok {
			msg = string(tmsg)
		}
	}

	return typ, msg, mv, nil
}
func getAwsQueryErrorCode(resp *smithyhttp.Response) string {
	header := resp.Header.Get("x-amzn-query-error")
	if header == "" {
		return ""
	}

	parts := strings.Split(header, ";")
	if len(parts) != 2 {
		return ""
	}

	return parts[0]
}
