package ssd

import "testing"

var ssdModelMap map[modelKey][]string
var correctMatches []pairs
var notConfident []string

type pairs struct {
	postTitle     string
	brandAndModel string
}

func init() {
	ssdModelMap = prepareProcessedData()
	correctMatches = []pairs{
		{postTitle: "[SSD] Crucial P3 4TB PCIe Gen3 3D NAND NVMe M.2 SSD, up to 3500MB/s (334.99 ATL) [Amazon.ca]", brandAndModel: "Crucial P3"},
		{postTitle: "[SSD] Crucial P3 Plus (blah) [Amazon.ca]", brandAndModel: "Crucial P3 Plus"},
		{postTitle: "[SSD] Western Digital 4TB WD Blue 3D NAND ($309.99) [Amazon/Newegg]", brandAndModel: "WD Blue 3D"},
		{postTitle: "[SSD] XPG 2TB GAMMIX S70 Blade - Gen4 NVMe SSD Up to 7,400 MB/s ($199) (Amazon)", brandAndModel: "ADATA S70/S70 Blade"},
		{postTitle: "[SSD] 1.92TB SSD 2.5\" Patriot SATA 3 Burst Elite ($107.99) [Amazon]", brandAndModel: "Patriot Burst Elite"},
		{postTitle: "[SSD] Patriot P210 SATA 3 2TB SSD 2.5 Inch ($109.99) [Amazon]", brandAndModel: "Patriot P210"},
		{postTitle: "[SSD] Crucial MX500 4TB (299.99) [Newegg]", brandAndModel: "Crucial MX500"},
		{postTitle: "[SSD] WD_BLACK 2TB SN770 NVMe Internal Gaming SSD Solid State Drive - Gen4 PCIe ($179) [Amazon]", brandAndModel: "WD SN770"},
		{postTitle: "[SSD] ADATA Legend 850 1TB PCIe Gen4 x4 NVMe 1.4 M.2 Internal Gaming SSD Up to 5,000 MB/s ($94) [Amazon]", brandAndModel: "ADATA Legend 850"},
		{postTitle: "[SSD] Kingston NV2 1TB M.2 2280 NVMe Internal SSD | PCIe 4.0 ($77) [Amazon]", brandAndModel: "Kingston NV2"},
		{postTitle: "[SSD] Lexar NM800 Pro 1 TB / 2 TB NVMe Gen 4 - 7500 / 6500 MB/s ($98 / $192) [Amazon]", brandAndModel: "Lexar NM800 (NM800PRO)"},
		{postTitle: "[SSD] Patriot P310 960GB Internal SSD - NVMe PCIe M.2 Gen3 x 4 ($69.99-$4=$65.99)", brandAndModel: "Patriot P310"},
		{postTitle: "[SSD] TEAMGROUP T-Create Classic 2TB M.2 PCIe 2280 NVMe 1.3 Internal SSD, Up to 2100MB/s Design for Creators Gen3x4 Solid State Drive, Terabyte Written TBW 1000TB ($136) [Amazon]", brandAndModel: "Team T-Create Classic (NVMe)"},
		{postTitle: "[Nvme] patriot p310 1.92TB ($137.99) [Amazon]", brandAndModel: "Patriot P310"},
		{postTitle: "[NVMe SSD] Lexar MB610 Pro 2TB M.2 2280 ($160 - $30 = $130) (F/S) [Canada Computers]", brandAndModel: "Lexar NM610 Pro"},
		{postTitle: "[NVMe SSD] Lexar NM610 Pro 2TB M.2 2280 ($160 - $30 = $130) (F/S) [Canada Computers]", brandAndModel: "Lexar NM610 Pro"},
		{postTitle: "[NVMe SSD] Lexar NM610 2TB M.2 2280 ($160 - $30 = $130) (F/S) [Canada Computers]", brandAndModel: "Lexar NM610"},
		{postTitle: "[SSD] TEAMGROUP MP33 2TB SLC Cache 3D NAND TLC NVMe ($139.99) (ATL) [Amazon.ca]", brandAndModel: "Team MP33"},
		{postTitle: "[SSD] TEAMGROUP AX2 2TB 3D NAND TLC 2.5 Inch SATA III SSD ($113) (Amazon)", brandAndModel: "Team AX2"},
		{postTitle: "[NVMe] Corsair MP600 CORE XT 4TB PCIe Gen4 x4 NVMe M.2 SSD QLC M.2 2280 Up to 5,000MB/sec ($355.64) [Amazon.ca]", brandAndModel: "Corsair MP600 Core XT"},
		{postTitle: "[ssd] corsair mp600 1tb nvme HMB TLC 4800 MB/s/ 3900 MB/s ($100-25=$75+fs) [amazon] ", brandAndModel: "Corsair MP600 GS"},
		{postTitle: "[SSD] MSI SPATIUM M371 2TB NVMe M.2 Read:2350MB/s Write:1900MB/s Solid State Drive(SM371N2TB)($129.99 - $20 = $109.99) [Canadacomputers] ", brandAndModel: "MSI M371"},
		{postTitle: "[NVME] Team Group MP44L M.2 2280 2TB PCIe 4.0 x4 TLC ($140-20-$10 coupon=$110) [Newegg.ca] ATL", brandAndModel: "Team MP44L"},
	}
}

func TestCorrectMatches(t *testing.T) {
	for _, pair := range correctMatches {
		score, matches := getMatching(pair.postTitle, ssdModelMap)
		if len(matches) != 1 {
			t.Errorf("Matches was not exactly 1 (was %d): %s,%s\n", len(matches), pair.brandAndModel, pair.postTitle)
			t.Error(matches)
			t.FailNow()
		}
		if score < 1000 {
			t.Errorf("Score was less than 100: %s,%s\n", pair.brandAndModel, pair.postTitle)
		}
		resultingTitle := matches[0][0] + " " + matches[0][1]
		if pair.brandAndModel != resultingTitle {
			t.Errorf("[%s] Did not match expected title [%s]", resultingTitle, pair.brandAndModel)
		}
	}
}
