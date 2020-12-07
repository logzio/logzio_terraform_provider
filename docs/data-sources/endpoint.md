# Endpoint Resource

Use this data source to access information about existing Logz.io notification endpoints and custom webhooks.

* Endpoints can be used to send alerts, automate reports, share logs and dashboards, directly from Logz.io. Learn more about endpoint integrations in the [Logz.io Docs](https://docs.logz.io/user-guide/integrations/endpoints.html)
* Learn more about available [APIs for managing Logz.io endpoints](https://docs.logz.io/api/#tag/Manage-notification-endpoints)

## Example Usage

```hcl
resource "logzio_endpoint" "my_endpoint" {
  title = "my_endpoint"
  description = "hello"
  endpoint_type = "Slack"
  slack {
    url = "https://this.is.com/some/url"
  }
}
```




## Argument Reference

* `endpoint_type` - (Required) Specifies the endpoint resource type: `custom`, `slack`, `pager_duty`, `big_panda`, `data_dog`, `victorops`. Use the appropriate parameters for your selected endpoint type.
* `title` - (Required) Name of the endpoint.
* `description` - (Required) Detailed description of the endpoint.
* `custom` - (Optional) Relevant when `endpoint_type` is `custom`. Manages a custom webhook for your integration of choice.
  * `url` - Specifies the URL destination.
  * `method` - Selects the HTTP request method.
  * `headers` - Header parameters for the request. Sent as comma-separated key-value pairs.
	* `body_template` - JSON object that serves as the template for the message body.
* `slack` - (Optional) Relevant when `endpoint_type` is `slack`. Manages a webhook to a specific Slack channel.
	  * `url` - Slack webhook URL to a specific Slack channel.
* `pager_duty` - (Optional) Relevant when `endpoint_type` is `pager_duty`. Manages a webhook to PagerDuty.
	* `service_key` - API key generated from PagerDuty for the purpose of the integration.
* `big_panda` - (Optional) Relevant when `endpoint_type` is `big_panda`. Manages a webhook to BigPanda.
	  * `api_token` - API authentication token from BigPanda.
  	* `app_key` - Application key from BigPanda.
* `data_dog` - (Optional) Relevant when `endpoint_type` is `data_dog`. Manages a webhook to Datadog.
  	* `api_key` - API key from Datadog.
* `victorops` - (Optional) Relevant when `endpoint_type` is `victorops`. Manages a webhook to VictorOps.
  * `routing_key` - Alert routing key from VictorOps.
  * `message_type` - VictorOps REST API `message_type`.
  * `service_api_key` - API key from VictorOps.


## Attribute Reference

* `id` - ID of the notification endpoint.


## Endpoints used

Logz.io integrates with:
* [Slack](https://docs.logz.io/api/#operation/updateSlack)
* [PagerDuty](https://docs.logz.io/api/#operation/updatePagerDuty)
* [BigPanda](https://docs.logz.io/api/#operation/updateBigPanda)
* [Datadog](https://docs.logz.io/api/#operation/updateDataDog)
* [VictorOps](https://docs.logz.io/api/#operation/updateVictorops)
* [Custom integration](https://docs.logz.io/api/#operation/updateCustom)

Other endpoints:
* [Get all endpoints](https://docs.logz.io/api/#operation/getAllEndpoints)
* [Get endpoint by ID](https://docs.logz.io/api/#operation/getEndpointById)