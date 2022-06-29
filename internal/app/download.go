package app

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ssl "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl/v20191205"
)

func (r *RequestParams) Download() {
	// get domain list from dnspod
	r.List()

	// filter domains from nginx profile
	domains, err := Filter("/etc/nginx/nginx.conf")
	if err != nil {
		log.Println(err)
		return
	}

	// compare domain list with domains for download certificates
	var tmpCertInfo []CertInfo
	for i := 0; i < len(domains); i++ {
		for j := 0; j < len(certs); j++ {
			if domains[i] == *certs[j].Domain {
				tmpCert := certs[j]
				tmpCertInfo = append(tmpCertInfo, tmpCert)
			}
		}
	}

	// Cyclically download all certificates in the nginx configuration file
	for i := 0; i < len(tmpCertInfo); i++ {
		r.DownloadReq.Request.CertificateId = tmpCertInfo[i].CertificateId
		credential := common.NewCredential(r.SecretId, r.SecretKey)
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "ssl.tencentcloudapi.com"
		client, _ := ssl.NewClient(credential, "", cpf)

		response, err := client.DownloadCertificate(r.DownloadReq.Request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			fmt.Printf("An API error has returned: %s", err)
			return
		}
		if err != nil {
			panic(err)
		}

		// decode base64 content
		data, err := base64.StdEncoding.DecodeString(*response.Response.Content)
		if err != nil {
			fmt.Println("decode content failed, error: ", err)
		}

		// TODO: determine certificate expiration time to load cert.
		UnzipWithSave(data, tmpCertInfo[i].Domain)
	}
}
