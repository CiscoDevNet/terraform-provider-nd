package provider

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/ciscoecosystem/mso-go-client/client"
	"github.com/ciscoecosystem/mso-go-client/container"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

func DoRestRequestEscapeHtml(ctx context.Context, diags *diag.Diagnostics, client *client.Client, path, method string, payload *container.Container, escapeHtml bool) *container.Container {
	// Ensure path starts with a slash to assure signature is created correctly
	if !strings.HasPrefix("/", path) {
		path = fmt.Sprintf("/%s", path)
	}
	var restRequest *http.Request
	var err error

	if escapeHtml {
		restRequest, err = client.MakeRestRequest(method, path, payload, true)
	} else {
		restRequest, err = client.MakeRestRequestRaw(method, path, payload.EncodeJSON(), true)
	}
	if err != nil {
		diags.AddError(
			"Creation of rest request failed",
			fmt.Sprintf("err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	cont, restResponse, err := client.Do(restRequest)

	// Return nil when the object is not found and ignore 404 not found error
	if restResponse.StatusCode == 404 {
		return nil
	}

	if restResponse != nil && cont.Data() != nil && (restResponse.StatusCode != 200 && restResponse.StatusCode != 201) {
		diags.AddError(
			fmt.Sprintf("The %s rest request failed inside status code check", strings.ToLower(method)),
			fmt.Sprintf("Code: %d Response: %s, err: %s. Please report this issue to the provider developers.", restResponse.StatusCode, cont.Data().(map[string]interface{})["errors"], err),
		)
		tflog.Debug(ctx, fmt.Sprintf("%v", cont.Search("errors")))
		return nil
	} else if err != nil {
		diags.AddError(
			fmt.Sprintf("The %s rest request failed else part of the != 200", strings.ToLower(method)),
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	return cont
}

func DoRestRequest(ctx context.Context, diags *diag.Diagnostics, client *client.Client, path, method string, payload *container.Container) *container.Container {
	return DoRestRequestEscapeHtml(ctx, diags, client, path, method, payload, true)
}
