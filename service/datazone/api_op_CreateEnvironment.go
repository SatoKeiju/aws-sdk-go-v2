// Code generated by smithy-go-codegen DO NOT EDIT.

package datazone

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datazone/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Create an Amazon DataZone environment.
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

	// The identifier of the Amazon DataZone domain in which the environment is
	// created.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the environment profile that is used to create this Amazon
	// DataZone environment.
	//
	// This member is required.
	EnvironmentProfileIdentifier *string

	// The name of the Amazon DataZone environment.
	//
	// This member is required.
	Name *string

	// The identifier of the Amazon DataZone project in which this environment is
	// created.
	//
	// This member is required.
	ProjectIdentifier *string

	// The description of the Amazon DataZone environment.
	Description *string

	// The glossary terms that can be used in this Amazon DataZone environment.
	GlossaryTerms []string

	// The user parameters of this Amazon DataZone environment.
	UserParameters []types.EnvironmentParameter

	noSmithyDocumentSerde
}

type CreateEnvironmentOutput struct {

	// The Amazon DataZone user who created this environment.
	//
	// This member is required.
	CreatedBy *string

	// The identifier of the Amazon DataZone domain in which the environment is
	// created.
	//
	// This member is required.
	DomainId *string

	// The ID of the environment profile with which this Amazon DataZone environment
	// was created.
	//
	// This member is required.
	EnvironmentProfileId *string

	// The name of this environment.
	//
	// This member is required.
	Name *string

	// The ID of the Amazon DataZone project in which this environment is created.
	//
	// This member is required.
	ProjectId *string

	// The provider of this Amazon DataZone environment.
	//
	// This member is required.
	Provider *string

	// The Amazon Web Services account in which the Amazon DataZone environment is
	// created.
	AwsAccountId *string

	// The Amazon Web Services region in which the Amazon DataZone environment is
	// created.
	AwsAccountRegion *string

	// The timestamp of when the environment was created.
	CreatedAt *time.Time

	// The deployment properties of this Amazon DataZone environment.
	DeploymentProperties *types.DeploymentProperties

	// The description of this Amazon DataZone environment.
	Description *string

	// The configurable actions of this Amazon DataZone environment.
	EnvironmentActions []types.ConfigurableEnvironmentAction

	// The ID of the blueprint with which this Amazon DataZone environment was created.
	EnvironmentBlueprintId *string

	// The glossary terms that can be used in this Amazon DataZone environment.
	GlossaryTerms []string

	// The ID of this Amazon DataZone environment.
	Id *string

	// The details of the last deployment of this Amazon DataZone environment.
	LastDeployment *types.Deployment

	// The provisioned resources of this Amazon DataZone environment.
	ProvisionedResources []types.Resource

	// The provisioning properties of this Amazon DataZone environment.
	ProvisioningProperties types.ProvisioningProperties

	// The status of this Amazon DataZone environment.
	Status types.EnvironmentStatus

	// The timestamp of when this environment was updated.
	UpdatedAt *time.Time

	// The user parameters of this Amazon DataZone environment.
	UserParameters []types.CustomParameter

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
