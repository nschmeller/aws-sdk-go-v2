// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes permissions from a profiling group's resource-based policy that are
// provided using an action group. The one supported action group that can be
// removed is agentPermission which grants ConfigureAgent and PostAgent
// permissions. For more information, see Resource-based policies in CodeGuru
// Profiler
// (https://docs.aws.amazon.com/codeguru/latest/profiler-ug/resource-based-policies.html)
// in the Amazon CodeGuru Profiler User Guide, ConfigureAgent
// (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_ConfigureAgent.html),
// and PostAgentProfile
// (https://docs.aws.amazon.com/codeguru/latest/profiler-api/API_PostAgentProfile.html).
func (c *Client) RemovePermission(ctx context.Context, params *RemovePermissionInput, optFns ...func(*Options)) (*RemovePermissionOutput, error) {
	if params == nil {
		params = &RemovePermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemovePermission", params, optFns, c.addOperationRemovePermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RemovePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the removePermissionRequest.
type RemovePermissionInput struct {

	// Specifies an action group that contains the permissions to remove from a
	// profiling group's resource-based policy. One action group is supported,
	// agentPermissions, which grants ConfigureAgent and PostAgentProfile permissions.
	//
	// This member is required.
	ActionGroup types.ActionGroup

	// The name of the profiling group.
	//
	// This member is required.
	ProfilingGroupName *string

	// A universally unique identifier (UUID) for the revision of the resource-based
	// policy from which you want to remove permissions.
	//
	// This member is required.
	RevisionId *string
}

// The structure representing the removePermissionResponse.
type RemovePermissionOutput struct {

	// The JSON-formatted resource-based policy on the profiling group after the
	// specified permissions were removed.
	//
	// This member is required.
	Policy *string

	// A universally unique identifier (UUID) for the revision of the resource-based
	// policy after the specified permissions were removed. The updated JSON-formatted
	// policy is in the policy element of the response.
	//
	// This member is required.
	RevisionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationRemovePermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRemovePermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRemovePermission{}, middleware.After)
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
	if err = addOpRemovePermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRemovePermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRemovePermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-profiler",
		OperationName: "RemovePermission",
	}
}
