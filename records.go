package dnspod

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	methodRecordList   = "Record.List"
	methodRecordCreate = "Record.Create"
	methodRecordInfo   = "Record.Info"
	methodRecordRemove = "Record.Remove"
	methodRecordModify = "Record.Modify"
)

// Record is the DNS record representation.
type Record struct {
	ID            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Line          string `json:"line,omitempty"`
	LineID        string `json:"line_id,omitempty"`
	Type          string `json:"type,omitempty"`
	TTL           string `json:"ttl,omitempty"`
	Value         string `json:"value,omitempty"`
	MX            string `json:"mx,omitempty"`
	Enabled       string `json:"enabled,omitempty"`
	Status        string `json:"status,omitempty"`
	MonitorStatus string `json:"monitor_status,omitempty"`
	Remark        string `json:"remark,omitempty"`
	UpdateOn      string `json:"updated_on,omitempty"`
	UseAQB        string `json:"use_aqb,omitempty"`
	Weight        *int   `json:"weight,omitempty"`
}

// RecordModify is the DNS record modify representation.
type RecordModify struct {
	ID     json.Number `json:"id,omitempty"`
	Name   string      `json:"name,omitempty"`
	Value  string      `json:"value,omitempty"`
	Status string      `json:"status,omitempty"`
}

type recordsWrapper struct {
	Status  Status     `json:"status"`
	Info    DomainInfo `json:"info"`
	Records []Record   `json:"records"`
}

type recordWrapper struct {
	Status Status     `json:"status"`
	Info   DomainInfo `json:"info"`
	Record Record     `json:"record"`
}

type recordModifyWrapper struct {
	Status Status       `json:"status"`
	Record RecordModify `json:"record"`
}

// RecordsService handles communication with the DNS records related methods of the dnspod API.
//
// DNSPod API docs:
// - https://www.dnspod.cn/docs/records.html
// - https://docs.dnspod.com/api/
type RecordsService struct {
	client *Client
}

// List List the domain records.
//
// DNSPod API docs:
// - https://www.dnspod.cn/docs/records.html#record-list
// - https://docs.dnspod.com/api/5fe19a7a6e336701a2111bb9/
func (s *RecordsService) List(domainID string, recordName string) ([]Record, *Response, error) {
	payload := s.client.CommonParams.toPayLoad()
	payload.Add("domain_id", domainID)
	if recordName != "" {
		payload.Add("sub_domain", recordName)
	}

	wrappedRecords := recordsWrapper{}

	res, err := s.client.post(methodRecordList, payload, &wrappedRecords)
	if err != nil {
		return nil, res, err
	}

	if wrappedRecords.Status.Code != "1" {
		return nil, nil, fmt.Errorf("could not get domains: %s", wrappedRecords.Status.Message)
	}

	return wrappedRecords.Records, res, nil
}

// Create Creates a domain record.
//
// DNSPod API docs:
// - https://www.dnspod.cn/docs/records.html#record-create
// - https://docs.dnspod.com/api/5fe19a3f6e336701a2111bb0/
func (s *RecordsService) Create(domain string, recordAttributes Record) (Record, *Response, error) {
	payload := s.client.CommonParams.toPayLoad()
	payload.Add("domain_id", domain)

	if recordAttributes.Name != "" {
		payload.Add("sub_domain", recordAttributes.Name)
	}

	if recordAttributes.Type != "" {
		payload.Add("record_type", recordAttributes.Type)
	}

	if recordAttributes.Line != "" {
		payload.Add("record_line", recordAttributes.Line)
	}

	if recordAttributes.LineID != "" {
		payload.Add("record_line_id", recordAttributes.LineID)
	}

	if recordAttributes.Value != "" {
		payload.Add("value", recordAttributes.Value)
	}

	if recordAttributes.MX != "" {
		payload.Add("mx", recordAttributes.MX)
	}

	if recordAttributes.TTL != "" {
		payload.Add("ttl", recordAttributes.TTL)
	}

	if recordAttributes.Status != "" {
		payload.Add("status", recordAttributes.Status)
	}

	if recordAttributes.Weight != nil {
		payload.Add("weight", strconv.Itoa(*recordAttributes.Weight))
	}

	returnedRecord := recordWrapper{}

	res, err := s.client.post(methodRecordCreate, payload, &returnedRecord)
	if err != nil {
		return Record{}, res, err
	}

	if returnedRecord.Status.Code != "1" {
		return returnedRecord.Record, nil, fmt.Errorf("could not get domains: %s", returnedRecord.Status.Message)
	}

	return returnedRecord.Record, res, nil
}

