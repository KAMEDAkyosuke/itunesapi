// Code generated by "stringer -type=StorefrontCode"; DO NOT EDIT

package country

import "fmt"

const (
	_StorefrontCode_name_0 = "USFRDEGBATBEFIGRIEITLUNLPTESCASENODKCHAUNZJPHKSGCNKRINMXRUTWVNZAMYPHTHIDPKPLSATRAEHUCLNPPALKROMVCZBDILUAKWHRCRSKLBQASIRSCOVEBRGTARSVPEDOECHNJMNIPYUYMOEGKZEELVLTMTLIMDAMBWBGCIJOKEMKMGMLMUNESNTNUGAIBSAGBBBMVGKYDMGDMSKNLCVCTTTCGYSRBZBOCYISBHBNNGOMDZAOBYUZ"
	_StorefrontCode_name_1 = "AZ"
	_StorefrontCode_name_2 = "YETZGH"
)

var (
	_StorefrontCode_index_0 = [...]uint8{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142, 144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164, 166, 168, 170, 172, 174, 176, 178, 180, 182, 184, 186, 188, 190, 192, 194, 196, 198, 200, 202, 204, 206, 208, 210, 212, 214, 216, 218, 220, 222, 224, 226, 228, 230, 232, 234, 236, 238, 240, 242, 244, 246, 248, 250, 252}
	_StorefrontCode_index_1 = [...]uint8{0, 2}
	_StorefrontCode_index_2 = [...]uint8{0, 2, 4, 6}
)

func (i StorefrontCode) String() string {
	switch {
	case 143441 <= i && i <= 143566:
		i -= 143441
		return _StorefrontCode_name_0[_StorefrontCode_index_0[i]:_StorefrontCode_index_0[i+1]]
	case i == 143568:
		return _StorefrontCode_name_1
	case 143571 <= i && i <= 143573:
		i -= 143571
		return _StorefrontCode_name_2[_StorefrontCode_index_2[i]:_StorefrontCode_index_2[i+1]]
	default:
		return fmt.Sprintf("StorefrontCode(%d)", i)
	}
}
