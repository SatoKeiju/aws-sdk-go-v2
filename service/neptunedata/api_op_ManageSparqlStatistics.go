// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Manages the generation and use of RDF graph statistics. When invoking this
// operation in a Neptune cluster that has IAM authentication enabled, the IAM user
// or role making the request must have a policy attached that allows the
// neptune-db:ManageStatistics (https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#managestatistics)
// IAM action in that cluster.
func (c *Client) ManageSparqlStatistics(ctx context.Context, params *ManageSparqlStatisticsInput, optFns ...func(*Options)) (*ManageSparqlStatisticsOutput, error) {
	if params == nil {
		params = &ManageSparqlStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ManageSparqlStatistics", params, optFns, c.addOperationManageSparqlStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ManageSparqlStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ManageSparqlStatisticsInput struct {

	// The statistics generation mode. One of: DISABLE_AUTOCOMPUTE , ENABLE_AUTOCOMPUTE
	// , or REFRESH , the last of which manually triggers DFE statistics generation.
	Mode types.StatisticsAutoGenerationMode

	noSmithyDocumentSerde
}

type ManageSparqlStatisticsOutput struct {

	// The HTTP return code of the request. If the request succeeded, the code is 200.
	//
	// This member is required.
	Status *string

	// This is only returned for refresh mode.
	Payload *types.RefreshStatisticsIdMap

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationManageSparqlStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpManageSparqlStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpManageSparqlStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ManageSparqlStatistics"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opManageSparqlStatistics(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opManageSparqlStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ManageSparqlStatistics",
	}
}
