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
)

// Deletes the specified manual snapshot. The snapshot must be in the available
// state, with no other users authorized to access the snapshot. Unlike automated
// snapshots, manual snapshots are retained even after you delete your cluster.
// Amazon Redshift does not delete your manual snapshots. You must delete manual
// snapshot explicitly to avoid getting charged. If other accounts are authorized
// to access the snapshot, you must revoke all of the authorizations before you can
// delete the snapshot.
func (c *Client) DeleteClusterSnapshot(ctx context.Context, params *DeleteClusterSnapshotInput, optFns ...func(*Options)) (*DeleteClusterSnapshotOutput, error) {
	if params == nil {
		params = &DeleteClusterSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteClusterSnapshot", params, optFns, c.addOperationDeleteClusterSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteClusterSnapshotInput struct {

	// The unique identifier of the manual snapshot to be deleted. Constraints: Must
	// be the name of an existing snapshot that is in the available , failed , or
	// cancelled state.
	//
	// This member is required.
	SnapshotIdentifier *string

	// The unique identifier of the cluster the snapshot was created from. This
	// parameter is required if your IAM user has a policy containing a snapshot
	// resource element that specifies anything other than * for the cluster name.
	// Constraints: Must be the name of valid cluster.
	SnapshotClusterIdentifier *string

	noSmithyDocumentSerde
}

type DeleteClusterSnapshotOutput struct {

	// Describes a snapshot.
	Snapshot *types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteClusterSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteClusterSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteClusterSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteClusterSnapshot"); err != nil {
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
	if err = addOpDeleteClusterSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteClusterSnapshot(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteClusterSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteClusterSnapshot",
	}
}
