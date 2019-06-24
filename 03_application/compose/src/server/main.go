package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/coreos/etcd/clientv3"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			list, err := kvs()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res, err := json.Marshal(list)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
		} else if r.Method == http.MethodPost {
			body, err := ioutil.ReadAll(r.Body)
			r.Body.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var msg map[string]interface{}
			err = json.Unmarshal(body, &msg)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			for k, v := range msg {
				err = addValue(k, v.(string))
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			io.WriteString(w, "OK")
		} else {
			http.Error(w, "unknown method", http.StatusNotFound)
			return
		}
	})
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println(err)
	}
}

var cfg = clientv3.Config{
	Endpoints:   []string{"http://etcd:2379"},
	DialTimeout: 3 * time.Second,
}

func addValue(key, value string) error {
	client, err := clientv3.New(cfg)
	if err != nil {
		return err
	}
	defer client.Close()

	_, err = client.Put(context.Background(), key, value)
	return err
}

func kvs() ([]string, error) {
	client, err := clientv3.New(cfg)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	resp, err := client.Get(context.Background(), "", clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	result := make([]string, len(resp.Kvs))
	for i, kv := range resp.Kvs {
		result[i] = fmt.Sprintf("%s:%s", kv.Key, kv.Value)
	}
	return result, nil
}
