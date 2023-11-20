// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/vpclattice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about the association between a service network and a VPC.
func (c *Client) GetServiceNetworkVpcAssociation(ctx context.Context, params *GetServiceNetworkVpcAssociationInput, optFns ...func(*Options)) (*GetServiceNetworkVpcAssociationOutput, error) {
	if params == nil {
		params = &GetServiceNetworkVpcAssociationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceNetworkVpcAssociation", params, optFns, c.addOperationGetServiceNetworkVpcAssociationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceNetworkVpcAssociationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceNetworkVpcAssociationInput struct {

	// The ID or Amazon Resource Name (ARN) of the association.
	//
	// This member is required.
	ServiceNetworkVpcAssociationIdentifier *string

	noSmithyDocumentSerde
}

type GetServiceNetworkVpcAssociationOutput struct {

	// The Amazon Resource Name (ARN) of the association.
	Arn *string

	// The date and time that the association was created, specified in ISO-8601
	// format.
	CreatedAt *time.Time

	// The account that created the association.
	CreatedBy *string

	// The failure code.
	FailureCode *string

	// The failure message.
	FailureMessage *string

	// The ID of the specified association between the service network and the VPC.
	Id *string

	// The date and time that the association was last updated, specified in ISO-8601
	// format.
	LastUpdatedAt *time.Time

	// The IDs of the security groups.
	SecurityGroupIds []string

	// The Amazon Resource Name (ARN) of the service network.
	ServiceNetworkArn *string

	// The ID of the service network.
	ServiceNetworkId *string

	// The name of the service network.
	ServiceNetworkName *string

	// The status of the association.
	Status types.ServiceNetworkVpcAssociationStatus

	// The ID of the VPC.
	VpcId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetServiceNetworkVpcAssociationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetServiceNetworkVpcAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetServiceNetworkVpcAssociation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetServiceNetworkVpcAssociation"); err != nil {
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
	if err = addOpGetServiceNetworkVpcAssociationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceNetworkVpcAssociation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetServiceNetworkVpcAssociation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetServiceNetworkVpcAssociation",
	}
}
