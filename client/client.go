package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"path"
)

// CreateCar generate CAR file
func CreateCar(ctx context.Context, carParam CreateCarModel) (metaDataPath string, err error) {
	args := make([]string, 0)
	args = append(args, "generate-car")
	args = append(args, "graphsplit")
	args = append(args, "car")
	args = append(args, "--import=false")
	args = append(args, fmt.Sprintf("--input-dir=%s", carParam.SourceFilePath))
	args = append(args, fmt.Sprintf("--out-dir=%s", carParam.OutPutPath))
	args = append(args, "--parent-path=true")
	if carParam.parallel != 0 {
		args = append(args, fmt.Sprintf("--parallel=%d", carParam.parallel))
	}
	if out, err := DoCmd(ctx, nil, args); err != nil {
		log.Println(string(out))
		return "", err
	}
	return path.Join(carParam.OutPutPath, JsonFileName), nil
}

// UploadCar upload car file to ipfs and return download url
func UploadCar(ctx context.Context, carPath string) (downloadUrl string, err error) {
	args := make([]string, 0)
	args = append(args, "upload")
	args = append(args, fmt.Sprintf("--input-dir=%s", carPath))
	if out, err := DoCmd(ctx, nil, args); err != nil {
		log.Println(string(out))
		return "", err
	}
	fileDescList, err := ReadMetaJson(ctx, path.Join(carPath, JsonFileName))
	if err != nil {
		return "", errors.Wrap(err, "read meta json failed")
	}

	if len(fileDescList) > 0 {
		downloadUrl = fileDescList[0].CarFileUrl
		return
	}
	return "", errors.New("not found meta json data.")
}

// SendDeal  send deal to swan
func SendDeal(ctx context.Context, deal SendDealModel) ([]MinerIdAndDealCid, error) {
	args := make([]string, 0)
	args = append(args, "task")
	args = append(args, fmt.Sprintf("--input-dir=%s", deal.MetaJsonPath))
	args = append(args, fmt.Sprintf("--out-dir=%s", deal.OutPutPath))

	if deal.MaxCopy != 0 {
		args = append(args, fmt.Sprintf("--max-copy-number=%d", deal.MaxCopy))
	}

	if deal.BidMode == BidAuto {
		args = append(args, "--auto-bid")
	} else if deal.BidMode == BidPrivate {
		args = append(args, fmt.Sprintf("--miners=%s", deal.MinerIds))
	} else {
		return nil, fmt.Errorf("not support BidMode: %d", deal.BidMode)
	}

	return ExecuteCmdAndLog(ctx, nil, args)
}

func ReadMetaJson(ctx context.Context, metaJsonPath string) ([]*FileDesc, error) {
	contents, err := ioutil.ReadFile(metaJsonPath)
	if err != nil {
		return nil, err
	}
	var fileDescList []*FileDesc
	err = json.Unmarshal(contents, &fileDescList)
	if err != nil {
		return nil, err
	}
	return fileDescList, nil
}
