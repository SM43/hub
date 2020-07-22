// Code generated by goa v3.2.0, DO NOT EDIT.
//
// resource views
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// ResourceCollection is the viewed result type that is projected based on a
// view.
type ResourceCollection struct {
	// Type to project
	Projected ResourceCollectionView
	// View to render
	View string
}

// Versions is the viewed result type that is projected based on a view.
type Versions struct {
	// Type to project
	Projected *VersionsView
	// View to render
	View string
}

// ResourceCollectionView is a type that runs validations on a projected type.
type ResourceCollectionView []*ResourceView

// ResourceView is a type that runs validations on a projected type.
type ResourceView struct {
	// ID is the unique id of the resource
	ID *uint
	// Name of resource
	Name *string
	// Type of catalog to which resource belongs
	Catalog *CatalogView
	// Type of resource
	Type *string
	// Latest version of resource
	LatestVersion *LatestVersionView
	// Tags related to resource
	Tags []*TagView
	// Rating of resource
	Rating *float64
}

// CatalogView is a type that runs validations on a projected type.
type CatalogView struct {
	// ID is the unique id of the catalog
	ID *uint
	// Type of catalog
	Type *string
}

// LatestVersionView is a type that runs validations on a projected type.
type LatestVersionView struct {
	// ID is the unique id of resource's version
	ID *uint
	// Version of resource
	Version *string
	// Display name of version
	DisplayName *string
	// Description of version
	Description *string
	// Minimum pipelines version the resource's version is compatible with
	MinPipelinesVersion *string
	// Raw URL of resource's yaml file of the version
	RawURL *string
	// Web URL of resource's yaml file of the version
	WebURL *string
	// Timestamp when version was last updated
	UpdatedAt *string
}

// TagView is a type that runs validations on a projected type.
type TagView struct {
	// ID is the unique id of tag
	ID *uint
	// Name of tag
	Name *string
}

// VersionsView is a type that runs validations on a projected type.
type VersionsView struct {
	// Latest Version of resource
	Latest *VersionView
	// List of all versions of resource
	Versions []*VersionView
}

// VersionView is a type that runs validations on a projected type.
type VersionView struct {
	// ID is the unique id of resource's version
	ID *uint
	// Version of resource
	Version *string
	// Raw URL of resource's yaml file of the version
	RawURL *string
	// Web URL of resource's yaml file of the version
	WebURL *string
}

var (
	// ResourceCollectionMap is a map of attribute names in result type
	// ResourceCollection indexed by view name.
	ResourceCollectionMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"catalog",
			"type",
			"latestVersion",
			"tags",
			"rating",
		},
	}
	// VersionsMap is a map of attribute names in result type Versions indexed by
	// view name.
	VersionsMap = map[string][]string{
		"default": []string{
			"latest",
			"versions",
		},
	}
	// ResourceMap is a map of attribute names in result type Resource indexed by
	// view name.
	ResourceMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"catalog",
			"type",
			"latestVersion",
			"tags",
			"rating",
		},
	}
	// VersionMap is a map of attribute names in result type Version indexed by
	// view name.
	VersionMap = map[string][]string{
		"default": []string{
			"id",
			"version",
			"rawURL",
			"webURL",
		},
	}
)

// ValidateResourceCollection runs the validations defined on the viewed result
// type ResourceCollection.
func ValidateResourceCollection(result ResourceCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateResourceCollectionView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateVersions runs the validations defined on the viewed result type
// Versions.
func ValidateVersions(result *Versions) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateVersionsView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateResourceCollectionView runs the validations defined on
// ResourceCollectionView using the "default" view.
func ValidateResourceCollectionView(result ResourceCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateResourceView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateResourceView runs the validations defined on ResourceView using the
// "default" view.
func ValidateResourceView(result *ResourceView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Catalog == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("catalog", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.LatestVersion == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("latestVersion", "result"))
	}
	if result.Tags == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("tags", "result"))
	}
	if result.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rating", "result"))
	}
	if result.Catalog != nil {
		if err2 := ValidateCatalogView(result.Catalog); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.LatestVersion != nil {
		if err2 := ValidateLatestVersionView(result.LatestVersion); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	for _, e := range result.Tags {
		if e != nil {
			if err2 := ValidateTagView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateCatalogView runs the validations defined on CatalogView.
func ValidateCatalogView(result *CatalogView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Type != nil {
		if !(*result.Type == "official" || *result.Type == "community") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []interface{}{"official", "community"}))
		}
	}
	return
}

// ValidateLatestVersionView runs the validations defined on LatestVersionView.
func ValidateLatestVersionView(result *LatestVersionView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Version == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("version", "result"))
	}
	if result.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "result"))
	}
	if result.DisplayName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("displayName", "result"))
	}
	if result.MinPipelinesVersion == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("minPipelinesVersion", "result"))
	}
	if result.RawURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rawURL", "result"))
	}
	if result.WebURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("webURL", "result"))
	}
	if result.UpdatedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("updatedAt", "result"))
	}
	if result.RawURL != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.rawURL", *result.RawURL, goa.FormatURI))
	}
	if result.WebURL != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.webURL", *result.WebURL, goa.FormatURI))
	}
	if result.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.updatedAt", *result.UpdatedAt, goa.FormatDateTime))
	}
	return
}

// ValidateTagView runs the validations defined on TagView.
func ValidateTagView(result *TagView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	return
}

// ValidateVersionsView runs the validations defined on VersionsView using the
// "default" view.
func ValidateVersionsView(result *VersionsView) (err error) {
	if result.Versions == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("versions", "result"))
	}
	for _, e := range result.Versions {
		if e != nil {
			if err2 := ValidateVersionView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Latest != nil {
		if err2 := ValidateVersionView(result.Latest); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateVersionView runs the validations defined on VersionView using the
// "default" view.
func ValidateVersionView(result *VersionView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Version == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("version", "result"))
	}
	if result.RawURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("rawURL", "result"))
	}
	if result.WebURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("webURL", "result"))
	}
	if result.RawURL != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.rawURL", *result.RawURL, goa.FormatURI))
	}
	if result.WebURL != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.webURL", *result.WebURL, goa.FormatURI))
	}
	return
}
