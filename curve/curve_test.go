package curve

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
	"testing"

	"github.com/anima-protocol/starknet.go/utils"
)

// BenchmarkPedersenHash benchmarks the performance of the PedersenHash function.
//
// The function takes a 2D slice of big.Int values as input and measures the time
// it takes to execute the PedersenHash function for each test case.
//
// Parameters:
// - b: a *testing.B value representing the testing context
// Returns:
//  none
func BenchmarkPedersenHash(b *testing.B) {
	suite := [][]*big.Int{
		{utils.HexToBN("0x12773"), utils.HexToBN("0x872362")},
		{utils.HexToBN("0x1277312773"), utils.HexToBN("0x872362872362")},
		{utils.HexToBN("0x1277312773"), utils.HexToBN("0xCD2a3d9F938E13CD947Ec05AbC7FE734Df8DD826")},
		{utils.HexToBN("0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB"), utils.HexToBN("0x872362872362")},
		{utils.HexToBN("0xCD2a3d9F938E13CD947Ec05AbC7FE734Df8DD826"), utils.HexToBN("0xbBbBBBBbbBBBbbbBbbBbbbbBBbBbbbbBbBbbBBbB")},
		{utils.HexToBN("0x7f15c38ea577a26f4f553282fcfe4f1feeb8ecfaad8f221ae41abf8224cbddd"), utils.HexToBN("0x13d41f388b8ea4db56c5aa6562f13359fab192b3db57651af916790f9debee9")},
		{utils.HexToBN("0x7f15c38ea577a26f4f553282fcfe4f1feeb8ecfaad8f221ae41abf8224cbddd"), utils.HexToBN("0x7f15c38ea577a26f4f553282fcfe4f1feeb8ecfaad8f221ae41abf8224cbdde")},
	}

	for _, test := range suite {
		b.Run(fmt.Sprintf("input_size_%d_%d", test[0].BitLen(), test[1].BitLen()), func(b *testing.B) {
			Curve.PedersenHash(test)
		})
	}
}

// BenchmarkCurveSign benchmarks the Curve.Sign function.
//
// Parameters:
// - b: a *testing.B value representing the testing context
// Returns:
//  none
func BenchmarkCurveSign(b *testing.B) {
	type data struct {
		MessageHash *big.Int
		PrivateKey  *big.Int
		Seed        *big.Int
	}

	dataSet := []data{}
	MessageHash := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(250), nil)
	PrivateKey := big.NewInt(0).Add(MessageHash, big.NewInt(1))
	Seed := big.NewInt(0)
	for i := int64(0); i < 20; i++ {
		dataSet = append(dataSet, data{
			MessageHash: big.NewInt(0).Add(MessageHash, big.NewInt(i)),
			PrivateKey:  big.NewInt(0).Add(PrivateKey, big.NewInt(i)),
			Seed:        big.NewInt(0).Add(Seed, big.NewInt(i)),
		})

		for _, test := range dataSet {
			Curve.Sign(test.MessageHash, test.PrivateKey, test.Seed)
		}
	}
}

// BenchmarkSignatureVerify benchmarks the SignatureVerify function.
//
// The function takes a testing.B object as a parameter and performs a series
// of operations to benchmark the SignatureVerify function. It generates a
// random private key, computes the corresponding public key, computes a hash
// using the PedersenHash function, signs the hash using the private key, and
// finally verifies the signature. The function runs two benchmarks: one for
// signing and one for verification. Each benchmark measures the time taken to
// perform the respective operation.
//
// Parameters:
// - b: a *testing.B value representing the testing context
// Returns:
//  none
func BenchmarkSignatureVerify(b *testing.B) {
	private, _ := Curve.GetRandomPrivateKey()
	x, y, _ := Curve.PrivateToPoint(private)

	hash, _ := Curve.PedersenHash(
		[]*big.Int{
			utils.HexToBN("0x7f15c38ea577a26f4f553282fcfe4f1feeb8ecfaad8f221ae41abf8224cbddd"),
			utils.HexToBN("0x7f15c38ea577a26f4f553282fcfe4f1feeb8ecfaad8f221ae41abf8224cbdde"),
		})

	r, s, _ := Curve.Sign(hash, private)

	b.Run(fmt.Sprintf("sign_input_size_%d", hash.BitLen()), func(b *testing.B) {
		Curve.Sign(hash, private)
	})
	b.Run(fmt.Sprintf("verify_input_size_%d", hash.BitLen()), func(b *testing.B) {
		Curve.Verify(hash, r, s, x, y)
	})
}

