package btc

import (
	"bytes"
	"testing"
	"math/big"
	"encoding/hex"
)

var _stealth_vecs = [][6]string { // x, y, k (z always 1) -> x,y,x
	{
		"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		"483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8",
		"a8d65dd8aaa1ab31c91fb18c8a18ac7a07bf649ec5ff15ef9255ad682912cada",
		"cda5a5dff4459558c8edc20e57bc5babc2a39dcba125fc75196bb5e0e04b4b54",
		"5d7612473f88d6cd522b09903ae08912aec39c4d7a508fa442382c6d57d58d24",
		"eaed65fd764f7d3d06068906b3402463785da475b995e284a48b4365b1c99151",
	},
	{
		"f46a67e20804f956a1ce64566d96a42658a9a7a4c9a0be924615bef881a4a3f2",
		"3a8218cdf4156c60585f5721189289cc89500eab79480a109eb1d0684e560996",
		"84e5f7d329c3dab1160dbf9cb0b1a3c82e6058c06260f4101b1660b865ce98c5",
		"eb8d6a5c12e70b0d5e05336e9103318e89ca4445004afc3640d3e47e488a4d0f",
		"8fabd090eade40104431906d3bc0c25d988270aa017bfa8ce3707c0d72649571",
		"631bcaec3954477e0dd17b3fa60d395a23ea7a47a7ad602c6372a6b6efb41ee3",
	},
	{
		"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		"483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8",
		"ac5e6a9ad86a9f31a37201887d9c9b0f3e183d230ffcf2e31137cb00acc1c105",
		"3628a0d7d8621b1d6c60293b3aaea5fcebc6360f4e3252094267dae1ec831a60",
		"e0c2e6fa2b164a0aa7ce31f37e7cc1d2d4431a13cef69559f0f62931066fdbff",
		"cc0764cb49940d20edda6b87178c102e1a949a44ea8524110657dce311e6270d",
	},
	{
		"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		"483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8",
		"a23b57ac32e4f90e20de6d011fbb4c628016b3436fdb860917e7e019d0bc7126",
		"06760669f905712bb497e561d57104f44eabce9716738dee764507913eda5ce8",
		"4f8bac8b946f46a9580614a9770987cae19eb9b0fac0590056e4f4d8e923ee5b",
		"db76a101c645e72e1367cd1d4bed7c205c1674735199ec253d13223959aa9f8c",
	},
	{
		"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		"483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8",
		"a0b769f8d15509488c34236491a5290514153dc55898c25e09ebb8f812e8d962",
		"4782a82cea022b3e0d4c5bd2b22433d3330b49df60ff71790510824050065d16",
		"782f7b1e5b5aa9339561587699741bbc9a74b0b4c29b79b126904d21bcf0f97c",
		"5d3358cd8d8ab4f93507fbe55c7af957a5e2e8927fa5dc9648521cd65836e7dc",
	},
	{
		"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		"483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8",
		"a23b57ac32e4f90e20de6d011fbb4c628016b3436fdb860917e7e019d0bc7126",
		"06760669f905712bb497e561d57104f44eabce9716738dee764507913eda5ce8",
		"4f8bac8b946f46a9580614a9770987cae19eb9b0fac0590056e4f4d8e923ee5b",
		"db76a101c645e72e1367cd1d4bed7c205c1674735199ec253d13223959aa9f8c",
	},
	{
		"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		"483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8",
		"3e0dff8baa536e7ed0005812ceb31f42e92650097a7425f1bc5e088368f46002",
		"3b168354a6f1c539683bc46ea95caa1752f103000266592ae968de5b7d6d267e",
		"e6bc807f4e12303f3f685290ac17b9a441b8fd42b4e9ce9fb2f7d24d11821609",
		"ea3e96224b96b7eeda982cdf2f344ef987004ea7c6749264478a1d56909d6b8b",
	},
	{
		"f46a67e20804f956a1ce64566d96a42658a9a7a4c9a0be924615bef881a4a3f2",
		"3a8218cdf4156c60585f5721189289cc89500eab79480a109eb1d0684e560996",
		"34936baac3c1bb6acfc65a595598d3eac1b6c47cbfd787b48ad465e40bab7d19",
		"3e094f13aea9f6d685148666acf35093583bee5864152263ff4603566968fe97",
		"91cd020d975267b9aa1657b108d83f4a33589d86aaaee3d3bd940bdaf06bb3b3",
		"9db345949444cfee35ff18f3a8d7393e043080fe8bd1a688ea909e71d2a44771",
	},
	{
		"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		"483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8",
		"34936baac3c1bb6acfc65a595598d3eac1b6c47cbfd787b48ad465e40bab7d19",
		"cdd4fa818a172d8d24307ebdc702f129ac2c5e4e8e7e2280f4cf53c17a149bd9",
		"d3a05479deebdac0fe3b60b1b3adf0d3452ed97ef770967259093cc869fce796",
		"68d73a2c4b6bb050d4c9cfbb4ba4b504e44f7847b41015e1c54553c0d007add0",
	},
	{
		"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		"483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8",
		"b122d72964cccf20ada6a0ff027f0d4cc27f634fa7303413aeecd626836e84c3",
		"2af1923c7cb97506feb0194028952872d2c47dca0920ef8033909fc70456405b",
		"4ac497dab5b2369bb838b73ab24297a2303add3f089cde4bcf8974c7b791445b",
		"609dfdbac6c252b39f95901d87e6920ebcd8829de5c45dd9d7f3fc2c7431b58d",
	},
	{
		"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		"483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8",
		"880d885eb89bd3f9b62fe36ba156d7173f2f50c06f943bc9c35ae5bdb702bd52",
		"503712dc1c4f429666c9680a0d388b07e662d8ac468a8512ac593839df0a5519",
		"d12c09e41918aed68432b85423f5d9eb3723a44a1d7e7781bbdb45e220695a1b",
		"50bf864cc3b2d632af55ac525bfb3bd47df598f86173be44de60e25f05f5da05",
	},
	{
		"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		"483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8",
		"cc5c48983181592a029f16c8e6d7ffa60e1125dafd056f22ac8dbc1b38e220ff",
		"efa5dd369b7d0a065c8b2b680ce949abf41919644d419bfd9bef0cdf64010b13",
		"926d1611bf8f45acf23c80157bd01a557f74210b683f201b9644c94c2f800c25",
		"e8f2cdfae8adf5ba8f9f9307de06d01b61fd384049b6350048bcb1cb0bd5c6b9",
	},
	{
		"79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		"483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8",
		"18c7b931d63f144b7e42489026f25219be661e0618d31cc1c5eae191c9a1bd07",
		"43c10a9910a260a041b911eaa33c1516754c8ec718aeb454b3a005a3e42f3fe0",
		"438bd9784b2fb80653e03c7d83be31b60bb724a4066fe0d86f5a810591f7462b",
		"1eea6b56b0b757e57d01b796882135ecd9ff6ecde0699c45de53b287edf9429c",
	},
}


