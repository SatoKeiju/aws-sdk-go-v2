// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a what-if forecast export created using the CreateWhatIfForecastExport
// operation. You can delete only what-if forecast exports that have a status of
// ACTIVE or CREATE_FAILED. To get the status, use the DescribeWhatIfForecastExport
// operation.
func (c *Client) DeleteWhatIfForecastExport(ctx context.Context, params *DeleteWhatIfForecastExportInput, optFns ...func(*Options)) (*DeleteWhatIfForecastExportOutput, error) {
	if params == nil {
		params = &DeleteWhatIfForecastExportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteWhatIfForecastExport", params, optFns, c.addOperationDeleteWhatIfForecastExportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteWhatIfForecastExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteWhatIfForecastExportInput struct {

	// The Amazon Resource Name (ARN) of the what-if forecast export that you want to
	// delete.
	//
	// This member is required.
	WhatIfForecastExportArn *string

	noSmithyDocumentSerde
}

type DeleteWhatIfForecastExportOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteWhatIfForecastExportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteWhatIfForecastExport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteWhatIfForecastExport{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteWhatIfForecastExportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteWhatIfForecastExport(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opDeleteWhatIfForecastExport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DeleteWhatIfForecastExport",
	}
}
