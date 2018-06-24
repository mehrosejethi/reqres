package reqres

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Default:     "",
				Description: descriptions["token"],
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"reqres_data_source_users_per_page": dataSourceUsersPerPage(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"reqres": resourceAddOrganisation(),
		},
	}
}
