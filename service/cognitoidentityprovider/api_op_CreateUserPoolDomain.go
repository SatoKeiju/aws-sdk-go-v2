// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new domain for a user pool. Amazon Cognito evaluates Identity and
// Access Management (IAM) policies in requests for this API operation. For this
// operation, you must use IAM credentials to authorize requests, and you must
// grant yourself the corresponding IAM permission in a policy. Learn more
//   - Signing Amazon Web Services API Requests (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_aws-signing.html)
//   - Using the Amazon Cognito user pools API and user pool endpoints (https://docs.aws.amazon.com/cognito/latest/developerguide/user-pools-API-operations.html)
func (c *Client) CreateUserPoolDomain(ctx context.Context, params *CreateUserPoolDomainInput, optFns ...func(*Options)) (*CreateUserPoolDomainOutput, error) {
	if params == nil {
		params = &CreateUserPoolDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUserPoolDomain", params, optFns, c.addOperationCreateUserPoolDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUserPoolDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUserPoolDomainInput struct {

	// The domain string. For custom domains, this is the fully-qualified domain name,
	// such as auth.example.com . For Amazon Cognito prefix domains, this is the prefix
	// alone, such as auth .
	//
	// This member is required.
	Domain *string

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The configuration for a custom domain that hosts the sign-up and sign-in
	// webpages for your application. Provide this parameter only if you want to use a
	// custom domain for your user pool. Otherwise, you can exclude this parameter and
	// use the Amazon Cognito hosted domain instead. For more information about the
	// hosted domain and custom domains, see Configuring a User Pool Domain (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-pools-assign-domain.html)
	// .
	CustomDomainConfig *types.CustomDomainConfigType

	noSmithyDocumentSerde
}

type CreateUserPoolDomainOutput struct {

	// The Amazon CloudFront endpoint that you use as the target of the alias that you
	// set up with your Domain Name Service (DNS) provider.
	CloudFrontDomain *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUserPoolDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUserPoolDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUserPoolDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateUserPoolDomain"); err != nil {
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
	if err = addOpCreateUserPoolDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUserPoolDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateUserPoolDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateUserPoolDomain",
	}
}
