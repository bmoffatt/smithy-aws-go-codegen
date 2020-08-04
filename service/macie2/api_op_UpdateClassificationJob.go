// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package macie2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Cancels a classification job.
type UpdateClassificationJobInput struct {
	_ struct{} `type:"structure"`

	// JobId is a required field
	JobId *string `location:"uri" locationName:"jobId" type:"string" required:"true"`

	// The current status of a classification job. Possible values are:
	//
	// JobStatus is a required field
	JobStatus JobStatus `locationName:"jobStatus" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s UpdateClassificationJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateClassificationJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateClassificationJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}
	if len(s.JobStatus) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("JobStatus"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateClassificationJobInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if len(s.JobStatus) > 0 {
		v := s.JobStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "jobStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.JobId != nil {
		v := *s.JobId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "jobId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type UpdateClassificationJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateClassificationJobOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateClassificationJobOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUpdateClassificationJob = "UpdateClassificationJob"

// UpdateClassificationJobRequest returns a request value for making API operation for
// Amazon Macie 2.
//
// Cancels a classification job.
//
//    // Example sending a request using UpdateClassificationJobRequest.
//    req := client.UpdateClassificationJobRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/macie2-2020-01-01/UpdateClassificationJob
func (c *Client) UpdateClassificationJobRequest(input *UpdateClassificationJobInput) UpdateClassificationJobRequest {
	op := &aws.Operation{
		Name:       opUpdateClassificationJob,
		HTTPMethod: "PATCH",
		HTTPPath:   "/jobs/{jobId}",
	}

	if input == nil {
		input = &UpdateClassificationJobInput{}
	}

	req := c.newRequest(op, input, &UpdateClassificationJobOutput{})

	return UpdateClassificationJobRequest{Request: req, Input: input, Copy: c.UpdateClassificationJobRequest}
}

// UpdateClassificationJobRequest is the request type for the
// UpdateClassificationJob API operation.
type UpdateClassificationJobRequest struct {
	*aws.Request
	Input *UpdateClassificationJobInput
	Copy  func(*UpdateClassificationJobInput) UpdateClassificationJobRequest
}

// Send marshals and sends the UpdateClassificationJob API request.
func (r UpdateClassificationJobRequest) Send(ctx context.Context) (*UpdateClassificationJobResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateClassificationJobResponse{
		UpdateClassificationJobOutput: r.Request.Data.(*UpdateClassificationJobOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateClassificationJobResponse is the response type for the
// UpdateClassificationJob API operation.
type UpdateClassificationJobResponse struct {
	*UpdateClassificationJobOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateClassificationJob request.
func (r *UpdateClassificationJobResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}