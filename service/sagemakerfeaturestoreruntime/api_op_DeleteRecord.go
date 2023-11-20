// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakerfeaturestoreruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemakerfeaturestoreruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a Record from a FeatureGroup in the OnlineStore . Feature Store supports
// both SoftDelete and HardDelete . For SoftDelete (default), feature columns are
// set to null and the record is no longer retrievable by GetRecord or
// BatchGetRecord . For HardDelete , the complete Record is removed from the
// OnlineStore . In both cases, Feature Store appends the deleted record marker to
// the OfflineStore with feature values set to null , is_deleted value set to True
// , and EventTime set to the delete input EventTime . Note that the EventTime
// specified in DeleteRecord should be set later than the EventTime of the
// existing record in the OnlineStore for that RecordIdentifer . If it is not, the
// deletion does not occur:
//   - For SoftDelete , the existing (undeleted) record remains in the OnlineStore
//     , though the delete record marker is still written to the OfflineStore .
//   - HardDelete returns EventTime : 400 ValidationException to indicate that the
//     delete operation failed. No delete record marker is written to the
//     OfflineStore .
func (c *Client) DeleteRecord(ctx context.Context, params *DeleteRecordInput, optFns ...func(*Options)) (*DeleteRecordOutput, error) {
	if params == nil {
		params = &DeleteRecordInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteRecord", params, optFns, c.addOperationDeleteRecordMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteRecordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteRecordInput struct {

	// Timestamp indicating when the deletion event occurred. EventTime can be used to
	// query data at a certain point in time.
	//
	// This member is required.
	EventTime *string

	// The name or Amazon Resource Name (ARN) of the feature group to delete the
	// record from.
	//
	// This member is required.
	FeatureGroupName *string

	// The value for the RecordIdentifier that uniquely identifies the record, in
	// string format.
	//
	// This member is required.
	RecordIdentifierValueAsString *string

	// The name of the deletion mode for deleting the record. By default, the deletion
	// mode is set to SoftDelete .
	DeletionMode types.DeletionMode

	// A list of stores from which you're deleting the record. By default, Feature
	// Store deletes the record from all of the stores that you're using for the
	// FeatureGroup .
	TargetStores []types.TargetStore

	noSmithyDocumentSerde
}

type DeleteRecordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteRecordMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteRecord{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteRecord{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteRecord"); err != nil {
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
	if err = addOpDeleteRecordValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteRecord(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteRecord(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteRecord",
	}
}
