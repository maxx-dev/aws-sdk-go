// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package support provides a client for AWS Support.
package support

import (
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/gen/endpoints"
)

// Support is a client for AWS Support.
type Support struct {
	client *aws.JSONClient
}

// New returns a new Support client.
func New(creds aws.Credentials, region string, client *http.Client) *Support {
	if client == nil {
		client = http.DefaultClient
	}

	service := "support"
	endpoint, service, region := endpoints.Lookup("support", region)

	return &Support{
		client: &aws.JSONClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			}, Client: client,
			Endpoint:     endpoint,
			JSONVersion:  "1.1",
			TargetPrefix: "AWSSupport_20130415",
		},
	}
}

// AddAttachmentsToSet adds one or more attachments to an attachment set.
// If an AttachmentSetId is not specified, a new attachment set is created,
// and the ID of the set is returned in the response. If an AttachmentSetId
// is specified, the attachments are added to the specified set, if it
// exists. An attachment set is a temporary container for attachments that
// are to be added to a case or case communication. The set is available
// for one hour after it is created; the ExpiryTime returned in the
// response indicates when the set expires. The maximum number of
// attachments in a set is 3, and the maximum size of any attachment in the
// set is 5
func (c *Support) AddAttachmentsToSet(req *AddAttachmentsToSetRequest) (resp *AddAttachmentsToSetResponse, err error) {
	resp = &AddAttachmentsToSetResponse{}
	err = c.client.Do("AddAttachmentsToSet", "POST", "/", req, resp)
	return
}

// AddCommunicationToCase adds additional customer communication to an AWS
// Support case. You use the CaseId value to identify the case to add
// communication to. You can list a set of email addresses to copy on the
// communication using the CcEmailAddresses value. The CommunicationBody
// value contains the text of the communication. The response indicates the
// success or failure of the request. This operation implements a subset of
// the features of the AWS Support Center.
func (c *Support) AddCommunicationToCase(req *AddCommunicationToCaseRequest) (resp *AddCommunicationToCaseResponse, err error) {
	resp = &AddCommunicationToCaseResponse{}
	err = c.client.Do("AddCommunicationToCase", "POST", "/", req, resp)
	return
}

// CreateCase creates a new case in the AWS Support Center. This operation
// is modeled on the behavior of the AWS Support Center Create Case page.
// Its parameters require you to specify the following information:
// IssueType. The type of issue for the case. You can specify either
// "customer-service" or "technical." If you do not indicate a value, the
// default is "technical." ServiceCode. The code for an AWS service. You
// obtain the ServiceCode by calling DescribeServices . CategoryCode. The
// category for the service defined for the ServiceCode value. You also
// obtain the category code for a service by calling DescribeServices .
// Each AWS service defines its own set of category codes. SeverityCode. A
// value that indicates the urgency of the case, which in turn determines
// the response time according to your service level agreement with AWS
// Support. You obtain the SeverityCode by calling DescribeSeverityLevels
// Subject. The Subject field on the AWS Support Center Create Case page.
// CommunicationBody. The Description field on the AWS Support Center
// Create Case page. AttachmentSetId. The ID of a set of attachments that
// has been created by using AddAttachmentsToSet Language. The human
// language in which AWS Support handles the case. English and Japanese are
// currently supported. CcEmailAddresses. The AWS Support Center field on
// the Create Case page. You can list email addresses to be copied on any
// correspondence about the case. The account that opens the case is
// already identified by passing the AWS Credentials in the method or in a
// method or function call from one of the programming languages supported
// by an AWS . To add additional communication or attachments to an
// existing case, use AddCommunicationToCase A successful CreateCase
// request returns an AWS Support case number. Case numbers are used by the
// DescribeCases operation to retrieve existing AWS Support cases.
func (c *Support) CreateCase(req *CreateCaseRequest) (resp *CreateCaseResponse, err error) {
	resp = &CreateCaseResponse{}
	err = c.client.Do("CreateCase", "POST", "/", req, resp)
	return
}

