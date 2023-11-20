// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a previously provisioned cluster. DeleteCacheCluster deletes all
// associated cache nodes, node endpoints and the cluster itself. When you receive
// a successful response from this operation, Amazon ElastiCache immediately begins
// deleting the cluster; you cannot cancel or revert this operation. This operation
// is not valid for:
//   - Redis (cluster mode enabled) clusters
//   - Redis (cluster mode disabled) clusters
//   - A cluster that is the last read replica of a replication group
//   - A cluster that is the primary node of a replication group
//   - A node group (shard) that has Multi-AZ mode enabled
//   - A cluster from a Redis (cluster mode enabled) replication group
//   - A cluster that is not in the available state
func (c *Client) DeleteCacheCluster(ctx context.Context, params *DeleteCacheClusterInput, optFns ...func(*Options)) (*DeleteCacheClusterOutput, error) {
	if params == nil {
		params = &DeleteCacheClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCacheCluster", params, optFns, c.addOperationDeleteCacheClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCacheClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DeleteCacheCluster operation.
type DeleteCacheClusterInput struct {

	// The cluster identifier for the cluster to be deleted. This parameter is not
	// case sensitive.
	//
	// This member is required.
	CacheClusterId *string

	// The user-supplied name of a final cluster snapshot. This is the unique name
	// that identifies the snapshot. ElastiCache creates the snapshot, and then deletes
	// the cluster immediately afterward.
	FinalSnapshotIdentifier *string

	noSmithyDocumentSerde
}

type DeleteCacheClusterOutput struct {

	// Contains all of the attributes of a specific cluster.
	CacheCluster *types.CacheCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteCacheClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteCacheCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteCacheCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteCacheCluster"); err != nil {
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
	if err = addOpDeleteCacheClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCacheCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteCacheCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteCacheCluster",
	}
}
