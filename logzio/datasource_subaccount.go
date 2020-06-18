package logzio

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/jonboydell/logzio_client/sub_accounts"
	"github.com/yyyogev/logzio_terraform_provider/logzio"
)

func dataSourceSubAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSubaccountRead,
		Schema: map[string]*schema.Schema{
			subAccountId: {
				Type:	schema.TypeInt,
			},
			subAccountEmail: {
				Type:	schema.TypeString,
			},
			subAccountToken: {
				Type:	schema.TypeString,
			},
			subAccountMaxDailyGB: {
				Type:	schema.TypeFloat,
			},
			subAccountRetentionDays: {
				Type:	schema.TypeInt,
			},
			subAccountSearchable: {
				Type:	schema.TypeBool,
			},
			subAccountDocSizeSetting: {
				Type:	schema.TypeBool,
			},
			subAccountSharingObjectsAccounts: {
				Type:	schema.TypeList,
			},
			subAccountUtilizationSettings: {
				Type:	schema.TypeMap,

			},
		},
	}
}

func dataSourceSubaccountRead(d *schema.ResourceData, m interface{}) error {
	var client *sub_accounts.SubAccountClient
	client, _ = sub_accounts.New(m.(logzio.Config).apiToken, m.(logzio.Config).baseUrl)

	accountId, ok := d.GetOk(subAccountId)
	if ok {
		subAccount, err := client.GetSubAccount(accountId.(int64))
		if err != nil {
			return err
		}
		setSubaccount(d, subAccount)
		return nil
	}

	accountToken, ok := d.GetOk(subAccountToken)
	if ok {
		list, err := client.ListSubAccounts()
		if err != nil {
			return err
		}
		for i := 0; i < len(list); i++ {
			subAccount := list[i]
			if subAccount.AccountToken == accountToken {
				setSubaccount(d, &subAccount)
				return nil
			}
		}
	}

	return fmt.Errorf("couldn't find sub-account with specified attributes")
}

func setSubaccount(data *schema.ResourceData, subAccount *sub_accounts.SubAccount) {
	data.SetId(fmt.Sprintf("%d", subAccount.Id))
	data.Set(subAccountName, subAccount.AccountName)
	data.Set(subAccountEmail, subAccount.Email)
	data.Set(subAccountToken, subAccount.AccountToken)
	data.Set(subAccountDocSizeSetting, subAccount.DocSizeSetting)
	data.Set(subAccountUtilizationSettings, subAccount.UtilizationSettings)
	data.Set(subAccountAccessible, subAccount.Accessible)
	data.Set(subAccountSearchable, subAccount.Searchable)
	data.Set(subAccountRetentionDays, subAccount.RetentionDays)
	data.Set(subAccountMaxDailyGB, subAccount.MaxDailyGB)
	data.Set(subAccountSharingObjectsAccounts, subAccount.SharingObjectAccounts)
}