// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package config

import "github.com/zeromicro/go-zero/rest"

//nolint:staticcheck
type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	Oss struct {
		Endpoint         string
		AccessKeyId      string
		AccessKeySecret  string
		BucketName       string
		ConnectTimeout   int64 
		ReadWriteTimeout int64 
	}
}