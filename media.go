// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/relayapi-dev/relay-go/internal/apiform"
	"github.com/relayapi-dev/relay-go/internal/apijson"
	"github.com/relayapi-dev/relay-go/internal/apiquery"
	"github.com/relayapi-dev/relay-go/internal/param"
	"github.com/relayapi-dev/relay-go/internal/requestconfig"
	"github.com/relayapi-dev/relay-go/option"
)

// MediaService contains methods and other services that help with interacting with
// the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaService] method instead.
type MediaService struct {
	Options []option.RequestOption
}

// NewMediaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMediaService(opts ...option.RequestOption) (r *MediaService) {
	r = &MediaService{}
	r.Options = opts
	return
}

// Get media details
func (r *MediaService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MediaGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/media/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete media
func (r *MediaService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("v1/media/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Generate a pre-signed URL for direct upload to R2. The client can PUT the file
// to the returned URL.
func (r *MediaService) GetPresignURL(ctx context.Context, body MediaGetPresignURLParams, opts ...option.RequestOption) (res *MediaGetPresignURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/media/presign"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Upload a raw file body. Pass the filename as a query parameter and set the
// Content-Type header.
func (r *MediaService) Upload(ctx context.Context, body io.Reader, params MediaUploadParams, opts ...option.RequestOption) (res *MediaUploadResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithRequestBody("application/octet-stream", body)}, opts...)
	path := "v1/media/upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type MediaGetResponse struct {
	// Media ID
	ID string `json:"id" api:"required"`
	// Upload timestamp
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Original filename
	Filename string `json:"filename" api:"required"`
	// MIME type
	MimeType string `json:"mime_type" api:"required"`
	// File size in bytes
	Size int64 `json:"size" api:"required"`
	// Public URL
	URL string `json:"url" api:"required,nullable" format:"uri"`
	// Duration in seconds (video/audio)
	Duration int64 `json:"duration" api:"nullable"`
	// Height in pixels
	Height int64 `json:"height" api:"nullable"`
	// Width in pixels
	Width int64                `json:"width" api:"nullable"`
	JSON  mediaGetResponseJSON `json:"-"`
}

// mediaGetResponseJSON contains the JSON metadata for the struct
// [MediaGetResponse]
type mediaGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Filename    apijson.Field
	MimeType    apijson.Field
	Size        apijson.Field
	URL         apijson.Field
	Duration    apijson.Field
	Height      apijson.Field
	Width       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MediaGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mediaGetResponseJSON) RawJSON() string {
	return r.raw
}

type MediaGetPresignURLResponse struct {
	// Seconds until the upload URL expires
	ExpiresIn int64 `json:"expires_in" api:"required"`
	// Pre-signed PUT URL for uploading
	UploadURL string `json:"upload_url" api:"required" format:"uri"`
	// Public URL after upload completes
	URL  string                         `json:"url" api:"required" format:"uri"`
	JSON mediaGetPresignURLResponseJSON `json:"-"`
}

// mediaGetPresignURLResponseJSON contains the JSON metadata for the struct
// [MediaGetPresignURLResponse]
type mediaGetPresignURLResponseJSON struct {
	ExpiresIn   apijson.Field
	UploadURL   apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MediaGetPresignURLResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mediaGetPresignURLResponseJSON) RawJSON() string {
	return r.raw
}

type MediaUploadResponse struct {
	// Original filename
	Filename string `json:"filename" api:"required"`
	// File size in bytes
	Size int64 `json:"size" api:"required"`
	// MIME type of the uploaded file
	Type string `json:"type" api:"required"`
	// Public URL of the uploaded file
	URL  string                  `json:"url" api:"required" format:"uri"`
	JSON mediaUploadResponseJSON `json:"-"`
}

// mediaUploadResponseJSON contains the JSON metadata for the struct
// [MediaUploadResponse]
type mediaUploadResponseJSON struct {
	Filename    apijson.Field
	Size        apijson.Field
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MediaUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mediaUploadResponseJSON) RawJSON() string {
	return r.raw
}

type MediaGetPresignURLParams struct {
	// MIME type of the file to upload
	ContentType param.Field[string] `json:"content_type" api:"required"`
	// Desired filename
	Filename param.Field[string] `json:"filename" api:"required"`
}

func (r MediaGetPresignURLParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MediaUploadParams struct {
	// Original filename
	Filename param.Field[string] `query:"filename" api:"required"`
}

func (r MediaUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [MediaUploadParams]'s query parameters as `url.Values`.
func (r MediaUploadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
