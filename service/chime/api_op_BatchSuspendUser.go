// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/BatchSuspendUserRequest
type BatchSuspendUserInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Chime account ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The request containing the user IDs to suspend.
	//
	// UserIdList is a required field
	UserIdList []string `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchSuspendUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchSuspendUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchSuspendUserInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.UserIdList == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserIdList"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchSuspendUserInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.UserIdList != nil {
		v := s.UserIdList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UserIdList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/BatchSuspendUserResponse
type BatchSuspendUserOutput struct {
	_ struct{} `type:"structure"`

	// If the BatchSuspendUser action fails for one or more of the user IDs in the
	// request, a list of the user IDs is returned, along with error codes and error
	// messages.
	UserErrors []UserError `type:"list"`
}

// String returns the string representation
func (s BatchSuspendUserOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s BatchSuspendUserOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.UserErrors != nil {
		v := s.UserErrors

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "UserErrors", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opBatchSuspendUser = "BatchSuspendUser"

// BatchSuspendUserRequest returns a request value for making API operation for
// Amazon Chime.
//
// Suspends up to 50 users from a Team or EnterpriseLWA Amazon Chime account.
// For more information about different account types, see Managing Your Amazon
// Chime Accounts (https://docs.aws.amazon.com/chime/latest/ag/manage-chime-account.html)
// in the Amazon Chime Administration Guide.
//
// Users suspended from a Team account are dissasociated from the account, but
// they can continue to use Amazon Chime as free users. To remove the suspension
// from suspended Team account users, invite them to the Team account again.
// You can use the InviteUsers action to do so.
//
// Users suspended from an EnterpriseLWA account are immediately signed out
// of Amazon Chime and can no longer sign in. To remove the suspension from
// suspended EnterpriseLWA account users, use the BatchUnsuspendUser action.
//
// To sign out users without suspending them, use the LogoutUser action.
//
//    // Example sending a request using BatchSuspendUserRequest.
//    req := client.BatchSuspendUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/BatchSuspendUser
func (c *Client) BatchSuspendUserRequest(input *BatchSuspendUserInput) BatchSuspendUserRequest {
	op := &aws.Operation{
		Name:       opBatchSuspendUser,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{accountId}/users?operation=suspend",
	}

	if input == nil {
		input = &BatchSuspendUserInput{}
	}

	req := c.newRequest(op, input, &BatchSuspendUserOutput{})
	return BatchSuspendUserRequest{Request: req, Input: input, Copy: c.BatchSuspendUserRequest}
}

// BatchSuspendUserRequest is the request type for the
// BatchSuspendUser API operation.
type BatchSuspendUserRequest struct {
	*aws.Request
	Input *BatchSuspendUserInput
	Copy  func(*BatchSuspendUserInput) BatchSuspendUserRequest
}

// Send marshals and sends the BatchSuspendUser API request.
func (r BatchSuspendUserRequest) Send(ctx context.Context) (*BatchSuspendUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchSuspendUserResponse{
		BatchSuspendUserOutput: r.Request.Data.(*BatchSuspendUserOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchSuspendUserResponse is the response type for the
// BatchSuspendUser API operation.
type BatchSuspendUserResponse struct {
	*BatchSuspendUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchSuspendUser request.
func (r *BatchSuspendUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