// DescribeAttachment returns the attachment that has the specified ID.
// Attachment IDs are generated by the case management system when you add
// an attachment to a case or case communication. Attachment IDs are
// returned in the AttachmentDetails objects that are returned by the
// DescribeCommunications operation.
func (c *Support) DescribeAttachment(req *DescribeAttachmentRequest) (resp *DescribeAttachmentResponse, err error) {
	resp = &DescribeAttachmentResponse{}
	err = c.client.Do("DescribeAttachment", "POST", "/", req, resp)
	return
}

// DescribeCases returns a list of cases that you specify by passing one or
// more case IDs. In addition, you can filter the cases by date by setting
// values for the AfterTime and BeforeTime request parameters. You can set
// values for the IncludeResolvedCases and IncludeCommunications request
// parameters to control how much information is returned. Case data is
// available for 12 months after creation. If a case was created more than
// 12 months ago, a request for data might cause an error. The response
// returns the following in format: One or more CaseDetails data types. One
// or more NextToken values, which specify where to paginate the returned
// records represented by the CaseDetails objects.
func (c *Support) DescribeCases(req *DescribeCasesRequest) (resp *DescribeCasesResponse, err error) {
	resp = &DescribeCasesResponse{}
	err = c.client.Do("DescribeCases", "POST", "/", req, resp)
	return
}

// DescribeCommunications returns communications (and attachments) for one
// or more support cases. You can use the AfterTime and BeforeTime
// parameters to filter by date. You can use the CaseId parameter to
// restrict the results to a particular case. Case data is available for 12
// months after creation. If a case was created more than 12 months ago, a
// request for data might cause an error. You can use the MaxResults and
// NextToken parameters to control the pagination of the result set. Set
// MaxResults to the number of cases you want displayed on each page, and
// use NextToken to specify the resumption of pagination.
func (c *Support) DescribeCommunications(req *DescribeCommunicationsRequest) (resp *DescribeCommunicationsResponse, err error) {
	resp = &DescribeCommunicationsResponse{}
	err = c.client.Do("DescribeCommunications", "POST", "/", req, resp)
	return
}

// DescribeServices returns the current list of AWS services and a list of
// service categories that applies to each one. You then use service names
// and categories in your CreateCase requests. Each AWS service has its own
// set of categories. The service codes and category codes correspond to
// the values that are displayed in the Service and Category drop-down
// lists on the AWS Support Center Create Case page. The values in those
// fields, however, do not necessarily match the service codes and
// categories returned by the DescribeServices request. Always use the
// service codes and categories obtained programmatically. This practice
// ensures that you always have the most recent set of service and category
// codes.
func (c *Support) DescribeServices(req *DescribeServicesRequest) (resp *DescribeServicesResponse, err error) {
	resp = &DescribeServicesResponse{}
	err = c.client.Do("DescribeServices", "POST", "/", req, resp)
	return
}

// DescribeSeverityLevels returns the list of severity levels that you can
// assign to an AWS Support case. The severity level for a case is also a
// field in the CaseDetails data type included in any CreateCase request.
func (c *Support) DescribeSeverityLevels(req *DescribeSeverityLevelsRequest) (resp *DescribeSeverityLevelsResponse, err error) {
	resp = &DescribeSeverityLevelsResponse{}
	err = c.client.Do("DescribeSeverityLevels", "POST", "/", req, resp)
	return
}

// DescribeTrustedAdvisorCheckRefreshStatuses returns the refresh status of
// the Trusted Advisor checks that have the specified check IDs. Check IDs
// can be obtained by calling DescribeTrustedAdvisorChecks
func (c *Support) DescribeTrustedAdvisorCheckRefreshStatuses(req *DescribeTrustedAdvisorCheckRefreshStatusesRequest) (resp *DescribeTrustedAdvisorCheckRefreshStatusesResponse, err error) {
	resp = &DescribeTrustedAdvisorCheckRefreshStatusesResponse{}
	err = c.client.Do("DescribeTrustedAdvisorCheckRefreshStatuses", "POST", "/", req, resp)
	return
}

