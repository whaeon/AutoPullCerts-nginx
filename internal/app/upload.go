package app

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ssl "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ssl/v20191205"
)

func (r *RequestParams) Upload() {
	// get file list
	files, err := os.ReadDir("/etc/letsencrypt/live")
	if err != nil {
		log.Println(err)
		return
	}

	// get cert list
	r.List()

	// compare cert list with file list, if the domain is equal, get projectID
	// if the domain cannot find in the certs, it will be not upload to the dnspod.
	var uploadCerts []CertInfo
	for i := 0; i < len(files); i++ {
		for j := 0; j < len(certs); j++ {
			if *certs[j].Domain == files[i].Name() {
				cert := NewCertInfo(
					certs[j].CertificateId,
					certs[j].Domain,
					certs[j].CertEndTime,
					certs[j].ProjectId,
				)
				uploadCerts = append(uploadCerts, cert)
				continue
			}
		}
	}

	// loop to get certificate content
	for i := 0; i < len(uploadCerts); i++ {
		// get certificate public key
		PublicKeyPath := "/etc/letsencrypt/live/" + *uploadCerts[i].Domain + "/fullchain.pem"
		pubContent, err := os.ReadFile(PublicKeyPath)
		PublicKeyContent := string(pubContent)
		if err != nil {
			log.Println(err)
			return
		}
		// PublicKeyFile, err := os.OpenFile(PublicKeyPath, os.O_RDONLY, 0666)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// PublicKeyReader := bufio.NewReader(PublicKeyFile)
		// var PublicKeyContent string
		// for {
		// 	content, _, err := PublicKeyReader.ReadLine()
		// 	PublicKeyContent = PublicKeyContent + string(content)
		// 	if err != nil {
		// 		if err == io.EOF {
		// 			break
		// 		}
		// 		log.Println(err)
		// 		break
		// 	}
		// }

		// get certificate private key
		PrivateKeyPath := "/etc/letsencrypt/live/" + *uploadCerts[i].Domain + "/privkey.pem"
		privContent, err := os.ReadFile(PrivateKeyPath)
		PrivateKeyContent := string(privContent)
		if err != nil {
			log.Println(err)
			return
		}
		// PrivateKeyFile, err := os.OpenFile(PrivateKeyPath, os.O_RDONLY, 0666)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// PrivateKeyReader := bufio.NewReader(PrivateKeyFile)
		// var PrivateKeyContent string
		// for {
		// 	content, _, err := PrivateKeyReader.ReadLine()
		// 	PrivateKeyContent = PrivateKeyContent + string(content)
		// 	if err != nil {
		// 		if err == io.EOF {
		// 			break
		// 		}
		// 		log.Println(err)
		// 		break
		// 	}
		// }

		// prepare to upload certificate
		credential := common.NewCredential(r.SecretId, r.SecretKey)
		cpf := profile.NewClientProfile()
		cpf.HttpProfile.Endpoint = "ssl.tencentcloudapi.com"
		client, _ := ssl.NewClient(credential, "", cpf)

		// prepare upload parameters
		r.UploadReq.Request.CertificatePublicKey = &PublicKeyContent
		r.UploadReq.Request.CertificatePrivateKey = &PrivateKeyContent
		ProjectId, _ := strconv.ParseUint(*uploadCerts[i].ProjectId, 10, 64)
		r.UploadReq.Request.ProjectId = &ProjectId
		r.UploadReq.Request.Alias = uploadCerts[i].Domain

		// execute upload
		response, err := client.UploadCertificate(r.UploadReq.Request)
		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			fmt.Printf("An API error has returned: %s", err)
			return
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s %v \n", response.ToJsonString(), *uploadCerts[i].Domain)

		// close open file

		// PublicKeyFile.Close()
		// PrivateKeyFile.Close()
	}
}
