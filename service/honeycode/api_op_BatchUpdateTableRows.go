// Code generated by smithy-go-codegen DO NOT EDIT.

package honeycode

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/honeycode/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The BatchUpdateTableRows API allows you to update one or more rows in a table
// in a workbook. You can specify the values to set in some or all of the columns
// in the table for the specified rows. If a column is not explicitly specified in
// a particular row, then that column will not be updated for that row. To clear
// out the data in a specific cell, you need to set the value as an empty string
// ("").
func (c *Client) BatchUpdateTableRows(ctx context.Context, params *BatchUpdateTableRowsInput, optFns ...func(*Options)) (*BatchUpdateTableRowsOutput, error) {
	if params == nil {
		params = &BatchUpdateTableRowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchUpdateTableRows", params, optFns, c.addOperationBatchUpdateTableRowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchUpdateTableRowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUpdateTableRowsInput struct {

	// The list of rows to update in the table. Each item in this list needs to
	// contain the row id to update along with the map of column id to cell values for
	// each column in that row that needs to be updated. You need to specify at least
	// one row in this list, and for each row, you need to specify at least one column
	// to update. Note that if one of the row or column ids in the request does not
	// exist in the table, then the request fails and no updates are made to the table.
	//
	// This member is required.
	RowsToUpdate []types.UpdateRowData

	// The ID of the table where the rows are being updated. If a table with the
	// specified id could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	TableId *string

	// The ID of the workbook where the rows are being updated. If a workbook with the
	// specified id could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	WorkbookId *string

	// The request token for performing the update action. Request tokens help to
	// identify duplicate requests. If a call times out or fails due to a transient
	// error like a failed network connection, you can retry the call with the same
	// request token. The service ensures that if the first call using that request
	// token is successfully performed, the second call will not perform the action
	// again. Note that request tokens are valid only for a few minutes. You cannot use
	// request tokens to dedupe requests spanning hours or days.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type BatchUpdateTableRowsOutput struct {

	// The updated workbook cursor after adding the new rows at the end of the table.
	//
	// This member is required.
	WorkbookCursor int64

	// The list of batch items in the request that could not be updated in the table.
	// Each element in this list contains one item from the request that could not be
	// updated in the table along with the reason why that item could not be updated.
	FailedBatchItems []types.FailedBatchItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchUpdateTableRowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchUpdateTableRows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUpdateTableRows{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchUpdateTableRows"); err != nil {
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
	if err = addOpBatchUpdateTableRowsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUpdateTableRows(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchUpdateTableRows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchUpdateTableRows",
	}
}
