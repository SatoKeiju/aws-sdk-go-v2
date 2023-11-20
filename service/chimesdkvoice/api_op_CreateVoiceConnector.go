// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Chime SDK Voice Connector. For more information about Voice
// Connectors, see Managing Amazon Chime SDK Voice Connector groups (https://docs.aws.amazon.com/chime-sdk/latest/ag/voice-connector-groups.html)
// in the Amazon Chime SDK Administrator Guide.
func (c *Client) CreateVoiceConnector(ctx context.Context, params *CreateVoiceConnectorInput, optFns ...func(*Options)) (*CreateVoiceConnectorOutput, error) {
	if params == nil {
		params = &CreateVoiceConnectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVoiceConnector", params, optFns, c.addOperationCreateVoiceConnectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVoiceConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVoiceConnectorInput struct {

	// The name of the Voice Connector.
	//
	// This member is required.
	Name *string

	// Enables or disables encryption for the Voice Connector.
	//
	// This member is required.
	RequireEncryption *bool

	// The AWS Region in which the Amazon Chime SDK Voice Connector is created.
	// Default value: us-east-1 .
	AwsRegion types.VoiceConnectorAwsRegion

	// The tags assigned to the Voice Connector.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateVoiceConnectorOutput struct {

	// The details of the Voice Connector.
	VoiceConnector *types.VoiceConnector

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVoiceConnectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVoiceConnector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVoiceConnector{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateVoiceConnector"); err != nil {
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
	if err = addOpCreateVoiceConnectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVoiceConnector(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVoiceConnector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateVoiceConnector",
	}
}