// DescribeTrustedAdvisorCheckResult returns the results of the Trusted
// Advisor check that has the specified check ID. Check IDs can be obtained
// by calling DescribeTrustedAdvisorChecks The response contains a
// TrustedAdvisorCheckResult object, which contains these three objects: In
// addition, the response contains these fields: Status. The alert status
// of the check: "ok" (green), "warning" (yellow), "error" (red), or
// "not_available". Timestamp. The time of the last refresh of the check.
// CheckId. The unique identifier for the check.
func (c *Support) DescribeTrustedAdvisorCheckResult(req *DescribeTrustedAdvisorCheckResultRequest) (resp *DescribeTrustedAdvisorCheckResultResponse, err error) {
	resp = &DescribeTrustedAdvisorCheckResultResponse{}
	err = c.client.Do("DescribeTrustedAdvisorCheckResult", "POST", "/", req, resp)
	return
}

// DescribeTrustedAdvisorCheckSummaries returns the summaries of the
// results of the Trusted Advisor checks that have the specified check IDs.
// Check IDs can be obtained by calling DescribeTrustedAdvisorChecks The
// response contains an array of TrustedAdvisorCheckSummary objects.
func (c *Support) DescribeTrustedAdvisorCheckSummaries(req *DescribeTrustedAdvisorCheckSummariesRequest) (resp *DescribeTrustedAdvisorCheckSummariesResponse, err error) {
	resp = &DescribeTrustedAdvisorCheckSummariesResponse{}
	err = c.client.Do("DescribeTrustedAdvisorCheckSummaries", "POST", "/", req, resp)
	return
}

// DescribeTrustedAdvisorChecks returns information about all available
// Trusted Advisor checks, including name, ID, category, description, and
// metadata. You must specify a language code; English ("en") and Japanese
// ("ja") are currently supported. The response contains a
// TrustedAdvisorCheckDescription for each check.
func (c *Support) DescribeTrustedAdvisorChecks(req *DescribeTrustedAdvisorChecksRequest) (resp *DescribeTrustedAdvisorChecksResponse, err error) {
	resp = &DescribeTrustedAdvisorChecksResponse{}
	err = c.client.Do("DescribeTrustedAdvisorChecks", "POST", "/", req, resp)
	return
}

// RefreshTrustedAdvisorCheck requests a refresh of the Trusted Advisor
// check that has the specified check ID. Check IDs can be obtained by
// calling DescribeTrustedAdvisorChecks The response contains a
// TrustedAdvisorCheckRefreshStatus object, which contains these fields:
// Status. The refresh status of the check: "none", "enqueued",
// "processing", "success", or "abandoned". MillisUntilNextRefreshable. The
// amount of time, in milliseconds, until the check is eligible for
// refresh. CheckId. The unique identifier for the check.
func (c *Support) RefreshTrustedAdvisorCheck(req *RefreshTrustedAdvisorCheckRequest) (resp *RefreshTrustedAdvisorCheckResponse, err error) {
	resp = &RefreshTrustedAdvisorCheckResponse{}
	err = c.client.Do("RefreshTrustedAdvisorCheck", "POST", "/", req, resp)
	return
}

// ResolveCase takes a CaseId and returns the initial state of the case
// along with the state of the case after the call to ResolveCase
// completed.
func (c *Support) ResolveCase(req *ResolveCaseRequest) (resp *ResolveCaseResponse, err error) {
	resp = &ResolveCaseResponse{}
	err = c.client.Do("ResolveCase", "POST", "/", req, resp)
	return
}

// AddAttachmentsToSetRequest is undocumented.
type AddAttachmentsToSetRequest struct {
	AttachmentSetID aws.StringValue `json:"attachmentSetId,omitempty"`
	Attachments     []Attachment    `json:"attachments"`
}

// AddAttachmentsToSetResponse is undocumented.
type AddAttachmentsToSetResponse struct {
	AttachmentSetID aws.StringValue `json:"attachmentSetId,omitempty"`
	ExpiryTime      aws.StringValue `json:"expiryTime,omitempty"`
}

// AddCommunicationToCaseRequest is undocumented.
type AddCommunicationToCaseRequest struct {
	AttachmentSetID   aws.StringValue `json:"attachmentSetId,omitempty"`
	CaseID            aws.StringValue `json:"caseId,omitempty"`
	CCEmailAddresses  []string        `json:"ccEmailAddresses,omitempty"`
	CommunicationBody aws.StringValue `json:"communicationBody"`
}

