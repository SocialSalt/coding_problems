package maximizepartitions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChooseLongest(t *testing.T) {
	s := "accca"
	n := maxPartitionsAfterOperations(s, 2)
	require.Equal(t, 3, n)
	fmt.Printf("%d, %d\n", countPartitions(s, 2), n)

	s = "aabaab"
	n = maxPartitionsAfterOperations(s, 3)
	require.Equal(t, 1, n)
	fmt.Printf("%d, %d\n", countPartitions(s, 3), n)

	s = "xxyz"
	n = maxPartitionsAfterOperations(s, 1)
	require.Equal(t, 4, n)
	fmt.Printf("%d, %d\n", countPartitions(s, 1), n)

	s = "noyynxgvtkhxsqdqcjyecjpwcawkgsrxmixokubliztvglyftkcrkpdfofwhaydetelrlyzirwmcjlnghqzsepsztnshfsanwezyrwugjtupaukeqhnqjuuyzlixhzewymafxyjasqlfvvabungssaylgcxydwvnwcayoogevdkpkxbvofwgohtjocqhtykbrpurqxqvwyxdxxqhstlbkcuohtkmlyqfdzcbatmshcpoeoqirqtyuabiwrtyprucmfpcezmawmjhsskexpzlnasejilkjtbwuylzdpunifykhyteoglauzfaljvndlpeubkxtmnisawrdlzfcvfljdrtnzwhyuelqdtbgjvrublexxslrckupnwznerwanngvfppxnayeorsgnozapmgnsbzuxmaeoyrfwhhsdnxsflqklbtopradhxgadzjrrdutduhiurdjaovkgtulcjndpcibywdzwxucxakouievplehkdkdhpnfgjqrrjcwdnwgfujzpkihjjvxrdtluuxdpzwwgdifhzvuuhpoe"
	n = maxPartitionsAfterOperations(s, 22)
	require.Equal(t, 11, n)
	fmt.Printf("%d, %d\n", countPartitions(s, 22), n)

	s = "ertiujdmcaijesimdzdnpvodcoteinkvmmoinjmzckcpgegjqjvurmyoxnvjqlqnegjpkmkiwzhkiggbohnwicpcizfmfijywewbgikyvtfhcafaghbognoklwdlpaoxqolxjvyxtzczchbgdufzdhdignxcvqbuvdxeplnwxvcdqdqbephouhqzgcbayqxurleubgahcqdhjkjdxbgcpcbtfglpqrpuonxwhskjanvgrdboujlkcydgsvdalzgqeclkraamvzamhjphhreqpbfwrrbyblgcixoxjresmgikpqgwaiairzghjtuobbsrwrildngpdqapvdvqsqbkqiempfvfffstrraayfahghhwasglkgpujfxfwdioahmfhxhxoasvdruawutyznuagatpagincaxyiywkylqxrgzqfeewknrnudoawojnsonsnaazzriifxobfrmoakpmffgcdtkbsdbfojpocimblnwpegggpybtxwjleygspjjpchorkgzoekvggipbzdijpiwplqmzmunrgibpgk"
	n = maxPartitionsAfterOperations(s, 24)
	require.Equal(t, 8, n)
	fmt.Printf("%d, %d\n", countPartitions(s, 24), n)

}
