package task

import (
	context "context"
	chunkrai "github.com/gitslim/chunkr-ai-sdk/sdk/go/chunkrai"
	core "github.com/gitslim/chunkr-ai-sdk/sdk/go/chunkrai/core"
	internal "github.com/gitslim/chunkr-ai-sdk/sdk/go/chunkrai/internal"
	option "github.com/gitslim/chunkr-ai-sdk/sdk/go/chunkrai/option"
	http "net/http"
)

type RawClient struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewRawClient(options *core.RequestOptions) *RawClient {
	return &RawClient{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

func (r *RawClient) CreateTaskRoute(
	ctx context.Context,
	request *chunkrai.CreateForm,
	opts ...option.RequestOption,
) (*core.Response[*chunkrai.TaskResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.chunkr.ai",
	)
	endpointURL := baseURL + "/api/v1/task/parse"
	headers := internal.MergeHeaders(
		r.header.Clone(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		500: func(apiError *core.APIError) error {
			return &chunkrai.InternalServerError{
				APIError: apiError,
			}
		},
	}
	var response *chunkrai.TaskResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*chunkrai.TaskResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}

func (r *RawClient) GetTaskRoute(
	ctx context.Context,
	// Id of the task to retrieve
	taskId *string,
	request *chunkrai.GetTaskRouteRequest,
	opts ...option.RequestOption,
) (*core.Response[*chunkrai.TaskResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.chunkr.ai",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/task/%v",
		taskId,
	)
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		r.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		500: func(apiError *core.APIError) error {
			return &chunkrai.InternalServerError{
				APIError: apiError,
			}
		},
	}
	var response *chunkrai.TaskResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*chunkrai.TaskResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}

func (r *RawClient) DeleteTaskRoute(
	ctx context.Context,
	// Id of the task to delete
	taskId *string,
	opts ...option.RequestOption,
) (*core.Response[any], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.chunkr.ai",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/task/%v",
		taskId,
	)
	headers := internal.MergeHeaders(
		r.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		500: func(apiError *core.APIError) error {
			return &chunkrai.InternalServerError{
				APIError: apiError,
			}
		},
	}
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodDelete,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[any]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       nil,
	}, nil
}

func (r *RawClient) CancelTaskRoute(
	ctx context.Context,
	// Id of the task to cancel
	taskId *string,
	opts ...option.RequestOption,
) (*core.Response[any], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.chunkr.ai",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/task/%v/cancel",
		taskId,
	)
	headers := internal.MergeHeaders(
		r.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		500: func(apiError *core.APIError) error {
			return &chunkrai.InternalServerError{
				APIError: apiError,
			}
		},
	}
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[any]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       nil,
	}, nil
}

func (r *RawClient) UpdateTaskRoute(
	ctx context.Context,
	taskId string,
	request *chunkrai.UpdateForm,
	opts ...option.RequestOption,
) (*core.Response[*chunkrai.TaskResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.chunkr.ai",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/api/v1/task/%v/parse",
		taskId,
	)
	headers := internal.MergeHeaders(
		r.header.Clone(),
		options.ToHeader(),
	)
	headers.Add("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		500: func(apiError *core.APIError) error {
			return &chunkrai.InternalServerError{
				APIError: apiError,
			}
		},
	}
	var response *chunkrai.TaskResponse
	raw, err := r.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPatch,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[*chunkrai.TaskResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}