// AddCommunicationToCaseResponse is undocumented.
type AddCommunicationToCaseResponse struct {
	Result aws.BooleanValue `json:"result,omitempty"`
}

// Attachment is undocumented.
type Attachment struct {
	Data     []byte          `json:"data,omitempty"`
	FileName aws.StringValue `json:"fileName,omitempty"`
}

// AttachmentDetails is undocumented.
type AttachmentDetails struct {
	AttachmentID aws.StringValue `json:"attachmentId,omitempty"`
	FileName     aws.StringValue `json:"fileName,omitempty"`
}

// CaseDetails is undocumented.
type CaseDetails struct {
	CaseID               aws.StringValue           `json:"caseId,omitempty"`
	CategoryCode         aws.StringValue           `json:"categoryCode,omitempty"`
	CCEmailAddresses     []string                  `json:"ccEmailAddresses,omitempty"`
	DisplayID            aws.StringValue           `json:"displayId,omitempty"`
	Language             aws.StringValue           `json:"language,omitempty"`
	RecentCommunications *RecentCaseCommunications `json:"recentCommunications,omitempty"`
	ServiceCode          aws.StringValue           `json:"serviceCode,omitempty"`
	SeverityCode         aws.StringValue           `json:"severityCode,omitempty"`
	Status               aws.StringValue           `json:"status,omitempty"`
	Subject              aws.StringValue           `json:"subject,omitempty"`
	SubmittedBy          aws.StringValue           `json:"submittedBy,omitempty"`
	TimeCreated          aws.StringValue           `json:"timeCreated,omitempty"`
}

// Category is undocumented.
type Category struct {
	Code aws.StringValue `json:"code,omitempty"`
	Name aws.StringValue `json:"name,omitempty"`
}

// Communication is undocumented.
type Communication struct {
	AttachmentSet []AttachmentDetails `json:"attachmentSet,omitempty"`
	Body          aws.StringValue     `json:"body,omitempty"`
	CaseID        aws.StringValue     `json:"caseId,omitempty"`
	SubmittedBy   aws.StringValue     `json:"submittedBy,omitempty"`
	TimeCreated   aws.StringValue     `json:"timeCreated,omitempty"`
}

// CreateCaseRequest is undocumented.
type CreateCaseRequest struct {
	AttachmentSetID   aws.StringValue `json:"attachmentSetId,omitempty"`
	CategoryCode      aws.StringValue `json:"categoryCode,omitempty"`
	CCEmailAddresses  []string        `json:"ccEmailAddresses,omitempty"`
	CommunicationBody aws.StringValue `json:"communicationBody"`
	IssueType         aws.StringValue `json:"issueType,omitempty"`
	Language          aws.StringValue `json:"language,omitempty"`
	ServiceCode       aws.StringValue `json:"serviceCode,omitempty"`
	SeverityCode      aws.StringValue `json:"severityCode,omitempty"`
	Subject           aws.StringValue `json:"subject"`
}

// CreateCaseResponse is undocumented.
type CreateCaseResponse struct {
	CaseID aws.StringValue `json:"caseId,omitempty"`
}

// DescribeAttachmentRequest is undocumented.
type DescribeAttachmentRequest struct {
	AttachmentID aws.StringValue `json:"attachmentId"`
}

// DescribeAttachmentResponse is undocumented.
type DescribeAttachmentResponse struct {
	Attachment *Attachment `json:"attachment,omitempty"`
}

// DescribeCasesRequest is undocumented.
type DescribeCasesRequest struct {
	AfterTime             aws.StringValue  `json:"afterTime,omitempty"`
	BeforeTime            aws.StringValue  `json:"beforeTime,omitempty"`
	CaseIDList            []string         `json:"caseIdList,omitempty"`
	DisplayID             aws.StringValue  `json:"displayId,omitempty"`
	IncludeCommunications aws.BooleanValue `json:"includeCommunications,omitempty"`
	IncludeResolvedCases  aws.BooleanValue `json:"includeResolvedCases,omitempty"`
	Language              aws.StringValue  `json:"language,omitempty"`
	MaxResults            aws.IntegerValue `json:"maxResults,omitempty"`
	NextToken             aws.StringValue  `json:"nextToken,omitempty"`
}

