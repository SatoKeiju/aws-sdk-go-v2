// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a user account for the specified Amazon Connect instance. Certain
// UserIdentityInfo (https://docs.aws.amazon.com/connect/latest/APIReference/API_UserIdentityInfo.html)
// parameters are required in some situations. For example, Email is required if
// you are using SAML for identity management. FirstName and LastName are required
// if you are using Amazon Connect or SAML for identity management. For information
// about how to create user accounts using the Amazon Connect console, see Add
// Users (https://docs.aws.amazon.com/connect/latest/adminguide/user-management.html)
// in the Amazon Connect Administrator Guide.
func (c *Client) CreateUser(ctx context.Context, params *CreateUserInput, optFns ...func(*Options)) (*CreateUserOutput, error) {
	if params == nil {
		params = &CreateUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUser", params, optFns, c.addOperationCreateUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUserInput struct {

	// The identifier of the Amazon Connect instance. You can find the instance ID (https://docs.aws.amazon.com/connect/latest/adminguide/find-instance-arn.html)
	// in the Amazon Resource Name (ARN) of the instance.
	//
	// This member is required.
	InstanceId *string

	// The phone settings for the user.
	//
	// This member is required.
	PhoneConfig *types.UserPhoneConfig

	// The identifier of the routing profile for the user.
	//
	// This member is required.
	RoutingProfileId *string

	// The identifier of the security profile for the user.
	//
	// This member is required.
	SecurityProfileIds []string

	// The user name for the account. For instances not using SAML for identity
	// management, the user name can include up to 20 characters. If you are using SAML
	// for identity management, the user name can include up to 64 characters from
	// [a-zA-Z0-9_-.\@]+.
	//
	// This member is required.
	Username *string

	// The identifier of the user account in the directory used for identity
	// management. If Amazon Connect cannot access the directory, you can specify this
	// identifier to authenticate users. If you include the identifier, we assume that
	// Amazon Connect cannot access the directory. Otherwise, the identity information
	// is used to authenticate users from your directory. This parameter is required if
	// you are using an existing directory for identity management in Amazon Connect
	// when Amazon Connect cannot access your directory to authenticate users. If you
	// are using SAML for identity management and include this parameter, an error is
	// returned.
	DirectoryUserId *string

	// The identifier of the hierarchy group for the user.
	HierarchyGroupId *string

	// The information about the identity of the user.
	IdentityInfo *types.UserIdentityInfo

	// The password for the user account. A password is required if you are using
	// Amazon Connect for identity management. Otherwise, it is an error to include a
	// password.
	Password *string

	// The tags used to organize, track, or control access for this resource. For
	// example, { "tags": {"key1":"value1", "key2":"value2"} }.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateUserOutput struct {

	// The Amazon Resource Name (ARN) of the user account.
	UserArn *string

	// The identifier of the user account.
	UserId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateUser{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateUser"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addV4DetectSkewMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUser(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateUser",
	}
}
