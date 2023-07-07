// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists the parts that have been uploaded for a specific multipart upload. This
// operation must include the upload ID, which you obtain by sending the initiate
// multipart upload request (see CreateMultipartUpload (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html)
// ). This request returns a maximum of 1,000 uploaded parts. The default number of
// parts returned is 1,000 parts. You can restrict the number of parts returned by
// specifying the max-parts request parameter. If your multipart upload consists
// of more than 1,000 parts, the response returns an IsTruncated field with the
// value of true, and a NextPartNumberMarker element. In subsequent ListParts
// requests you can include the part-number-marker query string parameter and set
// its value to the NextPartNumberMarker field value from the previous response.
// If the upload was created using a checksum algorithm, you will need to have
// permission to the kms:Decrypt action for the request to succeed. For more
// information on multipart uploads, see Uploading Objects Using Multipart Upload (https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html)
// . For information on permissions required to use the multipart upload API, see
// Multipart Upload and Permissions (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html)
// . The following operations are related to ListParts :
//   - CreateMultipartUpload (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html)
//   - UploadPart (https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html)
//   - CompleteMultipartUpload (https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html)
//   - AbortMultipartUpload (https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html)
//   - GetObjectAttributes (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html)
//   - ListMultipartUploads (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html)
func (c *Client) ListParts(ctx context.Context, params *ListPartsInput, optFns ...func(*Options)) (*ListPartsOutput, error) {
	if params == nil {
		params = &ListPartsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListParts", params, optFns, c.addOperationListPartsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPartsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPartsInput struct {

	// The name of the bucket to which the parts are being uploaded. When using this
	// action with an access point, you must direct requests to the access point
	// hostname. The access point hostname takes the form
	// AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com. When using this
	// action with an access point through the Amazon Web Services SDKs, you provide
	// the access point ARN in place of the bucket name. For more information about
	// access point ARNs, see Using access points (https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-access-points.html)
	// in the Amazon S3 User Guide. When you use this action with Amazon S3 on
	// Outposts, you must direct requests to the S3 on Outposts hostname. The S3 on
	// Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com . When you
	// use this action with S3 on Outposts through the Amazon Web Services SDKs, you
	// provide the Outposts access point ARN in place of the bucket name. For more
	// information about S3 on Outposts ARNs, see What is S3 on Outposts? (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html)
	// in the Amazon S3 User Guide.
	//
	// This member is required.
	Bucket *string

	// Object key for which the multipart upload was initiated.
	//
	// This member is required.
	Key *string

	// Upload ID identifying the multipart upload whose parts are being listed.
	//
	// This member is required.
	UploadId *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	// Sets the maximum number of parts to return.
	MaxParts int32

	// Specifies the part after which listing should begin. Only parts with higher
	// part numbers will be listed.
	PartNumberMarker *string

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from Requester Pays buckets, see Downloading Objects
	// in Requester Pays Buckets (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 User Guide.
	RequestPayer types.RequestPayer

	// The server-side encryption (SSE) algorithm used to encrypt the object. This
	// parameter is needed only when the object was created using a checksum algorithm.
	// For more information, see Protecting data using SSE-C keys (https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html)
	// in the Amazon S3 User Guide.
	SSECustomerAlgorithm *string

	// The server-side encryption (SSE) customer managed key. This parameter is needed
	// only when the object was created using a checksum algorithm. For more
	// information, see Protecting data using SSE-C keys (https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html)
	// in the Amazon S3 User Guide.
	SSECustomerKey *string

	// The MD5 server-side encryption (SSE) customer managed key. This parameter is
	// needed only when the object was created using a checksum algorithm. For more
	// information, see Protecting data using SSE-C keys (https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html)
	// in the Amazon S3 User Guide.
	SSECustomerKeyMD5 *string

	noSmithyDocumentSerde
}

type ListPartsOutput struct {

	// If the bucket has a lifecycle rule configured with an action to abort
	// incomplete multipart uploads and the prefix in the lifecycle rule matches the
	// object name in the request, then the response includes this header indicating
	// when the initiated multipart upload will become eligible for abort operation.
	// For more information, see Aborting Incomplete Multipart Uploads Using a Bucket
	// Lifecycle Configuration (https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config)
	// . The response will also include the x-amz-abort-rule-id header that will
	// provide the ID of the lifecycle configuration rule that defines this action.
	AbortDate *time.Time

	// This header is returned along with the x-amz-abort-date header. It identifies
	// applicable lifecycle configuration rule that defines the action to abort
	// incomplete multipart uploads.
	AbortRuleId *string

	// The name of the bucket to which the multipart upload was initiated. Does not
	// return the access point ARN or access point alias if used.
	Bucket *string

	// The algorithm that was used to create a checksum of the object.
	ChecksumAlgorithm types.ChecksumAlgorithm

	// Container element that identifies who initiated the multipart upload. If the
	// initiator is an Amazon Web Services account, this element provides the same
	// information as the Owner element. If the initiator is an IAM User, this element
	// provides the user ARN and display name.
	Initiator *types.Initiator

	// Indicates whether the returned list of parts is truncated. A true value
	// indicates that the list was truncated. A list can be truncated if the number of
	// parts exceeds the limit returned in the MaxParts element.
	IsTruncated bool

	// Object key for which the multipart upload was initiated.
	Key *string

	// Maximum number of parts that were allowed in the response.
	MaxParts int32

	// When a list is truncated, this element specifies the last part in the list, as
	// well as the value to use for the part-number-marker request parameter in a
	// subsequent request.
	NextPartNumberMarker *string

	// Container element that identifies the object owner, after the object is
	// created. If multipart upload is initiated by an IAM user, this element provides
	// the parent account ID and display name.
	Owner *types.Owner

	// When a list is truncated, this element specifies the last part in the list, as
	// well as the value to use for the part-number-marker request parameter in a
	// subsequent request.
	PartNumberMarker *string

	// Container for elements related to a particular part. A response can contain
	// zero or more Part elements.
	Parts []types.Part

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// Class of storage (STANDARD or REDUCED_REDUNDANCY) used to store the uploaded
	// object.
	StorageClass types.StorageClass

	// Upload ID identifying the multipart upload whose parts are being listed.
	UploadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPartsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListParts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListParts{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
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
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addListPartsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpListPartsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListParts(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListPartsAPIClient is a client that implements the ListParts operation.
type ListPartsAPIClient interface {
	ListParts(context.Context, *ListPartsInput, ...func(*Options)) (*ListPartsOutput, error)
}

var _ ListPartsAPIClient = (*Client)(nil)

// ListPartsPaginatorOptions is the paginator options for ListParts
type ListPartsPaginatorOptions struct {
	// Sets the maximum number of parts to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPartsPaginator is a paginator for ListParts
type ListPartsPaginator struct {
	options   ListPartsPaginatorOptions
	client    ListPartsAPIClient
	params    *ListPartsInput
	nextToken *string
	firstPage bool
}

// NewListPartsPaginator returns a new ListPartsPaginator
func NewListPartsPaginator(client ListPartsAPIClient, params *ListPartsInput, optFns ...func(*ListPartsPaginatorOptions)) *ListPartsPaginator {
	if params == nil {
		params = &ListPartsInput{}
	}

	options := ListPartsPaginatorOptions{}
	if params.MaxParts != 0 {
		options.Limit = params.MaxParts
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPartsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.PartNumberMarker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPartsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListParts page.
func (p *ListPartsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPartsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PartNumberMarker = p.nextToken

	params.MaxParts = p.options.Limit

	result, err := p.client.ListParts(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = nil
	if result.IsTruncated {
		p.nextToken = result.NextPartNumberMarker
	}

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListParts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "ListParts",
	}
}

// getListPartsBucketMember returns a pointer to string denoting a provided bucket
// member valueand a boolean indicating if the input has a modeled bucket name,
func getListPartsBucketMember(input interface{}) (*string, bool) {
	in := input.(*ListPartsInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addListPartsUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getListPartsBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}

type opListPartsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  BuiltInParameterResolver
}

func (*opListPartsResolveEndpointMiddleware) ID() string {
	return "opListPartsResolveEndpointMiddleware"
}

func (m *opListPartsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*ListPartsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	params.Bucket = input.Bucket

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "s3"
			signingRegion := m.BuiltInResolver.(*BuiltInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			ctx = s3cust.SetSignerVersion(ctx, internalauth.SigV4)
		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "s3"
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*BuiltInResolver).Region
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			ctx = s3cust.SetSignerVersion(ctx, v4Scheme.Name)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			var signingName string
			if v4aScheme.SigningName == nil {
				signingName = "s3"
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			ctx = s3cust.SetSignerVersion(ctx, v4aScheme.Name)
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addListPartsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opListPartsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &BuiltInResolver{
			Region:                         options.Region,
			UseFIPS:                        options.EndpointOptions.UseFIPSEndpoint,
			UseDualStack:                   options.EndpointOptions.UseDualStackEndpoint,
			Endpoint:                       options.BaseEndpoint,
			ForcePathStyle:                 options.UsePathStyle,
			Accelerate:                     options.UseAccelerate,
			DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
			UseArnRegion:                   options.UseARNRegion,
		},
	}, "ResolveEndpoint", middleware.After)
}
