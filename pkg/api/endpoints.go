package api

// func (c *Client) GetAbilityScores(ctx context.Context) (*ResourceList, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ability-scores", c.BaseURL), nil)
// 	if err != nil {
// 		return nil, oops.Wrapf(err, "unable to build request")
// 	}

// 	req = req.WithContext(ctx)

// 	res := ResourceList{}
// 	if err := c.sendRequest(req, &res); err != nil {
// 		return nil, oops.Wrapf(err, "unable to get ability-scores")
// 	}

// 	return &res, nil
// }

// func (c *Client) GetAbilityScoreByIndex(ctx context.Context, index string) (*apiAbilityScore, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ability-scores/%s", c.BaseURL, index), nil)
// 	if err != nil {
// 		return nil, oops.Wrapf(err, "unable to build request")
// 	}

// 	req = req.WithContext(ctx)

// 	res := apiAbilityScore{}
// 	if err := c.sendRequest(req, &res); err != nil {
// 		return nil, oops.Wrapf(err, "unable to get ability score")
// 	}

// 	return &res, nil
// }