// Get Fetches the domain record.
//
// DNSPod API docs:
// - https://www.dnspod.cn/docs/records.html#record-info
// - https://docs.dnspod.com/api/5fe1a2a06e336701a2111bcd/
func (s *RecordsService) Get(domain string, recordID string) (Record, *Response, error) {
	payload := s.client.CommonParams.toPayLoad()
	payload.Add("domain_id", domain)
	payload.Add("record_id", recordID)

	returnedRecord := recordWrapper{}

	res, err := s.client.post(methodRecordInfo, payload, &returnedRecord)
	if err != nil {
		return Record{}, res, err
	}

	if returnedRecord.Status.Code != "1" {
		return returnedRecord.Record, nil, fmt.Errorf("could not get domains: %s", returnedRecord.Status.Message)
	}

	return returnedRecord.Record, res, nil
}

// Update Updates a domain record.
//
// DNSPod API docs:
// - https://www.dnspod.cn/docs/records.html#record-modify
// - https://docs.dnspod.com/api/5fe1a5a16e336701a2111c76/
func (s *RecordsService) Update(domain string, recordID string, recordAttributes Record) (RecordModify, *Response, error) {
	payload := s.client.CommonParams.toPayLoad()
	payload.Add("domain_id", domain)
	payload.Add("record_id", recordID)

	if recordAttributes.Name != "" {
		payload.Add("sub_domain", recordAttributes.Name)
	}

	if recordAttributes.Type != "" {
		payload.Add("record_type", recordAttributes.Type)
	}

	if recordAttributes.Line != "" {
		payload.Add("record_line", recordAttributes.Line)
	}

	if recordAttributes.LineID != "" {
		payload.Add("record_line_id", recordAttributes.LineID)
	}

	if recordAttributes.Value != "" {
		payload.Add("value", recordAttributes.Value)
	}

	if recordAttributes.MX != "" {
		payload.Add("mx", recordAttributes.MX)
	}

	if recordAttributes.TTL != "" {
		payload.Add("ttl", recordAttributes.TTL)
	}

	if recordAttributes.Status != "" {
		payload.Add("status", recordAttributes.Status)
	}

	if recordAttributes.Weight != nil {
		payload.Add("weight", strconv.Itoa(*recordAttributes.Weight))
	}

	returnedRecord := recordModifyWrapper{}

	res, err := s.client.post(methodRecordModify, payload, &returnedRecord)
	if err != nil {
		return RecordModify{}, res, err
	}

	if returnedRecord.Status.Code != "1" {
		return returnedRecord.Record, nil, fmt.Errorf("could not get domains: %s", returnedRecord.Status.Message)
	}

	return returnedRecord.Record, res, nil
}

// Delete Deletes a domain record.
//
// DNSPod API docs:
// - https://www.dnspod.cn/docs/records.html#record-remove
// - https://docs.dnspod.com/api/5fe1a4576e336701a2111c24/
func (s *RecordsService) Delete(domain string, recordID string) (*Response, error) {
	payload := s.client.CommonParams.toPayLoad()
	payload.Add("domain_id", domain)
	payload.Add("record_id", recordID)

	returnedRecord := recordWrapper{}

	res, err := s.client.post(methodRecordRemove, payload, &returnedRecord)
	if err != nil {
		return res, err
	}

	if returnedRecord.Status.Code != "1" {
		return nil, fmt.Errorf("could not get domains: %s", returnedRecord.Status.Message)
	}

	return res, nil
}
