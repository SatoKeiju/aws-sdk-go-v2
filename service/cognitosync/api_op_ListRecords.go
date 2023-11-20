// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets paginated records, optionally changed after a particular sync count for a
// dataset and identity. With Amazon Cognito Sync, each identity has access only to
// its own data. Thus, the credentials used to make this API call need to have
// access to the identity data. ListRecords can be called with temporary user
// credentials provided by Cognito Identity or with developer credentials. You
// should use Cognito Identity credentials to make this API call. ListRecords The
// following examples have been edited for readability. POST / HTTP/1.1
// CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// b3d2e31e-d6b7-4612-8e84-c9ba288dab5d X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.ListRecords HOST:
// cognito-sync.us-east-1.amazonaws.com:443 X-AMZ-DATE: 20141111T183230Z
// AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;host;x-amz-date;x-amz-target;x-amzn-requestid,
// Signature= { "Operation": "com.amazonaws.cognito.sync.model#ListRecords",
// "Service": "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "IDENTITY_POOL_ID", "IdentityId": "IDENTITY_ID",
// "DatasetName": "newDataSet" } } 1.1 200 OK x-amzn-requestid:
// b3d2e31e-d6b7-4612-8e84-c9ba288dab5d content-type: application/json
// content-length: 623 date: Tue, 11 Nov 2014 18:32:30 GMT { "Output": { "__type":
// "com.amazonaws.cognito.sync.model#ListRecordsResponse", "Count": 0,
// "DatasetDeletedAfterRequestedSyncCount": false, "DatasetExists": false,
// "DatasetSyncCount": 0, "LastModifiedBy": null, "MergedDatasetNames": null,
// "NextToken": null, "Records": [], "SyncSessionToken": "SYNC_SESSION_TOKEN" },
// "Version": "1.0" }
func (c *Client) ListRecords(ctx context.Context, params *ListRecordsInput, optFns ...func(*Options)) (*ListRecordsOutput, error) {
	if params == nil {
		params = &ListRecordsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecords", params, optFns, c.addOperationListRecordsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecordsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request for a list of records.
type ListRecordsInput struct {

	// A string of up to 128 characters. Allowed characters are a-z, A-Z, 0-9, '_'
	// (underscore), '-' (dash), and '.' (dot).
	//
	// This member is required.
	DatasetName *string

	// A name-spaced GUID (for example,
	// us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito. GUID
	// generation is unique within a region.
	//
	// This member is required.
	IdentityId *string

	// A name-spaced GUID (for example,
	// us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito. GUID
	// generation is unique within a region.
	//
	// This member is required.
	IdentityPoolId *string

	// The last server sync count for this record.
	LastSyncCount *int64

	// The maximum number of results to be returned.
	MaxResults *int32

	// A pagination token for obtaining the next page of results.
	NextToken *string

	// A token containing a session ID, identity ID, and expiration.
	SyncSessionToken *string

	noSmithyDocumentSerde
}

// Returned for a successful ListRecordsRequest.
type ListRecordsOutput struct {

	// Total number of records.
	Count int32

	// A boolean value specifying whether to delete the dataset locally.
	DatasetDeletedAfterRequestedSyncCount bool

	// Indicates whether the dataset exists.
	DatasetExists bool

	// Server sync count for this dataset.
	DatasetSyncCount *int64

	// The user/device that made the last change to this record.
	LastModifiedBy *string

	// Names of merged datasets.
	MergedDatasetNames []string

	// A pagination token for obtaining the next page of results.
	NextToken *string

	// A list of all records.
	Records []types.Record

	// A token containing a session ID, identity ID, and expiration.
	SyncSessionToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRecordsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRecords{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRecords{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListRecords"); err != nil {
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
	if err = addOpListRecordsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecords(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListRecords(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListRecords",
	}
}
