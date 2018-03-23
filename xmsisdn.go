func (s *XMSISDNSuite) TestHandleUserSameXmsisdn() {
	// Preparing
	encDevID, err := s.xmsisdn.encryptDeviceID(dev.ID)
	s.Require().NoError(err)

	xmsisdn := msisdn.Msisdn
	changed, err := s.xmsisdn.Handle(encDevID, xmsisdn, "")
	s.Require().NoError(err)
	s.Require().True(changed, "state should have changed")
	s.Require().NoError(s.db.Reload(dev))
	s.Require().Equal(msisdn.Msisdn, *dev.XMSISDN, "msisdn should match xmsisdn")

	changed, err = s.xmsisdn.Handle(encDevID, xmsisdn, "")
	s.Require().NoError(err)
	s.Require().False(changed, "state should not have changed")
	// ...	

	changed, err = s.xmsisdn.Handle(encDevID, xmsisdn, "")
	// Next checks
}
