// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// This operation deletes an Amazon S3 on Outposts bucket's replication
// configuration. To delete an S3 bucket's replication configuration, see
// DeleteBucketReplication (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketReplication.html)
// in the Amazon S3 API Reference. Deletes the replication configuration from the
// specified S3 on Outposts bucket. To use this operation, you must have
// permissions to perform the s3-outposts:PutReplicationConfiguration action. The
// Outposts bucket owner has this permission by default and can grant it to others.
// For more information about permissions, see Setting up IAM with S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsIAM.html)
// and Managing access to S3 on Outposts buckets (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsBucketPolicy.html)
// in the Amazon S3 User Guide. It can take a while to propagate PUT or DELETE
// requests for a replication configuration to all S3 on Outposts systems.
// Therefore, the replication configuration that's returned by a GET request soon
// after a PUT or DELETE request might return a more recent result than what's on
// the Outpost. If an Outpost is offline, the delay in updating the replication
// configuration on that Outpost can be significant. All Amazon S3 on Outposts REST
// API requests for this action require an additional parameter of x-amz-outpost-id
// to be passed with the request. In addition, you must use an S3 on Outposts
// endpoint hostname prefix instead of s3-control . For an example of the request
// syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname
// prefix and the x-amz-outpost-id derived by using the access point ARN, see the
// Examples (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketReplication.html#API_control_DeleteBucketReplication_Examples)
// section. For information about S3 replication on Outposts configuration, see
// Replicating objects for S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsReplication.html)
// in the Amazon S3 User Guide. The following operations are related to
// DeleteBucketReplication :
//   - PutBucketReplication (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketReplication.html)
//   - GetBucketReplication (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketReplication.html)
func (c *Client) DeleteBucketReplication(ctx context.Context, params *DeleteBucketReplicationInput, optFns ...func(*Options)) (*DeleteBucketReplicationOutput, error) {
	if params == nil {
		params = &DeleteBucketReplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBucketReplication", params, optFns, c.addOperationDeleteBucketReplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBucketReplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteBucketReplicationInput struct {

	// The Amazon Web Services account ID of the Outposts bucket to delete the
	// replication configuration for.
	//
	// This member is required.
	AccountId *string

	// Specifies the S3 on Outposts bucket to delete the replication configuration
	// for. For using this parameter with Amazon S3 on Outposts with the REST API, you
	// must specify the name and the x-amz-outpost-id as well. For using this parameter
	// with S3 on Outposts with the Amazon Web Services SDK and CLI, you must specify
	// the ARN of the bucket accessed in the format
	// arn:aws:s3-outposts:::outpost//bucket/ . For example, to access the bucket
	// reports through Outpost my-outpost owned by account 123456789012 in Region
	// us-west-2 , use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports .
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string

	noSmithyDocumentSerde
}

func (in *DeleteBucketReplicationInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.Bucket = in.Bucket
	p.RequiresAccountId = ptr.Bool(true)
}

type DeleteBucketReplicationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteBucketReplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteBucketReplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteBucketReplication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteBucketReplication"); err != nil {
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
	if err = s3controlcust.AddUpdateOutpostARN(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addV4DetectSkewMiddleware(stack, options); err != nil {
		return err
	}
	if err = addEndpointPrefix_opDeleteBucketReplicationMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteBucketReplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBucketReplication(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addDeleteBucketReplicationUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addStashOperationInput(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = s3controlcust.AddDisableHostPrefixMiddleware(stack); err != nil {
		return err
	}
	return nil
}

func (m *DeleteBucketReplicationInput) GetARNMember() (*string, bool) {
	if m.Bucket == nil {
		return nil, false
	}
	return m.Bucket, true
}

func (m *DeleteBucketReplicationInput) SetARNMember(v string) error {
	m.Bucket = &v
	return nil
}

type endpointPrefix_opDeleteBucketReplicationMiddleware struct {
}

func (*endpointPrefix_opDeleteBucketReplicationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteBucketReplicationMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	opaqueInput := getOperationInput(ctx)
	input, ok := opaqueInput.(*DeleteBucketReplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", opaqueInput)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opDeleteBucketReplicationMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opDeleteBucketReplicationMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opDeleteBucketReplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteBucketReplication",
	}
}

func copyDeleteBucketReplicationInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*DeleteBucketReplicationInput)
	if !ok {
		return nil, fmt.Errorf("expect *DeleteBucketReplicationInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *DeleteBucketReplicationInput) copy() interface{} {
	v := *in
	return &v
}
func getDeleteBucketReplicationARNMember(input interface{}) (*string, bool) {
	in := input.(*DeleteBucketReplicationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setDeleteBucketReplicationARNMember(input interface{}, v string) error {
	in := input.(*DeleteBucketReplicationInput)
	in.Bucket = &v
	return nil
}
func backFillDeleteBucketReplicationAccountID(input interface{}, v string) error {
	in := input.(*DeleteBucketReplicationInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addDeleteBucketReplicationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getDeleteBucketReplicationARNMember,
			BackfillAccountID: backFillDeleteBucketReplicationAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setDeleteBucketReplicationARNMember,
			CopyInput:         copyDeleteBucketReplicationInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
