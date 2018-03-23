func (s *PromocodesSuite) TestSuccesfulActivation() {
	createPromocode(s.T(), films, films, startDate, finalDate, defaultUT, filmMech, 0)
	// preparing
	promocodes := map[string][]string{
		softtnb: {"popup_pcode_tnb_soft_info", "popup_pcode_tnb_soft_info_activated"},
		hardtnb: {"popup_pcode_tnb_hard_info", "popup_pcode_tnb_hard_info_activated"},
		film:    {"popup_pcode_film_info", "popup_pcode_film_info_activated"},
		films:   {"popup_pcode_film_multi_info", "popup_pcode_film_multi_info_activated"},
	}
	for code, results := range promocodes {
		url := fmt.Sprintf("/promocodes/%s", code)
		answer := GetWithCookie(s.T(), url, s.User1Cookie, 200).JSON(s.T()).(jsons.Object)
		s.Assert().Equal(results[0], answer["code"], "Something went wrong. %s", answer)
		answer = PostWithCookie(s.T(), url+"/activate", jsons.Parse("{}"), s.User1Cookie, 200).JSON(s.T()).(jsons.Object)
		s.Assert().Equal(results[1], answer["code"], "Something went wrong. %s", answer)
	}
}

