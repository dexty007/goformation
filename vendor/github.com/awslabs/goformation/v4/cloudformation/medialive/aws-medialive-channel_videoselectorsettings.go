package medialive

import (
	"github.com/anurocks1/goformation/v4/cloudformation/policies"
)

// Channel_VideoSelectorSettings AWS CloudFormation Resource (AWS::MediaLive::Channel.VideoSelectorSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselectorsettings.html
type Channel_VideoSelectorSettings struct {

	// VideoSelectorPid AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselectorsettings.html#cfn-medialive-channel-videoselectorsettings-videoselectorpid
	VideoSelectorPid *Channel_VideoSelectorPid `json:"VideoSelectorPid,omitempty"`

	// VideoSelectorProgramId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-medialive-channel-videoselectorsettings.html#cfn-medialive-channel-videoselectorsettings-videoselectorprogramid
	VideoSelectorProgramId *Channel_VideoSelectorProgramId `json:"VideoSelectorProgramId,omitempty"`

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
func (r *Channel_VideoSelectorSettings) AWSCloudFormationType() string {
	return "AWS::MediaLive::Channel.VideoSelectorSettings"
}
