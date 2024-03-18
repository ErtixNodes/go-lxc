package main

func (c *Client) InstallServer(name string, password string, port int) error {
	err := setPassword(name, password)
	if err != nil {
		return err
	}

	err = removeDefaultUser(name)
	if err != nil {
		return err
	}

	err = initSSH(name, port)
	if err != nil {
		return err
	}

	return nil
}
