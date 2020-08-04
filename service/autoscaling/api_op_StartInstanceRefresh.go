// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartInstanceRefreshInput struct {
	_ struct{} `type:"structure"`

	// The name of the Auto Scaling group.
	//
	// AutoScalingGroupName is a required field
	AutoScalingGroupName *string `min:"1" type:"string" required:"true"`

	// Set of preferences associated with the instance refresh request.
	//
	// If not provided, the default values are used. For MinHealthyPercentage, the
	// default value is 90. For InstanceWarmup, the default is to use the value
	// specified for the health check grace period for the Auto Scaling group.
	//
	// For more information, see RefreshPreferences (https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_RefreshPreferences.html)
	// in the Amazon EC2 Auto Scaling API Reference.
	Preferences *RefreshPreferences `type:"structure"`

	// The strategy to use for the instance refresh. The only valid value is Rolling.
	//
	// A rolling update is an update that is applied to all instances in an Auto
	// Scaling group until all instances have been updated. A rolling update can
	// fail due to failed health checks or if instances are on standby or are protected
	// from scale in. If the rolling update process fails, any instances that were
	// already replaced are not rolled back to their previous configuration.
	Strategy RefreshStrategy `type:"string" enum:"true"`
}

// String returns the string representation
func (s StartInstanceRefreshInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartInstanceRefreshInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartInstanceRefreshInput"}

	if s.AutoScalingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AutoScalingGroupName"))
	}
	if s.AutoScalingGroupName != nil && len(*s.AutoScalingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AutoScalingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartInstanceRefreshOutput struct {
	_ struct{} `type:"structure"`

	// A unique ID for tracking the progress of the request.
	InstanceRefreshId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s StartInstanceRefreshOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartInstanceRefresh = "StartInstanceRefresh"

// StartInstanceRefreshRequest returns a request value for making API operation for
// Auto Scaling.
//
// Starts a new instance refresh operation, which triggers a rolling replacement
// of all previously launched instances in the Auto Scaling group with a new
// group of instances.
//
// If successful, this call creates a new instance refresh request with a unique
// ID that you can use to track its progress. To query its status, call the
// DescribeInstanceRefreshes API. To describe the instance refreshes that have
// already run, call the DescribeInstanceRefreshes API. To cancel an instance
// refresh operation in progress, use the CancelInstanceRefresh API.
//
// For more information, see Replacing Auto Scaling Instances Based on an Instance
// Refresh (https://docs.aws.amazon.com/autoscaling/ec2/userguide/asg-instance-refresh.html).
//
//    // Example sending a request using StartInstanceRefreshRequest.
//    req := client.StartInstanceRefreshRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/autoscaling-2011-01-01/StartInstanceRefresh
func (c *Client) StartInstanceRefreshRequest(input *StartInstanceRefreshInput) StartInstanceRefreshRequest {
	op := &aws.Operation{
		Name:       opStartInstanceRefresh,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartInstanceRefreshInput{}
	}

	req := c.newRequest(op, input, &StartInstanceRefreshOutput{})

	return StartInstanceRefreshRequest{Request: req, Input: input, Copy: c.StartInstanceRefreshRequest}
}

// StartInstanceRefreshRequest is the request type for the
// StartInstanceRefresh API operation.
type StartInstanceRefreshRequest struct {
	*aws.Request
	Input *StartInstanceRefreshInput
	Copy  func(*StartInstanceRefreshInput) StartInstanceRefreshRequest
}

// Send marshals and sends the StartInstanceRefresh API request.
func (r StartInstanceRefreshRequest) Send(ctx context.Context) (*StartInstanceRefreshResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartInstanceRefreshResponse{
		StartInstanceRefreshOutput: r.Request.Data.(*StartInstanceRefreshOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartInstanceRefreshResponse is the response type for the
// StartInstanceRefresh API operation.
type StartInstanceRefreshResponse struct {
	*StartInstanceRefreshOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartInstanceRefresh request.
func (r *StartInstanceRefreshResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}