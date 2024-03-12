package qiniu

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
	"os"
	"resourceManager/rpc/internal/config"
)

func GetFile(path string) (file multipart.File, fileSize int64, err error) {
	file, err = os.Open(path)
	if err != nil {
		return nil, 0, err
	}

	fileInfo, err := os.Stat(path)

	if err != nil {
		return nil, 0, err
	}

	return file, fileInfo.Size(), nil
}

func UploadToQiNi(conf config.QiNiuConf, file multipart.File, fileSize int64) (url string, err error) {
	accessKey := conf.AccessKey
	secretKey := conf.SecretKey
	bucket := conf.Bucket
	imgUrl := conf.QiniuServer

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	//根据秘钥生成上传凭证
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	//定义上传内容的配置，区域是否使用cdn以及https
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	//根据配置生成formuploader对象，并且使用putwithoutkey上传文件
	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)

	if err != nil {
		return "", err
	}

	url = "http:/" + imgUrl + "/" + ret.Key

	return url, nil

}