// TestGeneral_PrivateToPoint tests the PrivateToPoint function.
//
// Parameters:
// - t: a *testing.T value representing the testing context
// Returns:
//  none
func TestGeneral_PrivateToPoint(t *testing.T) {
	x, _, err := Curve.PrivateToPoint(big.NewInt(2))
	if err != nil {
		t.Errorf("PrivateToPoint err %v", err)
	}
	expectedX, _ := new(big.Int).SetString("3324833730090626974525872402899302150520188025637965566623476530814354734325", 10)
	if x.Cmp(expectedX) != 0 {
		t.Errorf("Actual public key %v different from expected %v", x, expectedX)
	}
}

// TestGeneral_PedersenHash is a test function for the PedersenHash method in the General struct.
//
// The function tests the PedersenHash method by providing different test cases and comparing the computed hash with the expected hash.
// It uses the testing.T type from the testing package to report any errors encountered during the tests.
//
// Parameters:
// - t: a *testing.T value representing the testing context
// Returns:
//  none
func TestGeneral_PedersenHash(t *testing.T) {
	testPedersen := []struct {
		elements []*big.Int
		expected *big.Int
	}{
		{
			elements: []*big.Int{utils.HexToBN("0x12773"), utils.HexToBN("0x872362")},
			expected: utils.HexToBN("0x5ed2703dfdb505c587700ce2ebfcab5b3515cd7e6114817e6026ec9d4b364ca"),
		},
		{
			elements: []*big.Int{utils.HexToBN("0x13d41f388b8ea4db56c5aa6562f13359fab192b3db57651af916790f9debee9"), utils.HexToBN("0x537461726b4e6574204d61696c")},
			expected: utils.HexToBN("0x180c0a3d13c1adfaa5cbc251f4fc93cc0e26cec30ca4c247305a7ce50ac807c"),
		},
		{
			elements: []*big.Int{big.NewInt(100), big.NewInt(1000)},
			expected: utils.HexToBN("0x45a62091df6da02dce4250cb67597444d1f465319908486b836f48d0f8bf6e7"),
		},
	}

	for _, tt := range testPedersen {
		hash, err := Curve.PedersenHash(tt.elements)
		if err != nil {
			t.Errorf("Hashing err: %v\n", err)
		}
		if hash.Cmp(tt.expected) != 0 {
			t.Errorf("incorrect hash: got %v expected %v\n", hash, tt.expected)
		}
	}
}

// TestGeneral_DivMod tests the DivMod function.
//
// The function takes in a list of test cases which consist of inputs x, y, and the expected output.
// The inputs x and y are of type *big.Int and the expected output is also of type *big.Int.
// The function iterates through each test case and calls the DivMod function with the inputs x, y, and the prime number Curve.P.
// It then compares the output of DivMod with the expected output and throws an error if they are not equal.
// The function is used to test the correctness of the DivMod function.
//
// Parameters:
// - t: a *testing.T value representing the testing context
// Returns:
//  none
func TestGeneral_DivMod(t *testing.T) {
	testDivmod := []struct {
		x        *big.Int
		y        *big.Int
		expected *big.Int
	}{
		{
			x:        utils.StrToBig("311379432064974854430469844112069886938521247361583891764940938105250923060"),
			y:        utils.StrToBig("621253665351494585790174448601059271924288186997865022894315848222045687999"),
			expected: utils.StrToBig("2577265149861519081806762825827825639379641276854712526969977081060187505740"),
		},
		{
			x:        big.NewInt(1),
			y:        big.NewInt(2),
			expected: utils.HexToBN("0x0400000000000008800000000000000000000000000000000000000000000001"),
		},
	}

	for _, tt := range testDivmod {
		divR := DivMod(tt.x, tt.y, Curve.P)

		if divR.Cmp(tt.expected) != 0 {
			t.Errorf("DivMod Res %v does not == expected %v\n", divR, tt.expected)
		}
	}
}

