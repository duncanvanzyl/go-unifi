package unifi

import (
	"context"
	"fmt"
)

func (c *Client) ListDevice(ctx context.Context, site string) ([]Device, error) {
	return c.listDevice(ctx, site)
}

func (c *Client) GetDeviceByMAC(ctx context.Context, site, mac string) (*Device, error) {
	return c.getDevice(ctx, site, mac)
}

func (c *Client) DeleteDevice(ctx context.Context, site, id string) error {
	return c.deleteDevice(ctx, site, id)
}

func (c *Client) CreateDevice(ctx context.Context, site string, d *Device) (*Device, error) {
	return c.createDevice(ctx, site, d)
}

func (c *Client) UpdateDevice(ctx context.Context, site string, d *Device) (*Device, error) {
	return c.updateDevice(ctx, site, d)
}

func (c *Client) GetDevice(ctx context.Context, site, id string) (*Device, error) {
	devices, err := c.ListDevice(ctx, site)

	if err != nil {
		return nil, err
	}

	for _, d := range devices {
		if d.ID == id {
			return &d, nil
		}
	}

	return nil, &NotFoundError{}
}

func (c *Client) ProvisoionDeviceByMAC(ctx context.Context, site, mac string) error {
	reqBody := struct {
		MAC string `json:"mac"`
		CMD string `json:"cmd"`
	}{
		MAC: mac,
		CMD: "force-provision",
	}

	var respBody struct {
		Meta meta `json:"meta"`
	}

	err := c.do(ctx, "POST", fmt.Sprintf("s/%s/cmd/devmgr", site), reqBody, &respBody)
	if err != nil {
		return err
	}

	return respBody.Meta.error()
}
