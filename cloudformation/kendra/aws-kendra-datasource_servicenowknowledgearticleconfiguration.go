package kendra

import (
	"github.com/dexty007/goformation/v4/cloudformation/policies"
)

// DataSource_ServiceNowKnowledgeArticleConfiguration AWS CloudFormation Resource (AWS::Kendra::DataSource.ServiceNowKnowledgeArticleConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowknowledgearticleconfiguration.html
type DataSource_ServiceNowKnowledgeArticleConfiguration struct {

	// CrawlAttachments AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowknowledgearticleconfiguration.html#cfn-kendra-datasource-servicenowknowledgearticleconfiguration-crawlattachments
	CrawlAttachments bool `json:"CrawlAttachments,omitempty"`

	// DocumentDataFieldName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowknowledgearticleconfiguration.html#cfn-kendra-datasource-servicenowknowledgearticleconfiguration-documentdatafieldname
	DocumentDataFieldName string `json:"DocumentDataFieldName,omitempty"`

	// DocumentTitleFieldName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowknowledgearticleconfiguration.html#cfn-kendra-datasource-servicenowknowledgearticleconfiguration-documenttitlefieldname
	DocumentTitleFieldName string `json:"DocumentTitleFieldName,omitempty"`

	// ExcludeAttachmentFilePatterns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowknowledgearticleconfiguration.html#cfn-kendra-datasource-servicenowknowledgearticleconfiguration-excludeattachmentfilepatterns
	ExcludeAttachmentFilePatterns []string `json:"ExcludeAttachmentFilePatterns,omitempty"`

	// FieldMappings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowknowledgearticleconfiguration.html#cfn-kendra-datasource-servicenowknowledgearticleconfiguration-fieldmappings
	FieldMappings []DataSource_DataSourceToIndexFieldMapping `json:"FieldMappings,omitempty"`

	// IncludeAttachmentFilePatterns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-servicenowknowledgearticleconfiguration.html#cfn-kendra-datasource-servicenowknowledgearticleconfiguration-includeattachmentfilepatterns
	IncludeAttachmentFilePatterns []string `json:"IncludeAttachmentFilePatterns,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *DataSource_ServiceNowKnowledgeArticleConfiguration) AWSCloudFormationType() string {
	return "AWS::Kendra::DataSource.ServiceNowKnowledgeArticleConfiguration"
}
