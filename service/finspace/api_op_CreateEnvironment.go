// Code generated by smithy-go-codegen DO NOT EDIT.

package finspace

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/finspace/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new FinSpace environment.
//
// Deprecated: This method will be discontinued.
func (c *Client) CreateEnvironment(ctx context.Context, params *CreateEnvironmentInput, optFns ...func(*Options)) (*CreateEnvironmentOutput, error) {
	if params == nil {
		params = &CreateEnvironmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironment", params, optFns, c.addOperationCreateEnvironmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentInput struct {

	// The name of the FinSpace environment to be created.
	//
	// This member is required.
	Name *string

	// The list of Amazon Resource Names (ARN) of the data bundles to install.
	// Currently supported data bundle ARNs:
	//   - arn:aws:finspace:${Region}::data-bundle/capital-markets-sample - Contains
	//   sample Capital Markets datasets, categories and controlled vocabularies.
	//   - arn:aws:finspace:${Region}::data-bundle/taq (default) - Contains trades and
	//   quotes data in addition to sample Capital Markets data.
	DataBundles []string

	// The description of the FinSpace environment to be created.
	Description *string

	// Authentication mode for the environment.
	//   - FEDERATED - Users access FinSpace through Single Sign On (SSO) via your
	//   Identity provider.
	//   - LOCAL - Users access FinSpace via email and password managed within the
	//   FinSpace environment.
	FederationMode types.FederationMode

	// Configuration information when authentication mode is FEDERATED.
	FederationParameters *types.FederationParameters

	// The KMS key id to encrypt your data in the FinSpace environment.
	KmsKeyId *string

	// Configuration information for the superuser.
	SuperuserParameters *types.SuperuserParameters

	// Add tags to your FinSpace environment.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateEnvironmentOutput struct {

	// The Amazon Resource Name (ARN) of the FinSpace environment that you created.
	EnvironmentArn *string

	// The unique identifier for FinSpace environment that you created.
	EnvironmentId *string

	// The sign-in URL for the web application of the FinSpace environment you created.
	EnvironmentUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEnvironmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateEnvironment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEnvironment"); err != nil {
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
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
	if err = addOpCreateEnvironmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEnvironment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEnvironment",
	}
}