func TestStealthTwice(t *testing.T) {
	var a, exp, r xyz_t
	var k big.Int

	k.SetString("84e5f7d329c3dab1160dbf9cb0b1a3c82e6058c06260f4101b1660b865ce98c5", 16)
	a.x, _ = new(big.Int).SetString("f46a67e20804f956a1ce64566d96a42658a9a7a4c9a0be924615bef881a4a3f2", 16)
	a.y, _ = new(big.Int).SetString("3a8218cdf4156c60585f5721189289cc89500eab79480a109eb1d0684e560996", 16)
	a.z, _ = new(big.Int).SetString("01", 16)

	exp.x, _ = new(big.Int).SetString("7ca947676876381329cefdf6bb58b409d56438a4be0786b4c899ea43b1c99e4d", 16)
	exp.y, _ = new(big.Int).SetString("eb75fbe5b68e0ed0e36e959099dc9b992123cb7c58f3ee22b6894b35966bd1ad", 16)
	exp.z, _ = new(big.Int).SetString("4372423267b6452929646ab1307cb1756412d5dd2c962a2e12c51f658e4ed0b8", 16)

	a.twice(&r)
	if !r.equal(&exp) {
		t.Error("Twice() fail")
	}
}

func TestStealthAdd(t *testing.T) {
	var a, ad, exp, r xyz_t

	a.x, _ = new(big.Int).SetString("1064e1c5c1f77227c497fb8b45710321642de0d725b4683e0ab6c8dbb79fa474", 16)
	a.y, _ = new(big.Int).SetString("66dead94e701f140abf1fa87781fe26c5c1c5b30b8ec25e25a246d113c40ee45", 16)
	a.z, _ = new(big.Int).SetString("a48793ae3fd901a773b22ca8a642763d28717369df6ec97cfc0ea94d6dc75375", 16)

	ad.x, _ = new(big.Int).SetString("f46a67e20804f956a1ce64566d96a42658a9a7a4c9a0be924615bef881a4a3f2", 16)
	ad.y, _ = new(big.Int).SetString("3a8218cdf4156c60585f5721189289cc89500eab79480a109eb1d0684e560996", 16)
	ad.z, _ = new(big.Int).SetString("01", 16)

	exp.x, _ = new(big.Int).SetString("adc99888418c4ddc0aacc18650c98407b0fa02fe726fd0e07a81049a73a8cc7a", 16)
	exp.y, _ = new(big.Int).SetString("978815885cd7382b06345dd9c3fefeaa2fa24b2e78b72ad43633a513dec6b5eb", 16)
	exp.z, _ = new(big.Int).SetString("624637786832d2583e1e27ab53d06fdc749293db0097438ff7ed3c46f19f9ac6", 16)

	a.add(&r, &ad)
	if !r.equal(&exp) {
		t.Error("FpAdd() fail 1")
	}

	a.x, _ = new(big.Int).SetString("344325caaa8fcd06081c8b539b0daaf795a2e1de09f4ac915b55f6dfcc0f67f4", 16)
	a.y, _ = new(big.Int).SetString("1dcf49d655fd194150b1d5c3b606e04091a7b483acee8f696c8c5ac86af70c24", 16)
	a.z, _ = new(big.Int).SetString("32576efb35992c0794ab96913a4c0c7970e806087f9b2bb49d8ddcfa0cc61bfa", 16)

	ad.x, _ = new(big.Int).SetString("f46a67e20804f956a1ce64566d96a42658a9a7a4c9a0be924615bef881a4a3f2", 16)
	ad.y, _ = new(big.Int).SetString("c57de7320bea939fa7a0a8dee76d763376aff15486b7f5ef614e2f96b1a9f299", 16)
	ad.z, _ = new(big.Int).SetString("01", 16)

	exp.x, _ = new(big.Int).SetString("c3ef390a6079d8ab2ce3a44f0eb3ad7412271af3ae892725a58ba6ac76b3655e", 16)
	exp.y, _ = new(big.Int).SetString("23b3f3c6a210dcf5c92340787c0ce16b9ec4893ed3be075f3e7e1f63e85d93e2", 16)
	exp.z, _ = new(big.Int).SetString("366a5c7efd615197b8508d520d3f859d340e782c01ec917f675dda38cb8093c1", 16)

	a.add(&r, &ad)
	if !r.equal(&exp) {
		t.Error("FpAdd() fail 2")
	}
}


