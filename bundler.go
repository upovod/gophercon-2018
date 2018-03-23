package wrong 

func (b *BundlerSuite) TestAdditionalBundle() {
	rpBundle := modelstest.CreateBundleFreeTraffic(b.db, b.Require())
	bundle := modelstest.CreateBundleAdditional(b.db, b.Require())
	defer modelstest.DeleteBundles(b.db, b.Require(), rpBundle, bundle)
	rp := modelstest.CreateRatePlan(b.db, b.Require(), rpBundle)
	defer modelstest.DeleteRatePlans(b.db, b.Require(), rp)
	to := modelstest.CreateTariffOption(b.db, b.Require(), rp, bundle)
	defer modelstest.DeleteTariffOptions(b.db, b.Require(), to)
	msisdn := modelstest.CreateMsisdnMegafon(b.db, b.Require(), rp, to)
	hh := modelstest.CreateHouseholdMegafon(b.db, b.Require(), msisdn)
	defer modelstest.DeleteHouseholds(b.db, b.Require(), b.mdp, hh)

	svc := New(b.db, b.config, b.logger)

	b.Require().NoError(svc.ProcessTpTo(b.db.Querier, hh, msisdn))
	b.checkBundleAssignment(hh, bundle, models.HouseholdBundleAdditionalTariffOptionReason)
	// many checks
}
