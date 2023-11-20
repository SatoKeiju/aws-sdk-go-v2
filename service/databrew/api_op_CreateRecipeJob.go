// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databrew/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new job to transform input data, using steps defined in an existing
// Glue DataBrew recipe
func (c *Client) CreateRecipeJob(ctx context.Context, params *CreateRecipeJobInput, optFns ...func(*Options)) (*CreateRecipeJobOutput, error) {
	if params == nil {
		params = &CreateRecipeJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRecipeJob", params, optFns, c.addOperationCreateRecipeJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRecipeJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRecipeJobInput struct {

	// A unique name for the job. Valid characters are alphanumeric (A-Z, a-z, 0-9),
	// hyphen (-), period (.), and space.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// to be assumed when DataBrew runs the job.
	//
	// This member is required.
	RoleArn *string

	// One or more artifacts that represent the Glue Data Catalog output from running
	// the job.
	DataCatalogOutputs []types.DataCatalogOutput

	// Represents a list of JDBC database output objects which defines the output
	// destination for a DataBrew recipe job to write to.
	DatabaseOutputs []types.DatabaseOutput

	// The name of the dataset that this job processes.
	DatasetName *string

	// The Amazon Resource Name (ARN) of an encryption key that is used to protect the
	// job.
	EncryptionKeyArn *string

	// The encryption mode for the job, which can be one of the following:
	//   - SSE-KMS - Server-side encryption with keys managed by KMS.
	//   - SSE-S3 - Server-side encryption with keys managed by Amazon S3.
	EncryptionMode types.EncryptionMode

	// Enables or disables Amazon CloudWatch logging for the job. If logging is
	// enabled, CloudWatch writes one log stream for each job run.
	LogSubscription types.LogSubscription

	// The maximum number of nodes that DataBrew can consume when the job processes
	// data.
	MaxCapacity int32

	// The maximum number of times to retry the job after a job run fails.
	MaxRetries int32

	// One or more artifacts that represent the output from running the job.
	Outputs []types.Output

	// Either the name of an existing project, or a combination of a recipe and a
	// dataset to associate with the recipe.
	ProjectName *string

	// Represents the name and version of a DataBrew recipe.
	RecipeReference *types.RecipeReference

	// Metadata tags to apply to this job.
	Tags map[string]string

	// The job's timeout in minutes. A job that attempts to run longer than this
	// timeout period ends with a status of TIMEOUT .
	Timeout int32

	noSmithyDocumentSerde
}

type CreateRecipeJobOutput struct {

	// The name of the job that you created.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRecipeJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateRecipeJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRecipeJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRecipeJob"); err != nil {
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
	if err = addOpCreateRecipeJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRecipeJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateRecipeJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRecipeJob",
	}
}
