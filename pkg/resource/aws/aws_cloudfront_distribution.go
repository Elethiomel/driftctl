package aws

import (
	"github.com/cloudskiff/driftctl/pkg/resource"
)

const AwsCloudfrontDistributionResourceType = "aws_cloudfront_distribution"

func initAwsCloudfrontDistributionMetaData(resourceSchemaRepository resource.SchemaRepositoryInterface) {
	resourceSchemaRepository.SetNormalizeFunc(AwsCloudfrontDistributionResourceType, func(res *resource.Resource) {
		val := res.Attrs
		val.SafeDelete([]string{"etag"})
		val.SafeDelete([]string{"last_modified_time"})
		val.SafeDelete([]string{"retain_on_delete"})
		val.SafeDelete([]string{"status"})
		val.SafeDelete([]string{"wait_for_deployment"})
	})
	resourceSchemaRepository.SetFlags(AwsCloudfrontDistributionResourceType, resource.FlagDeepMode)
}
