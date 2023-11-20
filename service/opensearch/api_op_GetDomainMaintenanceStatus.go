// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// The status of the maintenance action.
func (c *Client) GetDomainMaintenanceStatus(ctx context.Context, params *GetDomainMaintenanceStatusInput, optFns ...func(*Options)) (*GetDomainMaintenanceStatusOutput, error) {
	if params == nil {
		params = &GetDomainMaintenanceStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDomainMaintenanceStatus", params, optFns, c.addOperationGetDomainMaintenanceStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDomainMaintenanceStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the GetDomainMaintenanceStatus operation.
type GetDomainMaintenanceStatusInput struct {

	// The name of the domain.
	//
	// This member is required.
	DomainName *string

	// The request ID of the maintenance action.
	//
	// This member is required.
	MaintenanceId *string

	noSmithyDocumentSerde
}

// The result of a GetDomainMaintenanceStatus request that information about the
// requested action.
type GetDomainMaintenanceStatusOutput struct {

	// The action name.
	Action types.MaintenanceType

	// The time at which the action was created.
	CreatedAt *time.Time

	// The node ID of the maintenance action.
	NodeId *string

	// The status of the maintenance action.
	Status types.MaintenanceStatus

	// The status message of the maintenance action.
	StatusMessage *string

	// The time at which the action was updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDomainMaintenanceStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDomainMaintenanceStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDomainMaintenanceStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDomainMaintenanceStatus"); err != nil {
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
	if err = addOpGetDomainMaintenanceStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDomainMaintenanceStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDomainMaintenanceStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDomainMaintenanceStatus",
	}
}
