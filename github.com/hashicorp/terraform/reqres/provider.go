package reqres

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{},

		DataSourcesMap: map[string]*schema.Resource{
			"reqres_users_per_page": dataSourceUsersPerPage(),
			"reqres_single_user":    dataSourceSingleUser(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"reqres": resourceAddOrganisation(),
		},
	}
}
