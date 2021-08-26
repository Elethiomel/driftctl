package aws_test

import (
	"testing"

	"github.com/cloudskiff/driftctl/test"

	"github.com/cloudskiff/driftctl/test/acceptance"
)

func TestAcc_AwsVPC(t *testing.T) {
	acceptance.Run(t, acceptance.AccTestCase{
		TerraformVersion: "0.15.5",
		Paths:            []string{"./testdata/acc/aws_vpc"},
		Args:             []string{"scan", "--filter", "Type=='aws_vpc'", "--deep"},
		Checks: []acceptance.AccCheck{
			{
				Env: map[string]string{
					"AWS_REGION": "us-east-1",
				},
				Check: func(result *test.ScanResult, stdout string, err error) {
					if err != nil {
						t.Fatal(err)
					}
					result.AssertInfrastructureIsInSync()
					result.AssertManagedCount(3)
				},
			},
		},
	})
}
