// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// List FeatureGroup s based on given filter and order.
func (c *Client) ListFeatureGroups(ctx context.Context, params *ListFeatureGroupsInput, optFns ...func(*Options)) (*ListFeatureGroupsOutput, error) {
	if params == nil {
		params = &ListFeatureGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFeatureGroups", params, optFns, c.addOperationListFeatureGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFeatureGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFeatureGroupsInput struct {

	// Use this parameter to search for FeatureGroups s created after a specific date
	// and time.
	CreationTimeAfter *time.Time

	// Use this parameter to search for FeatureGroups s created before a specific date
	// and time.
	CreationTimeBefore *time.Time

	// A FeatureGroup status. Filters by FeatureGroup status.
	FeatureGroupStatusEquals types.FeatureGroupStatus

	// The maximum number of results returned by ListFeatureGroups .
	MaxResults *int32

	// A string that partially matches one or more FeatureGroup s names. Filters
	// FeatureGroup s by name.
	NameContains *string

	// A token to resume pagination of ListFeatureGroups results.
	NextToken *string

	// An OfflineStore status. Filters by OfflineStore status.
	OfflineStoreStatusEquals types.OfflineStoreStatusValue

	// The value on which the feature group list is sorted.
	SortBy types.FeatureGroupSortBy

	// The order in which feature groups are listed.
	SortOrder types.FeatureGroupSortOrder

	noSmithyDocumentSerde
}

type ListFeatureGroupsOutput struct {

	// A summary of feature groups.
	//
	// This member is required.
	FeatureGroupSummaries []types.FeatureGroupSummary

	// A token to resume pagination of ListFeatureGroups results.
	//
	// This member is required.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListFeatureGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListFeatureGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFeatureGroups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListFeatureGroups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFeatureGroups(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListFeatureGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListFeatureGroups",
	}
}
