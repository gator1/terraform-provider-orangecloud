package hwcloud

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

// SKIP deprecated
func TestAccLBV1Pool_importBasic(t *testing.T) {
	resourceName := "hwcloud_lb_pool_v1.pool_1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheckDeprecated(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckLBV1PoolDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccLBV1Pool_basic,
			},

			resource.TestStep{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}