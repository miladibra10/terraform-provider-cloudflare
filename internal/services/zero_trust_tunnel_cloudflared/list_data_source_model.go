// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_tunnel_cloudflared

import (
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustTunnelCloudflaredsResultListDataSourceEnvelope struct {
	Result *[]*ZeroTrustTunnelCloudflaredsResultDataSourceModel `json:"result,computed"`
}

type ZeroTrustTunnelCloudflaredsDataSourceModel struct {
	AccountID     types.String                                         `tfsdk:"account_id" path:"account_id"`
	ExcludePrefix types.String                                         `tfsdk:"exclude_prefix" query:"exclude_prefix"`
	ExistedAt     timetypes.RFC3339                                    `tfsdk:"existed_at" query:"existed_at"`
	IncludePrefix types.String                                         `tfsdk:"include_prefix" query:"include_prefix"`
	IsDeleted     types.Bool                                           `tfsdk:"is_deleted" query:"is_deleted"`
	Name          types.String                                         `tfsdk:"name" query:"name"`
	Status        types.String                                         `tfsdk:"status" query:"status"`
	UUID          types.String                                         `tfsdk:"uuid" query:"uuid"`
	WasActiveAt   timetypes.RFC3339                                    `tfsdk:"was_active_at" query:"was_active_at"`
	WasInactiveAt timetypes.RFC3339                                    `tfsdk:"was_inactive_at" query:"was_inactive_at"`
	MaxItems      types.Int64                                          `tfsdk:"max_items"`
	Result        *[]*ZeroTrustTunnelCloudflaredsResultDataSourceModel `tfsdk:"result"`
}

type ZeroTrustTunnelCloudflaredsResultDataSourceModel struct {
}
