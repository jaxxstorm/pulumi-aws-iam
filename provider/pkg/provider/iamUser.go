// Copyright 2016-2021, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The set of arguments for creating a StaticPage component resource.
type IamUserArgs struct {
	Path string
}

// The IamUser component resource.
type IamUser struct {
	pulumi.ResourceState


}

// NewStaticPage creates a new StaticPage component resource.
func NewIamUser(ctx *pulumi.Context,
	name string, args *IamUserArgs, opts ...pulumi.ResourceOption) (*IamUser, error) {
	if args == nil {
		args = &IamUserArgs{}
	}

	component := &IamUser{}
	err := ctx.RegisterComponentResource("aws:index:IamUser", name, component, opts...)
	if err != nil {
		return nil, err
	}

	user, err := iam.NewUser(ctx, name, &iam.UserArgs{
		Path: ,
	})

	if err := ctx.RegisterResourceOutputs(component, pulumi.Map{
	}); err != nil {
		return nil, err
	}

	return component, nil
}
