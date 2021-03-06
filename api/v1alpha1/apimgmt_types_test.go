// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// These tests are written in BDD-style using Ginkgo framework. Refer to
// http://onsi.github.io/ginkgo to learn more.

var _ = Describe("AzureAPIManagement", func() {
	var (
		key              types.NamespacedName
		created, fetched *APIMgmtAPI
	)

	Context("Create AzureAPIManagement type", func() {

		It("should create and delete an APIManagement type successfully", func() {

			key = types.NamespacedName{
				Name:      "foo",
				Namespace: "default",
			}
			created = &APIMgmtAPI{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				},
				Spec: APIMgmtSpec{
					Location:      "westus",
					ResourceGroup: "foo-resourcegroup",
					Properties: APIProperties{
						APIVersion: "alphav1",
					},
				}}
			By("creating an APIManagement type")
			Expect(k8sClient.Create(context.TODO(), created)).To(Succeed())

			fetched = &APIMgmtAPI{}
			Expect(k8sClient.Get(context.TODO(), key, fetched)).To(Succeed())
			Expect(fetched).To(Equal(created))

			By("deleting the APIManagement")
			Expect(k8sClient.Delete(context.TODO(), created)).To(Succeed())
			Expect(k8sClient.Get(context.TODO(), key, created)).ToNot(Succeed())
		})

	})

})
