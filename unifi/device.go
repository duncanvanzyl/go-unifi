package unifi

import (
	"context"
	"fmt"
	"log"
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

func (c *Client) ProvisoionDeviceByMAC(ctx context.Context, site, mac string) (interface{}, error) {
	reqBody := struct {
		MAC string `json:"mac"`
		CMD string `json:"cmd"`
	}{
		MAC: mac,
		CMD: "force-provision",
	}

	// var respBody struct {
	// 	Meta meta   `json:"meta"`
	// 	Data []Site `json:"data"`
	// }
	respBody := make(map[string]string)

	// POST https://192.168.1.8:8443/api/s/default/cmd/devmgr
	// {"mac":"b4:fb:e4:27:87:65","cmd":"force-provision"}
	// err := c.do(ctx, "GET", fmt.Sprintf("s/%s/stat/user/%s", site, mac), nil, &respBody)

	err := c.do(ctx, "POST", fmt.Sprintf("s/%s/default/cmd/devmgr", site), reqBody, &respBody)
	if err != nil {
		return nil, err
	}

	log.Fatalf("%+v", respBody)

	// return respBody.Data, nil
	return nil, nil
}
