package health

import (
	bytes "bytes"
	context "context"
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

func (r *RawClient) Check(
	ctx context.Context,
	opts ...option.RequestOption,
) (*core.Response[string], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		r.baseURL,
		"https://api.chunkr.ai",
	)
	endpointURL := baseURL + "/health"
	headers := internal.MergeHeaders(
		r.header.Clone(),
		options.ToHeader(),
	)
	response := bytes.NewBuffer(nil)
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
			Response:        response,
		},
	)
	if err != nil {
		return nil, err
	}
	return &core.Response[string]{
		StatusCode: raw.StatusCode,
		Header:     raw.Header,
		Body:       response.String(),
	}, nil
}
