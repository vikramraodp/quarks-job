/*

Don't alter this file, it was generated.

*/
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "code.cloudfoundry.org/quarks-job/pkg/kube/apis/extendedjob/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ExtendedJobLister helps list ExtendedJobs.
type ExtendedJobLister interface {
	// List lists all ExtendedJobs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ExtendedJob, err error)
	// ExtendedJobs returns an object that can list and get ExtendedJobs.
	ExtendedJobs(namespace string) ExtendedJobNamespaceLister
	ExtendedJobListerExpansion
}

// extendedJobLister implements the ExtendedJobLister interface.
type extendedJobLister struct {
	indexer cache.Indexer
}

// NewExtendedJobLister returns a new ExtendedJobLister.
func NewExtendedJobLister(indexer cache.Indexer) ExtendedJobLister {
	return &extendedJobLister{indexer: indexer}
}

// List lists all ExtendedJobs in the indexer.
func (s *extendedJobLister) List(selector labels.Selector) (ret []*v1alpha1.ExtendedJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ExtendedJob))
	})
	return ret, err
}

// ExtendedJobs returns an object that can list and get ExtendedJobs.
func (s *extendedJobLister) ExtendedJobs(namespace string) ExtendedJobNamespaceLister {
	return extendedJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ExtendedJobNamespaceLister helps list and get ExtendedJobs.
type ExtendedJobNamespaceLister interface {
	// List lists all ExtendedJobs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ExtendedJob, err error)
	// Get retrieves the ExtendedJob from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ExtendedJob, error)
	ExtendedJobNamespaceListerExpansion
}

// extendedJobNamespaceLister implements the ExtendedJobNamespaceLister
// interface.
type extendedJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ExtendedJobs in the indexer for a given namespace.
func (s extendedJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ExtendedJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ExtendedJob))
	})
	return ret, err
}

// Get retrieves the ExtendedJob from the indexer for a given namespace and name.
func (s extendedJobNamespaceLister) Get(name string) (*v1alpha1.ExtendedJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("extendedjob"), name)
	}
	return obj.(*v1alpha1.ExtendedJob), nil
}