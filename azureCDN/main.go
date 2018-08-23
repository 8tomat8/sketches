package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/Azure/azure-storage-blob-go/2018-03-28/azblob"
	"github.com/VerizonDigital/ectoken/go-ectoken"
	"github.com/pkg/errors"
)

const (
	accountName = "accountName"
	accountKey = "accountKey"
	ecKey = "ecKey"
)

func main() {
	credential := azblob.NewSharedKeyCredential(accountName, accountKey)
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})

	// From the Azure portal, get your storage account blob service URL endpoint.
	URL, err := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net", accountName))
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse own url..."))
	}

	svc := azblob.NewServiceURL(*URL, p)

	// Create 3 containers to emulate 3 different apps with separate containers
	for i := 0; i < 3; i++ {
		containerName := fmt.Sprintf("%d", i)

		containerURL := svc.NewContainerURL(containerName)
		_, err = containerURL.Create(context.Background(), azblob.Metadata{"app": strconv.Itoa(i)}, azblob.PublicAccessBlob)
		if err != nil {
			rErr, ok := err.(azblob.StorageError)
			if !ok || rErr.ServiceCode() != azblob.ServiceCodeContainerAlreadyExists {
				log.Fatal(errors.Wrap(err, "failed to create container"))
			}
			log.Println("container already exist:", containerName)
		}
	}

	ctnrName := "0"
	blobName := "fileTest99999.ext"

	ctnr := svc.NewContainerURL(ctnrName)
	blobURL := ctnr.NewBlockBlobURL(blobName)

	props, err := blobURL.GetProperties(context.Background(), azblob.BlobAccessConditions{})
	if err != nil {
		rErr, ok := err.(azblob.ResponseError)
		if !ok || rErr.Response().StatusCode != http.StatusNotFound {
			log.Fatal(errors.Wrap(err, "failed to get blob properties"))
		}
		log.Println("blob already exist:", blobName)
	}

	// Upload only if file is not exist
	if props == nil {
		_, err := blobURL.Upload(context.Background(), bytes.NewReader([]byte("hw 8!")), azblob.BlobHTTPHeaders{}, azblob.Metadata{}, azblob.BlobAccessConditions{})
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to upload blob"))
		}
		fmt.Println("new blob was created")
	}

	u, err := url.Parse(fmt.Sprintf("https://%s.azureedge.net/%s/%s", accountName, ctnrName, blobName))
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to parse url"))
	}

	u.RawQuery = buildToken(u.Path)

	fmt.Println(u.String())
	rsp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to execute Get request"))
	}
	defer rsp.Body.Close()

	fmt.Println(rsp.StatusCode)

	_, err = io.Copy(os.Stdout, rsp.Body)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to copy file body to stdout"))
	}
	fmt.Println("Done.")
}

// buildToken will return auth token for specified path that will be valid for 4 hours
func buildToken(path string) string {
	return ectoken.EncryptV3(ecKey, fmt.Sprintf("ec_expire=%d&ec_url_allow=%s&ec_proto_allow=%s", time.Now().Add(time.Hour*4).Unix(), path, "https"))
}
