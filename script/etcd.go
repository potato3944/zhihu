package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// 定义命令行参数，方便灵活调整
	endpoints := flag.String("endpoints", "43.166.137.162:2379", "etcd endpoints, comma separated")
	username := flag.String("username", "root", "etcd username")
	password := flag.String("password", "my_strong_password", "etcd password")
	key := flag.String("key", "service.rpc", "the key to query")
	isPrefix := flag.Bool("prefix", false, "query keys with prefix")
	timeout := flag.Duration("timeout", 5*time.Second, "connection timeout")

	flag.Parse()

	log.Printf("Connecting to etcd endpoints: %s (Username: %s)...", *endpoints, *username)

	// 配置 etcd 客户端
	cfg := clientv3.Config{
		Endpoints:   []string{*endpoints},
		DialTimeout: *timeout,
		Username:    *username,
		Password:    *password,
	}

	// 初始化客户端
	cli, err := clientv3.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	}
	defer cli.Close()

	// 创建带超时的 Context
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	log.Printf("Getting key: %q (isPrefix: %t)...", *key, *isPrefix)

	// 查询选项
	var opts []clientv3.OpOption
	if *isPrefix {
		opts = append(opts, clientv3.WithPrefix())
	}

	// 获取键值
	resp, err := cli.Get(ctx, *key, opts...)
	if err != nil {
		log.Fatalf("Get key failed: %v", err)
	}

	if len(resp.Kvs) == 0 {
		fmt.Printf("No keys found for %q\n", *key)
		return
	}

	// 打印查询结果
	fmt.Printf("\n--- Found %d key-value pair(s) ---\n", len(resp.Kvs))
	for _, kv := range resp.Kvs {
		fmt.Printf("Key:   %s\nValue: %s\n----------------------------------\n", string(kv.Key), string(kv.Value))
	}
}
