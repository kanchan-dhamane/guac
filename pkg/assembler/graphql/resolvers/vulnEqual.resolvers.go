package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"strings"

	"github.com/guacsec/guac/pkg/assembler/backends/helper"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// IngestVulnEqual is the resolver for the ingestVulnEqual field.
func (r *mutationResolver) IngestVulnEqual(ctx context.Context, vulnerability model.VulnerabilityInputSpec, otherVulnerability model.VulnerabilityInputSpec, vulnEqual model.VulnEqualInputSpec) (string, error) {
	funcName := "IngestVulnEqual"
	err := helper.ValidateNoVul(vulnerability)
	if err != nil {
		return "", gqlerror.Errorf("%v ::  %s", funcName, err)
	}

	err = helper.ValidateVulnerabilityIDInputSpec(vulnerability)
	if err != nil {
		return "", gqlerror.Errorf("%v ::  %s", funcName, err)
	}

	err = helper.ValidateNoVul(otherVulnerability)
	if err != nil {
		return "", gqlerror.Errorf("%v ::  %s", funcName, err)
	}

	err = helper.ValidateVulnerabilityIDInputSpec(otherVulnerability)
	if err != nil {
		return "", gqlerror.Errorf("%v ::  %s", funcName, err)
	}

	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase
	return r.Backend.IngestVulnEqual(ctx,
		model.VulnerabilityInputSpec{Type: strings.ToLower(vulnerability.Type), VulnerabilityID: strings.ToLower(vulnerability.VulnerabilityID)},
		model.VulnerabilityInputSpec{Type: strings.ToLower(otherVulnerability.Type), VulnerabilityID: strings.ToLower(otherVulnerability.VulnerabilityID)},
		vulnEqual)
}

// IngestVulnEquals is the resolver for the ingestVulnEquals field.
func (r *mutationResolver) IngestVulnEquals(ctx context.Context, vulnerabilities []*model.VulnerabilityInputSpec, otherVulnerabilities []*model.VulnerabilityInputSpec, vulnEquals []*model.VulnEqualInputSpec) ([]string, error) {
	funcName := "IngestVulnEquals"

	if len(vulnerabilities) != len(otherVulnerabilities) {
		return []string{}, gqlerror.Errorf("%v :: uneven vulnerabilities and other vulnerabilities for ingestion", funcName)
	} else if len(vulnerabilities) != len(vulnEquals) {
		return []string{}, gqlerror.Errorf("%v :: uneven artifacts and hashEquals for ingestion", funcName)
	}

	var lowercaseVulnList []*model.VulnerabilityInputSpec
	var lowercaseOtherVulnList []*model.VulnerabilityInputSpec
	for i := range vulnEquals {
		err := helper.ValidateNoVul(*vulnerabilities[i])
		if err != nil {
			return []string{}, gqlerror.Errorf("%v ::  %s", funcName, err)
		}

		err = helper.ValidateVulnerabilityIDInputSpec(*vulnerabilities[i])
		if err != nil {
			return []string{}, gqlerror.Errorf("%v ::  %s", funcName, err)
		}

		err = helper.ValidateNoVul(*otherVulnerabilities[i])
		if err != nil {
			return []string{}, gqlerror.Errorf("%v ::  %s", funcName, err)
		}

		err = helper.ValidateVulnerabilityIDInputSpec(*otherVulnerabilities[i])
		if err != nil {
			return []string{}, gqlerror.Errorf("%v ::  %s", funcName, err)
		}

		lowercaseVulnInput := model.VulnerabilityInputSpec{
			Type:            strings.ToLower(vulnerabilities[i].Type),
			VulnerabilityID: strings.ToLower(vulnerabilities[i].VulnerabilityID),
		}
		lowercaseVulnList = append(lowercaseVulnList, &lowercaseVulnInput)

		lowercaseOtherVulnInput := model.VulnerabilityInputSpec{
			Type:            strings.ToLower(otherVulnerabilities[i].Type),
			VulnerabilityID: strings.ToLower(otherVulnerabilities[i].VulnerabilityID),
		}
		lowercaseOtherVulnList = append(lowercaseOtherVulnList, &lowercaseOtherVulnInput)
	}

	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase
	return r.Backend.IngestVulnEquals(ctx, lowercaseVulnList, lowercaseOtherVulnList, vulnEquals)
}

// VulnEqual is the resolver for the vulnEqual field.
func (r *queryResolver) VulnEqual(ctx context.Context, vulnEqualSpec model.VulnEqualSpec) ([]*model.VulnEqual, error) {
	// vulnerability input (type and vulnerability ID) will be enforced to be lowercase

	if vulnEqualSpec.Vulnerabilities != nil && len(vulnEqualSpec.Vulnerabilities) > 2 {
		return nil, gqlerror.Errorf("VulnEqual :: cannot specify more than 2 vulnerabilities in VulnEqual")
	}

	if len(vulnEqualSpec.Vulnerabilities) > 0 {
		var lowercaseVulnFilterList []*model.VulnerabilitySpec
		for _, v := range vulnEqualSpec.Vulnerabilities {
			var typeLowerCase *string = nil
			var vulnIDLowerCase *string = nil
			if v.Type != nil {
				lower := strings.ToLower(*v.Type)
				typeLowerCase = &lower
			}
			if v.VulnerabilityID != nil {
				lower := strings.ToLower(*v.VulnerabilityID)
				vulnIDLowerCase = &lower
			}

			lowercaseVulnFilter := model.VulnerabilitySpec{
				ID:              v.ID,
				Type:            typeLowerCase,
				VulnerabilityID: vulnIDLowerCase,
				NoVuln:          v.NoVuln,
			}
			lowercaseVulnFilterList = append(lowercaseVulnFilterList, &lowercaseVulnFilter)
		}

		lowercaseVulnEqualFilter := model.VulnEqualSpec{
			ID:              vulnEqualSpec.ID,
			Vulnerabilities: lowercaseVulnFilterList,
			Justification:   vulnEqualSpec.Justification,
			Origin:          vulnEqualSpec.Origin,
			Collector:       vulnEqualSpec.Collector,
		}
		return r.Backend.VulnEqual(ctx, &lowercaseVulnEqualFilter)
	} else {
		return r.Backend.VulnEqual(ctx, &vulnEqualSpec)
	}
}
