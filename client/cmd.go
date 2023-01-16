package client

import (
	"bufio"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func DoCmd(ctx context.Context, env []string, args []string) ([]byte, error) {
	ctx, cancelFunc := context.WithTimeout(ctx, 30*time.Second)
	defer cancelFunc()
	cmd := exec.CommandContext(ctx, "swan-client", args...)
	if env != nil {
		cmd.Env = append(os.Environ(), env...)
	}
	return cmd.CombinedOutput()
}

func ExecuteCmdAndLog(ctx context.Context, env []string, args []string) ([]MinerIdAndDealCid, error) {
	cmd := exec.CommandContext(ctx, "swan-client", args...)
	if env != nil {
		cmd.Env = append(os.Environ(), env...)
	}

	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if err != nil {
		return nil, err
	}

	if err = cmd.Start(); err != nil {
		return nil, errors.Wrap(err, "execute cmd failed")
	}

	minerIdAndDealCids := make([]MinerIdAndDealCid, 0)
	r := bufio.NewReader(stdout)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, errors.Wrap(err, "read stdout failed")
			}
		}
		fmt.Print(line)
		md, err := handleLog(line)
		if err == nil {
			minerIdAndDealCids = append(minerIdAndDealCids, md)
		}
	}
	return minerIdAndDealCids, nil
}

func handleLog(data string) (MinerIdAndDealCid, error) {
	var minerId, dealCid string
	if strings.Contains(data, "deal sent successfully") {
		minerReg := regexp.MustCompile(`miner:[^\s]+`)
		miner := minerReg.FindAllString(data, 1)
		if len(miner) > 0 {
			minerId = miner[0][6 : len(miner[0])-1]
		}

		dealReg := regexp.MustCompile(`dealCID\|dealUuid:[^\s]+`)
		deal := dealReg.FindAllString(data, 1)
		if len(deal) > 0 {
			dealCid = deal[0][17 : len(deal[0])-1]
		}
		return MinerIdAndDealCid{
			MinerId: minerId,
			DealCid: dealCid,
		}, nil
	}
	return MinerIdAndDealCid{}, errors.New("not found data")
}
