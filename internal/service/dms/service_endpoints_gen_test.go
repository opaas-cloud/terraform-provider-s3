// Code generated by internal/generate/serviceendpointtests/main.go; DO NOT EDIT.

package dms_test

import (
	"context"
	"fmt"
	"maps"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"testing"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	databasemigrationservice_sdkv1 "github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/aws-sdk-go-base/v2/servicemocks"
	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	terraformsdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/provider"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type endpointTestCase struct {
	with     []setupFunc
	expected caseExpectations
}

type caseSetup struct {
	config               map[string]any
	configFile           configFile
	environmentVariables map[string]string
}

type configFile struct {
	baseUrl    string
	serviceUrl string
}

type caseExpectations struct {
	diags    diag.Diagnostics
	endpoint string
	region   string
}

type apiCallParams struct {
	endpoint string
	region   string
}

type setupFunc func(setup *caseSetup)

type callFunc func(ctx context.Context, t *testing.T, meta *conns.AWSClient) apiCallParams

const (
	packageNameConfigEndpoint = "https://packagename-config.endpoint.test/"
	awsServiceEnvvarEndpoint  = "https://service-envvar.endpoint.test/"
	baseEnvvarEndpoint        = "https://base-envvar.endpoint.test/"
	serviceConfigFileEndpoint = "https://service-configfile.endpoint.test/"
	baseConfigFileEndpoint    = "https://base-configfile.endpoint.test/"

	aliasName0ConfigEndpoint = "https://aliasname0-config.endpoint.test/"
	aliasName1ConfigEndpoint = "https://aliasname1-config.endpoint.test/"
)

const (
	packageName = "dms"
	awsEnvVar   = "AWS_ENDPOINT_URL_DATABASE_MIGRATION_SERVICE"
	baseEnvVar  = "AWS_ENDPOINT_URL"
	configParam = "database_migration_service"

	aliasName0 = "databasemigration"
	aliasName1 = "databasemigrationservice"
)

const (
	expectedCallRegion = "us-west-2" //lintignore:AWSAT003
)