// DescribeCasesResponse is undocumented.
type DescribeCasesResponse struct {
	Cases     []CaseDetails   `json:"cases,omitempty"`
	NextToken aws.StringValue `json:"nextToken,omitempty"`
}

// DescribeCommunicationsRequest is undocumented.
type DescribeCommunicationsRequest struct {
	AfterTime  aws.StringValue  `json:"afterTime,omitempty"`
	BeforeTime aws.StringValue  `json:"beforeTime,omitempty"`
	CaseID     aws.StringValue  `json:"caseId"`
	MaxResults aws.IntegerValue `json:"maxResults,omitempty"`
	NextToken  aws.StringValue  `json:"nextToken,omitempty"`
}

// DescribeCommunicationsResponse is undocumented.
type DescribeCommunicationsResponse struct {
	Communications []Communication `json:"communications,omitempty"`
	NextToken      aws.StringValue `json:"nextToken,omitempty"`
}

// DescribeServicesRequest is undocumented.
type DescribeServicesRequest struct {
	Language        aws.StringValue `json:"language,omitempty"`
	ServiceCodeList []string        `json:"serviceCodeList,omitempty"`
}

// DescribeServicesResponse is undocumented.
type DescribeServicesResponse struct {
	Services []Service `json:"services,omitempty"`
}

// DescribeSeverityLevelsRequest is undocumented.
type DescribeSeverityLevelsRequest struct {
	Language aws.StringValue `json:"language,omitempty"`
}

// DescribeSeverityLevelsResponse is undocumented.
type DescribeSeverityLevelsResponse struct {
	SeverityLevels []SeverityLevel `json:"severityLevels,omitempty"`
}

// DescribeTrustedAdvisorCheckRefreshStatusesRequest is undocumented.
type DescribeTrustedAdvisorCheckRefreshStatusesRequest struct {
	CheckIDs []string `json:"checkIds"`
}

// DescribeTrustedAdvisorCheckRefreshStatusesResponse is undocumented.
type DescribeTrustedAdvisorCheckRefreshStatusesResponse struct {
	Statuses []TrustedAdvisorCheckRefreshStatus `json:"statuses"`
}

// DescribeTrustedAdvisorCheckResultRequest is undocumented.
type DescribeTrustedAdvisorCheckResultRequest struct {
	CheckID  aws.StringValue `json:"checkId"`
	Language aws.StringValue `json:"language,omitempty"`
}

// DescribeTrustedAdvisorCheckResultResponse is undocumented.
type DescribeTrustedAdvisorCheckResultResponse struct {
	Result *TrustedAdvisorCheckResult `json:"result,omitempty"`
}

// DescribeTrustedAdvisorCheckSummariesRequest is undocumented.
type DescribeTrustedAdvisorCheckSummariesRequest struct {
	CheckIDs []string `json:"checkIds"`
}

// DescribeTrustedAdvisorCheckSummariesResponse is undocumented.
type DescribeTrustedAdvisorCheckSummariesResponse struct {
	Summaries []TrustedAdvisorCheckSummary `json:"summaries"`
}

// DescribeTrustedAdvisorChecksRequest is undocumented.
type DescribeTrustedAdvisorChecksRequest struct {
	Language aws.StringValue `json:"language"`
}

// DescribeTrustedAdvisorChecksResponse is undocumented.
type DescribeTrustedAdvisorChecksResponse struct {
	Checks []TrustedAdvisorCheckDescription `json:"checks"`
}

// RecentCaseCommunications is undocumented.
type RecentCaseCommunications struct {
	Communications []Communication `json:"communications,omitempty"`
	NextToken      aws.StringValue `json:"nextToken,omitempty"`
}

// RefreshTrustedAdvisorCheckRequest is undocumented.
type RefreshTrustedAdvisorCheckRequest struct {
	CheckID aws.StringValue `json:"checkId"`
}

// RefreshTrustedAdvisorCheckResponse is undocumented.
type RefreshTrustedAdvisorCheckResponse struct {
	Status *TrustedAdvisorCheckRefreshStatus `json:"status"`
}

// ResolveCaseRequest is undocumented.
type ResolveCaseRequest struct {
	CaseID aws.StringValue `json:"caseId,omitempty"`
}

