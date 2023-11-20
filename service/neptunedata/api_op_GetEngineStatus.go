// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata/document"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the status of the graph database on the host. When invoking this
// operation in a Neptune cluster that has IAM authentication enabled, the IAM user
// or role making the request must have a policy attached that allows the
// neptune-db:GetEngineStatus (https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getenginestatus)
// IAM action in that cluster.
func (c *Client) GetEngineStatus(ctx context.Context, params *GetEngineStatusInput, optFns ...func(*Options)) (*GetEngineStatusOutput, error) {
	if params == nil {
		params = &GetEngineStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEngineStatus", params, optFns, c.addOperationGetEngineStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEngineStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEngineStatusInput struct {
	noSmithyDocumentSerde
}

type GetEngineStatusOutput struct {

	// Set to the Neptune engine version running on your DB cluster. If this engine
	// version has been manually patched since it was released, the version number is
	// prefixed by Patch- .
	DbEngineVersion *string

	// Set to enabled if the DFE engine is fully enabled, or to viaQueryHint (the
	// default) if the DFE engine is only used with queries that have the useDFE query
	// hint set to true .
	DfeQueryEngine *string

	// Contains status information about the features enabled on your DB cluster.
	Features map[string]document.Interface

	// Contains information about the Gremlin query language available on your
	// cluster. Specifically, it contains a version field that specifies the current
	// TinkerPop version being used by the engine.
	Gremlin *types.QueryLanguageVersion

	// Contains Lab Mode settings being used by the engine.
	LabMode map[string]string

	// Contains information about the openCypher query language available on your
	// cluster. Specifically, it contains a version field that specifies the current
	// operCypher version being used by the engine.
	Opencypher *types.QueryLanguageVersion

	// Set to reader if the instance is a read-replica, or to writer if the instance
	// is the primary instance.
	Role *string

	// If there are transactions being rolled back, this field is set to the number of
	// such transactions. If there are none, the field doesn't appear at all.
	RollingBackTrxCount *int32

	// Set to the start time of the earliest transaction being rolled back. If no
	// transactions are being rolled back, the field doesn't appear at all.
	RollingBackTrxEarliestStartTime *string

	// Contains information about the current settings on your DB cluster. For
	// example, contains the current cluster query timeout setting (
	// clusterQueryTimeoutInMs ).
	Settings map[string]string

	// Contains information about the SPARQL query language available on your cluster.
	// Specifically, it contains a version field that specifies the current SPARQL
	// version being used by the engine.
	Sparql *types.QueryLanguageVersion

	// Set to the UTC time at which the current server process started.
	StartTime *string

	// Set to healthy if the instance is not experiencing problems. If the instance is
	// recovering from a crash or from being rebooted and there are active transactions
	// running from the latest server shutdown, status is set to recovery .
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEngineStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEngineStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEngineStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetEngineStatus"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEngineStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEngineStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetEngineStatus",
	}
}
