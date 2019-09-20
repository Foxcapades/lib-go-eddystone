package eddystone

type UrlSuffix = byteValue

var (
	UrlSuffixComSlash  = UrlSuffix{0x00, ".com/"}
	UrlSuffixOrgSlash  = UrlSuffix{0x01, ".org/"}
	UrlSuffixEduSlash  = UrlSuffix{0x02, ".edu/"}
	UrlSuffixNetSlash  = UrlSuffix{0x03, ".net/"}
	UrlSuffixInfoSlash = UrlSuffix{0x04, ".info/"}
	UrlSuffixBizSlash  = UrlSuffix{0x05, ".biz/"}
	UrlSuffixGovSlash  = UrlSuffix{0x06, ".gov/"}

	UrlSuffixCom  = UrlSuffix{0x07, ".com"}
	UrlSuffixOrg  = UrlSuffix{0x08, ".org"}
	UrlSuffixEdu  = UrlSuffix{0x09, ".edu"}
	UrlSuffixNet  = UrlSuffix{0x0A, ".net"}
	UrlSuffixInfo = UrlSuffix{0x0B, ".info"}
	UrlSuffixBiz  = UrlSuffix{0x0C, ".biz"}
	UrlSuffixGov  = UrlSuffix{0x0D, ".gov"}
)

func IsUrlSuffix(i byte) bool {
	return i <= 0x0D
}

var suffixMap = [...]*UrlSuffix {
	&UrlSuffixComSlash,
	&UrlSuffixOrgSlash,
	&UrlSuffixEduSlash,
	&UrlSuffixNetSlash,
	&UrlSuffixInfoSlash,
	&UrlSuffixBizSlash,
	&UrlSuffixGovSlash,
	&UrlSuffixCom,
	&UrlSuffixOrg,
	&UrlSuffixEdu,
	&UrlSuffixNet,
	&UrlSuffixInfo,
	&UrlSuffixBiz,
	&UrlSuffixGov,
}

func ParseUrlSuffix(i byte) *UrlSuffix {
	return suffixMap[i]
}
