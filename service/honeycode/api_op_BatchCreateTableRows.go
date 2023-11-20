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

// The BatchCreateTableRows API allows you to create one or more rows at the end
// of a table in a workbook. The API allows you to specify the values to set in
// some or all of the columns in the new rows. If a column is not explicitly set in
// a specific row, then the column level formula specified in the table will be
// applied to the new row. If there is no column level formula but the last row of
// the table has a formula, then that formula will be copied down to the new row.
// If there is no column level formula and no formula in the last row of the table,
// then that column will be left blank for the new rows.
func (c *Client) BatchCreateTableRows(ctx context.Context, params *BatchCreateTableRowsInput, optFns ...func(*Options)) (*BatchCreateTableRowsOutput, error) {
	if params == nil {
		params = &BatchCreateTableRowsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchCreateTableRows", params, optFns, c.addOperationBatchCreateTableRowsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchCreateTableRowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchCreateTableRowsInput struct {

	// The list of rows to create at the end of the table. Each item in this list
	// needs to have a batch item id to uniquely identify the element in the request
	// and the cells to create for that row. You need to specify at least one item in
	// this list. Note that if one of the column ids in any of the rows in the request
	// does not exist in the table, then the request fails and no updates are made to
	// the table.
	//
	// This member is required.
	RowsToCreate []types.CreateRowData

	// The ID of the table where the new rows are being added. If a table with the
	// specified ID could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	TableId *string

	// The ID of the workbook where the new rows are being added. If a workbook with
	// the specified ID could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	WorkbookId *string

	// The request token for performing the batch create operation. Request tokens
	// help to identify duplicate requests. If a call times out or fails due to a
	// transient error like a failed network connection, you can retry the call with
	// the same request token. The service ensures that if the first call using that
	// request token is successfully performed, the second call will not perform the
	// operation again. Note that request tokens are valid only for a few minutes. You
	// cannot use request tokens to dedupe requests spanning hours or days.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type BatchCreateTableRowsOutput struct {

	// The map of batch item id to the row id that was created for that item.
	//
	// This member is required.
	CreatedRows map[string]string

	// The updated workbook cursor after adding the new rows at the end of the table.
	//
	// This member is required.
	WorkbookCursor int64

	// The list of batch items in the request that could not be added to the table.
	// Each element in this list contains one item from the request that could not be
	// added to the table along with the reason why that item could not be added.
	FailedBatchItems []types.FailedBatchItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchCreateTableRowsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchCreateTableRows{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchCreateTableRows{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchCreateTableRows"); err != nil {
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
	if err = addOpBatchCreateTableRowsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchCreateTableRows(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchCreateTableRows(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchCreateTableRows",
	}
}
