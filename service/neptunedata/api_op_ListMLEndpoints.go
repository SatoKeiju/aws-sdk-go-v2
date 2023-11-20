// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists existing inference endpoints. See Managing inference endpoints using the
// endpoints command (https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-api-endpoints.html)
// . When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the neptune-db:ListMLEndpoints (https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#listmlendpoints)
// IAM action in that cluster.
func (c *Client) ListMLEndpoints(ctx context.Context, params *ListMLEndpointsInput, optFns ...func(*Options)) (*ListMLEndpointsOutput, error) {
	if params == nil {
		params = &ListMLEndpointsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMLEndpoints", params, optFns, c.addOperationListMLEndpointsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMLEndpointsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMLEndpointsInput struct {

	// The maximum number of items to return (from 1 to 1024; the default is 10.
	MaxItems *int32

	// The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3
	// resources. This must be listed in your DB cluster parameter group or an error
	// will occur.
	NeptuneIamRoleArn *string

	noSmithyDocumentSerde
}

type ListMLEndpointsOutput struct {

	// A page from the list of inference endpoint IDs.
	Ids []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListMLEndpointsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListMLEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListMLEndpoints{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListMLEndpoints"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListMLEndpoints(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListMLEndpoints(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListMLEndpoints",
	}
}
