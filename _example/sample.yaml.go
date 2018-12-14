package main

import (
	"github.com/the-redback/go-oneliners"
)

func main() {
	v := map[string]interface{}{
		"str":   "foo",
		"num":   100,
		"bool":  false,
		"null":  nil,
		"array": []string{"foo", "<bar>", "baz"},
		"map": map[string]interface{}{
			"foo": "bar",
		},
	}
	oneliners.PrettyJson(v)


	str := map[string]interface{}{
		"Name": "Wednesday",
		"Age":  6,
		"Parents": []interface{}{
			"Gomez",
			"Morticia",
		},
	}
	oneliners.PrettyJson(str, "Asad")

	mongodb:=[]byte(`{ "metadata": {
    "creationTimestamp": "2018-12-13T06:24:37Z",
    "finalizers": [
      "kubedb.com"
    ],
    "generation": 1,
    "labels": {
      "app": "mongodb-e2e-qf6nso"
    },
    "name": "mongodb-rs-4r7yt6",
    "namespace": "mongodb-43aixt",
    "resourceVersion": "13215",
    "selfLink": "/apis/kubedb.com/v1alpha1/namespaces/mongodb-43aixt/mongodbs/mongodb-rs-4r7yt6",
    "uid": "c3aae769-fe9f-11e8-847a-080027ef36b1"
  },
  "spec": {
    "podTemplate": {
      "controller": {},
      "metadata": {},
      "spec": {
        "livenessProbe": {
          "exec": {
            "command": [
              "mongo",
              "--eval",
              "db.adminCommand('ping')"
            ]
          },
          "failureThreshold": 3,
          "periodSeconds": 10,
          "successThreshold": 1,
          "timeoutSeconds": 5
        },
        "readinessProbe": {
          "exec": {
            "command": [
              "mongo",
              "--eval",
              "db.adminCommand('ping')"
            ]
          },
          "failureThreshold": 3,
          "periodSeconds": 10,
          "successThreshold": 1,
          "timeoutSeconds": 1
        },
        "resources": {}
      }
    },
    "replicaSet": {
      "name": "mongodb-rs-4r7yt6"
    },
    "replicas": 3,
    "serviceTemplate": {
      "metadata": {},
      "spec": {}
    },
    "storage": {
      "dataSource": null,
      "resources": {
        "requests": {
          "storage": "1Gi"
        }
      },
      "storageClassName": "standard"
    },
    "storageType": "Durable",
    "terminationPolicy": "Pause",
    "updateStrategy": {
      "type": "RollingUpdate"
    },
    "version": "3.6-v1"
  },
  "status": {}
}`)
	oneliners.PrettyJson(mongodb,"mongodb")

}
