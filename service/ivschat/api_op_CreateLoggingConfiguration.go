// Code generated by smithy-go-codegen DO NOT EDIT.

package ivschat

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ivschat/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a logging configuration that allows clients to store and record sent
// messages.
func (c *Client) CreateLoggingConfiguration(ctx context.Context, params *CreateLoggingConfigurationInput, optFns ...func(*Options)) (*CreateLoggingConfigurationOutput, error) {
	if params == nil {
		params = &CreateLoggingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLoggingConfiguration", params, optFns, c.addOperationCreateLoggingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLoggingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLoggingConfigurationInput struct {

	// A complex type that contains a destination configuration for where chat content
	// will be logged. There can be only one type of destination ( cloudWatchLogs ,
	// firehose , or s3 ) in a destinationConfiguration .
	//
	// This member is required.
	DestinationConfiguration types.DestinationConfiguration

	// Logging-configuration name. The value does not need to be unique.
	Name *string

	// Tags to attach to the resource. Array of maps, each of the form string:string
	// (key:value) . See Tagging AWS Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// for details, including restrictions that apply to tags and "Tag naming limits
	// and requirements"; Amazon IVS Chat has no constraints on tags beyond what is
	// documented there.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateLoggingConfigurationOutput struct {

	// Logging-configuration ARN, assigned by the system.
	Arn *string

	// Time when the logging configuration was created. This is an ISO 8601 timestamp;
	// note that this is returned as a string.
	CreateTime *time.Time

	// A complex type that contains a destination configuration for where chat content
	// will be logged, from the request. There is only one type of destination (
	// cloudWatchLogs , firehose , or s3 ) in a destinationConfiguration .
	DestinationConfiguration types.DestinationConfiguration

	// Logging-configuration ID, generated by the system. This is a relative
	// identifier, the part of the ARN that uniquely identifies the logging
	// configuration.
	Id *string

	// Logging-configuration name, from the request (if specified).
	Name *string

	// The state of the logging configuration. When the state is ACTIVE , the
	// configuration is ready to log chat content.
	State types.CreateLoggingConfigurationState

	// Tags attached to the resource, from the request (if specified). Array of maps,
	// each of the form string:string (key:value) .
	Tags map[string]string

	// Time of the logging configuration’s last update. This is an ISO 8601 timestamp;
	// note that this is returned as a string.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLoggingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateLoggingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLoggingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLoggingConfiguration"); err != nil {
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
	if err = addOpCreateLoggingConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLoggingConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLoggingConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLoggingConfiguration",
	}
}
