// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers (or retrieves) a Cognito IdentityId and an OpenID Connect token for a
// user authenticated by your backend authentication process. Supplying multiple
// logins will create an implicit linked account. You can only specify one
// developer provider as part of the Logins map, which is linked to the identity
// pool. The developer provider is the "domain" by which Cognito will refer to your
// users. You can use GetOpenIdTokenForDeveloperIdentity to create a new identity
// and to link new logins (that is, user credentials issued by a public provider or
// developer provider) to an existing identity. When you want to create a new
// identity, the IdentityId should be null. When you want to associate a new login
// with an existing authenticated/unauthenticated identity, you can do so by
// providing the existing IdentityId . This API will create the identity in the
// specified IdentityPoolId . You must use AWS Developer credentials to call this
// API.
func (c *Client) GetOpenIdTokenForDeveloperIdentity(ctx context.Context, params *GetOpenIdTokenForDeveloperIdentityInput, optFns ...func(*Options)) (*GetOpenIdTokenForDeveloperIdentityOutput, error) {
	if params == nil {
		params = &GetOpenIdTokenForDeveloperIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOpenIdTokenForDeveloperIdentity", params, optFns, c.addOperationGetOpenIdTokenForDeveloperIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOpenIdTokenForDeveloperIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the GetOpenIdTokenForDeveloperIdentity action.
type GetOpenIdTokenForDeveloperIdentityInput struct {

	// An identity pool ID in the format REGION:GUID.
	//
	// This member is required.
	IdentityPoolId *string

	// A set of optional name-value pairs that map provider names to provider tokens.
	// Each name-value pair represents a user from a public provider or developer
	// provider. If the user is from a developer provider, the name-value pair will
	// follow the syntax "developer_provider_name": "developer_user_identifier" . The
	// developer provider is the "domain" by which Cognito will refer to your users;
	// you provided this domain while creating/updating the identity pool. The
	// developer user identifier is an identifier from your backend that uniquely
	// identifies a user. When you create an identity pool, you can specify the
	// supported logins.
	//
	// This member is required.
	Logins map[string]string

	// A unique identifier in the format REGION:GUID.
	IdentityId *string

	// Use this operation to configure attribute mappings for custom providers.
	PrincipalTags map[string]string

	// The expiration time of the token, in seconds. You can specify a custom
	// expiration time for the token so that you can cache it. If you don't provide an
	// expiration time, the token is valid for 15 minutes. You can exchange the token
	// with Amazon STS for temporary AWS credentials, which are valid for a maximum of
	// one hour. The maximum token duration you can set is 24 hours. You should take
	// care in setting the expiration time for a token, as there are significant
	// security implications: an attacker could use a leaked token to access your AWS
	// resources for the token's duration. Please provide for a small grace period,
	// usually no more than 5 minutes, to account for clock skew.
	TokenDuration *int64

	noSmithyDocumentSerde
}

// Returned in response to a successful GetOpenIdTokenForDeveloperIdentity request.
type GetOpenIdTokenForDeveloperIdentityOutput struct {

	// A unique identifier in the format REGION:GUID.
	IdentityId *string

	// An OpenID token.
	Token *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOpenIdTokenForDeveloperIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetOpenIdTokenForDeveloperIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOpenIdTokenForDeveloperIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetOpenIdTokenForDeveloperIdentity"); err != nil {
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
	if err = addOpGetOpenIdTokenForDeveloperIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOpenIdTokenForDeveloperIdentity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetOpenIdTokenForDeveloperIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetOpenIdTokenForDeveloperIdentity",
	}
}