// TestGeneral_Add tests the Add function in the General package.
//
// It tests the addition of two big integers and compares the result with the expected values.
// The function takes a slice of test cases, each containing two big integers and their expected sum.
// It iterates over the test cases, computes the sum using the Add function, and checks if it matches the expected sum.
// If the computed sum does not match the expected sum, an error is reported using the t.Errorf function.
//
// Parameters:
// - t: a *testing.T value representing the testing context
// Returns:
//  none
func TestGeneral_Add(t *testing.T) {
	testAdd := []struct {
		x         *big.Int
		y         *big.Int
		expectedX *big.Int
		expectedY *big.Int
	}{
		{
			x:         utils.StrToBig("1468732614996758835380505372879805860898778283940581072611506469031548393285"),
			y:         utils.StrToBig("1402551897475685522592936265087340527872184619899218186422141407423956771926"),
			expectedX: utils.StrToBig("2573054162739002771275146649287762003525422629677678278801887452213127777391"),
			expectedY: utils.StrToBig("3086444303034188041185211625370405120551769541291810669307042006593736192813"),
		},
		{
			x:         big.NewInt(1),
			y:         big.NewInt(2),
			expectedX: utils.StrToBig("225199957243206662471193729647752088571005624230831233470296838210993906468"),
			expectedY: utils.StrToBig("190092378222341939862849656213289777723812734888226565973306202593691957981"),
		},
	}

	for _, tt := range testAdd {
		resX, resY := Curve.Add(Curve.Gx, Curve.Gy, tt.x, tt.y)
		if resX.Cmp(tt.expectedX) != 0 {
			t.Errorf("ResX %v does not == expected %v\n", resX, tt.expectedX)

		}
		if resY.Cmp(tt.expectedY) != 0 {
			t.Errorf("ResY %v does not == expected %v\n", resY, tt.expectedY)
		}
	}
}

// TestGeneral_MultAir is a test function that tests the MultAir function in the General package.
//
// It tests the behavior of the MultAir function with different input values and checks if the
// returned values match the expected values.
// The MultAir function takes a big.Int value 'r', and two other big.Int values 'x' and 'y', and
// performs a mathematical operation on them to calculate new big.Int values 'x' and 'y'.
// The function then compares the calculated values with the expected values and reports any
// mismatches as test failures.
//
// Parameters:
// - t: a *testing.T value representing the testing context
// Returns:
//  none
func TestGeneral_MultAir(t *testing.T) {
	testMult := []struct {
		r         *big.Int
		x         *big.Int
		y         *big.Int
		expectedX *big.Int
		expectedY *big.Int
	}{
		{
			r:         utils.StrToBig("2458502865976494910213617956670505342647705497324144349552978333078363662855"),
			x:         utils.StrToBig("1468732614996758835380505372879805860898778283940581072611506469031548393285"),
			y:         utils.StrToBig("1402551897475685522592936265087340527872184619899218186422141407423956771926"),
			expectedX: utils.StrToBig("182543067952221301675635959482860590467161609552169396182763685292434699999"),
			expectedY: utils.StrToBig("3154881600662997558972388646773898448430820936643060392452233533274798056266"),
		},
	}

	for _, tt := range testMult {
		x, y, err := Curve.MimicEcMultAir(tt.r, tt.x, tt.y, Curve.Gx, Curve.Gy)
		if err != nil {
			t.Errorf("MultAirERR %v\n", err)
		}

		if x.Cmp(tt.expectedX) != 0 {
			t.Errorf("ResX %v does not == expected %v\n", x, tt.expectedX)

		}
		if y.Cmp(tt.expectedY) != 0 {
			t.Errorf("ResY %v does not == expected %v\n", y, tt.expectedY)
		}
	}
}

// TestGeneral_ComputeHashOnElements is a test function that verifies the correctness of the ComputeHashOnElements function in the General package.
//
// This function tests the ComputeHashOnElements function by passing in different arrays of big.Int elements and comparing the computed hash with the expected hash.
// It checks the behavior of the ComputeHashOnElements function when an empty array is passed as input, as well as when an array with multiple elements is passed.
// The expected hashes are precalculated using the utils.HexToBN function.
//
// Parameters:
// - t: a *testing.T value representing the testing context
// Returns:
//  none
func TestGeneral_ComputeHashOnElements(t *testing.T) {
	hashEmptyArray, err := Curve.ComputeHashOnElements([]*big.Int{})
	expectedHashEmmptyArray := utils.HexToBN("0x49ee3eba8c1600700ee1b87eb599f16716b0b1022947733551fde4050ca6804")
	if err != nil {
		t.Errorf("Could no hash an empty array %v\n", err)
	}
	if hashEmptyArray.Cmp(expectedHashEmmptyArray) != 0 {
		t.Errorf("Hash empty array wrong value. Expected %v got %v\n", expectedHashEmmptyArray, hashEmptyArray)
	}

	hashFilledArray, err := Curve.ComputeHashOnElements([]*big.Int{
		big.NewInt(123782376),
		big.NewInt(213984),
		big.NewInt(128763521321),
	})
	expectedHashFilledArray := utils.HexToBN("0x7b422405da6571242dfc245a43de3b0fe695e7021c148b918cd9cdb462cac59")

	if err != nil {
		t.Errorf("Could no hash an array with values %v\n", err)
	}
	if hashFilledArray.Cmp(expectedHashFilledArray) != 0 {
		t.Errorf("Hash filled array wrong value. Expected %v got %v\n", expectedHashFilledArray, hashFilledArray)
	}
}

