// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Provides options for retrieving a job list for an Amazon Glacier vault.
type ListJobsInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The state of the jobs to return. You can specify true or false.
	Completed *string `location:"querystring" locationName:"completed" type:"string"`

	// The maximum number of jobs to be returned. The default limit is 50. The number
	// of jobs returned might be fewer than the specified limit, but the number
	// of returned jobs never exceeds the limit.
	Limit *string `location:"querystring" locationName:"limit" type:"string"`

	// An opaque string used for pagination. This value specifies the job at which
	// the listing of jobs should begin. Get the marker value from a previous List
	// Jobs response. You only need to include the marker if you are continuing
	// the pagination of results started in a previous List Jobs request.
	Marker *string `location:"querystring" locationName:"marker" type:"string"`

	// The type of job status to return. You can specify the following values: InProgress,
	// Succeeded, or Failed.
	Statuscode *string `location:"querystring" locationName:"statuscode" type:"string"`

	// The name of the vault.
	//
	// VaultName is a required field
	VaultName *string `location:"uri" locationName:"vaultName" type:"string" required:"true"`
}

// String returns the string representation
func (s ListJobsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListJobsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListJobsInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if s.VaultName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VaultName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListJobsInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.AccountId != nil {
		v := *s.AccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "accountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VaultName != nil {
		v := *s.VaultName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "vaultName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Completed != nil {
		v := *s.Completed

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "completed", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Statuscode != nil {
		v := *s.Statuscode

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "statuscode", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Contains the Amazon Glacier response to your request.
type ListJobsOutput struct {
	_ struct{} `type:"structure"`

	// A list of job objects. Each job object contains metadata describing the job.
	JobList []GlacierJobDescription `type:"list"`

	// An opaque string used for pagination that specifies the job at which the
	// listing of jobs should begin. You get the marker value from a previous List
	// Jobs response. You only need to include the marker if you are continuing
	// the pagination of the results started in a previous List Jobs request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s ListJobsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListJobsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.JobList != nil {
		v := s.JobList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "JobList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListJobs = "ListJobs"

// ListJobsRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation lists jobs for a vault, including jobs that are in-progress
// and jobs that have recently finished. The List Job operation returns a list
// of these jobs sorted by job initiation time.
//
// Amazon Glacier retains recently completed jobs for a period before deleting
// them; however, it eventually removes completed jobs. The output of completed
// jobs can be retrieved. Retaining completed jobs for a period of time after
// they have completed enables you to get a job output in the event you miss
// the job completion notification or your first attempt to download it fails.
// For example, suppose you start an archive retrieval job to download an archive.
// After the job completes, you start to download the archive but encounter
// a network error. In this scenario, you can retry and download the archive
// while the job exists.
//
// The List Jobs operation supports pagination. You should always check the
// response Marker field. If there are no more jobs to list, the Marker field
// is set to null. If there are more jobs to list, the Marker field is set to
// a non-null value, which you can use to continue the pagination of the list.
// To return a list of jobs that begins at a specific job, set the marker request
// parameter to the Marker value for that job that you obtained from a previous
// List Jobs request.
//
// You can set a maximum limit for the number of jobs returned in the response
// by specifying the limit parameter in the request. The default limit is 50.
// The number of jobs returned might be fewer than the limit, but the number
// of returned jobs never exceeds the limit.
//
// Additionally, you can filter the jobs list returned by specifying the optional
// statuscode parameter or completed parameter, or both. Using the statuscode
// parameter, you can specify to return only jobs that match either the InProgress,
// Succeeded, or Failed status. Using the completed parameter, you can specify
// to return only jobs that were completed (true) or jobs that were not completed
// (false).
//
// For more information about using this operation, see the documentation for
// the underlying REST API List Jobs (http://docs.aws.amazon.com/amazonglacier/latest/dev/api-jobs-get.html).
//
//    // Example sending a request using ListJobsRequest.
//    req := client.ListJobsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListJobsRequest(input *ListJobsInput) ListJobsRequest {
	op := &aws.Operation{
		Name:       opListJobs,
		HTTPMethod: "GET",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/jobs",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListJobsInput{}
	}

	req := c.newRequest(op, input, &ListJobsOutput{})
	return ListJobsRequest{Request: req, Input: input, Copy: c.ListJobsRequest}
}

// ListJobsRequest is the request type for the
// ListJobs API operation.
type ListJobsRequest struct {
	*aws.Request
	Input *ListJobsInput
	Copy  func(*ListJobsInput) ListJobsRequest
}

// Send marshals and sends the ListJobs API request.
func (r ListJobsRequest) Send(ctx context.Context) (*ListJobsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListJobsResponse{
		ListJobsOutput: r.Request.Data.(*ListJobsOutput),
		response:       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListJobsRequestPaginator returns a paginator for ListJobs.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListJobsRequest(input)
//   p := glacier.NewListJobsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListJobsPaginator(req ListJobsRequest) ListJobsPaginator {
	return ListJobsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListJobsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListJobsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListJobsPaginator struct {
	aws.Pager
}

func (p *ListJobsPaginator) CurrentPage() *ListJobsOutput {
	return p.Pager.CurrentPage().(*ListJobsOutput)
}

// ListJobsResponse is the response type for the
// ListJobs API operation.
type ListJobsResponse struct {
	*ListJobsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListJobs request.
func (r *ListJobsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
