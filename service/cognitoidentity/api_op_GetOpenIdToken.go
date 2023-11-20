// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets an OpenID token, using a known Cognito ID. This known Cognito ID is
// returned by GetId . You can optionally add additional logins for the identity.
// Supplying multiple logins creates an implicit link. The OpenID token is valid
// for 10 minutes. This is a public API. You do not need any credentials to call
// this API.
func (c *Client) GetOpenIdToken(ctx context.Context, params *GetOpenIdTokenInput, optFns ...func(*Options)) (*GetOpenIdTokenOutput, error) {
	if params == nil {
		params = &GetOpenIdTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetOpenIdToken", params, optFns, c.addOperationGetOpenIdTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetOpenIdTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the GetOpenIdToken action.
type GetOpenIdTokenInput struct {

	// A unique identifier in the format REGION:GUID.
	//
	// This member is required.
	IdentityId *string

	// A set of optional name-value pairs that map provider names to provider tokens.
	// When using graph.facebook.com and www.amazon.com, supply the access_token
	// returned from the provider's authflow. For accounts.google.com, an Amazon
	// Cognito user pool provider, or any other OpenID Connect provider, always include
	// the id_token .
	Logins map[string]string

	noSmithyDocumentSerde
}

// Returned in response to a successful GetOpenIdToken request.
type GetOpenIdTokenOutput struct {

	// A unique identifier in the format REGION:GUID. Note that the IdentityId
	// returned may not match the one passed on input.
	IdentityId *string

	// An OpenID token, valid for 10 minutes.
	Token *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetOpenIdTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetOpenIdToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetOpenIdToken{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetOpenIdToken"); err != nil {
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
	if err = addOpGetOpenIdTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetOpenIdToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetOpenIdToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetOpenIdToken",
	}
}
