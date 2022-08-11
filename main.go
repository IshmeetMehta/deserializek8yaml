package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	k8Yaml "k8s.io/apimachinery/pkg/util/yaml"
	knative "knative.dev/serving/pkg/apis/serving/v1"
)

func main() {
	
	rev := &knative.Revision{}
	// spec := &knative.spec{}
	stream, err := ioutil.ReadFile("knative.yaml")
	if err != nil {
		fmt.Println("Cannot open the file", err)
	}
	dec := k8Yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(stream)), 1000)

	if err := dec.Decode(&rev); err != nil {
		fmt.Printf("error decoding the yaml: %v", err)
	}

	fmt.Printf("%+v", rev)
}