func TestEndpointConfiguration(t *testing.T) { //nolint:paralleltest // uses t.Setenv
	const providerRegion = "us-west-2" //lintignore:AWSAT003
	const expectedEndpointRegion = providerRegion

	testcases := map[string]endpointTestCase{
		"no config": {
			with:     []setupFunc{withNoConfig},
			expected: expectDefaultEndpoint(expectedEndpointRegion),
		},

		// Package name endpoint on Config

		"package name endpoint config": {
			with: []setupFunc{
				withPackageNameEndpointInConfig,
			},
			expected: expectPackageNameConfigEndpoint(),
		},

		"package name endpoint config overrides alias name 0 config": {
			with: []setupFunc{
				withPackageNameEndpointInConfig,
				withAliasName0EndpointInConfig,
			},
			expected: conflictsWith(expectPackageNameConfigEndpoint()),
		},

		"package name endpoint config overrides alias name 1 config": {
			with: []setupFunc{
				withPackageNameEndpointInConfig,
				withAliasName1EndpointInConfig,
			},
			expected: conflictsWith(expectPackageNameConfigEndpoint()),
		},

		"package name endpoint config overrides aws service envvar": {
			with: []setupFunc{
				withPackageNameEndpointInConfig,
				withAwsEnvVar,
			},
			expected: expectPackageNameConfigEndpoint(),
		},

		"package name endpoint config overrides base envvar": {
			with: []setupFunc{
				withPackageNameEndpointInConfig,
				withBaseEnvVar,
			},
			expected: expectPackageNameConfigEndpoint(),
		},

		"package name endpoint config overrides service config file": {
			with: []setupFunc{
				withPackageNameEndpointInConfig,
				withServiceEndpointInConfigFile,
			},
			expected: expectPackageNameConfigEndpoint(),
		},

		"package name endpoint config overrides base config file": {
			with: []setupFunc{
				withPackageNameEndpointInConfig,
				withBaseEndpointInConfigFile,
			},
			expected: expectPackageNameConfigEndpoint(),
		},

		// Alias name 0 endpoint on Config

		"alias name 0 endpoint config": {
			with: []setupFunc{
				withAliasName0EndpointInConfig,
			},
			expected: expectAliasName0ConfigEndpoint(),
		},

		"alias name 0 endpoint config overrides alias name 1 config": {
			with: []setupFunc{
				withAliasName0EndpointInConfig,
				withAliasName1EndpointInConfig,
			},
			expected: conflictsWith(expectAliasName0ConfigEndpoint()),
		},

		"alias name 0 endpoint config overrides aws service envvar": {
			with: []setupFunc{
				withAliasName0EndpointInConfig,
				withAwsEnvVar,
			},
			expected: expectAliasName0ConfigEndpoint(),
		},

		"alias name 0 endpoint config overrides base envvar": {
			with: []setupFunc{
				withAliasName0EndpointInConfig,
				withBaseEnvVar,
			},
			expected: expectAliasName0ConfigEndpoint(),
		},

		"alias name 0 endpoint config overrides service config file": {
			with: []setupFunc{
				withAliasName0EndpointInConfig,
				withServiceEndpointInConfigFile,
			},
			expected: expectAliasName0ConfigEndpoint(),
		},

		"alias name 0 endpoint config overrides base config file": {
			with: []setupFunc{
				withAliasName0EndpointInConfig,
				withBaseEndpointInConfigFile,
			},
			expected: expectAliasName0ConfigEndpoint(),
		},

		// Alias name 1 endpoint on Config

		"alias name 1 endpoint config": {
			with: []setupFunc{
				withAliasName1EndpointInConfig,
			},
			expected: expectAliasName1ConfigEndpoint(),
		},

		"alias name 1 endpoint config overrides aws service envvar": {
			with: []setupFunc{
				withAliasName1EndpointInConfig,
				withAwsEnvVar,
			},
			expected: expectAliasName1ConfigEndpoint(),
		},

		"alias name 1 endpoint config overrides base envvar": {
			with: []setupFunc{
				withAliasName1EndpointInConfig,
				withBaseEnvVar,
			},
			expected: expectAliasName1ConfigEndpoint(),
		},

		"alias name 1 endpoint config overrides service config file": {
			with: []setupFunc{
				withAliasName1EndpointInConfig,
				withServiceEndpointInConfigFile,
			},
			expected: expectAliasName1ConfigEndpoint(),
		},

		"alias name 1 endpoint config overrides base config file": {
			with: []setupFunc{
				withAliasName1EndpointInConfig,
				withBaseEndpointInConfigFile,
			},
			expected: expectAliasName1ConfigEndpoint(),
		},

		// Service endpoint in AWS envvar

		"service aws envvar": {
			with: []setupFunc{
				withAwsEnvVar,
			},
			expected: expectAwsEnvVarEndpoint(),
		},

		"service aws envvar overrides base envvar": {
			with: []setupFunc{
				withAwsEnvVar,
				withBaseEnvVar,
			},
			expected: expectAwsEnvVarEndpoint(),
		},

		"service aws envvar overrides service config file": {
			with: []setupFunc{
				withAwsEnvVar,
				withServiceEndpointInConfigFile,
			},
			expected: expectAwsEnvVarEndpoint(),
		},

		"service aws envvar overrides base config file": {
			with: []setupFunc{
				withAwsEnvVar,
				withBaseEndpointInConfigFile,
			},
			expected: expectAwsEnvVarEndpoint(),
		},

		// Base endpoint in envvar

		"base endpoint envvar": {
			with: []setupFunc{
				withBaseEnvVar,
			},
			expected: expectBaseEnvVarEndpoint(),
		},

		"base endpoint envvar overrides service config file": {
			with: []setupFunc{
				withBaseEnvVar,
				withServiceEndpointInConfigFile,
			},
			expected: expectBaseEnvVarEndpoint(),
		},

		"base endpoint envvar overrides base config file": {
			with: []setupFunc{
				withBaseEnvVar,
				withBaseEndpointInConfigFile,
			},
			expected: expectBaseEnvVarEndpoint(),
		},

		// Service endpoint in config file

		"service config file": {
			with: []setupFunc{
				withServiceEndpointInConfigFile,
			},
			expected: expectServiceConfigFileEndpoint(),
		},

		"service config file overrides base config file": {
			with: []setupFunc{
				withServiceEndpointInConfigFile,
				withBaseEndpointInConfigFile,
			},
			expected: expectServiceConfigFileEndpoint(),
		},

		// Base endpoint in config file

		"base endpoint config file": {
			with: []setupFunc{
				withBaseEndpointInConfigFile,
			},
			expected: expectBaseConfigFileEndpoint(),
		},

		// Use FIPS endpoint on Config

		"use fips config": {
			with: []setupFunc{
				withUseFIPSInConfig,
			},
			expected: expectDefaultFIPSEndpoint(expectedEndpointRegion),
		},

		"use fips config with package name endpoint config": {
			with: []setupFunc{
				withUseFIPSInConfig,
				withPackageNameEndpointInConfig,
			},
			expected: expectPackageNameConfigEndpoint(),
		},
	}

	for name, testcase := range testcases { //nolint:paralleltest // uses t.Setenv
		testcase := testcase

		t.Run(name, func(t *testing.T) {
			testEndpointCase(t, providerRegion, testcase, callService)
		})
	}
}

