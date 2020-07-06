// Code generated by goa v3.2.0, DO NOT EDIT.
//
// resource service
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package resource

import (
	"context"

	resourceviews "github.com/tektoncd/hub/api/gen/resource/views"
	goa "goa.design/goa/v3/pkg"
)

// The resource service provides details about all type of resources
type Service interface {
	// Find resources by a combination of name, type
	Query(context.Context, *QueryPayload) (res ResourceCollection, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "resource"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"Query"}

// QueryPayload is the payload type of the resource service Query method.
type QueryPayload struct {
	// Name of resource
	Name string
	// Type of resource
	Type string
	// Maximum number of resources to be returned
	Limit uint
}

// ResourceCollection is the result type of the resource service Query method.
type ResourceCollection []*Resource

// The resource type describes resource information.
type Resource struct {
	// ID is the unique id of the resource
	ID uint
	// Name of resource
	Name string
	// Type of catalog to which resource belongs
	Catalog *Catalog
	// Type of resource
	Type string
	// Latest version of resource
	LatestVersion *Version
	// Tags related to resource
	Tags []*Tag
	// Rating of resource
	Rating float64
}

type Catalog struct {
	// ID is the unique id of the catalog
	ID uint
	// Type of catalog
	Type string
}

type Version struct {
	// ID is the unique id of resource's version
	ID uint
	// Version of resource
	Version string
	// Display name of version
	DisplayName string
	// Description of version
	Description string
	// Minimum pipelines version the resource's version is compatible with
	MinPipelinesVersion string
	// Raw URL of resource's yaml file of the version
	RawURL string
	// Web URL of resource's yaml file of the version
	WebURL string
	// Timestamp when version was last updated
	UpdatedAt string
}

type Tag struct {
	// ID is the unique id of tag
	ID uint
	// Name of tag
	Name string
}

// MakeInternalError builds a goa.ServiceError from an error.
func MakeInternalError(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "internal-error",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not-found",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewResourceCollection initializes result type ResourceCollection from viewed
// result type ResourceCollection.
func NewResourceCollection(vres resourceviews.ResourceCollection) ResourceCollection {
	return newResourceCollection(vres.Projected)
}

// NewViewedResourceCollection initializes viewed result type
// ResourceCollection from result type ResourceCollection using the given view.
func NewViewedResourceCollection(res ResourceCollection, view string) resourceviews.ResourceCollection {
	p := newResourceCollectionView(res)
	return resourceviews.ResourceCollection{Projected: p, View: "default"}
}

// newResourceCollection converts projected type ResourceCollection to service
// type ResourceCollection.
func newResourceCollection(vres resourceviews.ResourceCollectionView) ResourceCollection {
	res := make(ResourceCollection, len(vres))
	for i, n := range vres {
		res[i] = newResource(n)
	}
	return res
}

// newResourceCollectionView projects result type ResourceCollection to
// projected type ResourceCollectionView using the "default" view.
func newResourceCollectionView(res ResourceCollection) resourceviews.ResourceCollectionView {
	vres := make(resourceviews.ResourceCollectionView, len(res))
	for i, n := range res {
		vres[i] = newResourceView(n)
	}
	return vres
}

// newResource converts projected type Resource to service type Resource.
func newResource(vres *resourceviews.ResourceView) *Resource {
	res := &Resource{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Type != nil {
		res.Type = *vres.Type
	}
	if vres.Rating != nil {
		res.Rating = *vres.Rating
	}
	if vres.Catalog != nil {
		res.Catalog = transformResourceviewsCatalogViewToCatalog(vres.Catalog)
	}
	if vres.LatestVersion != nil {
		res.LatestVersion = transformResourceviewsVersionViewToVersion(vres.LatestVersion)
	}
	if vres.Tags != nil {
		res.Tags = make([]*Tag, len(vres.Tags))
		for i, val := range vres.Tags {
			res.Tags[i] = transformResourceviewsTagViewToTag(val)
		}
	}
	return res
}

// newResourceView projects result type Resource to projected type ResourceView
// using the "default" view.
func newResourceView(res *Resource) *resourceviews.ResourceView {
	vres := &resourceviews.ResourceView{
		ID:     &res.ID,
		Name:   &res.Name,
		Type:   &res.Type,
		Rating: &res.Rating,
	}
	if res.Catalog != nil {
		vres.Catalog = transformCatalogToResourceviewsCatalogView(res.Catalog)
	}
	if res.LatestVersion != nil {
		vres.LatestVersion = transformVersionToResourceviewsVersionView(res.LatestVersion)
	}
	if res.Tags != nil {
		vres.Tags = make([]*resourceviews.TagView, len(res.Tags))
		for i, val := range res.Tags {
			vres.Tags[i] = transformTagToResourceviewsTagView(val)
		}
	}
	return vres
}

// transformResourceviewsCatalogViewToCatalog builds a value of type *Catalog
// from a value of type *resourceviews.CatalogView.
func transformResourceviewsCatalogViewToCatalog(v *resourceviews.CatalogView) *Catalog {
	if v == nil {
		return nil
	}
	res := &Catalog{
		ID:   *v.ID,
		Type: *v.Type,
	}

	return res
}

// transformResourceviewsVersionViewToVersion builds a value of type *Version
// from a value of type *resourceviews.VersionView.
func transformResourceviewsVersionViewToVersion(v *resourceviews.VersionView) *Version {
	if v == nil {
		return nil
	}
	res := &Version{
		ID:                  *v.ID,
		Version:             *v.Version,
		DisplayName:         *v.DisplayName,
		Description:         *v.Description,
		MinPipelinesVersion: *v.MinPipelinesVersion,
		RawURL:              *v.RawURL,
		WebURL:              *v.WebURL,
		UpdatedAt:           *v.UpdatedAt,
	}

	return res
}

// transformResourceviewsTagViewToTag builds a value of type *Tag from a value
// of type *resourceviews.TagView.
func transformResourceviewsTagViewToTag(v *resourceviews.TagView) *Tag {
	if v == nil {
		return nil
	}
	res := &Tag{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// transformCatalogToResourceviewsCatalogView builds a value of type
// *resourceviews.CatalogView from a value of type *Catalog.
func transformCatalogToResourceviewsCatalogView(v *Catalog) *resourceviews.CatalogView {
	res := &resourceviews.CatalogView{
		ID:   &v.ID,
		Type: &v.Type,
	}

	return res
}

// transformVersionToResourceviewsVersionView builds a value of type
// *resourceviews.VersionView from a value of type *Version.
func transformVersionToResourceviewsVersionView(v *Version) *resourceviews.VersionView {
	res := &resourceviews.VersionView{
		ID:                  &v.ID,
		Version:             &v.Version,
		DisplayName:         &v.DisplayName,
		Description:         &v.Description,
		MinPipelinesVersion: &v.MinPipelinesVersion,
		RawURL:              &v.RawURL,
		WebURL:              &v.WebURL,
		UpdatedAt:           &v.UpdatedAt,
	}

	return res
}

// transformTagToResourceviewsTagView builds a value of type
// *resourceviews.TagView from a value of type *Tag.
func transformTagToResourceviewsTagView(v *Tag) *resourceviews.TagView {
	res := &resourceviews.TagView{
		ID:   &v.ID,
		Name: &v.Name,
	}

	return res
}
