package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type ScnLoader struct {
	TrainList TrainList `xml:"trainList"`
}

type TrainList struct {
	TrainLoader TrainLoader `xml:"TrainLoader"`
}

type TrainLoader struct {
	UnitLoaderList UnitLoaderList `xml:"unitLoaderList"`
}

type UnitLoaderList struct {
	Vehicles []RailVehicle `xml:"RailVehicleStateClass"`
}

type RailVehicle struct {
	FileName string  `xml:"rvXMLfilename"`
	UnitType string  `xml:"unitType"`
	Weight   float64 `xml:"loadWeightUSTons"`
}

func parseTrain(filename string) []RailVehicle {
	xmlFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		log.Fatal(err)
	}
	var d ScnLoader
	xml.Unmarshal(bytes, &d)
	return d.TrainList.TrainLoader.UnitLoaderList.Vehicles
}

// These values are scraped using the parse.go tool in the tools/ directory.
func carWeight(car string) float64 {
	weights := map[string]float64{
		"R8_CoveredHopper_ACF2700_CSX1":         28,
		"R8_CoveredHopper_ACF2700_CSX2":         28,
		"R8_CoveredHopper_ACF2700_DTS01":        28,
		"R8_CoveredHopper_ACF2970_BN01":         28,
		"R8_CoveredHopper_ACF2970_BNSF01":       28,
		"R8_CoveredHopper_ACF2970_CNW01":        28,
		"R8_CoveredHopper_ACF2970_NW01":         28,
		"R8_CoveredHopper_ACF2970_SP01":         28,
		"R8_CoveredHopper_ACF2970_WW01":         28,
		"R8_CoveredHopper_Trin3281_CEFX":        26,
		"R8_CoveredHopper_Trin3281_CEMX":        26,
		"R8_CoveredHopper_Trin3281_GACX01":      27,
		"R8_CoveredHopper_Trin3281_MBKX01":      26,
		"R8_CoveredHopper_Trin3281_MCEX01":      26,
		"R8_CoveredHopper_Trin3281_NS":          26.0,
		"R8_CoveredHopper_Trin3281_TILX":        26.0,
		"R8_Autorack_ATSF_ETTX1":                50.0,
		"R8_AutoRack_Bi_CN01":                   50,
		"R8_AutoRack_Bi_CR01":                   50,
		"R8_AutoRack_Bi_CSX01":                  50,
		"R8_AutoRack_Bi_UP01":                   50,
		"R8_Autorack_CNW_TTGX1":                 50.0,
		"R8_Autorack_CP_TTGX1":                  50.0,
		"R8_Autorack_DRGW_TTGX1":                50.0,
		"R8_Autorack_GTW01":                     50.0,
		"R8_Autorack_NW_ETTX1":                  50.0,
		"R8_Autorack_SLSF_ETTX1":                50.0,
		"R8_AutoRack_Tri_BN01":                  53,
		"R8_AutoRack_Tri_BNSF01":                53,
		"R8_AutoRack_Tri_NS01":                  53,
		"R8_AutoRack_Tri_SP01":                  53,
		"R8_Autorack_WP_ETTX1":                  50.0,
		"R8_BallastHopper_BSC2200_UP01":         31,
		"R8_BallastHopper_PS2003_ATSF01":        25,
		"R8_BallastHopper_PS2003_BN01":          25,
		"R8_BallastHopper_PS2003_CSX01":         25,
		"R8_BallastHopper_PS2003_SPMW01":        25,
		"R8_Bethgon_BNSF01":                     25,
		"R8_Bethgon_CMO01":                      25,
		"R8_Bethgon_SEMX01":                     25,
		"R8_Boxcar_50ft_PlateF_AOK01":           37,
		"R8_Boxcar_50ft_PlateF_BN01":            37,
		"R8_Boxcar_50ft_PlateF_BNSF01":          37,
		"R8_Boxcar_50ft_PlateF_BNSF02":          37,
		"R8_Boxcar_50ft_PlateF_CAT01":           37,
		"R8_Boxcar_50ft_PlateF_CEFX01":          37,
		"R8_Boxcar_50ft_PlateF_CN01":            37,
		"R8_Boxcar_50ft_PlateF_CNW01":           37,
		"R8_Boxcar_50ft_PlateF_CR01":            37,
		"R8_Boxcar_50ft_PlateF_CSX01":           37,
		"R8_Boxcar_50ft_PlateF_IB01":            37,
		"R8_Boxcar_50ft_PlateF_NOKL01":          37,
		"R8_Boxcar_50ft_PlateF_R801":            37,
		"R8_Boxcar_50ft_PlateF_SP01":            37,
		"R8_Boxcar_50ft_PlateF_TTX01":           37,
		"R8_Boxcar_50ft_PlateF_WC01":            37,
		"R8_Boxcar_FMC5347_CAGY01":              31,
		"R8_Boxcar_FMC5347_CNW01":               31,
		"R8_Boxcar_FMC5347_EACH01":              31,
		"R8_Boxcar_FMC5347_MTW01":               31,
		"R8_Boxcar_FMC5347_NOPB01":              31,
		"R8_Boxcar_FMC5347_RBOX01":              31,
		"R8_Boxcar_FMC5347_SP01":                31,
		"R8_Boxcar_FMC5347_SSW01":               31,
		"R8_Caboose_c509_SP01":                  30,
		"R8_Caboose_c509_SP02":                  30,
		"R8_Caboose_c509_SP03":                  30,
		"R8_Caboose_ce11_ATSF01":                30,
		"R8_Caboose_ce11_ATSF02":                30,
		"R8_Caboose_ce11_ATSF03":                30,
		"R8_Centerbeam_BC01":                    31,
		"R8_Centerbeam_CN01":                    31,
		"R8_Centerbeam_IC01":                    31,
		"R8_Centerbeam_NOKL01":                  31,
		"R8_Centerbeam_TTZX01":                  31,
		"R8_Centerbeam_TTZX02":                  31,
		"R8_Hopper_BSC3350_LN02":                32,
		"R8_Hopper_BSC3420_ATSF":                29,
		"R8_Hopper_BSC3420_ATSFrock":            29,
		"R8_Hopper_BSC3420_BNSF":                29,
		"R8_Hopper_BSC3483_BN":                  34,
		"R8_Hopper_BSC3483_DRGW":                34,
		"R8_Hopper_BSC3483_DRGW02":              34,
		"R8_Hopper_BSC3600_CSX":                 34,
		"R8_Hopper_BSC3600_UP01":                34,
		"R8_Hopper_BSC3600_UP02":                34,
		"R8_C14Hopper_AGP01":                    30,
		"R8_C14Hopper_ASHX01":                   30,
		"R8_C14Hopper_BN01":                     30,
		"R8_C14Hopper_BNSF01":                   30,
		"R8_C14Hopper_BNSF02":                   30,
		"R8_C14Hopper_CAGX01":                   30,
		"R8_C14Hopper_CARG01":                   30,
		"R8_C14Hopper_CPR02":                    30,
		"R8_C14Hopper_CSX01":                    30,
		"R8_C14Hopper_DME01":                    30,
		"R8_C14Hopper_GAT01":                    30,
		"R8_C14Hopper_POTASH01":                 30,
		"R8_C14Hopper_Run8_01":                  30,
		"R8_C14Hopper_SOY01":                    30,
		"R8_C14Hopper_UP01":                     30,
		"R8_C14Hopper_UP02":                     30,
		"R8_Gon_ACF1995_EJE01":                  38,
		"R8_Gon_ALB2494_UP01":                   33,
		"R8_Gon_GSC3242_CR01":                   36,
		"R8_Gon_GSC3242_MP01":                   36,
		"R8_Gon_GSC3242_SP01":                   36,
		"R8_Gon_THR2494_GONX01":                 38,
		"R8_Gon_THR2743_CSX01":                  29,
		"R8_Gon_THR2743_DJJX01":                 29,
		"R8_Gon_TRIN1828_BNSF01":                31,
		"R8_Gon_TRIN1828_IHB01":                 38,
		"R8_Hyrail_MOW_BNSF_Ford":               2.5,
		"Run8_52ftwell_1China":                  27,
		"Run8_52ftwell_1Hanjin":                 27,
		"Run8_52ftwell_1JBHunt":                 27,
		"Run8_52ftwell_1Maersk":                 27,
		"Run8_52ftwell_1NFI":                    27,
		"Run8_52ftwell_1Schneider":              27,
		"Run8_52ftwell_2China":                  27,
		"Run8_52ftwell_2China_NFI":              27,
		"Run8_52ftwell_2Evergreen":              27,
		"Run8_52ftwell_2Hanjin":                 27,
		"Run8_52ftwell_2JBHunt":                 27,
		"Run8_52ftwell_2JBHunt_Schneider":       27,
		"Run8_52ftwell_2KLine":                  27,
		"Run8_52ftwell_2Maersk":                 27,
		"Run8_52ftwell_2Maersk_Hyundai":         27,
		"Run8_52ftwell_2NFI":                    27,
		"Run8_52ftwell_53_40_China":             27,
		"Run8_52ftwell_53_40_China_Hyundai":     27,
		"Run8_52ftwell_53_40_China_Maersk":      27,
		"Run8_52ftwell_53_40_China_NFI":         27,
		"Run8_52ftwell_53_40_Evergreen":         27,
		"Run8_52ftwell_53_40_Hanjin":            27,
		"Run8_52ftwell_53_40_HorizonHanjin":     27,
		"Run8_52ftwell_53_40_Hyundai":           27,
		"Run8_52ftwell_53_40_HyundaiChina":      27,
		"Run8_52ftwell_53_40_HyundaiNFI":        27,
		"Run8_52ftwell_53_40_JBHunt":            27,
		"Run8_52ftwell_53_40_JBHuntSchneider":   27,
		"Run8_52ftwell_53_40_JBHuntSeastar":     27,
		"Run8_52ftwell_53_40_Kline":             27,
		"Run8_52ftwell_53_40_Maersk":            27,
		"Run8_52ftwell_53_40_MaerskChina":       27,
		"Run8_52ftwell_53_40_MaerskHyundai":     27,
		"Run8_52ftwell_53_40_MaerskNFI":         27,
		"Run8_52ftwell_53_40_NFI":               27,
		"Run8_52ftwell_53_40_NFIChina":          27,
		"Run8_52ftwell_53_40_NFIHyundai":        27,
		"Run8_52ftwell_53_40_NFIMaersk":         27,
		"Run8_52ftwell_53_40_Schneider":         27,
		"Run8_52ftwell_53_40_SchneiderJBHunt":   27,
		"Run8_52ftwell_53_40_SeastarEvergreen":  27,
		"Run8_52ftwell_A_Empty":                 27,
		"Run8_52ftwell_Doublestack53":           27,
		"Run8_WellCar53AP_FEC01":                25,
		"Run8_WellCar53AP_FEC02":                25,
		"Run8_WellCar53AP_FEC03":                25,
		"Run8_WellCar53AP_FEC04":                25,
		"Run8_WellCar53AP_FEC08":                25,
		"Run8_WellCar53AP_FEC09":                25,
		"Run8_WellCar53AP_FEC10":                25,
		"Run8_WellCar53AP_FEC11":                25,
		"Run8_WellCar53AP_FEC12":                25,
		"Run8_WellCar53AP_FEC14":                25,
		"Run8_WellCar53AP_FEC15":                25,
		"Run8_WellCar53AP_FEC16":                25,
		"Run8_WellCar53AP_FEC17":                25,
		"Run8_WellCar53AP_FEC18":                25,
		"Run8_WellCar53_DTTX01":                 25,
		"Run8_WellCar53_DTTX02":                 25,
		"Run8_WellCar53_DTTX03":                 25,
		"Run8_WellCar53_DTTX04":                 25,
		"Run8_WellCar53_DTTX07":                 25,
		"Run8_WellCar53_DTTX08":                 25,
		"Run8_WellCar53_DTTX09":                 25,
		"Run8_WellCar53_DTTX10":                 25,
		"Run8_WellCar53_DTTX11":                 25,
		"R8_GP40-2_BNSF01":                      0,
		"R8_GP40-2_CON01":                       0,
		"R8_GP40-2_CSX01":                       0,
		"R8_GP40-2_CSX02":                       0,
		"R8_GP40-2_DRGW01":                      0,
		"R8_GP40-2_RUN802":                      0,
		"R8_GP40-2_SP01":                        0,
		"R8_GP40-2_UP01":                        0,
		"R8_GP40-2_WP01":                        0,
		"R8_SD40-2_ATSF01":                      0,
		"R8_SD40-2_ATSF02":                      0,
		"R8_SD40-2_BN01":                        0,
		"R8_SD40-2_BNSF01":                      0,
		"R8_SD40-2_BNSF02":                      0,
		"R8_SD40-2_BNSF03":                      0,
		"R8_SD40-2_BNSF04":                      0,
		"R8_SD40-2_CARG01":                      0,
		"R8_SD40-2_Chess":                       0,
		"R8_SD40-2_CN01":                        0,
		"R8_SD40-2_CON01":                       0,
		"R8_SD40-2_CSX01":                       0,
		"R8_SD40-2_CSX02":                       0,
		"R8_SD40-2_FEC01":                       0,
		"R8_SD40-2_GCFX01":                      0,
		"R8_SD40-2_NS01":                        0,
		"R8_SD40-2_RUNW01":                      0,
		"R8_SD40-2_SBS01":                       0,
		"R8_SD40-2_SOO02":                       0,
		"R8_SD40-2_SP01":                        0,
		"R8_SD40-2_UP01":                        0,
		"R8_SD40T-2_DRGW01":                     0,
		"R8_SD40T-2_R801":                       0,
		"R8_SD40T-2_SP01":                       0,
		"R8_SD40T-2_UP01":                       0,
		"R8_SD40T-2_UP02":                       0,
		"R8_SD40T-2_UP03":                       0,
		"R8_SD45-2_ACR01":                       0,
		"R8_SD45-2_ATSF01":                      0,
		"R8_SD45-2_ATSF02":                      0,
		"R8_SD45-2_ATSF03":                      0,
		"R8_SD45-2_ATSF04":                      0,
		"R8_SD45-2_BNSF01":                      0,
		"R8_SD45-2_BNSF02":                      0,
		"R8_SD45-2_CON01":                       0,
		"R8_SD45-2_CON02":                       0,
		"R8_SD45-2_CSX01":                       0,
		"R8_SD45-2_CSX02":                       0,
		"R8_SD45-2_CSX03":                       0,
		"R8_SD45-2_NS01":                        0,
		"R8_SD45-2_Run801":                      0,
		"R8_SD45-2_SBS01":                       0,
		"R8_SD45-2_SCL01":                       0,
		"R8_SD45-2_TRON01":                      0,
		"Run8_ES44DC":                           0,
		"Run8_ES44DC_BNSF":                      0,
		"Run8_ES44DC_BNSF02":                    0,
		"Run8_ES44DC_CSX":                       0,
		"Run8_ES44DC_NS":                        0,
		"Run8_ES44DC_SP":                        0,
		"Run8_ES44DC_UP":                        0,
		"Run8_ES44DC_UP01":                      0,
		"Run8_P42_Amtrak03":                     0,
		"Run8_P42_Amtrak04":                     0,
		"Run8_P42_Amtrak05":                     0,
		"Run8_SD70ACE_BNSF01":                   0,
		"Run8_SD70ACE_CSX01":                    0,
		"Run8_SD70ACE_KCS01":                    0,
		"Run8_SD70ACE_NS01":                     0,
		"Run8_SD70ACE_NS_JC01":                  0,
		"Run8_SD70ACE_RUNW01":                   0,
		"Run8_SD70ACE_RUNW02":                   0,
		"Run8_SD70ACE_UP01":                     0,
		"Run8_SD70ACE_UP02":                     0,
		"Run8_SD70ACE_UP_DRGW01":                0,
		"Run8_SD70ACE_UP_Katy01":                0,
		"Run8_SD70ACE_UP_SP":                    0,
		"Run8_SD70ACE_UP_WP01":                  0,
		"R8_Amtrak_AmfleetCoachIII01":           53,
		"R8_Amtrak_AmfleetCoachIV01":            53,
		"R8_Amtrak_AmfleetCoachIVb01":           53,
		"R8_Amtrak_AmfleetLoungeIII01":          53,
		"R8_Amtrak_AmfleetLoungeIV01":           53,
		"R8_Amtrak_AmfleetLoungeIVb01":          53,
		"R8_Amtrak_Baggage03":                   85,
		"R8_Amtrak_Baggage04":                   85,
		"R8_Amtrak_Baggage04b":                  85,
		"R8_Amtrak_CchBagPhsIII":                85,
		"R8_Amtrak_CchBagPhsIV":                 85,
		"R8_Amtrak_CchBagPhsV":                  85,
		"R8_Amtrak_CoachPhsIII":                 85,
		"R8_Amtrak_CoachPhsIV":                  85,
		"R8_Amtrak_CoachPhsV":                   85,
		"R8_Amtrak_DinerPhsIII":                 85,
		"R8_Amtrak_DinerPhsIV":                  85,
		"R8_Amtrak_DinerPhsV":                   85,
		"R8_Amtrak_LoungePhsIII":                85,
		"R8_Amtrak_LoungePhsIV":                 85,
		"R8_Amtrak_LoungePhsV":                  85,
		"R8_Amtrak_MHC03":                       57,
		"R8_Amtrak_MHC04":                       57,
		"R8_Amtrak_SleeperPhsIII":               85,
		"R8_Amtrak_SleeperPhsIV":                85,
		"R8_Amtrak_SleeperPhsV":                 85,
		"R8_Amtrak_TransPhsIII":                 85,
		"R8_Amtrak_TransPhsIV":                  85,
		"R8_Amtrak_TransPhsV":                   85,
		"R8_AutoRack_Amtrak01":                  41.3,
		"R8_AutoRack_Amtrak02":                  41.3,
		"R8_Pig_ATSF01":                         31,
		"R8_Pig_KTTX01":                         35,
		"R8_Pig_KTTX02":                         35,
		"R8_Pig_RTTX01":                         36,
		"R8_Pig_RTTX02":                         36,
		"R8_Pig_RTTX03":                         36,
		"R8_Pig_RTTX04":                         36,
		"R8_Pig_RTTX05":                         36,
		"R8_Pig_RTTX06":                         36,
		"R8_Pig_RTTX07":                         36,
		"R8_Pig_TTAX01":                         35,
		"R8_Pig_TTWX01":                         35,
		"R8_Pig_TTWX02":                         35,
		"R8_Pig_TTWX03":                         35,
		"R8_Pig_TTWX04":                         35,
		"R8_Pig_TTWX05":                         35,
		"R8_Pig_TTWX06":                         35,
		"R8_PlasticPelletHopper_ACF5250_ACFX01": 33,
		"R8_PlasticPelletHopper_ACF5250_ACFX02": 33,
		"R8_PlasticPelletHopper_ACF5250_AMCX01": 33,
		"R8_PlasticPelletHopper_ACF5250_BPRX01": 33,
		"R8_PlasticPelletHopper_ACF5250_ETCX01": 33,
		"R8_PlasticPelletHopper_ACF5250_KCIX01": 33,
		"R8_Reefer_Greenbrier_72_CRYX":          50,
		"R8_Reefer_Millenium_64_TPIX":           48,
		"R8_Reefer_PCF_57_rbld_ARMN":            44,
		"R8_Reefer_PCF_57_SFRC":                 44,
		"R8_Reefer_PCF_57_SPFE":                 44,
		"R8_Reefer_PCF_57_UPFE":                 44,
		"R8_Reefer_TrinCool_64_ARMN":            46,
		"R8_Reefer_TrinCool_72_BNSF":            53,
		"Run8_Tank105_Acid":                     36,
		"Run8_Tank105_AcidDirt":                 36,
		"Run8_Tank105_AGFX":                     36,
		"Run8_Tank105_BNSF":                     36,
		"Run8_Tank105_Cstar":                    36,
		"Run8_Tank105_UTLX":                     36,
		"Run8_Tank107AGP01":                     40,
		"Run8_Tank107BN01":                      40,
		"Run8_Tank107BNSF01":                    40,
		"Run8_Tank107FRS01":                     40,
		"Run8_Tank107GATX02":                    40,
		"Run8_Tank107GATX03":                    40,
		"Run8_Tank107PROC02":                    40,
		"Run8_Tank107PROCR01":                   40,
		"Run8_Tank107RC01":                      40,
		"Run8_Tank107SanFe01":                   40,
		"Run8_Tank107UTLXblk":                   40,
		"Run8_Tank107UTLXwht":                   40,
		"R8_ChipHopper_GSC7000_CSXT01":          41,
		"R8_ChipHopper_GSC7000_GPSX01":          41,
		"R8_ChipHopper_GSC7000_IFRX01":          41,
	}

	return weights[car]
}

