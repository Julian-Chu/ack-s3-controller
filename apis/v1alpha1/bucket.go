// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BucketSpec defines the desired state of Bucket
type BucketSpec struct {
	// The canned ACL to apply to the bucket.
	ACL *string `json:"acl,omitempty"`
	// The configuration information for the bucket.
	CreateBucketConfiguration *CreateBucketConfiguration `json:"createBucketConfiguration,omitempty"`
	// Allows grantee the read, write, read ACP, and write ACP permissions on the
	// bucket.
	GrantFullControl *string `json:"grantFullControl,omitempty"`
	// Allows grantee to list the objects in the bucket.
	GrantRead *string `json:"grantRead,omitempty"`
	// Allows grantee to read the bucket ACL.
	GrantReadACP *string `json:"grantReadACP,omitempty"`
	// Allows grantee to create, overwrite, and delete any object in the bucket.
	GrantWrite *string `json:"grantWrite,omitempty"`
	// Allows grantee to write the ACL for the applicable bucket.
	GrantWriteACP *string `json:"grantWriteACP,omitempty"`
	// The name of the bucket to create.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// Specifies whether you want S3 Object Lock to be enabled for the new bucket.
	ObjectLockEnabledForBucket *bool `json:"objectLockEnabledForBucket,omitempty"`
}

// BucketStatus defines the observed state of Bucket
type BucketStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// Specifies the Region where the bucket will be created. If you are creating
	// a bucket on the US East (N. Virginia) Region (us-east-1), you do not need
	// to specify the location.
	Location *string `json:"location,omitempty"`
}

// Bucket is the Schema for the Buckets API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Bucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BucketSpec   `json:"spec,omitempty"`
	Status            BucketStatus `json:"status,omitempty"`
}

// BucketList contains a list of Bucket
// +kubebuilder:object:root=true
type BucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Bucket `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Bucket{}, &BucketList{})
}
