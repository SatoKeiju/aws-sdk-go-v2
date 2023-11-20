// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a schema by the SchemaDefinition . The schema definition is sent to
// the Schema Registry, canonicalized, and hashed. If the hash is matched within
// the scope of the SchemaName or ARN (or the default registry, if none is
// supplied), that schema’s metadata is returned. Otherwise, a 404 or NotFound
// error is returned. Schema versions in Deleted statuses will not be included in
// the results.
func (c *Client) GetSchemaByDefinition(ctx context.Context, params *GetSchemaByDefinitionInput, optFns ...func(*Options)) (*GetSchemaByDefinitionOutput, error) {
	if params == nil {
		params = &GetSchemaByDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSchemaByDefinition", params, optFns, c.addOperationGetSchemaByDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSchemaByDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSchemaByDefinitionInput struct {

	// The definition of the schema for which schema details are required.
	//
	// This member is required.
	SchemaDefinition *string

	// This is a wrapper structure to contain schema identity fields. The structure
	// contains:
	//   - SchemaId$SchemaArn: The Amazon Resource Name (ARN) of the schema. One of
	//   SchemaArn or SchemaName has to be provided.
	//   - SchemaId$SchemaName: The name of the schema. One of SchemaArn or SchemaName
	//   has to be provided.
	//
	// This member is required.
	SchemaId *types.SchemaId

	noSmithyDocumentSerde
}

type GetSchemaByDefinitionOutput struct {

	// The date and time the schema was created.
	CreatedTime *string

	// The data format of the schema definition. Currently AVRO , JSON and PROTOBUF
	// are supported.
	DataFormat types.DataFormat

	// The Amazon Resource Name (ARN) of the schema.
	SchemaArn *string

	// The schema ID of the schema version.
	SchemaVersionId *string

	// The status of the schema version.
	Status types.SchemaVersionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSchemaByDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSchemaByDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSchemaByDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSchemaByDefinition"); err != nil {
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
	if err = addOpGetSchemaByDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSchemaByDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSchemaByDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSchemaByDefinition",
	}
}
