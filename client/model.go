package client

type BidMode int

const (
	BidAuto = iota
	BidPrivate
	JsonFileName = "car.json"
)

type CreateCarModel struct {
	SourceFilePath string
	OutPutPath     string
	//ParentPath     bool
	parallel int
}

type SendDealModel struct {
	MetaJsonPath string
	OutPutPath   string
	BidMode      BidMode
	MaxCopy      int
	MinerIds     string
}

type MinerIdAndDealCid struct {
	MinerId string
	DealCid string
}

type FileDesc struct {
	Uuid           string
	SourceFileName string
	SourceFilePath string
	SourceFileMd5  string
	SourceFileSize int64
	CarFileName    string
	CarFilePath    string
	CarFileMd5     string
	CarFileUrl     string
	CarFileSize    int64
	PayloadCid     string
	PieceCid       string
	StartEpoch     *int64
	SourceId       *int
	Deals          []*DealInfo
}

type DealInfo struct {
	DealCid    string
	MinerFid   string
	StartEpoch int
	Cost       string
}
