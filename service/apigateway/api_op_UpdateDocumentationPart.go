// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a documentation part.
func (c *Client) UpdateDocumentationPart(ctx context.Context, params *UpdateDocumentationPartInput, optFns ...func(*Options)) (*UpdateDocumentationPartOutput, error) {
	if params == nil {
		params = &UpdateDocumentationPartInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDocumentationPart", params, optFns, c.addOperationUpdateDocumentationPartMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDocumentationPartOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Updates an existing documentation part of a given API.
type UpdateDocumentationPartInput struct {

	// The identifier of the to-be-updated documentation part.
	//
	// This member is required.
	DocumentationPartId *string

	// The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// For more information about supported patch operations, see Patch Operations (https://docs.aws.amazon.com/apigateway/latest/api/patch-operations.html)
	// .
	PatchOperations []types.PatchOperation

	noSmithyDocumentSerde
}

// A documentation part for a targeted API entity.
type UpdateDocumentationPartOutput struct {

	// The DocumentationPart identifier, generated by API Gateway when the
	// DocumentationPart is created.
	Id *string

	// The location of the API entity to which the documentation applies. Valid fields
	// depend on the targeted API entity type. All the valid location fields are not
	// required. If not explicitly specified, a valid location field is treated as a
	// wildcard and associated documentation content may be inherited by matching
	// entities, unless overridden.
	Location *types.DocumentationPartLocation

	// A content map of API-specific key-value pairs describing the targeted API
	// entity. The map must be encoded as a JSON string, e.g., "{ \"description\":
	// \"The API does ...\" }" . Only OpenAPI-compliant documentation-related fields
	// from the properties map are exported and, hence, published as part of the API
	// entity definitions, while the original documentation parts are exported in a
	// OpenAPI extension of x-amazon-apigateway-documentation .
	Properties *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDocumentationPartMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDocumentationPart{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDocumentationPart{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateDocumentationPart"); err != nil {
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
	if err = addOpUpdateDocumentationPartValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDocumentationPart(options.Region), middleware.Before); err != nil {
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
	if err = addAcceptHeader(stack); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDocumentationPart(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateDocumentationPart",
	}
}