func defaultEndpoint(region string) string {
	r := endpoints.DefaultResolver()

	ep, err := r.EndpointFor(databasemigrationservice_sdkv1.EndpointsID, region)
	if err != nil {
		return err.Error()
	}

	url, _ := url.Parse(ep.URL)

	if url.Path == "" {
		url.Path = "/"
	}

	return url.String()
}

func defaultFIPSEndpoint(region string) string {
	r := endpoints.DefaultResolver()

	ep, err := r.EndpointFor(databasemigrationservice_sdkv1.EndpointsID, region, func(opt *endpoints.Options) {
		opt.UseFIPSEndpoint = endpoints.FIPSEndpointStateEnabled
	})
	if err != nil {
		return err.Error()
	}

	url, _ := url.Parse(ep.URL)

	if url.Path == "" {
		url.Path = "/"
	}

	return url.String()
}

func callService(ctx context.Context, t *testing.T, meta *conns.AWSClient) apiCallParams {
	t.Helper()

	client := meta.DMSConn(ctx)

	req, _ := client.DescribeCertificatesRequest(&databasemigrationservice_sdkv1.DescribeCertificatesInput{})

	req.HTTPRequest.URL.Path = "/"

	return apiCallParams{
		endpoint: req.HTTPRequest.URL.String(),
		region:   aws_sdkv1.StringValue(client.Config.Region),
	}
}

func withNoConfig(_ *caseSetup) {
	// no-op
}

func withPackageNameEndpointInConfig(setup *caseSetup) {
	if _, ok := setup.config[names.AttrEndpoints]; !ok {
		setup.config[names.AttrEndpoints] = []any{
			map[string]any{},
		}
	}
	endpoints := setup.config[names.AttrEndpoints].([]any)[0].(map[string]any)
	endpoints[packageName] = packageNameConfigEndpoint
}

func withAliasName0EndpointInConfig(setup *caseSetup) {
	if _, ok := setup.config[names.AttrEndpoints]; !ok {
		setup.config[names.AttrEndpoints] = []any{
			map[string]any{},
		}
	}
	endpoints := setup.config[names.AttrEndpoints].([]any)[0].(map[string]any)
	endpoints[aliasName0] = aliasName0ConfigEndpoint
}

func withAliasName1EndpointInConfig(setup *caseSetup) {
	if _, ok := setup.config[names.AttrEndpoints]; !ok {
		setup.config[names.AttrEndpoints] = []any{
			map[string]any{},
		}
	}
	endpoints := setup.config[names.AttrEndpoints].([]any)[0].(map[string]any)
	endpoints[aliasName1] = aliasName1ConfigEndpoint
}

func conflictsWith(e caseExpectations) caseExpectations {
	e.diags = append(e.diags, provider.ConflictingEndpointsWarningDiag(
		cty.GetAttrPath(names.AttrEndpoints).IndexInt(0),
		packageName,
		aliasName0,
		aliasName1,
	))
	return e
}

func withAwsEnvVar(setup *caseSetup) {
	setup.environmentVariables[awsEnvVar] = awsServiceEnvvarEndpoint
}

func withBaseEnvVar(setup *caseSetup) {
	setup.environmentVariables[baseEnvVar] = baseEnvvarEndpoint
}

func withServiceEndpointInConfigFile(setup *caseSetup) {
	setup.configFile.serviceUrl = serviceConfigFileEndpoint
}

func withBaseEndpointInConfigFile(setup *caseSetup) {
	setup.configFile.baseUrl = baseConfigFileEndpoint
}

func withUseFIPSInConfig(setup *caseSetup) {
	setup.config["use_fips_endpoint"] = true
}

func expectDefaultEndpoint(region string) caseExpectations {
	return caseExpectations{
		endpoint: defaultEndpoint(region),
		region:   expectedCallRegion,
	}
}

func expectDefaultFIPSEndpoint(region string) caseExpectations {
	return caseExpectations{
		endpoint: defaultFIPSEndpoint(region),
		region:   "us-west-2",
	}
}

func expectPackageNameConfigEndpoint() caseExpectations {
	return caseExpectations{
		endpoint: packageNameConfigEndpoint,
		region:   expectedCallRegion,
	}
}

func expectAliasName0ConfigEndpoint() caseExpectations {
	return caseExpectations{
		endpoint: aliasName0ConfigEndpoint,
		region:   expectedCallRegion,
	}
}

