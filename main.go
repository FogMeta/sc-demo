package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"sc-demo/client"
)

func main() {
	ctx := context.TODO()

	currentPwd, _ := os.Getwd()
	sourceFilePath := path.Join(currentPwd, "test/sources")
	outputPath := path.Join(currentPwd, "test/car")

	metaDataPath, err := client.CreateCar(ctx, client.CreateCarModel{
		SourceFilePath: sourceFilePath,
		OutPutPath:     outputPath,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("CreateCar-->  metaDataPath: %s \n", metaDataPath)

	downloadUrl, err := client.UploadCar(ctx, outputPath)
	if err != nil {
		log.Fatalln(err)
	}
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("UploadCar-->  downloadUrl: %s \n", downloadUrl)

	println("=======private task =======")
	minerIdAndDealCids, err := client.SendDeal(ctx, client.SendDealModel{
		MetaJsonPath: metaDataPath,
		OutPutPath:   path.Join(outputPath, "deal"),
		BidMode:      client.BidPrivate,
		MinerIds:     "t2430",
		MaxCopy:      1,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("SendDeal-->  minerIdAndDealCids: %+v", minerIdAndDealCids)

	println("=======auto deal =======")
	minerIdAndDealCids2, err := client.SendDeal(ctx, client.SendDealModel{
		MetaJsonPath: metaDataPath,
		OutPutPath:   path.Join(outputPath, "deal"),
		BidMode:      client.BidAuto,
		MaxCopy:      1,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("SendDeal-->  minerIdAndDealCids: %+v", minerIdAndDealCids2)

}