// ResolveCaseResponse is undocumented.
type ResolveCaseResponse struct {
	FinalCaseStatus   aws.StringValue `json:"finalCaseStatus,omitempty"`
	InitialCaseStatus aws.StringValue `json:"initialCaseStatus,omitempty"`
}

// Service is undocumented.
type Service struct {
	Categories []Category      `json:"categories,omitempty"`
	Code       aws.StringValue `json:"code,omitempty"`
	Name       aws.StringValue `json:"name,omitempty"`
}

// SeverityLevel is undocumented.
type SeverityLevel struct {
	Code aws.StringValue `json:"code,omitempty"`
	Name aws.StringValue `json:"name,omitempty"`
}

// TrustedAdvisorCategorySpecificSummary is undocumented.
type TrustedAdvisorCategorySpecificSummary struct {
	CostOptimizing *TrustedAdvisorCostOptimizingSummary `json:"costOptimizing,omitempty"`
}

// TrustedAdvisorCheckDescription is undocumented.
type TrustedAdvisorCheckDescription struct {
	Category    aws.StringValue `json:"category"`
	Description aws.StringValue `json:"description"`
	ID          aws.StringValue `json:"id"`
	Metadata    []string        `json:"metadata"`
	Name        aws.StringValue `json:"name"`
}

// TrustedAdvisorCheckRefreshStatus is undocumented.
type TrustedAdvisorCheckRefreshStatus struct {
	CheckID                    aws.StringValue `json:"checkId"`
	MillisUntilNextRefreshable aws.LongValue   `json:"millisUntilNextRefreshable"`
	Status                     aws.StringValue `json:"status"`
}

// TrustedAdvisorCheckResult is undocumented.
type TrustedAdvisorCheckResult struct {
	CategorySpecificSummary *TrustedAdvisorCategorySpecificSummary `json:"categorySpecificSummary"`
	CheckID                 aws.StringValue                        `json:"checkId"`
	FlaggedResources        []TrustedAdvisorResourceDetail         `json:"flaggedResources"`
	ResourcesSummary        *TrustedAdvisorResourcesSummary        `json:"resourcesSummary"`
	Status                  aws.StringValue                        `json:"status"`
	Timestamp               aws.StringValue                        `json:"timestamp"`
}

// TrustedAdvisorCheckSummary is undocumented.
type TrustedAdvisorCheckSummary struct {
	CategorySpecificSummary *TrustedAdvisorCategorySpecificSummary `json:"categorySpecificSummary"`
	CheckID                 aws.StringValue                        `json:"checkId"`
	HasFlaggedResources     aws.BooleanValue                       `json:"hasFlaggedResources,omitempty"`
	ResourcesSummary        *TrustedAdvisorResourcesSummary        `json:"resourcesSummary"`
	Status                  aws.StringValue                        `json:"status"`
	Timestamp               aws.StringValue                        `json:"timestamp"`
}

// TrustedAdvisorCostOptimizingSummary is undocumented.
type TrustedAdvisorCostOptimizingSummary struct {
	EstimatedMonthlySavings        aws.DoubleValue `json:"estimatedMonthlySavings"`
	EstimatedPercentMonthlySavings aws.DoubleValue `json:"estimatedPercentMonthlySavings"`
}

// TrustedAdvisorResourceDetail is undocumented.
type TrustedAdvisorResourceDetail struct {
	IsSuppressed aws.BooleanValue `json:"isSuppressed,omitempty"`
	Metadata     []string         `json:"metadata"`
	Region       aws.StringValue  `json:"region"`
	ResourceID   aws.StringValue  `json:"resourceId"`
	Status       aws.StringValue  `json:"status"`
}

// TrustedAdvisorResourcesSummary is undocumented.
type TrustedAdvisorResourcesSummary struct {
	ResourcesFlagged    aws.LongValue `json:"resourcesFlagged"`
	ResourcesIgnored    aws.LongValue `json:"resourcesIgnored"`
	ResourcesProcessed  aws.LongValue `json:"resourcesProcessed"`
	ResourcesSuppressed aws.LongValue `json:"resourcesSuppressed"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
