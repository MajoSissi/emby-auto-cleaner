package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type EmbyLibrary struct {
	Name string `json:"Name"`
	ID   string `json:"ItemId"`
	Path string `json:"Path"`
}

func (c *EmbyClient) GetAllLibraries() ([]EmbyLibrary, error) {
	apiURL := fmt.Sprintf("%s/emby/Library/VirtualFolders", c.baseURL)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Emby-Token", c.authToken)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var libraries []EmbyLibrary
	if err := json.Unmarshal(body, &libraries); err != nil {
		return nil, err
	}

	return libraries, nil
}

func (c *EmbyClient) GetLibraryPath(libraryID string) (string, error) {
	libraries, err := c.GetAllLibraries()
	if err != nil {
		return "", err
	}

	for _, lib := range libraries {
		if lib.ID == libraryID && len(lib.Path) > 0 {
			return lib.Path, nil
		}
	}

	return "", fmt.Errorf("library not found: %s", libraryID)
}