// In order, the ints in the 6-item arrays are the rows in which we will
// find engine counts for each engine in the front, middle, and rear
// engine groups, first for counting axles, then for counting brakes.  This
// is really ugly and can definitely be cleaned up if this thing persists.
func bnsfCells(model string) [6]int {
	var cells [6]int
	switch model {
	case "ES44DC":
		cells = [6]int{3, 10, 17, 30, 37, 44}
	case "SD70ACE":
		cells = [6]int{4, 11, 18, 31, 38, 45}
	case "SD40-2":
		cells = [6]int{5, 12, 19, 32, 39, 46}
	case "SD40T-2":
		cells = [6]int{5, 12, 19, 41, 48, 55}
	case "SD45-2":
		cells = [6]int{6, 13, 20, 33, 40, 47}
	case "GP40-2":
		cells = [6]int{7, 14, 21, 34, 41, 48}
	default:
		cells = [6]int{0, 0, 0, 0, 0, 0}
	}
	return cells

}

func upCells(model string) [6]int {
	var cells [6]int
	switch model {
	case "ES44DC":
		cells = [6]int{3, 10, 17, 39, 46, 53}
	case "SD70ACE":
		cells = [6]int{4, 11, 18, 40, 47, 54}
	case "SD40-2":
		cells = [6]int{5, 12, 19, 41, 48, 55}
	case "SD40T-2":
		cells = [6]int{5, 12, 19, 41, 48, 55}
	case "SD45-2":
		cells = [6]int{6, 13, 20, 42, 49, 56}
	case "GP40-2":
		cells = [6]int{7, 14, 21, 43, 50, 57}
	default:
		cells = [6]int{0, 0, 0, 0, 0, 0}
	}
	return cells

}
