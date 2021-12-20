package apiserver

import (
	"context"

	rndv1 "github.com/glrf/kube-rnd-api/api/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/runtime"
	genericregistry "k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/registry/rest"
	restbuilder "sigs.k8s.io/apiserver-runtime/pkg/builder/rest"
)

type numberStorage struct {
}

func (r *numberStorage) New() runtime.Object {
	return &rndv1.Number{}
}

func NewNumberStorageProvider() restbuilder.ResourceHandlerProvider {
	return func(s *runtime.Scheme, g genericregistry.RESTOptionsGetter) (rest.Storage, error) {
		return &numberStorage{}, nil
	}
}

var _ rest.Scoper = &numberStorage{}

func (r *numberStorage) NamespaceScoped() bool {
	return false
}

var _ rest.Getter = &numberStorage{}

func (r *numberStorage) Get(ctx context.Context, name string, options *metav1.GetOptions) (runtime.Object, error) {
  return &rndv1.Number{
  	TypeMeta:   metav1.TypeMeta{
  		Kind:       options.Kind,
  		APIVersion: options.APIVersion,
  	},
  	ObjectMeta: metav1.ObjectMeta{Name: name},
  	Spec:       rndv1.NumberSpec{Int: 4},
  }, nil
}