// TestGeneral_HashAndSign is a test function that verifies the hashing and signing process.
//
// Parameters:
// - t: The testing.T object for running the test.
// Returns:
//   none
func TestGeneral_HashAndSign(t *testing.T) {
	hashy, err := Curve.HashElements([]*big.Int{
		big.NewInt(1953658213),
		big.NewInt(126947999705460),
		big.NewInt(1953658213),
	})
	if err != nil {
		t.Errorf("Hasing elements: %v\n", err)
	}

	priv, _ := Curve.GetRandomPrivateKey()
	x, y, err := Curve.PrivateToPoint(priv)
	if err != nil {
		t.Errorf("Could not convert random private key to point: %v\n", err)
	}

	r, s, err := Curve.Sign(hashy, priv)
	if err != nil {
		t.Errorf("Could not convert gen signature: %v\n", err)
	}

	if !Curve.Verify(hashy, r, s, x, y) {
		t.Errorf("Verified bad signature %v %v\n", r, s)
	}
}

// TestGeneral_ComputeFact tests the ComputeFact function.
//
// It tests the ComputeFact function by providing a set of test cases
// and comparing the computed hash with the expected hash.
// The test cases consist of program hashes, program outputs,
// and expected hash values.
//
// Parameters:
// - t: The testing.T object for running the test
// Returns:
//   none
func TestGeneral_ComputeFact(t *testing.T) {
	testFacts := []struct {
		programHash   *big.Int
		programOutput []*big.Int
		expected      *big.Int
	}{
		{
			programHash:   utils.HexToBN("0x114952172aed91e59f870a314e75de0a437ff550e4618068cec2d832e48b0c7"),
			programOutput: []*big.Int{big.NewInt(289)},
			expected:      utils.HexToBN("0xe6168c0a865aa80d724ad05627fa65fbcfe4b1d66a586e9f348f461b076072c4"),
		},
		{
			programHash:   utils.HexToBN("0x79920d895101ad1fbdea9adf141d8f362fdea9ee35f33dfcd07f38e4a589bab"),
			programOutput: []*big.Int{utils.StrToBig("2754806153357301156380357983574496185342034785016738734224771556919270737441")},
			expected:      utils.HexToBN("0x1d174fa1443deea9aab54bbca8d9be308dd14a0323dd827556c173bd132098db"),
		},
	}

	for _, tt := range testFacts {
		hash := utils.ComputeFact(tt.programHash, tt.programOutput)
		if hash.Cmp(tt.expected) != 0 {
			t.Errorf("Fact does not equal ex %v %v\n", hash, tt.expected)
		}
	}
}

// TestGeneral_BadSignature tests the behavior of the function that checks for bad signatures.
//
// Parameters:
// - t: The testing.T object for running the test
// Returns:
//   none
func TestGeneral_BadSignature(t *testing.T) {
	hash, err := Curve.PedersenHash([]*big.Int{utils.HexToBN("0x12773"), utils.HexToBN("0x872362")})
	if err != nil {
		t.Errorf("Hashing err: %v\n", err)
	}

	priv, _ := Curve.GetRandomPrivateKey()
	x, y, err := Curve.PrivateToPoint(priv)
	if err != nil {
		t.Errorf("Could not convert random private key to point: %v\n", err)
	}

	r, s, err := Curve.Sign(hash, priv)
	if err != nil {
		t.Errorf("Could not convert gen signature: %v\n", err)
	}

	badR := new(big.Int).Add(r, big.NewInt(1))
	if Curve.Verify(hash, badR, s, x, y) {
		t.Errorf("Verified bad signature %v %v\n", r, s)
	}

	badS := new(big.Int).Add(s, big.NewInt(1))
	if Curve.Verify(hash, r, badS, x, y) {
		t.Errorf("Verified bad signature %v %v\n", r, s)
	}

	badHash := new(big.Int).Add(hash, big.NewInt(1))
	if Curve.Verify(badHash, r, s, x, y) {
		t.Errorf("Verified bad signature %v %v\n", r, s)
	}
}

