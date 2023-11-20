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

// Configures actions on detected risks. To delete the risk configuration for
// UserPoolId or ClientId , pass null values for all four configuration types. To
// activate Amazon Cognito advanced security features, update the user pool to
// include the UserPoolAddOns key AdvancedSecurityMode .
func (c *Client) SetRiskConfiguration(ctx context.Context, params *SetRiskConfigurationInput, optFns ...func(*Options)) (*SetRiskConfigurationOutput, error) {
	if params == nil {
		params = &SetRiskConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetRiskConfiguration", params, optFns, c.addOperationSetRiskConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetRiskConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetRiskConfigurationInput struct {

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string

	// The account takeover risk configuration.
	AccountTakeoverRiskConfiguration *types.AccountTakeoverRiskConfigurationType

	// The app client ID. If ClientId is null, then the risk configuration is mapped
	// to userPoolId . When the client ID is null, the same risk configuration is
	// applied to all the clients in the userPool. Otherwise, ClientId is mapped to
	// the client. When the client ID isn't null, the user pool configuration is
	// overridden and the risk configuration for the client is used instead.
	ClientId *string

	// The compromised credentials risk configuration.
	CompromisedCredentialsRiskConfiguration *types.CompromisedCredentialsRiskConfigurationType

	// The configuration to override the risk decision.
	RiskExceptionConfiguration *types.RiskExceptionConfigurationType

	noSmithyDocumentSerde
}

type SetRiskConfigurationOutput struct {

	// The risk configuration.
	//
	// This member is required.
	RiskConfiguration *types.RiskConfigurationType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetRiskConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetRiskConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetRiskConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetRiskConfiguration"); err != nil {
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
	if err = addOpSetRiskConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetRiskConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetRiskConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetRiskConfiguration",
	}
}
