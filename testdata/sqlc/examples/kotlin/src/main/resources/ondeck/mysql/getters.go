// Code generated by sqlcgetters; DO NOT EDIT.

package ondeck

import (
	"database/sql"
	"time"
)

func (x CreateCityParams) GetName() string {
	return x.Name
}

func (x CreateCityParams) GetSlug() string {
	return x.Slug
}

func (x UpdateCityNameParams) GetName() string {
	return x.Name
}

func (x UpdateCityNameParams) GetSlug() string {
	return x.Slug
}

func (x City) GetSlug() string {
	return x.Slug
}

func (x City) GetName() string {
	return x.Name
}

func (x Venue) GetID() int64 {
	return x.ID
}

func (x Venue) GetStatus() VenuesStatus {
	return x.Status
}

func (x Venue) GetStatuses() sql.NullString {
	return x.Statuses
}

func (x Venue) GetSlug() string {
	return x.Slug
}

func (x Venue) GetName() string {
	return x.Name
}

func (x Venue) GetCity() string {
	return x.City
}

func (x Venue) GetSpotifyPlaylist() string {
	return x.SpotifyPlaylist
}

func (x Venue) GetSongkickID() sql.NullString {
	return x.SongkickID
}

func (x Venue) GetTags() sql.NullString {
	return x.Tags
}

func (x Venue) GetCreatedAt() time.Time {
	return x.CreatedAt
}

func (x CreateVenueParams) GetSlug() string {
	return x.Slug
}

func (x CreateVenueParams) GetName() string {
	return x.Name
}

func (x CreateVenueParams) GetCity() string {
	return x.City
}

func (x CreateVenueParams) GetSpotifyPlaylist() string {
	return x.SpotifyPlaylist
}

func (x CreateVenueParams) GetStatus() VenuesStatus {
	return x.Status
}

func (x CreateVenueParams) GetStatuses() sql.NullString {
	return x.Statuses
}

func (x CreateVenueParams) GetTags() sql.NullString {
	return x.Tags
}

func (x DeleteVenueParams) GetSlug() string {
	return x.Slug
}

func (x DeleteVenueParams) GetSlug_2() string {
	return x.Slug_2
}

func (x GetVenueParams) GetSlug() string {
	return x.Slug
}

func (x GetVenueParams) GetCity() string {
	return x.City
}

func (x UpdateVenueNameParams) GetName() string {
	return x.Name
}

func (x UpdateVenueNameParams) GetSlug() string {
	return x.Slug
}

func (x VenueCountByCityRow) GetCity() string {
	return x.City
}

func (x VenueCountByCityRow) GetCount() int64 {
	return x.Count
}
