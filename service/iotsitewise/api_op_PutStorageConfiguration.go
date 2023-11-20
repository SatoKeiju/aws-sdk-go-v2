// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Configures storage settings for IoT SiteWise.
func (c *Client) PutStorageConfiguration(ctx context.Context, params *PutStorageConfigurationInput, optFns ...func(*Options)) (*PutStorageConfigurationOutput, error) {
	if params == nil {
		params = &PutStorageConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutStorageConfiguration", params, optFns, c.addOperationPutStorageConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutStorageConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutStorageConfigurationInput struct {

	// The storage tier that you specified for your data. The storageType parameter
	// can be one of the following values:
	//   - SITEWISE_DEFAULT_STORAGE – IoT SiteWise saves your data into the hot tier.
	//   The hot tier is a service-managed database.
	//   - MULTI_LAYER_STORAGE – IoT SiteWise saves your data in both the cold tier and
	//   the hot tier. The cold tier is a customer-managed Amazon S3 bucket.
	//
	// This member is required.
	StorageType types.StorageType

	// Contains the storage configuration for time series (data streams) that aren't
	// associated with asset properties. The disassociatedDataStorage can be one of
	// the following values:
	//   - ENABLED – IoT SiteWise accepts time series that aren't associated with asset
	//   properties. After the disassociatedDataStorage is enabled, you can't disable
	//   it.
	//   - DISABLED – IoT SiteWise doesn't accept time series (data streams) that
	//   aren't associated with asset properties.
	// For more information, see Data streams (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/data-streams.html)
	// in the IoT SiteWise User Guide.
	DisassociatedDataStorage types.DisassociatedDataStorageState

	// Identifies a storage destination. If you specified MULTI_LAYER_STORAGE for the
	// storage type, you must specify a MultiLayerStorage object.
	MultiLayerStorage *types.MultiLayerStorage

	// How many days your data is kept in the hot tier. By default, your data is kept
	// indefinitely in the hot tier.
	RetentionPeriod *types.RetentionPeriod

	noSmithyDocumentSerde
}

type PutStorageConfigurationOutput struct {

	// Contains current status information for the configuration.
	//
	// This member is required.
	ConfigurationStatus *types.ConfigurationStatus

	// The storage tier that you specified for your data. The storageType parameter
	// can be one of the following values:
	//   - SITEWISE_DEFAULT_STORAGE – IoT SiteWise saves your data into the hot tier.
	//   The hot tier is a service-managed database.
	//   - MULTI_LAYER_STORAGE – IoT SiteWise saves your data in both the cold tier and
	//   the hot tier. The cold tier is a customer-managed Amazon S3 bucket.
	//
	// This member is required.
	StorageType types.StorageType

	// Contains the storage configuration for time series (data streams) that aren't
	// associated with asset properties. The disassociatedDataStorage can be one of
	// the following values:
	//   - ENABLED – IoT SiteWise accepts time series that aren't associated with asset
	//   properties. After the disassociatedDataStorage is enabled, you can't disable
	//   it.
	//   - DISABLED – IoT SiteWise doesn't accept time series (data streams) that
	//   aren't associated with asset properties.
	// For more information, see Data streams (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/data-streams.html)
	// in the IoT SiteWise User Guide.
	DisassociatedDataStorage types.DisassociatedDataStorageState

	// Contains information about the storage destination.
	MultiLayerStorage *types.MultiLayerStorage

	// How many days your data is kept in the hot tier. By default, your data is kept
	// indefinitely in the hot tier.
	RetentionPeriod *types.RetentionPeriod

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutStorageConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutStorageConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutStorageConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutStorageConfiguration"); err != nil {
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
	if err = addEndpointPrefix_opPutStorageConfigurationMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutStorageConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutStorageConfiguration(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opPutStorageConfigurationMiddleware struct {
}

func (*endpointPrefix_opPutStorageConfigurationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opPutStorageConfigurationMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opPutStorageConfigurationMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opPutStorageConfigurationMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opPutStorageConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutStorageConfiguration",
	}
}
