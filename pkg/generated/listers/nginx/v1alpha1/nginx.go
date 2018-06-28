// Code generated by lister-gen. DO NOT EDIT.

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/tsuru/nginx-operator/pkg/apis/nginx/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NginxLister helps list Nginxes.
type NginxLister interface {
	// List lists all Nginxes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Nginx, err error)
	// Nginxes returns an object that can list and get Nginxes.
	Nginxes(namespace string) NginxNamespaceLister
	NginxListerExpansion
}

// nginxLister implements the NginxLister interface.
type nginxLister struct {
	indexer cache.Indexer
}

// NewNginxLister returns a new NginxLister.
func NewNginxLister(indexer cache.Indexer) NginxLister {
	return &nginxLister{indexer: indexer}
}

// List lists all Nginxes in the indexer.
func (s *nginxLister) List(selector labels.Selector) (ret []*v1alpha1.Nginx, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Nginx))
	})
	return ret, err
}

// Nginxes returns an object that can list and get Nginxes.
func (s *nginxLister) Nginxes(namespace string) NginxNamespaceLister {
	return nginxNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NginxNamespaceLister helps list and get Nginxes.
type NginxNamespaceLister interface {
	// List lists all Nginxes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Nginx, err error)
	// Get retrieves the Nginx from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Nginx, error)
	NginxNamespaceListerExpansion
}

// nginxNamespaceLister implements the NginxNamespaceLister
// interface.
type nginxNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Nginxes in the indexer for a given namespace.
func (s nginxNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Nginx, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Nginx))
	})
	return ret, err
}

// Get retrieves the Nginx from the indexer for a given namespace and name.
func (s nginxNamespaceLister) Get(name string) (*v1alpha1.Nginx, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("nginx"), name)
	}
	return obj.(*v1alpha1.Nginx), nil
}