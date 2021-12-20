package main

import (
	"log"

	"sigs.k8s.io/apiserver-runtime/pkg/builder"

	rndv1 "github.com/glrf/kube-rnd-api/api/v1"
	rndapiserver "github.com/glrf/kube-rnd-api/apiserver"
)

func main() {
	err := builder.APIServer.
		WithResourceAndHandler(&rndv1.Number{}, rndapiserver.NewNumberStorageProvider()).
		//WithResourceAndHandler(&rndv1.Number{},filepath.NewJSONFilepathStorageProvider(&rndv1.Number{}, "data")).
		WithLocalDebugExtension().
		WithoutEtcd().
		Execute()
	if err != nil {
		log.Fatal(err)
	}
}
