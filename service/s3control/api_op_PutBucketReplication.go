// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3controlcust "github.com/aws/aws-sdk-go-v2/service/s3control/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// This action creates an Amazon S3 on Outposts bucket's replication
// configuration. To create an S3 bucket's replication configuration, see
// PutBucketReplication (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketReplication.html)
// in the Amazon S3 API Reference. Creates a replication configuration or replaces
// an existing one. For information about S3 replication on Outposts configuration,
// see Replicating objects for S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsReplication.html)
// in the Amazon S3 User Guide. It can take a while to propagate PUT or DELETE
// requests for a replication configuration to all S3 on Outposts systems.
// Therefore, the replication configuration that's returned by a GET request soon
// after a PUT or DELETE request might return a more recent result than what's on
// the Outpost. If an Outpost is offline, the delay in updating the replication
// configuration on that Outpost can be significant. Specify the replication
// configuration in the request body. In the replication configuration, you provide
// the following information:
//   - The name of the destination bucket or buckets where you want S3 on Outposts
//     to replicate objects
//   - The Identity and Access Management (IAM) role that S3 on Outposts can
//     assume to replicate objects on your behalf
//   - Other relevant information, such as replication rules
//
// A replication configuration must include at least one rule and can contain a
// maximum of 100. Each rule identifies a subset of objects to replicate by
// filtering the objects in the source Outposts bucket. To choose additional
// subsets of objects to replicate, add a rule for each subset. To specify a subset
// of the objects in the source Outposts bucket to apply a replication rule to, add
// the Filter element as a child of the Rule element. You can filter objects based
// on an object key prefix, one or more object tags, or both. When you add the
// Filter element in the configuration, you must also add the following elements:
// DeleteMarkerReplication , Status , and Priority . Using PutBucketReplication on
// Outposts requires that both the source and destination buckets must have
// versioning enabled. For information about enabling versioning on a bucket, see
// Managing S3 Versioning for your S3 on Outposts bucket (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsManagingVersioning.html)
// . For information about S3 on Outposts replication failure reasons, see
// Replication failure reasons (https://docs.aws.amazon.com/AmazonS3/latest/userguide/outposts-replication-eventbridge.html#outposts-replication-failure-codes)
// in the Amazon S3 User Guide. Handling Replication of Encrypted Objects Outposts
// buckets are encrypted at all times. All the objects in the source Outposts
// bucket are encrypted and can be replicated. Also, all the replicas in the
// destination Outposts bucket are encrypted with the same encryption key as the
// objects in the source Outposts bucket. Permissions To create a
// PutBucketReplication request, you must have
// s3-outposts:PutReplicationConfiguration permissions for the bucket. The Outposts
// bucket owner has this permission by default and can grant it to others. For more
// information about permissions, see Setting up IAM with S3 on Outposts (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsIAM.html)
// and Managing access to S3 on Outposts buckets (https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsBucketPolicy.html)
// . To perform this operation, the user or role must also have the iam:CreateRole
// and iam:PassRole permissions. For more information, see Granting a user
// permissions to pass a role to an Amazon Web Services service (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html)
// . All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the Examples (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketReplication.html#API_control_PutBucketReplication_Examples)
// section. The following operations are related to PutBucketReplication :
//   - GetBucketReplication (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketReplication.html)
//   - DeleteBucketReplication (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketReplication.html)
func (c *Client) PutBucketReplication(ctx context.Context, params *PutBucketReplicationInput, optFns ...func(*Options)) (*PutBucketReplicationOutput, error) {
	if params == nil {
		params = &PutBucketReplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketReplication", params, optFns, c.addOperationPutBucketReplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketReplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketReplicationInput struct {

	// The Amazon Web Services account ID of the Outposts bucket.
	//
	// This member is required.
	AccountId *string

	// Specifies the S3 on Outposts bucket to set the configuration for. For using
	// this parameter with Amazon S3 on Outposts with the REST API, you must specify
	// the name and the x-amz-outpost-id as well. For using this parameter with S3 on
	// Outposts with the Amazon Web Services SDK and CLI, you must specify the ARN of
	// the bucket accessed in the format arn:aws:s3-outposts:::outpost//bucket/ . For
	// example, to access the bucket reports through Outpost my-outpost owned by
	// account 123456789012 in Region us-west-2 , use the URL encoding of
	// arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports .
	// The value must be URL encoded.
	//
	// This member is required.
	Bucket *string

	//
	//
	// This member is required.
	ReplicationConfiguration *types.ReplicationConfiguration

	noSmithyDocumentSerde
}

func (in *PutBucketReplicationInput) bindEndpointParams(p *EndpointParameters) {
	p.AccountId = in.AccountId
	p.Bucket = in.Bucket
	p.RequiresAccountId = ptr.Bool(true)
}

type PutBucketReplicationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutBucketReplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketReplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketReplication{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutBucketReplication"); err != nil {
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
	if err = smithyhttp.AddContentChecksumMiddleware(stack); err != nil {
		return err
	}
	if err = addEndpointPrefix_opPutBucketReplicationMiddleware(stack); err != nil {
		return err
	}
	if err = addOpPutBucketReplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketReplication(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addPutBucketReplicationUpdateEndpoint(stack, options); err != nil {
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

func (m *PutBucketReplicationInput) GetARNMember() (*string, bool) {
	if m.Bucket == nil {
		return nil, false
	}
	return m.Bucket, true
}

func (m *PutBucketReplicationInput) SetARNMember(v string) error {
	m.Bucket = &v
	return nil
}

type endpointPrefix_opPutBucketReplicationMiddleware struct {
}

func (*endpointPrefix_opPutBucketReplicationMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opPutBucketReplicationMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
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
	input, ok := opaqueInput.(*PutBucketReplicationInput)
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
func addEndpointPrefix_opPutBucketReplicationMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opPutBucketReplicationMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opPutBucketReplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutBucketReplication",
	}
}

func copyPutBucketReplicationInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*PutBucketReplicationInput)
	if !ok {
		return nil, fmt.Errorf("expect *PutBucketReplicationInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func (in *PutBucketReplicationInput) copy() interface{} {
	v := *in
	return &v
}
func getPutBucketReplicationARNMember(input interface{}) (*string, bool) {
	in := input.(*PutBucketReplicationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func setPutBucketReplicationARNMember(input interface{}, v string) error {
	in := input.(*PutBucketReplicationInput)
	in.Bucket = &v
	return nil
}
func backFillPutBucketReplicationAccountID(input interface{}, v string) error {
	in := input.(*PutBucketReplicationInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addPutBucketReplicationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: getPutBucketReplicationARNMember,
			BackfillAccountID: backFillPutBucketReplicationAccountID,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    setPutBucketReplicationARNMember,
			CopyInput:         copyPutBucketReplicationInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}
