package rinth

import (
	"context"
	"fmt"
	"net/http"
)

// Mod represents a Modrinth's mod page.
type Mod struct {
	ID          *string   `json:"id,omitempty"`          //The ID of the mod, encoded as a base62 string
	Slug        *string   `json:"slug,omitempty"`        //The slug of a mod, used for vanity URLs
	Team        *string   `json:"team,omitempty"`        //The id of the team that has ownership of this mod
	Title       *string   `json:"title,omitempty"`       //The title or name of the mod
	Description *string   `json:"description,omitempty"` //A short description of the mod
	Body        *string   `json:"body,omitempty"`        //A long form description of the mod.
	BodyUrl     *string   `json:"body_url,omitempty"`    //DEPRECATED The link to the long description of the mod (Optional)
	Published   *string   `json:"published,omitempty"`   //The date at which the mod was first published
	Updated     *string   `json:"updated,omitempty"`     //The date at which the mod was updated
	Status      *string   `json:"status,omitempty"`      //The status of the mod - approved, rejected, draft, unlisted, processing, or unknown
	ClientSide  *string   `json:"client_side,omitempty"` //The support range for the client mod - required, optional, unsupported, or unknown
	ServerSide  *string   `json:"server_side,omitempty"` //The support range for the server mod - required, optional, unsupported, or unknown
	Downloads   *int      `json:"downloads,omitempty"`   //The total number of downloads the mod has
	Categories  *[]string `json:"categories,omitempty"`  //A list of the categories that the mod is in
	Versions    *[]string `json:"versions,omitempty"`    //A list of ids for versions of the mod
	IconUrl     *string   `json:"icon_url,omitempty"`    //The URL of the icon of the mod (Optional)
	IssuesUrl   *string   `json:"issues_url,omitempty"`  //An optional link to where to submit bugs or issues with the mod (Optional)
	SourceUrl   *string   `json:"source_url,omitempty"`  //An optional link to the source code for the mod (Optional)
	WikiUrl     *string   `json:"wiki_url,omitempty"`    //An optional link to the mod's wiki page or other relevant information (Optional)
	DiscordUrl  *string   `json:"discord_url,omitempty"` //An optional link to the mod's discord (Optional)
}

// GetMod gets a Mod for a given Mod ID.
func (c *Client) GetMod(ctx context.Context, id string) (*Mod, *http.Response, error) {
	mod := new(Mod)
	path := fmt.Sprintf("mod/%s", id)
	rsp, err := c.Do(ctx, path, mod)
	if err != nil {
		return &Mod{}, rsp, err
	}
	return mod, rsp, nil
}
