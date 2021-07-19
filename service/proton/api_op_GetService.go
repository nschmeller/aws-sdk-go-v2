// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"errors"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Get detail data for a service.
func (c *Client) GetService(ctx context.Context, params *GetServiceInput, optFns ...func(*Options)) (*GetServiceOutput, error) {
	if params == nil {
		params = &GetServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetService", params, optFns, c.addOperationGetServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceInput struct {

	// The name of the service that you want to get the detail data for.
	//
	// This member is required.
	Name *string
}

type GetServiceOutput struct {

	// The service detail data that's returned by AWS Proton.
	Service *types.Service

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationGetServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetService{}, middleware.After)
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
	if err = addOpGetServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetService(options.Region), middleware.Before); err != nil {
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

// GetServiceAPIClient is a client that implements the GetService operation.
type GetServiceAPIClient interface {
	GetService(context.Context, *GetServiceInput, ...func(*Options)) (*GetServiceOutput, error)
}

var _ GetServiceAPIClient = (*Client)(nil)

// ServiceCreatedWaiterOptions are waiter options for ServiceCreatedWaiter
type ServiceCreatedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ServiceCreatedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ServiceCreatedWaiter will use default max delay of 4999 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetServiceInput, *GetServiceOutput, error) (bool, error)
}

// ServiceCreatedWaiter defines the waiters for ServiceCreated
type ServiceCreatedWaiter struct {
	client GetServiceAPIClient

	options ServiceCreatedWaiterOptions
}

