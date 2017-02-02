package api

import (
	"github.com/zballs/envoke/crypto/ed25519"
)

type PartnerInfo struct {
	PartnerId string              `json:"partner_id"`
	Privkey   *ed25519.PrivateKey `json:"private_key"`
	Pubkey    *ed25519.PublicKey  `json:"public_key"`
}

func NewPartnerInfo(partnerId string, priv *ed25519.PrivateKey, pub *ed25519.PublicKey) *PartnerInfo {
	return &PartnerInfo{
		PartnerId: partnerId,
		Privkey:   priv,
		Pubkey:    pub,
	}
}

type Login struct {
	Message     string `json:"message"`
	PartnerType string `json:"user_type"`
}

func NewLogin(_type string) *Login {
	return &Login{
		Message:     "Logged in!",
		PartnerType: _type,
	}
}

type AlbumInfo struct {
	AlbumId  string   `json:"album_id"`
	TrackIds []string `json:"track_ids"`
}

func NewAlbumInfo(albumId string, songIds []string) *AlbumInfo {
	return &AlbumInfo{
		AlbumId:  albumId,
		TrackIds: songIds,
	}
}

type Stream struct {
	Artist      string `json:"artist"`
	ProjecTitle string `json:"album_title"`
	TrackTitle  string `json:"track_title"`
	URL         string `json:"url"`
}

func NewStream(artist, albumTitle, trackTitle, url string) *Stream {
	return &Stream{
		Artist:      artist,
		ProjecTitle: albumTitle,
		TrackTitle:  trackTitle,
		URL:         url,
	}
}
