module github.com/bmoffatt/smithy-aws-go-codegen/hack/endpoints/v2

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.18.0
	github.com/aws/smithy-go v1.13.5
	github.com/google/go-cmp v0.5.8
)

replace github.com/aws/aws-sdk-go-v2 => ../../../
