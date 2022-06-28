package app

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ssl "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl/v20191205"
)

var certs []CertInfo

func (r *RequestParams) List() {
	credential := common.NewCredential(r.SecretId, r.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ssl.tencentcloudapi.com"
	client, _ := ssl.NewClient(credential, "", cpf)

	response, _ := client.DescribeCertificates(r.ListReq.Request)

	// first query to get the total of certificates, and then set the limit to TotalCount
	r.ListReq.Request.Limit = response.Response.TotalCount
	response, err := client.DescribeCertificates(r.ListReq.Request)

	for i := 0; i < int(*response.Response.TotalCount); i++ {
		cert := NewCertInfo(
			response.Response.Certificates[i].CertificateId,
			response.Response.Certificates[i].Domain,
			response.Response.Certificates[i].CertEndTime,
			response.Response.Certificates[i].ProjectId,
		)
		certs = append(certs, cert)
	}
	// fmt.Println(len(certs))
	// for i := 0; i < len(certs); i++ {
	// 	fmt.Printf("%v \n", *certs[i].Domain)
	// }
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	if r.Display {
		fmt.Printf("%s", response.ToJsonString())
	}
}
