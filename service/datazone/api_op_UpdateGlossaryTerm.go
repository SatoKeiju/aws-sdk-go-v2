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
)

// Updates a business glossary term in Amazon DataZone.
func (c *Client) UpdateGlossaryTerm(ctx context.Context, params *UpdateGlossaryTermInput, optFns ...func(*Options)) (*UpdateGlossaryTermOutput, error) {
	if params == nil {
		params = &UpdateGlossaryTermInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGlossaryTerm", params, optFns, c.addOperationUpdateGlossaryTermMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGlossaryTermOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGlossaryTermInput struct {

	// The identifier of the Amazon DataZone domain in which a business glossary term
	// is to be updated.
	//
	// This member is required.
	DomainIdentifier *string

	// The identifier of the business glossary term that is to be updated.
	//
	// This member is required.
	Identifier *string

	// The identifier of the business glossary in which a term is to be updated.
	GlossaryIdentifier *string

	// The long description to be updated as part of the UpdateGlossaryTerm action.
	LongDescription *string

	// The name to be updated as part of the UpdateGlossaryTerm action.
	Name *string

	// The short description to be updated as part of the UpdateGlossaryTerm action.
	ShortDescription *string

	// The status to be updated as part of the UpdateGlossaryTerm action.
	Status types.GlossaryTermStatus

	// The term relations to be updated as part of the UpdateGlossaryTerm action.
	TermRelations *types.TermRelations

	noSmithyDocumentSerde
}

type UpdateGlossaryTermOutput struct {

	// The identifier of the Amazon DataZone domain in which a business glossary term
	// is to be updated.
	//
	// This member is required.
	DomainId *string

	// The identifier of the business glossary in which a term is to be updated.
	//
	// This member is required.
	GlossaryId *string

	// The identifier of the business glossary term that is to be updated.
	//
	// This member is required.
	Id *string

	// The name to be updated as part of the UpdateGlossaryTerm action.
	//
	// This member is required.
	Name *string

	// The status to be updated as part of the UpdateGlossaryTerm action.
	//
	// This member is required.
	Status types.GlossaryTermStatus

	// The long description to be updated as part of the UpdateGlossaryTerm action.
	LongDescription *string

	// The short description to be updated as part of the UpdateGlossaryTerm action.
	ShortDescription *string

	// The term relations to be updated as part of the UpdateGlossaryTerm action.
	TermRelations *types.TermRelations

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGlossaryTermMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGlossaryTerm{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGlossaryTerm{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateGlossaryTerm"); err != nil {
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
	if err = addOpUpdateGlossaryTermValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGlossaryTerm(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateGlossaryTerm(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateGlossaryTerm",
	}
}
