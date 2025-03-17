package euler1

import (
	"math/big"
	"strconv"
)

// largeSum returns an int whose digits are the first ten digits of the sum of the following one-hundred 50-digit numbers.
func largeSum() int {
	numbers := make([]big.Int, 100)

	numbers[0].SetString("37107287533902102798797998220837590246510135740250", 10)
	numbers[1].SetString("46376937677490009712648124896970078050417018260538", 10)
	numbers[2].SetString("74324986199524741059474233309513058123726617309629", 10)
	numbers[3].SetString("91942213363574161572522430563301811072406154908250", 10)
	numbers[4].SetString("23067588207539346171171980310421047513778063246676", 10)
	numbers[5].SetString("89261670696623633820136378418383684178734361726757", 10)
	numbers[6].SetString("28112879812849979408065481931592621691275889832738", 10)
	numbers[7].SetString("44274228917432520321923589422876796487670272189318", 10)
	numbers[8].SetString("47451445736001306439091167216856844588711603153276", 10)
	numbers[9].SetString("70386486105843025439939619828917593665686757934951", 10)
	numbers[10].SetString("62176457141856560629502157223196586755079324193331", 10)
	numbers[11].SetString("64906352462741904929101432445813822663347944758178", 10)
	numbers[12].SetString("92575867718337217661963751590579239728245598838407", 10)
	numbers[13].SetString("58203565325359399008402633568948830189458628227828", 10)
	numbers[14].SetString("80181199384826282014278194139940567587151170094390", 10)
	numbers[15].SetString("35398664372827112653829987240784473053190104293586", 10)
	numbers[16].SetString("86515506006295864861532075273371959191420517255829", 10)
	numbers[17].SetString("71693888707715466499115593487603532921714970056938", 10)
	numbers[18].SetString("54370070576826684624621495650076471787294438377604", 10)
	numbers[19].SetString("53282654108756828443191190634694037855217779295145", 10)
	numbers[20].SetString("36123272525000296071075082563815656710885258350721", 10)
	numbers[21].SetString("45876576172410976447339110607218265236877223636045", 10)
	numbers[22].SetString("17423706905851860660448207621209813287860733969412", 10)
	numbers[23].SetString("81142660418086830619328460811191061556940512689692", 10)
	numbers[24].SetString("51934325451728388641918047049293215058642563049483", 10)
	numbers[25].SetString("62467221648435076201727918039944693004732956340691", 10)
	numbers[26].SetString("15732444386908125794514089057706229429197107928209", 10)
	numbers[27].SetString("55037687525678773091862540744969844508330393682126", 10)
	numbers[28].SetString("18336384825330154686196124348767681297534375946515", 10)
	numbers[29].SetString("80386287592878490201521685554828717201219257766954", 10)
	numbers[30].SetString("78182833757993103614740356856449095527097864797581", 10)
	numbers[31].SetString("16726320100436897842553539920931837441497806860984", 10)
	numbers[32].SetString("48403098129077791799088218795327364475675590848030", 10)
	numbers[33].SetString("87086987551392711854517078544161852424320693150332", 10)
	numbers[34].SetString("59959406895756536782107074926966537676326235447210", 10)
	numbers[35].SetString("69793950679652694742597709739166693763042633987085", 10)
	numbers[36].SetString("41052684708299085211399427365734116182760315001271", 10)
	numbers[37].SetString("65378607361501080857009149939512557028198746004375", 10)
	numbers[38].SetString("35829035317434717326932123578154982629742552737307", 10)
	numbers[39].SetString("94953759765105305946966067683156574377167401875275", 10)
	numbers[40].SetString("88902802571733229619176668713819931811048770190271", 10)
	numbers[41].SetString("25267680276078003013678680992525463401061632866526", 10)
	numbers[42].SetString("36270218540497705585629946580636237993140746255962", 10)
	numbers[43].SetString("24074486908231174977792365466257246923322810917141", 10)
	numbers[44].SetString("91430288197103288597806669760892938638285025333403", 10)
	numbers[45].SetString("34413065578016127815921815005561868836468420090470", 10)
	numbers[46].SetString("23053081172816430487623791969842487255036638784583", 10)
	numbers[47].SetString("11487696932154902810424020138335124462181441773470", 10)
	numbers[48].SetString("63783299490636259666498587618221225225512486764533", 10)
	numbers[49].SetString("67720186971698544312419572409913959008952310058822", 10)
	numbers[50].SetString("95548255300263520781532296796249481641953868218774", 10)
	numbers[51].SetString("76085327132285723110424803456124867697064507995236", 10)
	numbers[52].SetString("37774242535411291684276865538926205024910326572967", 10)
	numbers[53].SetString("23701913275725675285653248258265463092207058596522", 10)
	numbers[54].SetString("29798860272258331913126375147341994889534765745501", 10)
	numbers[55].SetString("18495701454879288984856827726077713721403798879715", 10)
	numbers[56].SetString("38298203783031473527721580348144513491373226651381", 10)
	numbers[57].SetString("34829543829199918180278916522431027392251122869539", 10)
	numbers[58].SetString("40957953066405232632538044100059654939159879593635", 10)
	numbers[59].SetString("29746152185502371307642255121183693803580388584903", 10)
	numbers[60].SetString("41698116222072977186158236678424689157993532961922", 10)
	numbers[61].SetString("62467957194401269043877107275048102390895523597457", 10)
	numbers[62].SetString("23189706772547915061505504953922979530901129967519", 10)
	numbers[63].SetString("86188088225875314529584099251203829009407770775672", 10)
	numbers[64].SetString("11306739708304724483816533873502340845647058077308", 10)
	numbers[65].SetString("82959174767140363198008187129011875491310547126581", 10)
	numbers[66].SetString("97623331044818386269515456334926366572897563400500", 10)
	numbers[67].SetString("42846280183517070527831839425882145521227251250327", 10)
	numbers[68].SetString("55121603546981200581762165212827652751691296897789", 10)
	numbers[69].SetString("32238195734329339946437501907836945765883352399886", 10)
	numbers[70].SetString("75506164965184775180738168837861091527357929701337", 10)
	numbers[71].SetString("62177842752192623401942399639168044983993173312731", 10)
	numbers[72].SetString("32924185707147349566916674687634660915035914677504", 10)
	numbers[73].SetString("99518671430235219628894890102423325116913619626622", 10)
	numbers[74].SetString("73267460800591547471830798392868535206946944540724", 10)
	numbers[75].SetString("76841822524674417161514036427982273348055556214818", 10)
	numbers[76].SetString("97142617910342598647204516893989422179826088076852", 10)
	numbers[77].SetString("87783646182799346313767754307809363333018982642090", 10)
	numbers[78].SetString("10848802521674670883215120185883543223812876952786", 10)
	numbers[79].SetString("71329612474782464538636993009049310363619763878039", 10)
	numbers[80].SetString("62184073572399794223406235393808339651327408011116", 10)
	numbers[81].SetString("66627891981488087797941876876144230030984490851411", 10)
	numbers[82].SetString("60661826293682836764744779239180335110989069790714", 10)
	numbers[83].SetString("85786944089552990653640447425576083659976645795096", 10)
	numbers[84].SetString("66024396409905389607120198219976047599490197230297", 10)
	numbers[85].SetString("64913982680032973156037120041377903785566085089252", 10)
	numbers[86].SetString("16730939319872750275468906903707539413042652315011", 10)
	numbers[87].SetString("94809377245048795150954100921645863754710598436791", 10)
	numbers[88].SetString("78639167021187492431995700641917969777599028300699", 10)
	numbers[89].SetString("15368713711936614952811305876380278410754449733078", 10)
	numbers[90].SetString("40789923115535562561142322423255033685442488917353", 10)
	numbers[91].SetString("44889911501440648020369068063960672322193204149535", 10)
	numbers[92].SetString("41503128880339536053299340368006977710650566631954", 10)
	numbers[93].SetString("81234880673210146739058568557934581403627822703280", 10)
	numbers[94].SetString("82616570773948327592232845941706525094512325230608", 10)
	numbers[95].SetString("22918802058777319719839450180888072429661980811197", 10)
	numbers[96].SetString("77158542502016545090413245809786882778948721859617", 10)
	numbers[97].SetString("72107838435069186155435662884062257473692284509516", 10)
	numbers[98].SetString("20849603980134001723930671666823555245252804609722", 10)
	numbers[99].SetString("53503534226472524250874054075591789781264330331690", 10)

	// Sum.
	sum := big.NewInt(0)
	for _, val := range numbers {
		sum.Add(sum, &val)
	}

	digits, err := strconv.Atoi(sum.String()[:10]) // Get first 10 'digits' (Characters of this string can be manipulated easily as a slice).
	if err != nil {
		panic(err)
	}

	return digits
}

// Go makes this quite easy. This returns an int instead of individual ints so it is easier to test.
// Possible to use maths.GetDigitsBig but this is easier to get the digits as one number and no arithmetic on the digits is needed.
