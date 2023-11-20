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

// Gets the status of a specified Gremlin query. When invoking this operation in a
// Neptune cluster that has IAM authentication enabled, the IAM user or role making
// the request must have a policy attached that allows the
// neptune-db:GetQueryStatus (https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getquerystatus)
// IAM action in that cluster. Note that the neptune-db:QueryLanguage:Gremlin (https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html#iam-neptune-condition-keys)
// IAM condition key can be used in the policy document to restrict the use of
// Gremlin queries (see Condition keys available in Neptune IAM data-access policy
// statements (https://docs.aws.amazon.com/neptune/latest/userguide/iam-data-condition-keys.html)
// ).
func (c *Client) GetGremlinQueryStatus(ctx context.Context, params *GetGremlinQueryStatusInput, optFns ...func(*Options)) (*GetGremlinQueryStatusOutput, error) {
	if params == nil {
		params = &GetGremlinQueryStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGremlinQueryStatus", params, optFns, c.addOperationGetGremlinQueryStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGremlinQueryStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGremlinQueryStatusInput struct {

	// The unique identifier that identifies the Gremlin query.
	//
	// This member is required.
	QueryId *string

	noSmithyDocumentSerde
}

type GetGremlinQueryStatusOutput struct {

	// The evaluation status of the Gremlin query.
	QueryEvalStats *types.QueryEvalStats

	// The ID of the query for which status is being returned.
	QueryId *string

	// The Gremlin query string.
	QueryString *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGremlinQueryStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetGremlinQueryStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGremlinQueryStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetGremlinQueryStatus"); err != nil {
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
	if err = addOpGetGremlinQueryStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGremlinQueryStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetGremlinQueryStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetGremlinQueryStatus",
	}
}
