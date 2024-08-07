package remoteconfigworker

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// fetch calls the supergood /config endpoint and returns a marshalled config object
func (rc *RemoteConfigWorker) fetch() ([]TenantConfig, error) {
	url, err := url.JoinPath(rc.baseURL, "/proxyconfig")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("sg-admin-api-key", rc.adminClientKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return nil, fmt.Errorf("supergood: invalid ADMIN_CLIENT_KEY")
	} else if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body, _ := io.ReadAll(resp.Body)
		message := string(body)
		return nil, fmt.Errorf("supergood: got HTTP %v posting to /config with error: %s", resp.Status, message)
	}

	var remoteConfigArray []TenantConfig
	err = json.NewDecoder(resp.Body).Decode(&remoteConfigArray)
	if err != nil {
		return nil, err
	}

	return remoteConfigArray, nil
}
