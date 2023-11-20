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

// The StartTableDataImportJob API allows you to start an import job on a table.
// This API will only return the id of the job that was started. To find out the
// status of the import request, you need to call the DescribeTableDataImportJob
// API.
func (c *Client) StartTableDataImportJob(ctx context.Context, params *StartTableDataImportJobInput, optFns ...func(*Options)) (*StartTableDataImportJobOutput, error) {
	if params == nil {
		params = &StartTableDataImportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartTableDataImportJob", params, optFns, c.addOperationStartTableDataImportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartTableDataImportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartTableDataImportJobInput struct {

	// The request token for performing the update action. Request tokens help to
	// identify duplicate requests. If a call times out or fails due to a transient
	// error like a failed network connection, you can retry the call with the same
	// request token. The service ensures that if the first call using that request
	// token is successfully performed, the second call will not perform the action
	// again. Note that request tokens are valid only for a few minutes. You cannot use
	// request tokens to dedupe requests spanning hours or days.
	//
	// This member is required.
	ClientRequestToken *string

	// The format of the data that is being imported. Currently the only option
	// supported is "DELIMITED_TEXT".
	//
	// This member is required.
	DataFormat types.ImportSourceDataFormat

	// The source of the data that is being imported. The size of source must be no
	// larger than 100 MB. Source must have no more than 100,000 cells and no more than
	// 1,000 rows.
	//
	// This member is required.
	DataSource *types.ImportDataSource

	// The ID of the table where the rows are being imported. If a table with the
	// specified id could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	DestinationTableId *string

	// The options for customizing this import request.
	//
	// This member is required.
	ImportOptions *types.ImportOptions

	// The ID of the workbook where the rows are being imported. If a workbook with
	// the specified id could not be found, this API throws ResourceNotFoundException.
	//
	// This member is required.
	WorkbookId *string

	noSmithyDocumentSerde
}

type StartTableDataImportJobOutput struct {

	// The id that is assigned to this import job. Future requests to find out the
	// status of this import job need to send this id in the appropriate parameter in
	// the request.
	//
	// This member is required.
	JobId *string

	// The status of the import job immediately after submitting the request.
	//
	// This member is required.
	JobStatus types.TableDataImportJobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartTableDataImportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartTableDataImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartTableDataImportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartTableDataImportJob"); err != nil {
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
	if err = addOpStartTableDataImportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartTableDataImportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartTableDataImportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartTableDataImportJob",
	}
}
