// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsdata

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Execute SQL Request
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-data-2018-08-01/ExecuteSqlRequest
type ExecuteSqlInput struct {
	_ struct{} `type:"structure"`

	// ARN of the db credentials in AWS Secret Store or the friendly secret name
	//
	// AwsSecretStoreArn is a required field
	AwsSecretStoreArn *string `locationName:"awsSecretStoreArn" type:"string" required:"true"`

	// Target DB name
	Database *string `locationName:"database" type:"string"`

	// ARN of the target db cluster or instance
	//
	// DbClusterOrInstanceArn is a required field
	DbClusterOrInstanceArn *string `locationName:"dbClusterOrInstanceArn" type:"string" required:"true"`

	// Target Schema name
	Schema *string `locationName:"schema" type:"string"`

	// SQL statement(s) to be executed. Statements can be chained by using semicolons
	//
	// SqlStatements is a required field
	SqlStatements *string `locationName:"sqlStatements" type:"string" required:"true"`
}

// String returns the string representation
func (s ExecuteSqlInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ExecuteSqlInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ExecuteSqlInput"}

	if s.AwsSecretStoreArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsSecretStoreArn"))
	}

	if s.DbClusterOrInstanceArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DbClusterOrInstanceArn"))
	}

	if s.SqlStatements == nil {
		invalidParams.Add(aws.NewErrParamRequired("SqlStatements"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ExecuteSqlInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.AwsSecretStoreArn != nil {
		v := *s.AwsSecretStoreArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "awsSecretStoreArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Database != nil {
		v := *s.Database

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "database", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DbClusterOrInstanceArn != nil {
		v := *s.DbClusterOrInstanceArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "dbClusterOrInstanceArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Schema != nil {
		v := *s.Schema

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "schema", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SqlStatements != nil {
		v := *s.SqlStatements

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "sqlStatements", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Execute SQL response
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-data-2018-08-01/ExecuteSqlResponse
type ExecuteSqlOutput struct {
	_ struct{} `type:"structure"`

	// Results returned by executing the sql statement(s)
	//
	// SqlStatementResults is a required field
	SqlStatementResults []SqlStatementResult `locationName:"sqlStatementResults" type:"list" required:"true"`
}

// String returns the string representation
func (s ExecuteSqlOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ExecuteSqlOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SqlStatementResults != nil {
		v := s.SqlStatementResults

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "sqlStatementResults", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opExecuteSql = "ExecuteSql"

// ExecuteSqlRequest returns a request value for making API operation for
// AWS RDS DataService.
//
// Executes any SQL statement on the target database synchronously
//
//    // Example sending a request using ExecuteSqlRequest.
//    req := client.ExecuteSqlRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-data-2018-08-01/ExecuteSql
func (c *Client) ExecuteSqlRequest(input *ExecuteSqlInput) ExecuteSqlRequest {
	op := &aws.Operation{
		Name:       opExecuteSql,
		HTTPMethod: "POST",
		HTTPPath:   "/ExecuteSql",
	}

	if input == nil {
		input = &ExecuteSqlInput{}
	}

	req := c.newRequest(op, input, &ExecuteSqlOutput{})
	return ExecuteSqlRequest{Request: req, Input: input, Copy: c.ExecuteSqlRequest}
}

// ExecuteSqlRequest is the request type for the
// ExecuteSql API operation.
type ExecuteSqlRequest struct {
	*aws.Request
	Input *ExecuteSqlInput
	Copy  func(*ExecuteSqlInput) ExecuteSqlRequest
}

// Send marshals and sends the ExecuteSql API request.
func (r ExecuteSqlRequest) Send(ctx context.Context) (*ExecuteSqlResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ExecuteSqlResponse{
		ExecuteSqlOutput: r.Request.Data.(*ExecuteSqlOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ExecuteSqlResponse is the response type for the
// ExecuteSql API operation.
type ExecuteSqlResponse struct {
	*ExecuteSqlOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ExecuteSql request.
func (r *ExecuteSqlResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
