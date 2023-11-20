// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Modifies the maintenance settings of a cluster.
func (c *Client) ModifyClusterMaintenance(ctx context.Context, params *ModifyClusterMaintenanceInput, optFns ...func(*Options)) (*ModifyClusterMaintenanceOutput, error) {
	if params == nil {
		params = &ModifyClusterMaintenanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyClusterMaintenance", params, optFns, c.addOperationModifyClusterMaintenanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyClusterMaintenanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyClusterMaintenanceInput struct {

	// A unique identifier for the cluster.
	//
	// This member is required.
	ClusterIdentifier *string

	// A boolean indicating whether to enable the deferred maintenance window.
	DeferMaintenance *bool

	// An integer indicating the duration of the maintenance window in days. If you
	// specify a duration, you can't specify an end time. The duration must be 45 days
	// or less.
	DeferMaintenanceDuration *int32

	// A timestamp indicating end time for the deferred maintenance window. If you
	// specify an end time, you can't specify a duration.
	DeferMaintenanceEndTime *time.Time

	// A unique identifier for the deferred maintenance window.
	DeferMaintenanceIdentifier *string

	// A timestamp indicating the start time for the deferred maintenance window.
	DeferMaintenanceStartTime *time.Time

	noSmithyDocumentSerde
}

type ModifyClusterMaintenanceOutput struct {

	// Describes a cluster.
	Cluster *types.Cluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyClusterMaintenanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyClusterMaintenance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyClusterMaintenance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyClusterMaintenance"); err != nil {
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
	if err = addOpModifyClusterMaintenanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyClusterMaintenance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyClusterMaintenance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyClusterMaintenance",
	}
}