func TestStealthMult(t *testing.T) {
	var a, exp, r xyz_t
	var k *big.Int

	a.x = new(big.Int)
	a.y = new(big.Int)
	a.z = new(big.Int)
	k = new(big.Int)
	exp.x = new(big.Int)
	exp.y = new(big.Int)
	exp.z = new(big.Int)

	for i := range _stealth_vecs {
		a.x.SetString(_stealth_vecs[i][0], 16)
		a.y.SetString(_stealth_vecs[i][1], 16)
		a.z.SetString("01", 16)
		k.SetString(_stealth_vecs[i][2], 16)
		exp.x.SetString(_stealth_vecs[i][3], 16)
		exp.y.SetString(_stealth_vecs[i][4], 16)
		exp.z.SetString(_stealth_vecs[i][5], 16)
		a.mul(&r, k)
		if !r.equal(&exp) {
			t.Error("FpMult() fail at", i)
		}
	}
}


func TestStealthDiff(t *testing.T) {
	for i := range _stealth_vecs {
		x, _ := hex.DecodeString(_stealth_vecs[i][0])
		y, _ := hex.DecodeString(_stealth_vecs[i][1])
		k, _ := hex.DecodeString(_stealth_vecs[i][2])
		exp, _ := hex.DecodeString(_stealth_vecs[i][3])

		res := StealthDiff(append([]byte{0x04}, append(x, y...)...), k)
		if !bytes.Equal(exp, res) {
			println(hex.EncodeToString(exp))
			println(hex.EncodeToString(res))
			t.Error("StealthDiff() fail at", i)
		}
	}
}
