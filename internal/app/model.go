package app

import (
	"os"

	ssl "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl/v20191205"
)

type Opts interface {
	List()
	Upload()
	Download()
}

type RequestParams struct {
	UploadReq   *UploadCertInfo
	DownloadReq *DownloadCertInfo
	ListReq     *ListCertInfo
	SecretId    string
	SecretKey   string
	Display     bool
	Domain      string
}

func NewRequestParams(display bool, domain string, offset uint64, limit uint64, upload int64) Opts {
	res := &RequestParams{
		SecretId:    os.Getenv("SecretId"),
		SecretKey:   os.Getenv("SecretKey"),
		Display:     display,
		Domain:      domain,
		ListReq:     NewListCertInfo(),
		DownloadReq: NewDownloadCertInfo(),
		UploadReq:   NewUploadCertInfo(),
	}
	res.ListReq.Request.Upload = &upload
	res.ListReq.Request.Limit = &limit
	res.ListReq.Request.Offset = &offset
	return res
}

type CertInfo struct {
	CertificateId *string
	Domain        *string
	CertEndTime   *string
	ProjectId     *string
}

func NewCertInfo(cid *string, domain *string, endtime *string, pid *string) CertInfo {
	return CertInfo{
		CertificateId: cid,
		Domain:        domain,
		CertEndTime:   endtime,
		ProjectId:     pid,
	}
}

type ListCertInfo struct {
	Request *ssl.DescribeCertificatesRequest
}

func NewListCertInfo() *ListCertInfo {
	return &ListCertInfo{
		Request: ssl.NewDescribeCertificatesRequest(),
	}
}

type DownloadCertInfo struct {
	Request *ssl.DownloadCertificateRequest
}

func NewDownloadCertInfo() *DownloadCertInfo {
	return &DownloadCertInfo{
		Request: ssl.NewDownloadCertificateRequest(),
	}
}

type UploadCertInfo struct {
	Request *ssl.UploadCertificateRequest
}

func NewUploadCertInfo() *UploadCertInfo {
	return &UploadCertInfo{
		Request: ssl.NewUploadCertificateRequest(),
	}
}
