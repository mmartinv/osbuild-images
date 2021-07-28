// Package cloudapi provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package cloudapi

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// AWSS3UploadRequestOptions defines model for AWSS3UploadRequestOptions.
type AWSS3UploadRequestOptions struct {
	Region string                    `json:"region"`
	S3     AWSUploadRequestOptionsS3 `json:"s3"`
}

// AWSS3UploadStatus defines model for AWSS3UploadStatus.
type AWSS3UploadStatus struct {
	Url string `json:"url"`
}

// AWSUploadRequestOptions defines model for AWSUploadRequestOptions.
type AWSUploadRequestOptions struct {
	Ec2    AWSUploadRequestOptionsEc2 `json:"ec2"`
	Region string                     `json:"region"`
	S3     AWSUploadRequestOptionsS3  `json:"s3"`
}

// AWSUploadRequestOptionsEc2 defines model for AWSUploadRequestOptionsEc2.
type AWSUploadRequestOptionsEc2 struct {
	AccessKeyId       string    `json:"access_key_id"`
	SecretAccessKey   string    `json:"secret_access_key"`
	ShareWithAccounts *[]string `json:"share_with_accounts,omitempty"`
	SnapshotName      *string   `json:"snapshot_name,omitempty"`
}

// AWSUploadRequestOptionsS3 defines model for AWSUploadRequestOptionsS3.
type AWSUploadRequestOptionsS3 struct {
	AccessKeyId     string `json:"access_key_id"`
	Bucket          string `json:"bucket"`
	SecretAccessKey string `json:"secret_access_key"`
}

// AWSUploadStatus defines model for AWSUploadStatus.
type AWSUploadStatus struct {
	Ami    string `json:"ami"`
	Region string `json:"region"`
}

// AzureUploadRequestOptions defines model for AzureUploadRequestOptions.
type AzureUploadRequestOptions struct {

	// Name of the uploaded image. It must be unique in the given resource group.
	// If name is omitted from the request, a random one based on a UUID is
	// generated.
	ImageName *string `json:"image_name,omitempty"`

	// Location where the image should be uploaded and registered. This link explain
	// how to list all locations:
	// https://docs.microsoft.com/en-us/cli/azure/account?view=azure-cli-latest#az_account_list_locations'
	Location string `json:"location"`

	// Name of the resource group where the image should be uploaded.
	ResourceGroup string `json:"resource_group"`

	// ID of subscription where the image should be uploaded.
	SubscriptionId string `json:"subscription_id"`

	// ID of the tenant where the image should be uploaded. This link explains how
	// to find it in the Azure Portal:
	// https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-how-to-find-tenant
	TenantId string `json:"tenant_id"`
}

// AzureUploadStatus defines model for AzureUploadStatus.
type AzureUploadStatus struct {
	ImageName string `json:"image_name"`
}

// ComposeMetadata defines model for ComposeMetadata.
type ComposeMetadata struct {

	// ID (hash) of the built commit
	OstreeCommit *string `json:"ostree_commit,omitempty"`

	// Package list including NEVRA
	Packages *[]PackageMetadata `json:"packages,omitempty"`
}

// ComposeRequest defines model for ComposeRequest.
type ComposeRequest struct {
	Customizations *Customizations `json:"customizations,omitempty"`
	Distribution   string          `json:"distribution"`
	ImageRequests  []ImageRequest  `json:"image_requests"`
}

// ComposeResult defines model for ComposeResult.
type ComposeResult struct {
	Id string `json:"id"`
}

// ComposeStatus defines model for ComposeStatus.
type ComposeStatus struct {
	ImageStatus ImageStatus `json:"image_status"`
}

