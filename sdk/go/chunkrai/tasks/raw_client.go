package tasks

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

func (r *RawClient) GetTasksRoute(
	ctx context.Context,
	request *chunkrai.GetTasksRouteRequest,
	opts ...option.RequestOption,
) (*core.Response[[]*chunkrai.TaskResponse], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.chunkr.ai",
	)
	endpointURL := baseURL + "/api/v1/tasks"
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
	var response []*chunkrai.TaskResponse
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
	return &core.Response[[]*chunkrai.TaskResponse]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response,
	}, nil
}
