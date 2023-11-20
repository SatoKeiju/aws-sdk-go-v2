// Code generated by smithy-go-codegen DO NOT EDIT.

package pcaconnectorad

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a connector for Active Directory. You must provide the Amazon Resource
// Name (ARN) of the connector that you want to delete. You can find the ARN by
// calling the
// https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_ListConnectors (https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_ListConnectors)
// action. Deleting a connector does not deregister your directory with Amazon Web
// Services Private CA. You can deregister your directory by calling the
// https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_DeleteDirectoryRegistration (https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_DeleteDirectoryRegistration)
// action.
func (c *Client) DeleteConnector(ctx context.Context, params *DeleteConnectorInput, optFns ...func(*Options)) (*DeleteConnectorOutput, error) {
	if params == nil {
		params = &DeleteConnectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteConnector", params, optFns, c.addOperationDeleteConnectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteConnectorInput struct {

	// The Amazon Resource Name (ARN) that was returned when you called CreateConnector (https://docs.aws.amazon.com/pca-connector-ad/latest/APIReference/API_CreateConnector.html)
	// .
	//
	// This member is required.
	ConnectorArn *string

	noSmithyDocumentSerde
}

type DeleteConnectorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteConnectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteConnector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteConnector{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteConnector"); err != nil {
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
	if err = addOpDeleteConnectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteConnector(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteConnector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteConnector",
	}
}
