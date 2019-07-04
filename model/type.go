package model

type Case struct {
	Folder string
	Pf     string
	Belong string
}

const (
	USA   string = "USA"
	CAN   string = "CAN"
	CAB   string = "CAB"
	LNC   string = "LNC"
	PRINT string = "PRINT"
	OTHER string = "OTHER"
)

func GetCase(option string) Case {
	var targetCase Case
	switch option {
	case USA:
		UsaCase := Case{
			Folder: "WMT_xxxxxx_xxx 做稿",
			Pf:     "U180XXX_XXX_DetailList_W.xlsx",
			Belong: USA,
		}
		targetCase = UsaCase
	case CAN:
		CanCase := Case{
			Folder: "CWMT_xxxxxx_xxx 做稿",
			Pf:     "CXXXXXX_XXX_DetailList_W.xlsx",
			Belong: CAN,
		}
		targetCase = CanCase
	case CAB:
		CabCase := Case{
			Folder: "CAB_18xxxxx_xxx 做稿",
			Pf:     "B180xxx_xxx_DetailList_W.xlsx",
			Belong: CAB,
		}
		targetCase = CabCase
	case LNC:
		LncCase := Case{
			Folder: "CAB_18xxxxx_xxx 做稿",
			Pf:     "B180xxx_xxx_DetailList_W.xlsx",
			Belong: LNC,
		}
		targetCase = LncCase
	case OTHER:
		OtherCase := Case{
			Folder: "Non WMT_18xxxx_xxx 做稿",
			Pf:     "BXXXXXX_XXX_DetailList_W.xlsx",
			Belong: OTHER,
		}
		targetCase = OtherCase
	case PRINT:
		PrintCase := Case{
			Folder: "印刷/P15xxxx_xxx",
			Pf:     "BXXXXXX_XXX_DetailList_W.xlsx",
			Belong: PRINT,
		}
		targetCase = PrintCase
	}
	return targetCase
}