// Customizations defines model for Customizations.
type Customizations struct {
	Packages     *[]string     `json:"packages,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
	Users        *[]User       `json:"users,omitempty"`
}

// GCPUploadRequestOptions defines model for GCPUploadRequestOptions.
type GCPUploadRequestOptions struct {

	// Name of an existing STANDARD Storage class Bucket.
	Bucket string `json:"bucket"`

	// The name to use for the imported and shared Compute Engine image.
	// The image name must be unique within the GCP project, which is used
	// for the OS image upload and import. If not specified a random
	// 'composer-api-<uuid>' string is used as the image name.
	ImageName *string `json:"image_name,omitempty"`

	// The GCP region where the OS image will be imported to and shared from.
	// The value must be a valid GCP location. See https://cloud.google.com/storage/docs/locations.
	// If not specified, the multi-region location closest to the source
	// (source Storage Bucket location) is chosen automatically.
	Region *string `json:"region,omitempty"`

	// List of valid Google accounts to share the imported Compute Engine image with.
	// Each string must contain a specifier of the account type. Valid formats are:
	//   - 'user:{emailid}': An email address that represents a specific
	//     Google account. For example, 'alice@example.com'.
	//   - 'serviceAccount:{emailid}': An email address that represents a
	//     service account. For example, 'my-other-app@appspot.gserviceaccount.com'.
	//   - 'group:{emailid}': An email address that represents a Google group.
	//     For example, 'admins@example.com'.
	//   - 'domain:{domain}': The G Suite domain (primary) that represents all
	//     the users of that domain. For example, 'google.com' or 'example.com'.
	// If not specified, the imported Compute Engine image is not shared with any
	// account.
	ShareWithAccounts *[]string `json:"share_with_accounts,omitempty"`
}

// GCPUploadStatus defines model for GCPUploadStatus.
type GCPUploadStatus struct {
	ImageName string `json:"image_name"`
	ProjectId string `json:"project_id"`
}

// ImageRequest defines model for ImageRequest.
type ImageRequest struct {
	Architecture  string        `json:"architecture"`
	ImageType     string        `json:"image_type"`
	Ostree        *OSTree       `json:"ostree,omitempty"`
	Repositories  []Repository  `json:"repositories"`
	UploadRequest UploadRequest `json:"upload_request"`
}

// ImageStatus defines model for ImageStatus.
type ImageStatus struct {
	Status       ImageStatusValue `json:"status"`
	UploadStatus *UploadStatus    `json:"upload_status,omitempty"`
}

// ImageStatusValue defines model for ImageStatusValue.
type ImageStatusValue string

// List of ImageStatusValue
const (
	ImageStatusValue_building    ImageStatusValue = "building"
	ImageStatusValue_failure     ImageStatusValue = "failure"
	ImageStatusValue_pending     ImageStatusValue = "pending"
	ImageStatusValue_registering ImageStatusValue = "registering"
	ImageStatusValue_success     ImageStatusValue = "success"
	ImageStatusValue_uploading   ImageStatusValue = "uploading"
)

// OSTree defines model for OSTree.
type OSTree struct {
	Ref *string `json:"ref,omitempty"`
	Url *string `json:"url,omitempty"`
}

// PackageMetadata defines model for PackageMetadata.
type PackageMetadata struct {
	Arch      string  `json:"arch"`
	Epoch     *string `json:"epoch,omitempty"`
	Name      string  `json:"name"`
	Release   string  `json:"release"`
	Sigmd5    string  `json:"sigmd5"`
	Signature *string `json:"signature,omitempty"`
	Type      string  `json:"type"`
	Version   string  `json:"version"`
}

// Repository defines model for Repository.
type Repository struct {
	Baseurl    *string `json:"baseurl,omitempty"`
	Metalink   *string `json:"metalink,omitempty"`
	Mirrorlist *string `json:"mirrorlist,omitempty"`
	Rhsm       bool    `json:"rhsm"`
}

// Subscription defines model for Subscription.
type Subscription struct {
	ActivationKey string `json:"activation-key"`
	BaseUrl       string `json:"base-url"`
	Insights      bool   `json:"insights"`
	Organization  int    `json:"organization"`
	ServerUrl     string `json:"server-url"`
}

// UploadRequest defines model for UploadRequest.
type UploadRequest struct {
	Options interface{} `json:"options"`
	Type    UploadTypes `json:"type"`
}

// UploadStatus defines model for UploadStatus.
type UploadStatus struct {
	Options interface{} `json:"options"`
	Status  string      `json:"status"`
	Type    UploadTypes `json:"type"`
}

// UploadTypes defines model for UploadTypes.
type UploadTypes string

// List of UploadTypes
const (
	UploadTypes_aws    UploadTypes = "aws"
	UploadTypes_aws_s3 UploadTypes = "aws.s3"
	UploadTypes_azure  UploadTypes = "azure"
	UploadTypes_gcp    UploadTypes = "gcp"
)

// User defines model for User.
type User struct {
	Groups *[]string `json:"groups,omitempty"`
	Key    *string   `json:"key,omitempty"`
	Name   string    `json:"name"`
}

// Version defines model for Version.
type Version struct {
	Version string `json:"version"`
}

// ComposeJSONBody defines parameters for Compose.
type ComposeJSONBody ComposeRequest

// ComposeRequestBody defines body for Compose for application/json ContentType.
type ComposeJSONRequestBody ComposeJSONBody

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = http.DefaultClient
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditor = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// Compose request  with any body
	ComposeWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error)

	Compose(ctx context.Context, body ComposeJSONRequestBody) (*http.Response, error)

	// ComposeStatus request
	ComposeStatus(ctx context.Context, id string) (*http.Response, error)

	// ComposeMetadata request
	ComposeMetadata(ctx context.Context, id string) (*http.Response, error)

	// GetOpenapiJson request
	GetOpenapiJson(ctx context.Context) (*http.Response, error)

	// GetVersion request
	GetVersion(ctx context.Context) (*http.Response, error)
}

func (c *Client) ComposeWithBody(ctx context.Context, contentType string, body io.Reader) (*http.Response, error) {
	req, err := NewComposeRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) Compose(ctx context.Context, body ComposeJSONRequestBody) (*http.Response, error) {
	req, err := NewComposeRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) ComposeStatus(ctx context.Context, id string) (*http.Response, error) {
	req, err := NewComposeStatusRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) ComposeMetadata(ctx context.Context, id string) (*http.Response, error) {
	req, err := NewComposeMetadataRequest(c.Server, id)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetOpenapiJson(ctx context.Context) (*http.Response, error) {
	req, err := NewGetOpenapiJsonRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

func (c *Client) GetVersion(ctx context.Context) (*http.Response, error) {
	req, err := NewGetVersionRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if c.RequestEditor != nil {
		err = c.RequestEditor(ctx, req)
		if err != nil {
			return nil, err
		}
	}
	return c.Client.Do(req)
}

// NewComposeRequest calls the generic Compose builder with application/json body
func NewComposeRequest(server string, body ComposeJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewComposeRequestWithBody(server, "application/json", bodyReader)
}

// NewComposeRequestWithBody generates requests for Compose with any type of body
func NewComposeRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/compose")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryUrl.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)
	return req, nil
}

// NewComposeStatusRequest generates requests for ComposeStatus
func NewComposeStatusRequest(server string, id string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "id", id)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/compose/%s", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewComposeMetadataRequest generates requests for ComposeMetadata
func NewComposeMetadataRequest(server string, id string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParam("simple", false, "id", id)
	if err != nil {
		return nil, err
	}

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/compose/%s/metadata", pathParam0)
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetOpenapiJsonRequest generates requests for GetOpenapiJson
func NewGetOpenapiJsonRequest(server string) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/openapi.json")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetVersionRequest generates requests for GetVersion
func NewGetVersionRequest(server string) (*http.Request, error) {
	var err error

	queryUrl, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	basePath := fmt.Sprintf("/version")
	if basePath[0] == '/' {
		basePath = basePath[1:]
	}

	queryUrl, err = queryUrl.Parse(basePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// Compose request  with any body
	ComposeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*ComposeResponse, error)

	ComposeWithResponse(ctx context.Context, body ComposeJSONRequestBody) (*ComposeResponse, error)

	// ComposeStatus request
	ComposeStatusWithResponse(ctx context.Context, id string) (*ComposeStatusResponse, error)

	// ComposeMetadata request
	ComposeMetadataWithResponse(ctx context.Context, id string) (*ComposeMetadataResponse, error)

	// GetOpenapiJson request
	GetOpenapiJsonWithResponse(ctx context.Context) (*GetOpenapiJsonResponse, error)

	// GetVersion request
	GetVersionWithResponse(ctx context.Context) (*GetVersionResponse, error)
}

type ComposeResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON201      *ComposeResult
}

// Status returns HTTPResponse.Status
func (r ComposeResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ComposeResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ComposeStatusResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ComposeStatus
}

// Status returns HTTPResponse.Status
func (r ComposeStatusResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ComposeStatusResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ComposeMetadataResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *ComposeMetadata
}

// Status returns HTTPResponse.Status
func (r ComposeMetadataResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ComposeMetadataResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetOpenapiJsonResponse struct {
	Body         []byte
	HTTPResponse *http.Response
}

// Status returns HTTPResponse.Status
func (r GetOpenapiJsonResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetOpenapiJsonResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetVersionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Version
}

// Status returns HTTPResponse.Status
func (r GetVersionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetVersionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ComposeWithBodyWithResponse request with arbitrary body returning *ComposeResponse
func (c *ClientWithResponses) ComposeWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader) (*ComposeResponse, error) {
	rsp, err := c.ComposeWithBody(ctx, contentType, body)
	if err != nil {
		return nil, err
	}
	return ParseComposeResponse(rsp)
}

func (c *ClientWithResponses) ComposeWithResponse(ctx context.Context, body ComposeJSONRequestBody) (*ComposeResponse, error) {
	rsp, err := c.Compose(ctx, body)
	if err != nil {
		return nil, err
	}
	return ParseComposeResponse(rsp)
}

// ComposeStatusWithResponse request returning *ComposeStatusResponse
func (c *ClientWithResponses) ComposeStatusWithResponse(ctx context.Context, id string) (*ComposeStatusResponse, error) {
	rsp, err := c.ComposeStatus(ctx, id)
	if err != nil {
		return nil, err
	}
	return ParseComposeStatusResponse(rsp)
}

// ComposeMetadataWithResponse request returning *ComposeMetadataResponse
func (c *ClientWithResponses) ComposeMetadataWithResponse(ctx context.Context, id string) (*ComposeMetadataResponse, error) {
	rsp, err := c.ComposeMetadata(ctx, id)
	if err != nil {
		return nil, err
	}
	return ParseComposeMetadataResponse(rsp)
}

// GetOpenapiJsonWithResponse request returning *GetOpenapiJsonResponse
func (c *ClientWithResponses) GetOpenapiJsonWithResponse(ctx context.Context) (*GetOpenapiJsonResponse, error) {
	rsp, err := c.GetOpenapiJson(ctx)
	if err != nil {
		return nil, err
	}
	return ParseGetOpenapiJsonResponse(rsp)
}

// GetVersionWithResponse request returning *GetVersionResponse
func (c *ClientWithResponses) GetVersionWithResponse(ctx context.Context) (*GetVersionResponse, error) {
	rsp, err := c.GetVersion(ctx)
	if err != nil {
		return nil, err
	}
	return ParseGetVersionResponse(rsp)
}

// ParseComposeResponse parses an HTTP response from a ComposeWithResponse call
func ParseComposeResponse(rsp *http.Response) (*ComposeResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ComposeResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 201:
		var dest ComposeResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON201 = &dest

	}

	return response, nil
}

// ParseComposeStatusResponse parses an HTTP response from a ComposeStatusWithResponse call
func ParseComposeStatusResponse(rsp *http.Response) (*ComposeStatusResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ComposeStatusResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ComposeStatus
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseComposeMetadataResponse parses an HTTP response from a ComposeMetadataWithResponse call
func ParseComposeMetadataResponse(rsp *http.Response) (*ComposeMetadataResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &ComposeMetadataResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest ComposeMetadata
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseGetOpenapiJsonResponse parses an HTTP response from a GetOpenapiJsonWithResponse call
func ParseGetOpenapiJsonResponse(rsp *http.Response) (*GetOpenapiJsonResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetOpenapiJsonResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	}

	return response, nil
}

// ParseGetVersionResponse parses an HTTP response from a GetVersionWithResponse call
func ParseGetVersionResponse(rsp *http.Response) (*GetVersionResponse, error) {
	bodyBytes, err := ioutil.ReadAll(rsp.Body)
	defer rsp.Body.Close()
	if err != nil {
		return nil, err
	}

	response := &GetVersionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Version
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Create compose
	// (POST /compose)
	Compose(w http.ResponseWriter, r *http.Request)
	// The status of a compose
	// (GET /compose/{id})
	ComposeStatus(w http.ResponseWriter, r *http.Request, id string)
	// Get the metadata for a compose.
	// (GET /compose/{id}/metadata)
	ComposeMetadata(w http.ResponseWriter, r *http.Request, id string)
	// get the openapi json specification
	// (GET /openapi.json)
	GetOpenapiJson(w http.ResponseWriter, r *http.Request)
	// get the service version
	// (GET /version)
	GetVersion(w http.ResponseWriter, r *http.Request)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// Compose operation middleware
func (siw *ServerInterfaceWrapper) Compose(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	siw.Handler.Compose(w, r.WithContext(ctx))
}

// ComposeStatus operation middleware
func (siw *ServerInterfaceWrapper) ComposeStatus(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	siw.Handler.ComposeStatus(w, r.WithContext(ctx), id)
}

// ComposeMetadata operation middleware
func (siw *ServerInterfaceWrapper) ComposeMetadata(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	siw.Handler.ComposeMetadata(w, r.WithContext(ctx), id)
}

// GetOpenapiJson operation middleware
func (siw *ServerInterfaceWrapper) GetOpenapiJson(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	siw.Handler.GetOpenapiJson(w, r.WithContext(ctx))
}

// GetVersion operation middleware
func (siw *ServerInterfaceWrapper) GetVersion(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	siw.Handler.GetVersion(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerFromMux(si, chi.NewRouter())
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	r.Group(func(r chi.Router) {
		r.Post("/compose", wrapper.Compose)
	})
	r.Group(func(r chi.Router) {
		r.Get("/compose/{id}", wrapper.ComposeStatus)
	})
	r.Group(func(r chi.Router) {
		r.Get("/compose/{id}/metadata", wrapper.ComposeMetadata)
	})
	r.Group(func(r chi.Router) {
		r.Get("/openapi.json", wrapper.GetOpenapiJson)
	})
	r.Group(func(r chi.Router) {
		r.Get("/version", wrapper.GetVersion)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
