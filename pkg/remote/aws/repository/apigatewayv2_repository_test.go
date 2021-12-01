package repository

import (
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigatewayv2"
	"github.com/cloudskiff/driftctl/pkg/remote/cache"
	awstest "github.com/cloudskiff/driftctl/test/aws"
	"github.com/pkg/errors"

	"github.com/r3labs/diff/v2"
	"github.com/stretchr/testify/assert"
)

func Test_apigatewayv2Repository_ListAllApis(t *testing.T) {
	apis := []*apigatewayv2.Api{
		{ApiId: aws.String("api1")},
		{ApiId: aws.String("api2")},
		{ApiId: aws.String("api3")},
		{ApiId: aws.String("api4")},
		{ApiId: aws.String("api5")},
		{ApiId: aws.String("api6")},
	}

	remoteError := errors.New("remote error")

	tests := []struct {
		name    string
		mocks   func(client *awstest.MockFakeApiGatewayV2, store *cache.MockCache)
		want    []*apigatewayv2.Api
		wantErr error
	}{
		{
			name: "list multiple apis",
			mocks: func(client *awstest.MockFakeApiGatewayV2, store *cache.MockCache) {
				client.On("GetApis",
					&apigatewayv2.GetApisInput{}).Return(&apigatewayv2.GetApisOutput{Items: apis}, nil).Once()

				store.On("GetAndLock", "apigatewayv2ListAllApis").Return(nil).Times(1)
				store.On("Unlock", "apigatewayv2ListAllApis").Times(1)
				store.On("Put", "apigatewayv2ListAllApis", apis).Return(false).Times(1)
			},
			want: apis,
		},
		{
			name: "should hit cache",
			mocks: func(client *awstest.MockFakeApiGatewayV2, store *cache.MockCache) {
				store.On("GetAndLock", "apigatewayv2ListAllApis").Return(apis).Times(1)
				store.On("Unlock", "apigatewayv2ListAllApis").Times(1)
			},
			want: apis,
		},
		{
			name: "should return remote error",
			mocks: func(client *awstest.MockFakeApiGatewayV2, store *cache.MockCache) {
				client.On("GetApis",
					&apigatewayv2.GetApisInput{}).Return(nil, remoteError).Once()

				store.On("GetAndLock", "apigatewayv2ListAllApis").Return(nil).Times(1)
				store.On("Unlock", "apigatewayv2ListAllApis").Times(1)
			},
			wantErr: remoteError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &cache.MockCache{}
			client := &awstest.MockFakeApiGatewayV2{}
			tt.mocks(client, store)
			r := &apigatewayv2Repository{
				client: client,
				cache:  store,
			}
			got, err := r.ListAllApis()
			assert.Equal(t, tt.wantErr, err)

			changelog, err := diff.Diff(got, tt.want)
			assert.Nil(t, err)
			if len(changelog) > 0 {
				for _, change := range changelog {
					t.Errorf("%s: %s -> %s", strings.Join(change.Path, "."), change.From, change.To)
				}
				t.Fail()
			}
			store.AssertExpectations(t)
			client.AssertExpectations(t)
		})
	}
}

