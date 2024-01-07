package day3

import "testing"

var (
	lightInput = []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	largeInput = []string{
		"jVTBgVbgJQVrTLRRsLvRzWcZvnDs",
		"dhtmhfdfNlNNldfqmPCflqGbNZDHsDWcRzvczWsczZNzHz",
		"tmwwwCCfbJSMbwMb",
		"hsrZZhHlhrHmPPbMbDFDQdnQgLfMFDdDQQ",
		"GpBtwtqrcCcjgnLgqfgDDgRn",
		"cJwVwpCpGJctJtBcCrSCGrVJhlsbvSvTvbmHmmsWmHslmsHm",
		"gCtWJvmfmGGwVVMhJw",
		"nzRSpZbSVFFRDFSDzcplddqplqMhQMclMp",
		"zFLszzRTDnZnbTZTRZsVNgCjrvfvgtvNmtfvLW",
		"glRQRpQQtQtGtQws",
		"TnmbLqvBFRFFLPBFnPbvRBhshTtHWhwzdwtHdsdzWhws",
		"qmCLPNmCFnLBnmPPqVbFLRrJjVggDgJjlZVVDjDlDD",
		"vRRgpWvPQFdTFDDNQs",
		"bqtCmltmlbwqLVLZqwtmLBBTMcGBddTTBgFNGcZGMD",
		"bbtmJmjlVlwblwwbwzbbvrrznvzShgRhRvhfWrWn",
		"ZMhThfNcpbbMNNjsHpmpsRqsPmRs",
		"wQjDgggQDPqqDlsD",
		"SCwSzvLVCSVtQVgLnrccfdGdTdZfcZMtjJhG",
		"wNnNmNHnNPPwwPGCrLSZZvdVVZvBtMMvdm",
		"WQzlhzjzbBtMMlBrMl",
		"szbgWhJjTTcsWTqgzsqcsGHfwNcwfwnHHrCGCPPGwr",
		"CNsbpFCMSrmDhQHNNGmH",
		"fQPPPcqvljQzjVDDgRBhGGqDgqqD",
		"ZctlcVzcfltQtnrndbQMCM",
		"NQjQjQvZvZjcvrrrNjgTQgBQwTJsJswJlbGstqqtmGhmwhqw",
		"PWpHRzRnPHHSCnPFwlqhbtqGZClJqGqG",
		"ZzVpMpWPHnVzzpWzDRzSZrcdDQdrQNrcQgQfjcNfjf",
		"BSZMtdtZBzMFvhCbBJDbhDDC",
		"qcqVVmccrmVcjrlHqTrjDJRPQhWvPWWfPvblffDf",
		"cHTrbGwpmGwVjdFMnzpLznMztd",
		"DGDGGbNgTgJQQLMRMMTNVzvPsPbdsfPVsdVVZfPf",
		"lcCmmlpwwnwSjCHtFpCpHzCrZsVrBZZzdvDfVrsfZd",
		"lSHwnjpFmppHqppttttcFmnhMLGNRLhTMDqJLQNLLQNLhJ",
		"WqWfDWBjBjLbfcbbqGbWqQsrFFztsppMFCzgCJzJCrFpCM",
		"VRlhdHZTZVRRmZwlmFrJwFpMMDDrrrtJwg",
		"PPVddvvDZmRmHvndTHmLcbcjSjLQBGWGGQWSqP",
		"HHvgvwHMPMwHwmcRfJNFchFGNNCm",
		"BsTDsjzBCCBJGffN",
		"jznSnSdqzqnQTbdDljQjQSMHWWvvgWvtZlrWpPfWwtPr",
		"DJCJJCNjCDRfDfDRhDnNhfjFPgbGbddBTjFdTFTPbgGSdF",
		"cmMcMVqBZVcwsGFgGqgWdFqb",
		"vcZVzwmHVcrtrwMMvvmBHwNNCDpQRfhLhrNJNfJDDDRC",
		"VWSSScsncpDRdnsWsVncVzTwMMMHtMTrLTLMMVHM",
		"fjJvQqBCQNhfQlZZnqmFHLFMTFTzwzLzrHMB",
		"CqfjjQZJPjjvpRppDPgbnSnP",
		"tLnjNwjRWttdCwRLsfGzfzPbzbWbQQPT",
		"cvvBvlBrFrlTffsbfTqZ",
		"vmrFpsrvFMMMNwhtDwhN",
		"wzgRNqwtgzMWtqGwCssBBSBZnSRsrQQS",
		"bbjLTmpjpHcpVncVdmffPCDrjZDSDsZffDQC",
		"pbVFncvvbpbbHJJHVdvqMGMJNwWlNlJMwqzlNM",
		"TdszlzFsRQqFdRzqwwQGlFsGmmSBBdmdVVmgSZdSPZBBBmSD",
		"HMCCMbJHJJLHSmSvZLBcSDBP",
		"NWbhPjfrbbNfWCPjhNnGsNzsqpQszTQGRwGQzR",
		"ppQpTNCPBTlNBVmNQTNrrrrtqsrWbGrVFWqhZb",
		"DMvDHnjRvMDLghhhgbqZWbqFSS",
		"vWvMRRJDLMnHLJjDwWnnndcllNTzBmClQBpBBzQQzCBzCJ",
		"VsNZfPMnrCnlCnWtbvBbvwtTbZwT",
		"JhJQdQhmRHDRdLmHJhljhphgGwtttmcvbwwTWtvtbGBttc",
		"jqLpjSLDHlhnffCMzsqCPP",
		"JzJdLRmmzJwrLLwLJwLTWwBrMrpHlFnScVVqccrBHSlc",
		"tvQDQhjhvbqFqFSpjHSV",
		"QbgZtsgfbNNQbgbsbQRfmWRJwJGWWTPRwqzR",
		"MwvDgpwszSpSJPsssMccTQfTDfTQRTljfRfT",
		"PmGbWhbVtZWPTqRrcTrrjTrt",
		"NbZbWHBmVGmLbZVBVNLPsMBzwSSzMJMCsCzvnvSp",
		"cTpgTnpzbZlJHTZm",
		"jRrLVtqtvtrjqsHZhvpmBfJHmhhp",
		"FqDsFLCwCVRDqwgNQcQMwnncnMpg",
		"sCCtssdZdZJmMbNJDmtmJzLSrcGfHGLTHDHnGLrHTSSn",
		"gQwRRWqwqgpggFWPjPpFnSTSnHNfLLQrccrLGHLL",
		"qqlpqhjFwPgPwvvRgFlPdslZCzMZtVdtVdVMtJlN",
		"BWVmPtRVRRWDNtZBVQzCfdscmfjcfdfzzSfz",
		"MhgJLbGggHbqpGgpgJbrFJdjwsCChCzllzCdjDSszCCD",
		"grpDLFrbgpJnqJrHGFHGqMHBBPRWQTPWZBQVTQWnPNVvBt",
		"gJfggdmHDgfJJWzCcbqvcqcmTG",
		"hLZlRBZNlBlrpprzrWqvtHvvGtbvHT",
		"sNLRjHZLNVnZZVLppnhNsMFJSjFPPQMjwgMPPfDMfS",
		"mDDzHfrPBvJRJhpBRg",
		"SsTSTwJwcbCtwssGFVWppgZRbLpLRNhpLhbLvg",
		"ttVCGjTGrfJJMfjD",
		"ffhcsTjnfqHLqvZSHvHB",
		"CsgmPVmstsQVpRbHBvFbHBDMvZLZGB",
		"ppgpmmRVpVPwPrrrTNlcrwrlWs",
		"NDtssPNBjQtCtCcT",
		"ZZqncfqhZqhJZFTJTllCGVljSl",
		"ndrWbfZwWhngwbqbZnMZcwWhDmmLDmzvHPvzNLdmzsvBBHHP",
		"FQtptwMppSFQRRQfSfvTrTJJJTrvLjMMJbgJ",
		"CcDqWBWzbldcchDGCWBCBhdGrvhnrsrjjsJvLvghrjjnnjsJ",
		"GDBldPDGqPGWVBBcwHFmSmbpSpNRwHFV",
		"CwHwCFwtCmdCDflHDpwFnnvzhhNJhJNzmhzhhNMM",
		"sTbZssPqScTrqSWSShdMvgMRRzdRRQ",
		"brBrsWPTrdrWsPcBcTcGqDLDDHjFfDClwLLBfFlFlw",
		"lNptNFWpbVMjlBgQgvdNBRQLrd",
		"TsDCDfSCQqqQQqDq",
		"TwTmwPPPzZSCccScJwpWZHljtbWMFFQVHpnj",
		"fhFmwbrgnCcSnwtS",
		"vZVVZvQZVPZsMnNSccMHPN",
		"ddzvQJvQzBzWRTJzdRVSWGqbRbgmfbFFbmbgmpmlff",
		"wdslVdQtdlBVHDrHBcBc",
		"TJWvpncCcJzCWcRfWvJRRpfZHBDZZHmDZPBjZHjZrSfb",
		"TRJFvvWpTRRWCpCgGtgtGQdlcNtsGlwg",
		"rCHvRLJtCjpbRCLpptHMVCQgGlMndVlQGNcCMc",
		"zZfzSmsfSsMzMccQMVcN",
		"ssSmPmSmhwVhZBZsTBRJrtRjtbDvjvrTHLtJ",
		"wswRCNHHhsrWFsGfGWFBGb",
		"lLtngDPLgLJPttgWzQWlbCfBlmBFCW",
		"DcDnVpcnnJPngjjcdpRhwNZSCZNdNdvr",
		"NWNzWpfMRHfwsRNznPdjtdjJtPVPHVdJ",
		"CSLTZBrGbSmClvBVdGzFVVcFdjjjnV",
		"SrZhZbTvmmbbLvwQzMhsfhqqWwQR",
		"vvZqwFBZvzZzrqltPsQstrGGpMcbbR",
		"bhJdjJJmTRQMTMMPPM",
		"mCmgLbdVVVLhVnJmLgJhBWZFDlqFvwZDlZBnBvWB",
		"sMrcmQcHHsLLrSHZhvdCddvtJJFl",
		"jzjjWplWpPRPDzPzfRGjqvdJqqqCtqdqDvdDqtwC",
		"pPNWVpGGVVffPlnSMnMsTmgsQVLr",
		"zGMMRbpGmqqqNRmmzbNfbzPRPlvThCTrHPnrHSVPlTHR",
		"LSjjwgsjtTjhPhHhvC",
		"dJZtJwFgtBLDZZbmNMGzMMqS",
		"ZrnstppPWccnsFWpnZnRdjRtjbCtSSRjjjLLbG",
		"JgBQPJvPgHHJvmmzwGGLdjHGSjShGHqdHL",
		"wTJvPzmTJvzJQBfwNmTmlPsrVfnVFpcMVZprDpFrsrVc",
		"ssGCtltsSdJJtQjPdvHvfbfvqLHqZtBfVb",
		"pzRwwDwDTgzbqVTVvHqVWB",
		"grzHnRpmFpDMnmzFhplJCjlsPdGJSjsFCdFs",
		"CTGBBGCBlSTTSsnTMrQbNrBMtpVzNddWHWzVpHVtdHmwhphm",
		"JqPZgMRFFvFZvDZMDFcFFfDchdtdPWpzpWhtHWwwhVzWpddW",
		"ZfLZMcgqRDjgjcLccDRDLsbbrSBBlQGQBCTrCBnT",
		"VjVGVqSqFLFVSqCjFJSsbfPprHbCCRRPccRWPW",
		"wwnQmtwlvNmpZRsPsWWNbZ",
		"DBhhhznhddldnvjMqJMTqDMGPSjF",
		"TTJbsJPPBDsBVbJJGJBGWLfmWzgmDmzmLqmmLddQ",
		"jHwVZZjwFZlzLzWZLMdLqM",
		"HHHhjHplrHSpcCSvjlsNPbNRbRnVTrRsNNNJ",
		"NzMMLZtwRmbwFnVDhnqD",
		"SlsJsSdSJdNJnFphVHFjph",
		"vlrsTlGPTgvvSBScGcBvfzmmLLCtMWNZRBQMzftf",
		"tRFLmZZRLrtRvtvvrvvGgvtLNfwzMzNwgdznMpwwpnMpqHdz",
		"hcsBsWTcQJhjbHMncNwzMqnzwl",
		"VWSWWTJhWBVDTJsTVWQWTVVbZtCPSvrrGSLCRFRGtRZLPmmq",
		"hbdFhdShGsFSGBlQhNhQMLLLlLJCvLLDtVJDLlwtwD",
		"WcqWsmcWrmqcmtDHJjVHrJCjjt",
		"WPWcnggssmzqzzzgfzZnWRqqdGpSNFdMSFdBMSFZhSFBZZhh",
		"GNFNtRQMGjDjwfgDZjmz",
		"bqDsPWWbsqVsdvvBJvBfmgfLLzSwzLcmzmfB",
		"WPWrrVrrVJhVWJDVsqnPdRQllRQRNhQGMlFQFMplGl",
		"mChChmSQGSGJrjPHCpPFtwgsFZjtFVZfgVwtdV",
		"MBMqvDWMlTbzlRWzllDzblZfsrgdZffgrtwrswZfdZTd",
		"DlcMvBbBMnqMDqcRqDBMWvLmCGrGLSJhnLHCCHHnPmSQ",
		"tscfGqftGfDnnppJGDGCZLbzMVMwPPhsblzbjzzMMz",
		"TWPTWWmHTmFRSBSvgBPwwwlVzhMdwblhbVjRLL",
		"TTQNNWWgHWQHPBWTNPWNcqJtGCnfqfpQZJCZpQZn",
		"CqzCGDQJCzzftfRqRzzMdvFpFpccdZFvFMtMbd",
		"rNHwmHVmswsHVsPTLnbFSTbZZpFcpvSZZgpg",
		"vPLhLLwHLhVVNVvQhBqDCBfhDzCffG",
		"WsBbBbsWNhsPsCNssRWLPLpmLDDQHlJlnlzFnDDnzF",
		"gcGfqggMqfjjGwrdDZzQFmpJQzZlFDMF",
		"qdwqvjdTrwfvvBCPtpVvtR",
		"RQdbbRHtHRvqZtwVcmwVwV",
		"WnLNnqFDlDWzzNLLrjClwCZCGZcglGCccZ",
		"LnTFfWBpffrfrhBqdsSBqhBJHb",
		"sqsgJpDMrNQgGsFMsPCfjCPChPWjqSWSjh",
		"RBnRnZVcwZllLwbwwLbZVLclhpjRttSdtjPhjWShWdphjCPd",
		"HwHwVVwVnBVBmmGppNFzrrsgJJ",
		"WpmDFlQlzmmgMMLMLQVTvTTSwNbmTVwTtHHw",
		"jrhPDnDnnZfjPTNtHSVTcjccwt",
		"CJqPfGnhGZCBfnPJCrBCqdLqMlQqDLlLRgLQFFRW",
		"vZVvDZsvhDZhZvzgVcgVqPqmwWMqcw",
		"bdTbdBBFQcdCdcGmcP",
		"bmjQpHfbpzDJRjNhJZ",
		"twRrwjFptprQjjjtQRdWCmNJTlNSCmZQcNlmlSST",
		"VVDzMWDnDHMzLZDNNJSqqCJZZD",
		"HHhhsfVHGbnwgjfrgdpgWj",
		"GmszZGMrLmnmsfGVRcVlwtwccc",
		"SSCgbNqSTgCCJMHCJtlVcwVbVljlclVfwf",
		"HTCgHSgHQThMqWQQSgDnvBdsFDvrzdZsLdmLZZ",
		"PRlMlBPPctVBlstzVLhsdwCqCdCNDjSDdWmMqdNw",
		"fZrQQHFffgGFprSJSgvZrfnqjmWdnndCNGWmWDwNNDCG",
		"rTvvpJZZrfpFSbQQvrrsssLRVPPtlRtRPThzTR",
		"FqgHFFMRTRjRFRpRRjFtNdCtJCMnNNdrdMLdrQ",
		"VhWSmwGwWVbGbvlwwlLZJLZSdJZtNCtnTSCJ",
		"BmWwWWzhvVfwWhmhwlmvwlRqHpscHRTpTDRFfTsjDPjD",
		"MJMgGqMFLPGgWVjpcmjZTpmZjZpJ",
		"hdSzzlCtzNdtWSdndttflBmjpBRmvvpjnjcvBjmvHj",
		"zdhCrfztrDSzfWzdChrhlhgPPPGDgqFFLGGFVVqqPVQw",
		"ZQZNQRZzFdCCgfLcCGDfScjDcG",
		"vsmwVHTmTfGcSHjcDS",
		"tMsmMVlSWWzdWNnQNJ",
		"GPRcQnwwQWwFFnrnnldSqzMfSCdfdldrJf",
		"LpTsjmZTsBZjpmzhLLMdSJJSMhCC",
		"ZsZBssBsDpDmmVBjCmDZHgnVNWvPQPNPGvNQncQPcRWn",
		"pznzpzlGFrvGHGrnnMvDmBMfgfTmsBsTTghDsg",
		"LJtWVCWLCNPcSbdcShWBgBThgTfjzwfhhz",
		"VZZCdLCVNNCVbCLzFnRqGqlHprQZHRqr",
		"dFTsQPdMFsMnWFPdSnwBltftttvlflNN",
		"VLZhZLqghCgzqgrLrcgVgbCvtDDtwpqNpDtlBRflDwNqDpNw",
		"cLmVhVcgbZrVhrhCLhczhQdJHTmPJJTjvTJsPFWFTW",
		"SSwNPNHldNJSngHqBssQvBfccB",
		"mMppmDprWpFGRGWmWrDrtGzQfcvvQZBBzzczqRQgRqRT",
		"WtFtFvhvpLphNJJVSCbSNP",
		"chpGMMzcwSSGnQFRQQFWcFWn",
		"sgddTfjLqsWrRtLvQnJr",
		"CsmlZgssbRdMhzCHDGpGGG",
		"vHBrTzpMPTHMtbBRRJGtDsNB",
		"QJWWVwnCZmQlWQWlLWCCmmLwRgtDdDbgngqtRdRGbDNDGtGq",
		"LwmWQLJcWmwPrpvpjjrcjT",
		"fcsWnWzhWcWgcbfbvtbHTRvpvHttmLtR",
		"lNSjwjrDFjlFhlZlQTpLHvSHptvTSRtLmJ",
		"NwjwrQDwFCZBCfWzqhqqzc",
		"fgNJNRcvvWRfWRrZFldlwlFwfFllflDH",
		"spQshQhpqhJsLpnQVLqBqlFwddHSdBFFjSFmwlFmwl",
		"hppLsqPVLnpnzJPtqtPPJTCgNcrcNbrrGNcgRNrzcZbG",
		"PWFdgDGCFPGhMtQqHBrpJqqW",
		"nlllLNmnVNNLllVbVRLRsQMqqpccQQJcJtqcJcnBrc",
		"NmvZLsNrbZNjNZVNGvdFFfwdDhPFdDzC",
		"LpZpwgLsLSzDdjVGpS",
		"bWBlHqqBhNJWNbJQFzGtCtDtGGjNGDgtGC",
		"RBRbJggbWRRmhWqTcnnfnmZMTTsTcP",
		"JJgzvfzpdzzJjJhgdfhvqfdScNsLwwGsrRbwRLbcbVrVRp",
		"WDFBTTDtHTntltnCnnntCDwGlGGbSwNVSssrbwGVsVLs",
		"nCHMDtCWWTCCHmmPDnZQgvfQgZNJJdvdMZQg",
		"lFDgvlsGvvZGDsFZWgGvWrPqnmwwtqmMVSWrSqMM",
		"hRpJhLQHhdTTVPmVSrqwtHmV",
		"ctJJJfjLpjglZDGCGljF",
		"CnnVMbhVRbQQZjBP",
		"rlfsLFLtmLSJscttFfsdjZwZNNwBPWRjRNZBZBfQ",
		"tFrmDFlDtmcltFvqVzDqQzGvCVCG",
		"JzzJzVrmzJpCCmTFmjZS",
		"HtDDtggWssqWfDgwDWvsfDBBchZjSnGGpCFjSFZjpGjFShZg",
		"sbvbvfbbDblWtrNVNRzRlJPCzM",
		"nlFnFWsWhrctWVdJPDPTnTNJPJQJ",
		"fHqvHSqRqSHjBmqvmqqHCtJSTZGdMQJDPQDPPMdNTMZN",
		"jzqmbtRRztmbLHcFpgWVsFphpcLV",
		"PHZFZFVZZfHgpwjFtmgjtq",
		"rpTrTNzzNdrTJwgMwqCBJzJz",
		"vsTWbvccRcdRbrRnRRbRrcvVpGlGHZPspSVhSPPQGGZVHV",
		"HWnDDjfPFccDPhfchnMMVWGzzpvGszCCGWWV",
		"JNBtBTQJNwJQjTpRVRMvpLsQzsRR",
		"JltNSrBjmrHfdPHnDlHg",
		"ldCJHlZFspjzHMnp",
		"zvcLQBQcQvhBwmcDppqbNpjMLnLDDn",
		"BvcmQhWWRzPJJzWWWg",
		"ggSTPZBwTPTPSTRwZPBnwPMLdVvBqzsqLzqqtVzqtzBszN",
		"QRmhQhffCQhJcDfmpChQWJmJNNtzvqtLdvNsGtLqNGzvWqvG",
		"HQFCDhDFCCQJQmZTZwTSSwrZnRFS",
		"QbFlsMbgPWPlJWzsJsJZntvnvZtctHBfZvBZlD",
		"VTqpTqmSrhVLqrpjNppgntfBgfjddHffZBdtcB",
		"VCwpqqqNVgNVgMJJwbsWWGMGFR",
		"GCwRjQlsCQrPrGMQPsRvpdvgnjgmVVmSStptgJ",
		"DzNcZNHZZhzzHhDzDTLWhDzSdJSSpnnVSTSvJJmgvSdmmn",
		"pfLNDppNWHqDWfWbzcHPRslCGbwRCPwQrGQPCF",
		"hVLnDgCmbhltrmDlhbhVmcgFBWdSBSZZBFPwBLBPfWdPWZ",
		"TNTjJNpjqGMTRRsTTCZddfWwFHFWHSSJSFPS",
		"CGQqvQRTNzTpQsDtmcrgDQllhnQl",
		"SWrtcHWjcWrPcwWrBwSPffnJNsqfMNCNqNfJFfSq",
		"QLQvhBpbbvdvTdvpTdGDbDQqqqslsNMJlMMCqMfQnlfC",
		"gThhpmDLbzBmGLptrmRZcwHWHZjVrW",
		"gQvzQRvSSbvvJvQgfRrfbSpGqBPGwqwVjPBBwwjpRmjB",
		"TcNHHVVtNsDHcMcdMBjqpBnGnwlGTwlPBl",
		"MHsMDHFMdDtZddZdFdFrhhgbFFVvffrzzffrrS",
		"QSFmrDSSQrqlfmDDHPRTdrrTPRbsTPTsNN",
		"wBcclhhgwgMhWLLtVMgVvzRTNNjvbszzTdjNdsRbjp",
		"nJgVMtBhwLBctmQHQlnCGqmQZZ",
		"sggtjzzggfGmPbCMvRCMvTmT",
		"RDqqhdQdDlcDpqVlLbSbZFSTTPllbCMv",
		"hqpQWRhhdncDQcBsrwzjnfgtgGfgHs",
		"MTrzlgMNQNggrrrPlzQDPCsFRfscTfFVhVftRsFFsScS",
		"nWWHZHhZWJjjwjLjwbLbwHGSFppVfFVcGcFLfFSVttpf",
		"jWhZZBhqBbwvZvBjZNNQrQQQzPMlzzglPB",
		"jLVhJZQjwFCLRjQhPRZFLDzrGDHpDGsGqztGHststC",
		"SBNlNmWnfvdLmlnvfNSbzrDTTzprdqrDpGqqpHrG",
		"bNmSBnBmcfMmcmnlfcnNSnLhZPVVJhwjZRJPwjZMPjMZ",
		"VpzBDgGTGVNNpSGzppMdTQwcvFdFMQdcdFwc",
		"ZDsqfRftWtllmlWbLLtjFFMwMrjCcwfvMFvwcC",
		"PHbDLHZtZJSJPpVgnV",
		"bRvTdswLLSTvwswSbDhsDTvFmmGRVmJGZJnRcGGfFVJcqn",
		"MWllQMllWHrjWPNplrllQMPZJmnnVmqJcNmVFFnJNFqVGC",
		"PgpQrQjjzQWHzpBdvtwhwdSShBZTsT",
		"MhTwjMTsTRFStjmSMqqppBrHpDrzHtPqbD",
		"dllNcZWgldLvcsvvvvgvWddlHffqBHBBfPzbbqBpPHrpHNbq",
		"VddvgWWdCZhnhsCSSCGT",
		"LPjqHnDNqqHNllqLpqPCZCGRCssGdGrGFrPrgr",
		"WVBztWTQSQMBQrGgRwwwCGgtwg",
		"QJMTgvbTTWSzWWvSbVJTzJllJlHHhLpNqpHqjNjNHpjq",
		"PCHCbPPPHPlTThBhjGTTNhMNTh",
		"FrmfLqdqgfmfttqtWqfrqdhshchDBshDllDBcGDhGDWs",
		"mqgdpvFmmdLdqqQCPHZZblvwZQZl",
		"bQGqmngwwgSNrBWJWdHZmjfZWB",
		"FlpRLCptFLMlLPRLlCCcCCMpjJZJHShWdWvBHHdcZdBWZvSv",
		"FRPCDTTtFptVTnQnrGbwbS",
		"LsdmnDMTLbzsbNtqcb",
		"lJjCnHSvQRRwQQjRRHQbgWbqctNJPbcWrcPPgc",
		"RhGSQGwBvvGShnGlHClwjmfpmfdmVmfFDBLDDZpmMf",
		"ppDnPmwvNDjTjjcssT",
		"qqfRHzdCPHWfhHHtTjjbbLLGZr",
		"MhzqWdJCzqJWSJnpnpvvPSPP",
		"NGWdQgDDHGJgQLznzzsJFFzvzB",
		"twRCpZVjVWqVSqVwwjtZfrrfntfvznBssBncfLrc",
		"jRRwCqwCZhlhZRpSZpjSqWwqmDMQdMmHPQQMHGdlHdTldNGd",
	}
)

