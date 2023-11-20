// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates and starts a new SPICE ingestion for a dataset. You can manually
// refresh datasets in an Enterprise edition account 32 times in a 24-hour period.
// You can manually refresh datasets in a Standard edition account 8 times in a
// 24-hour period. Each 24-hour period is measured starting 24 hours before the
// current date and time. Any ingestions operating on tagged datasets inherit the
// same tags automatically for use in access control. For an example, see How do I
// create an IAM policy to control access to Amazon EC2 resources using tags? (http://aws.amazon.com/premiumsupport/knowledge-center/iam-ec2-resource-tags/)
// in the Amazon Web Services Knowledge Center. Tags are visible on the tagged
// dataset, but not on the ingestion resource.
func (c *Client) CreateIngestion(ctx context.Context, params *CreateIngestionInput, optFns ...func(*Options)) (*CreateIngestionOutput, error) {
	if params == nil {
		params = &CreateIngestionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateIngestion", params, optFns, c.addOperationCreateIngestionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateIngestionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateIngestionInput struct {

	// The Amazon Web Services account ID.
	//
	// This member is required.
	AwsAccountId *string

	// The ID of the dataset used in the ingestion.
	//
	// This member is required.
	DataSetId *string

	// An ID for the ingestion.
	//
	// This member is required.
	IngestionId *string

	// The type of ingestion that you want to create.
	IngestionType types.IngestionType

	noSmithyDocumentSerde
}

type CreateIngestionOutput struct {

	// The Amazon Resource Name (ARN) for the data ingestion.
	Arn *string

	// An ID for the ingestion.
	IngestionId *string

	// The ingestion status.
	IngestionStatus types.IngestionStatus

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateIngestionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateIngestion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateIngestion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateIngestion"); err != nil {
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
	if err = addOpCreateIngestionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateIngestion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateIngestion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateIngestion",
	}
}
