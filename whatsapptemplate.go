// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaygo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/relayapi-dev/relay-go/internal/apijson"
	"github.com/relayapi-dev/relay-go/internal/apiquery"
	"github.com/relayapi-dev/relay-go/internal/param"
	"github.com/relayapi-dev/relay-go/internal/requestconfig"
	"github.com/relayapi-dev/relay-go/option"
)

// WhatsappTemplateService contains methods and other services that help with
// interacting with the relay API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappTemplateService] method instead.
type WhatsappTemplateService struct {
	Options []option.RequestOption
}

// NewWhatsappTemplateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWhatsappTemplateService(opts ...option.RequestOption) (r *WhatsappTemplateService) {
	r = &WhatsappTemplateService{}
	r.Options = opts
	return
}

// Create a message template
func (r *WhatsappTemplateService) New(ctx context.Context, body WhatsappTemplateNewParams, opts ...option.RequestOption) (res *WhatsappTemplateNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whatsapp/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get template details
func (r *WhatsappTemplateService) Get(ctx context.Context, templateName string, query WhatsappTemplateGetParams, opts ...option.RequestOption) (res *WhatsappTemplateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateName == "" {
		err = errors.New("missing required template_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/whatsapp/templates/%s", templateName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List message templates
func (r *WhatsappTemplateService) List(ctx context.Context, query WhatsappTemplateListParams, opts ...option.RequestOption) (res *WhatsappTemplateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/whatsapp/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a message template
func (r *WhatsappTemplateService) Delete(ctx context.Context, templateName string, body WhatsappTemplateDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if templateName == "" {
		err = errors.New("missing required template_name parameter")
		return err
	}
	path := fmt.Sprintf("v1/whatsapp/templates/%s", templateName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

type WhatsappTemplateNewResponse struct {
	// Template category
	Category   WhatsappTemplateNewResponseCategory    `json:"category" api:"required"`
	Components []WhatsappTemplateNewResponseComponent `json:"components" api:"required"`
	// Template language code
	Language string `json:"language" api:"required"`
	// Template name
	Name string `json:"name" api:"required"`
	// Approval status
	Status WhatsappTemplateNewResponseStatus `json:"status" api:"required"`
	JSON   whatsappTemplateNewResponseJSON   `json:"-"`
}

// whatsappTemplateNewResponseJSON contains the JSON metadata for the struct
// [WhatsappTemplateNewResponse]
type whatsappTemplateNewResponseJSON struct {
	Category    apijson.Field
	Components  apijson.Field
	Language    apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappTemplateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappTemplateNewResponseJSON) RawJSON() string {
	return r.raw
}

// Template category
type WhatsappTemplateNewResponseCategory string

const (
	WhatsappTemplateNewResponseCategoryMarketing      WhatsappTemplateNewResponseCategory = "MARKETING"
	WhatsappTemplateNewResponseCategoryUtility        WhatsappTemplateNewResponseCategory = "UTILITY"
	WhatsappTemplateNewResponseCategoryAuthentication WhatsappTemplateNewResponseCategory = "AUTHENTICATION"
)

func (r WhatsappTemplateNewResponseCategory) IsKnown() bool {
	switch r {
	case WhatsappTemplateNewResponseCategoryMarketing, WhatsappTemplateNewResponseCategoryUtility, WhatsappTemplateNewResponseCategoryAuthentication:
		return true
	}
	return false
}

type WhatsappTemplateNewResponseComponent struct {
	// Component type
	Type    WhatsappTemplateNewResponseComponentsType     `json:"type" api:"required"`
	Buttons []WhatsappTemplateNewResponseComponentsButton `json:"buttons"`
	// Header format (TEXT, IMAGE, etc.)
	Format string `json:"format"`
	// Component text
	Text string                                   `json:"text"`
	JSON whatsappTemplateNewResponseComponentJSON `json:"-"`
}

// whatsappTemplateNewResponseComponentJSON contains the JSON metadata for the
// struct [WhatsappTemplateNewResponseComponent]
type whatsappTemplateNewResponseComponentJSON struct {
	Type        apijson.Field
	Buttons     apijson.Field
	Format      apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappTemplateNewResponseComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappTemplateNewResponseComponentJSON) RawJSON() string {
	return r.raw
}

// Component type
type WhatsappTemplateNewResponseComponentsType string

const (
	WhatsappTemplateNewResponseComponentsTypeHeader  WhatsappTemplateNewResponseComponentsType = "HEADER"
	WhatsappTemplateNewResponseComponentsTypeBody    WhatsappTemplateNewResponseComponentsType = "BODY"
	WhatsappTemplateNewResponseComponentsTypeFooter  WhatsappTemplateNewResponseComponentsType = "FOOTER"
	WhatsappTemplateNewResponseComponentsTypeButtons WhatsappTemplateNewResponseComponentsType = "BUTTONS"
)

func (r WhatsappTemplateNewResponseComponentsType) IsKnown() bool {
	switch r {
	case WhatsappTemplateNewResponseComponentsTypeHeader, WhatsappTemplateNewResponseComponentsTypeBody, WhatsappTemplateNewResponseComponentsTypeFooter, WhatsappTemplateNewResponseComponentsTypeButtons:
		return true
	}
	return false
}

type WhatsappTemplateNewResponseComponentsButton struct {
	// Button text
	Text string `json:"text" api:"required"`
	// Button type
	Type        string                                          `json:"type" api:"required"`
	PhoneNumber string                                          `json:"phone_number"`
	URL         string                                          `json:"url"`
	JSON        whatsappTemplateNewResponseComponentsButtonJSON `json:"-"`
}

// whatsappTemplateNewResponseComponentsButtonJSON contains the JSON metadata for
// the struct [WhatsappTemplateNewResponseComponentsButton]
type whatsappTemplateNewResponseComponentsButtonJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	PhoneNumber apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappTemplateNewResponseComponentsButton) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappTemplateNewResponseComponentsButtonJSON) RawJSON() string {
	return r.raw
}

// Approval status
type WhatsappTemplateNewResponseStatus string

const (
	WhatsappTemplateNewResponseStatusApproved WhatsappTemplateNewResponseStatus = "APPROVED"
	WhatsappTemplateNewResponseStatusPending  WhatsappTemplateNewResponseStatus = "PENDING"
	WhatsappTemplateNewResponseStatusRejected WhatsappTemplateNewResponseStatus = "REJECTED"
)

func (r WhatsappTemplateNewResponseStatus) IsKnown() bool {
	switch r {
	case WhatsappTemplateNewResponseStatusApproved, WhatsappTemplateNewResponseStatusPending, WhatsappTemplateNewResponseStatusRejected:
		return true
	}
	return false
}

type WhatsappTemplateGetResponse struct {
	// Template category
	Category   WhatsappTemplateGetResponseCategory    `json:"category" api:"required"`
	Components []WhatsappTemplateGetResponseComponent `json:"components" api:"required"`
	// Template language code
	Language string `json:"language" api:"required"`
	// Template name
	Name string `json:"name" api:"required"`
	// Approval status
	Status WhatsappTemplateGetResponseStatus `json:"status" api:"required"`
	JSON   whatsappTemplateGetResponseJSON   `json:"-"`
}

// whatsappTemplateGetResponseJSON contains the JSON metadata for the struct
// [WhatsappTemplateGetResponse]
type whatsappTemplateGetResponseJSON struct {
	Category    apijson.Field
	Components  apijson.Field
	Language    apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappTemplateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappTemplateGetResponseJSON) RawJSON() string {
	return r.raw
}

// Template category
type WhatsappTemplateGetResponseCategory string

const (
	WhatsappTemplateGetResponseCategoryMarketing      WhatsappTemplateGetResponseCategory = "MARKETING"
	WhatsappTemplateGetResponseCategoryUtility        WhatsappTemplateGetResponseCategory = "UTILITY"
	WhatsappTemplateGetResponseCategoryAuthentication WhatsappTemplateGetResponseCategory = "AUTHENTICATION"
)

func (r WhatsappTemplateGetResponseCategory) IsKnown() bool {
	switch r {
	case WhatsappTemplateGetResponseCategoryMarketing, WhatsappTemplateGetResponseCategoryUtility, WhatsappTemplateGetResponseCategoryAuthentication:
		return true
	}
	return false
}

type WhatsappTemplateGetResponseComponent struct {
	// Component type
	Type    WhatsappTemplateGetResponseComponentsType     `json:"type" api:"required"`
	Buttons []WhatsappTemplateGetResponseComponentsButton `json:"buttons"`
	// Header format (TEXT, IMAGE, etc.)
	Format string `json:"format"`
	// Component text
	Text string                                   `json:"text"`
	JSON whatsappTemplateGetResponseComponentJSON `json:"-"`
}

// whatsappTemplateGetResponseComponentJSON contains the JSON metadata for the
// struct [WhatsappTemplateGetResponseComponent]
type whatsappTemplateGetResponseComponentJSON struct {
	Type        apijson.Field
	Buttons     apijson.Field
	Format      apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappTemplateGetResponseComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappTemplateGetResponseComponentJSON) RawJSON() string {
	return r.raw
}

// Component type
type WhatsappTemplateGetResponseComponentsType string

const (
	WhatsappTemplateGetResponseComponentsTypeHeader  WhatsappTemplateGetResponseComponentsType = "HEADER"
	WhatsappTemplateGetResponseComponentsTypeBody    WhatsappTemplateGetResponseComponentsType = "BODY"
	WhatsappTemplateGetResponseComponentsTypeFooter  WhatsappTemplateGetResponseComponentsType = "FOOTER"
	WhatsappTemplateGetResponseComponentsTypeButtons WhatsappTemplateGetResponseComponentsType = "BUTTONS"
)

func (r WhatsappTemplateGetResponseComponentsType) IsKnown() bool {
	switch r {
	case WhatsappTemplateGetResponseComponentsTypeHeader, WhatsappTemplateGetResponseComponentsTypeBody, WhatsappTemplateGetResponseComponentsTypeFooter, WhatsappTemplateGetResponseComponentsTypeButtons:
		return true
	}
	return false
}

type WhatsappTemplateGetResponseComponentsButton struct {
	// Button text
	Text string `json:"text" api:"required"`
	// Button type
	Type        string                                          `json:"type" api:"required"`
	PhoneNumber string                                          `json:"phone_number"`
	URL         string                                          `json:"url"`
	JSON        whatsappTemplateGetResponseComponentsButtonJSON `json:"-"`
}

// whatsappTemplateGetResponseComponentsButtonJSON contains the JSON metadata for
// the struct [WhatsappTemplateGetResponseComponentsButton]
type whatsappTemplateGetResponseComponentsButtonJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	PhoneNumber apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappTemplateGetResponseComponentsButton) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappTemplateGetResponseComponentsButtonJSON) RawJSON() string {
	return r.raw
}

// Approval status
type WhatsappTemplateGetResponseStatus string

const (
	WhatsappTemplateGetResponseStatusApproved WhatsappTemplateGetResponseStatus = "APPROVED"
	WhatsappTemplateGetResponseStatusPending  WhatsappTemplateGetResponseStatus = "PENDING"
	WhatsappTemplateGetResponseStatusRejected WhatsappTemplateGetResponseStatus = "REJECTED"
)

func (r WhatsappTemplateGetResponseStatus) IsKnown() bool {
	switch r {
	case WhatsappTemplateGetResponseStatusApproved, WhatsappTemplateGetResponseStatusPending, WhatsappTemplateGetResponseStatusRejected:
		return true
	}
	return false
}

type WhatsappTemplateListResponse struct {
	Data []WhatsappTemplateListResponseData `json:"data" api:"required"`
	JSON whatsappTemplateListResponseJSON   `json:"-"`
}

// whatsappTemplateListResponseJSON contains the JSON metadata for the struct
// [WhatsappTemplateListResponse]
type whatsappTemplateListResponseJSON struct {
	Data        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappTemplateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappTemplateListResponseJSON) RawJSON() string {
	return r.raw
}

type WhatsappTemplateListResponseData struct {
	// Template category
	Category   WhatsappTemplateListResponseDataCategory    `json:"category" api:"required"`
	Components []WhatsappTemplateListResponseDataComponent `json:"components" api:"required"`
	// Template language code
	Language string `json:"language" api:"required"`
	// Template name
	Name string `json:"name" api:"required"`
	// Approval status
	Status WhatsappTemplateListResponseDataStatus `json:"status" api:"required"`
	JSON   whatsappTemplateListResponseDataJSON   `json:"-"`
}

// whatsappTemplateListResponseDataJSON contains the JSON metadata for the struct
// [WhatsappTemplateListResponseData]
type whatsappTemplateListResponseDataJSON struct {
	Category    apijson.Field
	Components  apijson.Field
	Language    apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappTemplateListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappTemplateListResponseDataJSON) RawJSON() string {
	return r.raw
}

// Template category
type WhatsappTemplateListResponseDataCategory string

const (
	WhatsappTemplateListResponseDataCategoryMarketing      WhatsappTemplateListResponseDataCategory = "MARKETING"
	WhatsappTemplateListResponseDataCategoryUtility        WhatsappTemplateListResponseDataCategory = "UTILITY"
	WhatsappTemplateListResponseDataCategoryAuthentication WhatsappTemplateListResponseDataCategory = "AUTHENTICATION"
)

func (r WhatsappTemplateListResponseDataCategory) IsKnown() bool {
	switch r {
	case WhatsappTemplateListResponseDataCategoryMarketing, WhatsappTemplateListResponseDataCategoryUtility, WhatsappTemplateListResponseDataCategoryAuthentication:
		return true
	}
	return false
}

type WhatsappTemplateListResponseDataComponent struct {
	// Component type
	Type    WhatsappTemplateListResponseDataComponentsType     `json:"type" api:"required"`
	Buttons []WhatsappTemplateListResponseDataComponentsButton `json:"buttons"`
	// Header format (TEXT, IMAGE, etc.)
	Format string `json:"format"`
	// Component text
	Text string                                        `json:"text"`
	JSON whatsappTemplateListResponseDataComponentJSON `json:"-"`
}

// whatsappTemplateListResponseDataComponentJSON contains the JSON metadata for the
// struct [WhatsappTemplateListResponseDataComponent]
type whatsappTemplateListResponseDataComponentJSON struct {
	Type        apijson.Field
	Buttons     apijson.Field
	Format      apijson.Field
	Text        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappTemplateListResponseDataComponent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappTemplateListResponseDataComponentJSON) RawJSON() string {
	return r.raw
}

// Component type
type WhatsappTemplateListResponseDataComponentsType string

const (
	WhatsappTemplateListResponseDataComponentsTypeHeader  WhatsappTemplateListResponseDataComponentsType = "HEADER"
	WhatsappTemplateListResponseDataComponentsTypeBody    WhatsappTemplateListResponseDataComponentsType = "BODY"
	WhatsappTemplateListResponseDataComponentsTypeFooter  WhatsappTemplateListResponseDataComponentsType = "FOOTER"
	WhatsappTemplateListResponseDataComponentsTypeButtons WhatsappTemplateListResponseDataComponentsType = "BUTTONS"
)

func (r WhatsappTemplateListResponseDataComponentsType) IsKnown() bool {
	switch r {
	case WhatsappTemplateListResponseDataComponentsTypeHeader, WhatsappTemplateListResponseDataComponentsTypeBody, WhatsappTemplateListResponseDataComponentsTypeFooter, WhatsappTemplateListResponseDataComponentsTypeButtons:
		return true
	}
	return false
}

type WhatsappTemplateListResponseDataComponentsButton struct {
	// Button text
	Text string `json:"text" api:"required"`
	// Button type
	Type        string                                               `json:"type" api:"required"`
	PhoneNumber string                                               `json:"phone_number"`
	URL         string                                               `json:"url"`
	JSON        whatsappTemplateListResponseDataComponentsButtonJSON `json:"-"`
}

// whatsappTemplateListResponseDataComponentsButtonJSON contains the JSON metadata
// for the struct [WhatsappTemplateListResponseDataComponentsButton]
type whatsappTemplateListResponseDataComponentsButtonJSON struct {
	Text        apijson.Field
	Type        apijson.Field
	PhoneNumber apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhatsappTemplateListResponseDataComponentsButton) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whatsappTemplateListResponseDataComponentsButtonJSON) RawJSON() string {
	return r.raw
}

// Approval status
type WhatsappTemplateListResponseDataStatus string

const (
	WhatsappTemplateListResponseDataStatusApproved WhatsappTemplateListResponseDataStatus = "APPROVED"
	WhatsappTemplateListResponseDataStatusPending  WhatsappTemplateListResponseDataStatus = "PENDING"
	WhatsappTemplateListResponseDataStatusRejected WhatsappTemplateListResponseDataStatus = "REJECTED"
)

func (r WhatsappTemplateListResponseDataStatus) IsKnown() bool {
	switch r {
	case WhatsappTemplateListResponseDataStatusApproved, WhatsappTemplateListResponseDataStatusPending, WhatsappTemplateListResponseDataStatusRejected:
		return true
	}
	return false
}

type WhatsappTemplateNewParams struct {
	// WhatsApp account ID
	AccountID param.Field[string] `json:"account_id" api:"required"`
	// Template category
	Category param.Field[WhatsappTemplateNewParamsCategory] `json:"category" api:"required"`
	// Template components
	Components param.Field[[]WhatsappTemplateNewParamsComponent] `json:"components" api:"required"`
	// Template language code
	Language param.Field[string] `json:"language" api:"required"`
	// Template name
	Name param.Field[string] `json:"name" api:"required"`
}

func (r WhatsappTemplateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Template category
type WhatsappTemplateNewParamsCategory string

const (
	WhatsappTemplateNewParamsCategoryMarketing      WhatsappTemplateNewParamsCategory = "MARKETING"
	WhatsappTemplateNewParamsCategoryUtility        WhatsappTemplateNewParamsCategory = "UTILITY"
	WhatsappTemplateNewParamsCategoryAuthentication WhatsappTemplateNewParamsCategory = "AUTHENTICATION"
)

func (r WhatsappTemplateNewParamsCategory) IsKnown() bool {
	switch r {
	case WhatsappTemplateNewParamsCategoryMarketing, WhatsappTemplateNewParamsCategoryUtility, WhatsappTemplateNewParamsCategoryAuthentication:
		return true
	}
	return false
}

type WhatsappTemplateNewParamsComponent struct {
	// Component type
	Type    param.Field[WhatsappTemplateNewParamsComponentsType]     `json:"type" api:"required"`
	Buttons param.Field[[]WhatsappTemplateNewParamsComponentsButton] `json:"buttons"`
	// Header format (TEXT, IMAGE, etc.)
	Format param.Field[string] `json:"format"`
	// Component text
	Text param.Field[string] `json:"text"`
}

func (r WhatsappTemplateNewParamsComponent) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Component type
type WhatsappTemplateNewParamsComponentsType string

const (
	WhatsappTemplateNewParamsComponentsTypeHeader  WhatsappTemplateNewParamsComponentsType = "HEADER"
	WhatsappTemplateNewParamsComponentsTypeBody    WhatsappTemplateNewParamsComponentsType = "BODY"
	WhatsappTemplateNewParamsComponentsTypeFooter  WhatsappTemplateNewParamsComponentsType = "FOOTER"
	WhatsappTemplateNewParamsComponentsTypeButtons WhatsappTemplateNewParamsComponentsType = "BUTTONS"
)

func (r WhatsappTemplateNewParamsComponentsType) IsKnown() bool {
	switch r {
	case WhatsappTemplateNewParamsComponentsTypeHeader, WhatsappTemplateNewParamsComponentsTypeBody, WhatsappTemplateNewParamsComponentsTypeFooter, WhatsappTemplateNewParamsComponentsTypeButtons:
		return true
	}
	return false
}

type WhatsappTemplateNewParamsComponentsButton struct {
	// Button text
	Text param.Field[string] `json:"text" api:"required"`
	// Button type
	Type        param.Field[string] `json:"type" api:"required"`
	PhoneNumber param.Field[string] `json:"phone_number"`
	URL         param.Field[string] `json:"url"`
}

func (r WhatsappTemplateNewParamsComponentsButton) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WhatsappTemplateGetParams struct {
	// WhatsApp account ID
	AccountID param.Field[string] `query:"account_id" api:"required"`
}

// URLQuery serializes [WhatsappTemplateGetParams]'s query parameters as
// `url.Values`.
func (r WhatsappTemplateGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WhatsappTemplateListParams struct {
	// WhatsApp account ID
	AccountID param.Field[string] `query:"account_id" api:"required"`
}

// URLQuery serializes [WhatsappTemplateListParams]'s query parameters as
// `url.Values`.
func (r WhatsappTemplateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WhatsappTemplateDeleteParams struct {
	// WhatsApp account ID
	AccountID param.Field[string] `query:"account_id" api:"required"`
}

// URLQuery serializes [WhatsappTemplateDeleteParams]'s query parameters as
// `url.Values`.
func (r WhatsappTemplateDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