// NewServiceCreatedWaiter constructs a ServiceCreatedWaiter.
func NewServiceCreatedWaiter(client GetServiceAPIClient, optFns ...func(*ServiceCreatedWaiterOptions)) *ServiceCreatedWaiter {
	options := ServiceCreatedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 4999 * time.Second
	options.Retryable = serviceCreatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ServiceCreatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ServiceCreated waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ServiceCreatedWaiter) Wait(ctx context.Context, params *GetServiceInput, maxWaitDur time.Duration, optFns ...func(*ServiceCreatedWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 4999 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetService(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for ServiceCreated waiter")
}

func serviceCreatedStateRetryable(ctx context.Context, input *GetServiceInput, output *GetServiceOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("service.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ACTIVE"
		value, ok := pathValue.(types.ServiceStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ServiceStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("service.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATE_FAILED_CLEANUP_COMPLETE"
		value, ok := pathValue.(types.ServiceStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ServiceStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("service.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATE_FAILED_CLEANUP_FAILED"
		value, ok := pathValue.(types.ServiceStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ServiceStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("service.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "CREATE_FAILED"
		value, ok := pathValue.(types.ServiceStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ServiceStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// ServiceUpdatedWaiterOptions are waiter options for ServiceUpdatedWaiter
type ServiceUpdatedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ServiceUpdatedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ServiceUpdatedWaiter will use default max delay of 4999 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetServiceInput, *GetServiceOutput, error) (bool, error)
}

// ServiceUpdatedWaiter defines the waiters for ServiceUpdated
type ServiceUpdatedWaiter struct {
	client GetServiceAPIClient

	options ServiceUpdatedWaiterOptions
}

// NewServiceUpdatedWaiter constructs a ServiceUpdatedWaiter.
func NewServiceUpdatedWaiter(client GetServiceAPIClient, optFns ...func(*ServiceUpdatedWaiterOptions)) *ServiceUpdatedWaiter {
	options := ServiceUpdatedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 4999 * time.Second
	options.Retryable = serviceUpdatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ServiceUpdatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ServiceUpdated waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ServiceUpdatedWaiter) Wait(ctx context.Context, params *GetServiceInput, maxWaitDur time.Duration, optFns ...func(*ServiceUpdatedWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 4999 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetService(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for ServiceUpdated waiter")
}

func serviceUpdatedStateRetryable(ctx context.Context, input *GetServiceInput, output *GetServiceOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("service.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "ACTIVE"
		value, ok := pathValue.(types.ServiceStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ServiceStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("service.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "UPDATE_FAILED_CLEANUP_COMPLETE"
		value, ok := pathValue.(types.ServiceStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ServiceStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("service.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "UPDATE_FAILED_CLEANUP_FAILED"
		value, ok := pathValue.(types.ServiceStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ServiceStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("service.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "UPDATE_FAILED"
		value, ok := pathValue.(types.ServiceStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ServiceStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("service.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "UPDATE_COMPLETE_CLEANUP_FAILED"
		value, ok := pathValue.(types.ServiceStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ServiceStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// ServiceDeletedWaiterOptions are waiter options for ServiceDeletedWaiter
type ServiceDeletedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ServiceDeletedWaiter will use default minimum delay of 5 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ServiceDeletedWaiter will use default max delay of 4999 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetServiceInput, *GetServiceOutput, error) (bool, error)
}

// ServiceDeletedWaiter defines the waiters for ServiceDeleted
type ServiceDeletedWaiter struct {
	client GetServiceAPIClient

	options ServiceDeletedWaiterOptions
}

// NewServiceDeletedWaiter constructs a ServiceDeletedWaiter.
func NewServiceDeletedWaiter(client GetServiceAPIClient, optFns ...func(*ServiceDeletedWaiterOptions)) *ServiceDeletedWaiter {
	options := ServiceDeletedWaiterOptions{}
	options.MinDelay = 5 * time.Second
	options.MaxDelay = 4999 * time.Second
	options.Retryable = serviceDeletedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ServiceDeletedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ServiceDeleted waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ServiceDeletedWaiter) Wait(ctx context.Context, params *GetServiceInput, maxWaitDur time.Duration, optFns ...func(*ServiceDeletedWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 4999 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetService(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for ServiceDeleted waiter")
}

func serviceDeletedStateRetryable(ctx context.Context, input *GetServiceInput, output *GetServiceOutput, err error) (bool, error) {

	if err != nil {
		var errorType *types.ResourceNotFoundException
		if errors.As(err, &errorType) {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("service.status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "DELETE_FAILED"
		value, ok := pathValue.(types.ServiceStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.ServiceStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

// ServicePipelineDeployedWaiterOptions are waiter options for
// ServicePipelineDeployedWaiter
type ServicePipelineDeployedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ServicePipelineDeployedWaiter will use default minimum delay of 10 seconds. Note
	// that MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ServicePipelineDeployedWaiter will use default max delay of 3600
	// seconds. Note that MaxDelay must resolve to value greater than or equal to the
	// MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *GetServiceInput, *GetServiceOutput, error) (bool, error)
}

// ServicePipelineDeployedWaiter defines the waiters for ServicePipelineDeployed
type ServicePipelineDeployedWaiter struct {
	client GetServiceAPIClient

	options ServicePipelineDeployedWaiterOptions
}

// NewServicePipelineDeployedWaiter constructs a ServicePipelineDeployedWaiter.
func NewServicePipelineDeployedWaiter(client GetServiceAPIClient, optFns ...func(*ServicePipelineDeployedWaiterOptions)) *ServicePipelineDeployedWaiter {
	options := ServicePipelineDeployedWaiterOptions{}
	options.MinDelay = 10 * time.Second
	options.MaxDelay = 3600 * time.Second
	options.Retryable = servicePipelineDeployedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ServicePipelineDeployedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ServicePipelineDeployed waiter. The
// maxWaitDur is the maximum wait duration the waiter will wait. The maxWaitDur is
// required and must be greater than zero.
func (w *ServicePipelineDeployedWaiter) Wait(ctx context.Context, params *GetServiceInput, maxWaitDur time.Duration, optFns ...func(*ServicePipelineDeployedWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 3600 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.GetService(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for ServicePipelineDeployed waiter")
}

func servicePipelineDeployedStateRetryable(ctx context.Context, input *GetServiceInput, output *GetServiceOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("service.pipeline.deploymentStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "SUCCEEDED"
		value, ok := pathValue.(types.DeploymentStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.DeploymentStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("service.pipeline.deploymentStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "FAILED"
		value, ok := pathValue.(types.DeploymentStatus)
		if !ok {
			return false, fmt.Errorf("waiter comparator expected types.DeploymentStatus value, got %T", pathValue)
		}

		if string(value) == expectedValue {
			return false, fmt.Errorf("waiter state transitioned to Failure")
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opGetService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "GetService",
	}
}
