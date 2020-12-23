package model

import (
	"context"
	"mime/multipart"

	"github.com/panda8z/eeyoo/utils"
	"github.com/panda8z/eeyoo/utils/errors"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

var AccessKey = utils.AccessKey
var SecretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.ImgUrl

func UploadFile(file multipart.File, fileSize int64) (string, int) {
	// new policy obj
	putPolicy := storage.PutPolicy{
		Scope: Bucket,
	}
	// new mac obj
	mac := qbox.NewMac(AccessKey, SecretKey)
	// get token string
	upToken := putPolicy.UploadToken(mac)

	// new storage config obj
	conf := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	// new extra info obj
	putExtra := storage.PutExtra{}

	// new uploader obj
	formUploader := storage.NewFormUploader(&conf)

	// define ret obj
	ret := storage.PutRet{}

	// do upload
	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errors.ERROR_FILE_UPLOAD
	}

	// receive the key
	url := ImgUrl + ret.Key
	return url, errors.SUCCESS
}
