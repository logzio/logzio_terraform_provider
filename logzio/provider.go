package logzio

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

const (
	providerApiToken       = "api_token"
	providerBaseUrl        = "base_url"
	providerRegion         = "region"
	resourceAlertType      = "logzio_alert"
	resourceEndpointType   = "logzio_endpoint"
	resourceUserType       = "logzio_user"
	resourceSubAccountType = "logzio_subaccount"
	envLogzioApiToken      = "LOGZIO_API_TOKEN"
	envLogzioRegion        = "LOGZIO_REGION"
	envLogzioBaseURL       = "LOGZIO_BASE_URL"
	envLogzioAccountId     = "LOGZIO_ACCOUNT_ID"

	baseUrl = "https://api%s.logz.io"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			providerApiToken: {
				Type:        schema.TypeString,
				Required:    true,
				Description: descriptions[providerApiToken],
				DefaultFunc: schema.EnvDefaultFunc(envLogzioApiToken, nil),
				Sensitive:   true,
			},
			providerRegion: {
				Type:        schema.TypeString,
				Optional:    true,
				Description: descriptions[providerRegion],
				DefaultFunc: schema.EnvDefaultFunc(envLogzioRegion, ""),
				Sensitive:   false,
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			resourceAlertType:      dataSourceAlert(),
			resourceEndpointType:   dataSourceEndpoint(),
			resourceUserType:       dataSourceUser(),
			resourceSubAccountType: dataSourceSubAccount(),
		},
		ResourcesMap: map[string]*schema.Resource{
			resourceAlertType:      resourceAlert(),
			resourceEndpointType:   resourceEndpoint(),
			resourceUserType:       resourceUser(),
			resourceSubAccountType: resourceSubAccount(),
		},
		ConfigureFunc: providerConfigure,
	}
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{providerApiToken: "Your API token", providerRegion: "Your logz.io region"}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	apiToken, ok := d.GetOk(providerApiToken)
	if !ok {
		return nil, fmt.Errorf("can't find the %s, either set it in the provider or set the %s env var", providerApiToken, envLogzioApiToken)
	}
	region := d.Get(providerRegion).(string)
	regionCode := ""
	if region != "" && region != "us" {
		regionCode = fmt.Sprintf("-%s", region)
	}
	apiUrl := fmt.Sprintf(baseUrl, regionCode)

	config := Config{
		apiToken: apiToken.(string),
		baseUrl:  apiUrl,
	}
	return config, nil
}