func Test_apigatewayv2Repository_ListAllVpcLinks(t *testing.T) {
	vpcLinks := []*apigatewayv2.VpcLink{
		{VpcLinkId: aws.String("vpcLink1")},
		{VpcLinkId: aws.String("vpcLink2")},
		{VpcLinkId: aws.String("vpcLink3")},
		{VpcLinkId: aws.String("vpcLink4")},
		{VpcLinkId: aws.String("vpcLink5")},
		{VpcLinkId: aws.String("vpcLink6")},
	}

	remoteError := errors.New("remote error")

	tests := []struct {
		name    string
		mocks   func(client *awstest.MockFakeApiGatewayV2, store *cache.MockCache)
		want    []*apigatewayv2.VpcLink
		wantErr error
	}{
		{
			name: "list multiple vpc links",
			mocks: func(client *awstest.MockFakeApiGatewayV2, store *cache.MockCache) {
				client.On("GetVpcLinks",
					&apigatewayv2.GetVpcLinksInput{}).Return(&apigatewayv2.GetVpcLinksOutput{Items: vpcLinks}, nil).Once()

				store.On("Get", "apigatewayv2ListAllVpcLinks").Return(nil).Times(1)
				store.On("Put", "apigatewayv2ListAllVpcLinks", vpcLinks).Return(false).Times(1)
			},
			want: vpcLinks,
		},
		{
			name: "should hit cache",
			mocks: func(client *awstest.MockFakeApiGatewayV2, store *cache.MockCache) {
				store.On("Get", "apigatewayv2ListAllVpcLinks").Return(vpcLinks).Times(1)
			},
			want: vpcLinks,
		},
		{
			name: "should return remote error",
			mocks: func(client *awstest.MockFakeApiGatewayV2, store *cache.MockCache) {
				client.On("GetVpcLinks",
					&apigatewayv2.GetVpcLinksInput{}).Return(nil, remoteError).Once()

				store.On("Get", "apigatewayv2ListAllVpcLinks").Return(nil).Times(1)
			},
			wantErr: remoteError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &cache.MockCache{}
			client := &awstest.MockFakeApiGatewayV2{}
			tt.mocks(client, store)
			r := &apigatewayv2Repository{
				client: client,
				cache:  store,
			}
			got, err := r.ListAllVpcLinks()
			assert.Equal(t, tt.wantErr, err)

			changelog, err := diff.Diff(got, tt.want)
			assert.Nil(t, err)
			if len(changelog) > 0 {
				for _, change := range changelog {
					t.Errorf("%s: %s -> %s", strings.Join(change.Path, "."), change.From, change.To)
				}
				t.Fail()
			}
			store.AssertExpectations(t)
			client.AssertExpectations(t)
		})
	}
}

func Test_apigatewayv2Repository_ListAllApiAuthorizers(t *testing.T) {
	api := &apigatewayv2.Api{
		ApiId: aws.String("api1"),
	}

	apiAuthorizers := []*apigatewayv2.Authorizer{
		{AuthorizerId: aws.String("authorizer1")},
		{AuthorizerId: aws.String("authorizer2")},
		{AuthorizerId: aws.String("authorizer3")},
		{AuthorizerId: aws.String("authorizer4")},
	}

	remoteError := errors.New("remote error")

	tests := []struct {
		name    string
		mocks   func(client *awstest.MockFakeApiGatewayV2, store *cache.MockCache)
		want    []*apigatewayv2.Authorizer
		wantErr error
	}{
		{
			name: "list multiple api authorizers",
			mocks: func(client *awstest.MockFakeApiGatewayV2, store *cache.MockCache) {
				client.On("GetAuthorizers",
					&apigatewayv2.GetAuthorizersInput{
						ApiId: aws.String("api1"),
					}).Return(&apigatewayv2.GetAuthorizersOutput{Items: apiAuthorizers}, nil).Once()

				store.On("Get", "apigatewayv2ListAllApiAuthorizers_api_api1").Return(nil).Times(1)
				store.On("Put", "apigatewayv2ListAllApiAuthorizers_api_api1", apiAuthorizers).Return(false).Times(1)
			},
			want: apiAuthorizers,
		},
		{
			name: "should hit cache",
			mocks: func(client *awstest.MockFakeApiGatewayV2, store *cache.MockCache) {
				store.On("Get", "apigatewayv2ListAllApiAuthorizers_api_api1").Return(apiAuthorizers).Times(1)
			},
			want: apiAuthorizers,
		},
		{
			name: "should return remote error",
			mocks: func(client *awstest.MockFakeApiGatewayV2, store *cache.MockCache) {
				client.On("GetAuthorizers",
					&apigatewayv2.GetAuthorizersInput{
						ApiId: aws.String("api1"),
					}).Return(nil, remoteError).Once()

				store.On("Get", "apigatewayv2ListAllApiAuthorizers_api_api1").Return(nil).Times(1)
			},
			wantErr: remoteError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := &cache.MockCache{}
			client := &awstest.MockFakeApiGatewayV2{}
			tt.mocks(client, store)
			r := &apigatewayv2Repository{
				client: client,
				cache:  store,
			}
			got, err := r.ListAllApiAuthorizers(*api.ApiId)
			assert.Equal(t, tt.wantErr, err)

			changelog, err := diff.Diff(got, tt.want)
			assert.Nil(t, err)
			if len(changelog) > 0 {
				for _, change := range changelog {
					t.Errorf("%s: %s -> %s", strings.Join(change.Path, "."), change.From, change.To)
				}
				t.Fail()
			}
			store.AssertExpectations(t)
			client.AssertExpectations(t)
		})
	}
}
