package main_test

import (
	"encoding/json"
	"github.com/anurocks1/goformation/v4/cloudformation/rds"

	"github.com/anurocks1/goformation/v4/cloudformation/ec2"
	"github.com/anurocks1/goformation/v4/cloudformation/s3"
	"github.com/anurocks1/goformation/v4/cloudformation/serverless"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resource", func() {

	Context("with a resource that contains a mix of primitive, custom and polymorphic properties", func() {

		Context("described as Go structs", func() {

			Context("with a simple primitive used for a polymorphic property", func() {

				codeuri := "s3://bucket/key"
				resource := &serverless.Function{
					Runtime: "nodejs6.10",
					CodeUri: &serverless.Function_CodeUri{
						String: &codeuri,
					},
				}

				expected := []byte(`{"Type":"AWS::Serverless::Function","Properties":{"CodeUri":"s3://bucket/key","Runtime":"nodejs6.10"}}`)

				result, err := json.Marshal(resource)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("with a custom type used for a polymorphic property", func() {

				resource := &serverless.Function{
					Runtime: "nodejs6.10",
					CodeUri: &serverless.Function_CodeUri{
						S3Location: &serverless.Function_S3Location{
							Bucket:  "test-bucket",
							Key:     "test-key",
							Version: 123,
						},
					},
				}

				expected := []byte(`{"Type":"AWS::Serverless::Function","Properties":{"CodeUri":{"Bucket":"test-bucket","Key":"test-key","Version":123},"Runtime":"nodejs6.10"}}`)

				result, err := json.Marshal(resource)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

		})

	})

	Context("with a resource that has some resource attributes defined", func() {

		Context("described as Go structs", func() {

			Context("with a dependency on another resource", func() {

				resource := &ec2.Instance{
					ImageId: "ami-0123456789",
				}
				resource.AWSCloudFormationDependsOn = []string{"MyDependency"}

				expected := []byte(`{"Type":"AWS::EC2::Instance","Properties":{"ImageId":"ami-0123456789"},"DependsOn":["MyDependency"]}`)

				result, err := json.Marshal(resource)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("with a metadata attribute", func() {

				resource := &s3.Bucket{
					BucketName: "MyBucket",
				}
				resource.AWSCloudFormationMetadata = map[string]interface{}{"Object1": "Location1", "Object2": "Location2"}

				expected := []byte(`{"Type":"AWS::S3::Bucket","Properties":{"BucketName":"MyBucket"},"Metadata":{"Object1":"Location1","Object2":"Location2"}}`)

				result, err := json.Marshal(resource)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("with a condition attribute", func() {

				resource := &rds.DBCluster{
					DatabaseName: "MyDatabase",
				}
				resource.AWSCloudFormationCondition = "MyCondition"

				expected := []byte(`{"Type":"AWS::RDS::DBCluster","Properties":{"DatabaseName":"MyDatabase"},"Condition":"MyCondition"}`)

				result, err := json.Marshal(resource)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

		})

		Context("specified as JSON", func() {

			Context("with a dependency on another resource", func() {

				property := []byte(`{"Type":"AWS::EC2::Instance","Properties":{"ImageId":"ami-0123456789"},"DependsOn":["MyDependency"]}`)
				expected := &ec2.Instance{
					ImageId: "ami-0123456789",
				}
				expected.AWSCloudFormationDependsOn = []string{"MyDependency"}

				result := &ec2.Instance{}
				err := json.Unmarshal(property, result)
				It("should unmarshal to a Go struct successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("with a metadata attribute", func() {

				property := []byte(`{"Type":"AWS::S3::Bucket","Properties":{"BucketName":"MyBucket"},"Metadata":{"Object1":"Location1","Object2":"Location2"}}`)
				expected := &s3.Bucket{
					BucketName: "MyBucket",
				}
				expected.AWSCloudFormationMetadata = map[string]interface{}{"Object1": "Location1", "Object2": "Location2"}

				result := &s3.Bucket{}
				err := json.Unmarshal(property, result)
				It("should unmarshal to a Go struct successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("with a condition attribute", func() {

				property := []byte(`{"Type":"AWS::RDS::DBCluster","Properties":{"DatabaseName":"MyDatabase"},"Condition":"MyCondition"}`)
				expected := &rds.DBCluster{
					DatabaseName: "MyDatabase",
				}
				expected.AWSCloudFormationCondition = "MyCondition"

				result := &rds.DBCluster{}
				err := json.Unmarshal(property, result)
				It("should unmarshal to a Go struct successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

		})

	})

	Context("with a custom property type resource", func() {

		Context("described as Go structs", func() {

			Context("with a list type", func() {

				subproperty := &serverless.Function_S3Event{
					Bucket: "my-bucket",
					Events: &serverless.Function_Events{
						StringArray: &[]string{"s3:ObjectCreated:*", "s3:ObjectRemoved:*"},
					},
				}

				expected := []byte(`{"Bucket":"my-bucket","Events":["s3:ObjectCreated:*","s3:ObjectRemoved:*"]}`)

				result, err := json.Marshal(subproperty)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

			Context("with a primitive type", func() {

				event := "s3:ObjectCreated:*"
				subproperty := &serverless.Function_S3Event{
					Bucket: "my-bucket",
					Events: &serverless.Function_Events{
						String: &event,
					},
				}

				expected := []byte(`{"Bucket":"my-bucket","Events":"s3:ObjectCreated:*"}`)

				result, err := json.Marshal(subproperty)
				It("should marshal to JSON successfully", func() {
					Expect(result).To(Equal(expected))
					Expect(err).To(BeNil())
				})

			})

		})
	})
})
