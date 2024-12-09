/*
Copyright 2018 The Rook Authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	scheme "github.com/rook/rook/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// CephFilesystemMirrorsGetter has a method to return a CephFilesystemMirrorInterface.
// A group's client should implement this interface.
type CephFilesystemMirrorsGetter interface {
	CephFilesystemMirrors(namespace string) CephFilesystemMirrorInterface
}

// CephFilesystemMirrorInterface has methods to work with CephFilesystemMirror resources.
type CephFilesystemMirrorInterface interface {
	Create(ctx context.Context, cephFilesystemMirror *v1.CephFilesystemMirror, opts metav1.CreateOptions) (*v1.CephFilesystemMirror, error)
	Update(ctx context.Context, cephFilesystemMirror *v1.CephFilesystemMirror, opts metav1.UpdateOptions) (*v1.CephFilesystemMirror, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CephFilesystemMirror, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CephFilesystemMirrorList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CephFilesystemMirror, err error)
	CephFilesystemMirrorExpansion
}

// cephFilesystemMirrors implements CephFilesystemMirrorInterface
type cephFilesystemMirrors struct {
	*gentype.ClientWithList[*v1.CephFilesystemMirror, *v1.CephFilesystemMirrorList]
}

// newCephFilesystemMirrors returns a CephFilesystemMirrors
func newCephFilesystemMirrors(c *CephV1Client, namespace string) *cephFilesystemMirrors {
	return &cephFilesystemMirrors{
		gentype.NewClientWithList[*v1.CephFilesystemMirror, *v1.CephFilesystemMirrorList](
			"cephfilesystemmirrors",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.CephFilesystemMirror { return &v1.CephFilesystemMirror{} },
			func() *v1.CephFilesystemMirrorList { return &v1.CephFilesystemMirrorList{} }),
	}
}
