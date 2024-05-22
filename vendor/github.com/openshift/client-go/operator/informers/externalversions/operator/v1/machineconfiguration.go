// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	operatorv1 "github.com/openshift/api/operator/v1"
	versioned "github.com/openshift/client-go/operator/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/operator/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/operator/listers/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MachineConfigurationInformer provides access to a shared informer and lister for
// MachineConfigurations.
type MachineConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.MachineConfigurationLister
}

type machineConfigurationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewMachineConfigurationInformer constructs a new informer for MachineConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMachineConfigurationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMachineConfigurationInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredMachineConfigurationInformer constructs a new informer for MachineConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMachineConfigurationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().MachineConfigurations().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().MachineConfigurations().Watch(context.TODO(), options)
			},
		},
		&operatorv1.MachineConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *machineConfigurationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMachineConfigurationInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *machineConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operatorv1.MachineConfiguration{}, f.defaultInformer)
}

func (f *machineConfigurationInformer) Lister() v1.MachineConfigurationLister {
	return v1.NewMachineConfigurationLister(f.Informer().GetIndexer())
}