// TestGeneral_Signature tests the Signature function.
//
// testSignature is a test struct containing private, publicX, publicY,
// hash, rIn, sIn, and raw fields.
// The function iterates over the testSignature slice and performs various
// operations including signing, verifying, converting, and initializing
// variables.
//
// Parameters:
// - t: The testing.T object for running the test
// Returns:
//   none
func TestGeneral_Signature(t *testing.T) {
	testSignature := []struct {
		private *big.Int
		publicX *big.Int
		publicY *big.Int
		hash    *big.Int
		rIn     *big.Int
		sIn     *big.Int
		raw     string
	}{
		{
			private: utils.StrToBig("104397037759416840641267745129360920341912682966983343798870479003077644689"),
			publicX: utils.StrToBig("1913222325711601599563860015182907040361852177892954047964358042507353067365"),
			publicY: utils.StrToBig("798905265292544287704154888908626830160713383708400542998012716235575472365"),
			hash:    utils.StrToBig("2680576269831035412725132645807649347045997097070150916157159360688041452746"),
			rIn:     utils.StrToBig("607684330780324271206686790958794501662789535258258105407533051445036595885"),
			sIn:     utils.StrToBig("453590782387078613313238308551260565642934039343903827708036287031471258875"),
		},
		{
			hash: utils.HexToBN("0x7f15c38ea577a26f4f553282fcfe4f1feeb8ecfaad8f221ae41abf8224cbddd"),
			rIn:  utils.StrToBig("2458502865976494910213617956670505342647705497324144349552978333078363662855"),
			sIn:  utils.StrToBig("3439514492576562277095748549117516048613512930236865921315982886313695689433"),
			raw:  "04033f45f07e1bd1a51b45fc24ec8c8c9908db9e42191be9e169bfcac0c0d997450319d0f53f6ca077c4fa5207819144a2a4165daef6ee47a7c1d06c0dcaa3e456",
		},
		{
			hash:    utils.HexToBN("0x324df642fcc7d98b1d9941250840704f35b9ac2e3e2b58b6a034cc09adac54c"),
			publicX: utils.HexToBN("0x4e52f2f40700e9cdd0f386c31a1f160d0f310504fc508a1051b747a26070d10"),
			rIn:     utils.StrToBig("2849277527182985104629156126825776904262411756563556603659114084811678482647"),
			sIn:     utils.StrToBig("3156340738553451171391693475354397094160428600037567299774561739201502791079"),
		},
	}

	var err error
	for _, tt := range testSignature {
		if tt.raw != "" {
			h, _ := utils.HexToBytes(tt.raw)
			tt.publicX, tt.publicY = elliptic.Unmarshal(Curve, h)
		} else if tt.private != nil {
			tt.publicX, tt.publicY, err = Curve.PrivateToPoint(tt.private)
			if err != nil {
				t.Errorf("Could not convert random private key to point: %v\n", err)
			}
		} else if tt.publicX != nil {
			tt.publicY = Curve.GetYCoordinate(tt.publicX)
		}

		if tt.rIn == nil && tt.private != nil {
			tt.rIn, tt.sIn, err = Curve.Sign(tt.hash, tt.private)
			if err != nil {
				t.Errorf("Could not sign good hash: %v\n", err)
			}
		}

		if !Curve.Verify(tt.hash, tt.rIn, tt.sIn, tt.publicX, tt.publicY) {
			t.Errorf("successful signature did not verify\n")
		}
	}
}

// TestGeneral_SplitFactStr is a test function that tests the SplitFactStr function.
//
// It verifies the behavior of the SplitFactStr function by providing different inputs and checking the output.
// The function takes no parameters and returns no values.
//
// Parameters:
// - t: The testing.T object for running the test
// Returns:
//   none
func TestGeneral_SplitFactStr(t *testing.T) {
	data := []map[string]string{
		{"input": "0x3", "h": "0x0", "l": "0x3"},
		{"input": "0x300000000000000000000000000000000", "h": "0x3", "l": "0x0"},
	}
	for _, d := range data {
		l, h := utils.SplitFactStr(d["input"]) // 0x3
		if l != d["l"] {
			t.Errorf("expected %s, got %s", d["l"], l)
		}
		if h != d["h"] {
			t.Errorf("expected %s, got %s", d["h"], h)
		}
	}
}