func priority(c byte) int {
	if c >= 'a' && c <= 'z' {
		return int(c-'a') + 1
	}
	if c >= 'A' && c <= 'Z' {
		return int(c-'A') + 1 + 26
	}
	panic("invalud input")

}

func linePriority(l string) int {
	compartment := make(map[byte]bool)
	prio := 0
	for i := 0; i < len(l); i++ {
		c := l[i]
		if i < len(l)/2 {
			compartment[c] = true
		} else if compartment[c] {
			prio += priority(c)
			delete(compartment, c)
		}
	}
	return prio
}

func TestDay3Phase1(t *testing.T) {
	r := 0
	for _, l := range lightInput {
		r += linePriority(l)
	}
	if r != 157 {
		t.Fail()
	}

	r = 0
	for _, l := range largeInput {
		r += linePriority(l)
	}
	if r != 8109 {
		t.Fail()
	}
}

func badgePriority(lines []string) int {
	content := make(map[byte]byte)
	for i := 0; i < 3; i++ {
		l := lines[i]
		for j := 0; j < len(l); j++ {
			c := l[j]
			content[c] = content[c] | (1 << i)
		}
	}
	for k, v := range content {
		if v == 7 { // 4 + 2 + 1
			return priority(k)
		}
	}
	panic("inalid ruck sacks")
}

func TestDay3Phase2(t *testing.T) {
	res := 0
	for i := 0; i < len(lightInput); i += 3 {
		res += badgePriority(lightInput[i : i+3])
	}
	if res != 70 {
		t.Fail()
	}

	res = 0
	for i := 0; i < len(largeInput); i += 3 {
		res += badgePriority(largeInput[i : i+3])
	}
	if res != 2738 {
		t.Fail()
	}
}
