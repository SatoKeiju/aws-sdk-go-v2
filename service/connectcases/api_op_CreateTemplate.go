// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcases

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connectcases/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a template in the Cases domain. This template is used to define the
// case object model (that is, to define what data can be captured on cases) in a
// Cases domain. A template must have a unique name within a domain, and it must
// reference existing field IDs and layout IDs. Additionally, multiple fields with
// same IDs are not allowed within the same Template. A template can be either
// Active or Inactive, as indicated by its status. Inactive templates cannot be
// used to create cases.
func (c *Client) CreateTemplate(ctx context.Context, params *CreateTemplateInput, optFns ...func(*Options)) (*CreateTemplateOutput, error) {
	if params == nil {
		params = &CreateTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTemplate", params, optFns, c.addOperationCreateTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTemplateInput struct {

	// The unique identifier of the Cases domain.
	//
	// This member is required.
	DomainId *string

	// A name for the template. It must be unique per domain.
	//
	// This member is required.
	Name *string

	// A brief description of the template.
	Description *string

	// Configuration of layouts associated to the template.
	LayoutConfiguration *types.LayoutConfiguration

	// A list of fields that must contain a value for a case to be successfully
	// created with this template.
	RequiredFields []types.RequiredField

	// The status of the template.
	Status types.TemplateStatus

	noSmithyDocumentSerde
}

type CreateTemplateOutput struct {

	// The Amazon Resource Name (ARN) of the newly created template.
	//
	// This member is required.
	TemplateArn *string

	// A unique identifier of a template.
	//
	// This member is required.
	TemplateId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateTemplate"); err != nil {
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
	if err = addOpCreateTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateTemplate",
	}
}
