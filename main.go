package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/build-on-aws/aws-redis-iam-auth-golang/auth"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
)

var (
	clusterEndpoint string
	username        string
	client          *redis.ClusterClient
	region          string
	clusterName     string
	serviceName     string
)

const defaultRegion = "us-east-1"

func init() {

	serviceName = os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("SERVICE_NAME env var missing")
	}

	if serviceName != "elasticache" && serviceName != "memorydb" {
		log.Fatal("specify elasticache or memorydb as value for SERVICE_NAME env var")
	}

	clusterName = os.Getenv("CLUSTER_NAME")
	if clusterName == "" {
		log.Fatal("CLUSTER_NAME env var missing")
	}

	fmt.Println("cluster name", clusterName)

	clusterEndpoint = os.Getenv("CLUSTER_ENDPOINT")
	if clusterEndpoint == "" {
		log.Fatal("CLUSTER_ENDPOINT env var missing")
	}

	fmt.Println("cluster endpoint", clusterEndpoint)

	username = os.Getenv("USERNAME")
	if username == "" {
		log.Fatal("USERNAME env var missing")
	}

	fmt.Println("username", username)

	region = os.Getenv("AWS_REGION")
	if region == "" {
		region = defaultRegion
	}

	fmt.Println("connecting to cluster....", clusterEndpoint)

	generator, err := auth.New(serviceName, clusterName, username, region)
	if err != nil {
		log.Fatal("failed to initialise token generator", err)
	}

	client = redis.NewClusterClient(
		&redis.ClusterOptions{
			Username: username,
			Addrs:    []string{clusterEndpoint},
			NewClient: func(opt *redis.Options) *redis.Client {

				return redis.NewClient(&redis.Options{
					Addr: opt.Addr,
					CredentialsProvider: func() (username string, password string) {

						token, err := generator.Generate()
						if err != nil {
							log.Fatal("failed to generate auth token", err)
						}

						fmt.Println("auth token generated successfully")

						return opt.Username, token
					},
					TLSConfig: &tls.Config{InsecureSkipVerify: true},
				})
			},
		})

	err = client.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("failed to connect to memorydb -", err)
	}

	fmt.Println("successfully connected to cluster", clusterEndpoint)

}
func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", set).Methods(http.MethodPost)
	r.HandleFunc("/{key}", get).Methods(http.MethodGet)

	fmt.Println("started HTTP server....")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["key"]

	fmt.Println("getting value for", key)

	val, err := client.Get(context.Background(), key).Result()
	if err != nil {

		if errors.Is(err, redis.Nil) {
			w.WriteHeader(http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(KV{Key: key, Value: val})

}

func set(w http.ResponseWriter, req *http.Request) {
	var kv KV

	err := json.NewDecoder(req.Body).Decode(&kv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("setting %s=%s\n", kv.Key, kv.Value)

	err = client.Set(context.Background(), kv.Key, kv.Value, 0).Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("successfully set %s=%s\n", kv.Key, kv.Value)
}

type KV struct {
	Key   string
	Value string
}
