package hetzner

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func LoadOrCreateSSHKeys(ctx context.Context, client *hcloud.Client, paths []string) ([]*hcloud.SSHKey, error) {
	var keys []*hcloud.SSHKey

	for _, p := range paths {
		expanded := expandPath(p)

		content, err := os.ReadFile(expanded)
		if err != nil {
			return nil, fmt.Errorf("read ssh key %s: %w", expanded, err)
		}

		keyName := filepath.Base(expanded)

		// Try to find existing key
		sshKey, _, _ := client.SSHKey.GetByName(ctx, keyName)

		if sshKey == nil {
			fmt.Println("Uploading SSH key:", keyName)

			sshKey, _, err = client.SSHKey.Create(ctx, hcloud.SSHKeyCreateOpts{
				Name:      keyName,
				PublicKey: string(content),
				Labels: map[string]string{
					"managed-by": "hetzner-compose",
				},
			})
			if err != nil {
				return nil, err
			}
		}

		keys = append(keys, sshKey)
	}

	return keys, nil
}

func expandPath(p string) string {
	if len(p) > 0 && p[0] == '~' {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, p[2:])
	}
	return p
}