func expectAliasName1ConfigEndpoint() caseExpectations {
	return caseExpectations{
		endpoint: aliasName1ConfigEndpoint,
		region:   expectedCallRegion,
	}
}

func expectAwsEnvVarEndpoint() caseExpectations {
	return caseExpectations{
		endpoint: awsServiceEnvvarEndpoint,
		region:   expectedCallRegion,
	}
}

func expectBaseEnvVarEndpoint() caseExpectations {
	return caseExpectations{
		endpoint: baseEnvvarEndpoint,
		region:   expectedCallRegion,
	}
}

func expectServiceConfigFileEndpoint() caseExpectations {
	return caseExpectations{
		endpoint: serviceConfigFileEndpoint,
		region:   expectedCallRegion,
	}
}

func expectBaseConfigFileEndpoint() caseExpectations {
	return caseExpectations{
		endpoint: baseConfigFileEndpoint,
		region:   expectedCallRegion,
	}
}

func testEndpointCase(t *testing.T, region string, testcase endpointTestCase, callF callFunc) {
	t.Helper()

	ctx := context.Background()

	setup := caseSetup{
		config:               map[string]any{},
		environmentVariables: map[string]string{},
	}

	for _, f := range testcase.with {
		f(&setup)
	}

	config := map[string]any{
		names.AttrAccessKey:                 servicemocks.MockStaticAccessKey,
		names.AttrSecretKey:                 servicemocks.MockStaticSecretKey,
		names.AttrRegion:                    region,
		names.AttrSkipCredentialsValidation: true,
		names.AttrSkipRequestingAccountID:   true,
	}

	maps.Copy(config, setup.config)

	if setup.configFile.baseUrl != "" || setup.configFile.serviceUrl != "" {
		config[names.AttrProfile] = "default"
		tempDir := t.TempDir()
		writeSharedConfigFile(t, &config, tempDir, generateSharedConfigFile(setup.configFile))
	}

	for k, v := range setup.environmentVariables {
		t.Setenv(k, v)
	}

	p, err := provider.New(ctx)
	if err != nil {
		t.Fatal(err)
	}

	expectedDiags := testcase.expected.diags
	expectedDiags = append(
		expectedDiags,
		errs.NewWarningDiagnostic(
			"AWS account ID not found for provider",
			"See https://registry.terraform.io/providers/hashicorp/aws/latest/docs#skip_requesting_account_id for implications.",
		),
	)

	diags := p.Configure(ctx, terraformsdk.NewResourceConfigRaw(config))

	if diff := cmp.Diff(diags, expectedDiags, cmp.Comparer(sdkdiag.Comparer)); diff != "" {
		t.Errorf("unexpected diagnostics difference: %s", diff)
	}

	if diags.HasError() {
		return
	}

	meta := p.Meta().(*conns.AWSClient)

	callParams := callF(ctx, t, meta)

	if e, a := testcase.expected.endpoint, callParams.endpoint; e != a {
		t.Errorf("expected endpoint %q, got %q", e, a)
	}

	if e, a := testcase.expected.region, callParams.region; e != a {
		t.Errorf("expected region %q, got %q", e, a)
	}
}

func generateSharedConfigFile(config configFile) string {
	var buf strings.Builder

	buf.WriteString(`
[default]
aws_access_key_id = DefaultSharedCredentialsAccessKey
aws_secret_access_key = DefaultSharedCredentialsSecretKey
`)
	if config.baseUrl != "" {
		buf.WriteString(fmt.Sprintf("endpoint_url = %s\n", config.baseUrl))
	}

	if config.serviceUrl != "" {
		buf.WriteString(fmt.Sprintf(`
services = endpoint-test

[services endpoint-test]
%[1]s =
  endpoint_url = %[2]s
`, configParam, serviceConfigFileEndpoint))
	}

	return buf.String()
}

func writeSharedConfigFile(t *testing.T, config *map[string]any, tempDir, content string) string {
	t.Helper()

	file, err := os.Create(filepath.Join(tempDir, "aws-sdk-go-base-shared-configuration-file"))
	if err != nil {
		t.Fatalf("creating shared configuration file: %s", err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		t.Fatalf(" writing shared configuration file: %s", err)
	}

	if v, ok := (*config)[names.AttrSharedConfigFiles]; !ok {
		(*config)[names.AttrSharedConfigFiles] = []any{file.Name()}
	} else {
		(*config)[names.AttrSharedConfigFiles] = append(v.([]any), file.Name())
	}

	return file.Name()
}
