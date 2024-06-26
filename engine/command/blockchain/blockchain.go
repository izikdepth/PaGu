package blockchain

import (
	"fmt"
	"strconv"

	"github.com/pactus-project/pactus/types/amount"
	"github.com/pagu-project/Pagu/client"
	"github.com/pagu-project/Pagu/engine/command"
	"github.com/pagu-project/Pagu/utils"
)

const (
	CommandName           = "blockchain"
	CalcRewardCommandName = "reward-calc"
	CalcFeeCommandName    = "fee-calc"
	HelpCommandName       = "help"
)

type Blockchain struct {
	clientMgr *client.Mgr
}

func NewBlockchain(
	clientMgr *client.Mgr,
) Blockchain {
	return Blockchain{
		clientMgr: clientMgr,
	}
}

func (bc *Blockchain) GetCommand() command.Command {
	subCmdCalcReward := command.Command{
		Name: CalcRewardCommandName,
		Desc: "Calculate how many PAC coins you will earn with your validator stake",
		Help: "Provide a stake amount between 1 to 100, please avoid using float numbers like: 1.9 or PAC suffix",
		Args: []command.Args{
			{
				Name:     "stake",
				Desc:     "Amount of stake in your validator (1-1000)",
				Optional: false,
			},
			{
				Name:     "time",
				Desc:     "After one: day/month/year",
				Optional: false,
			},
		},
		SubCommands: nil,
		AppIDs:      command.AllAppIDs(),
		Handler:     bc.calcRewardHandler,
	}

	subCmdCalcFee := command.Command{
		Name: CalcFeeCommandName,
		Desc: "Calculate fee of a transaction with providing amount",
		Help: "Provide your amount in PAC, please avoid using float numbers like: 1.9 or PAC suffix",
		Args: []command.Args{
			{
				Name:     "amount",
				Desc:     "Amount of transaction",
				Optional: false,
			},
		},
		SubCommands: nil,
		AppIDs:      command.AllAppIDs(),
		Handler:     bc.calcFeeHandler,
	}

	cmdBlockchain := command.Command{
		Name:        CommandName,
		Desc:        "Blockchain information and tools",
		Help:        "",
		Args:        nil,
		AppIDs:      command.AllAppIDs(),
		SubCommands: make([]command.Command, 0),
		Handler:     nil,
	}

	cmdBlockchain.AddSubCommand(subCmdCalcReward)
	cmdBlockchain.AddSubCommand(subCmdCalcFee)

	return cmdBlockchain
}

func (bc *Blockchain) calcRewardHandler(cmd command.Command, _ command.AppID, _ string, args ...string) command.CommandResult {
	stake, err := strconv.Atoi(args[0])
	if err != nil {
		return cmd.ErrorResult(err)
	}

	time := args[1]

	if stake < 1 || stake > 1_000 {
		return cmd.ErrorResult(fmt.Errorf("%v is invalid amount; minimum stake amount is 1 PAC and maximum is 1,000 PAC", stake))
	}

	var blocks int
	switch time {
	case "day":
		blocks = 8640
	case "month":
		blocks = 259200
	case "year":
		blocks = 3110400
	default:
		blocks = 8640
		time = "day"
	}

	bi, err := bc.clientMgr.GetBlockchainInfo()
	if err != nil {
		return cmd.ErrorResult(err)
	}

	reward := int64(stake*blocks) / int64(amount.Amount(bi.TotalPower).ToPAC())

	return cmd.SuccessfulResult("Approximately you earn %v PAC reward, with %v PAC stake 🔒 on your validator in one %s ⏰ with %s total power ⚡ of committee."+
		"\n\n> Note📝: This number is just an estimation. It will vary depending on your stake amount and total network power.",
		utils.FormatNumber(reward), utils.FormatNumber(int64(stake)), time, utils.FormatNumber(int64(amount.Amount(bi.TotalPower).ToPAC())))
}

func (bc *Blockchain) calcFeeHandler(cmd command.Command, _ command.AppID, _ string, args ...string) command.CommandResult {
	amt, err := amount.FromString(args[0])
	if err != nil {
		return cmd.ErrorResult(err)
	}

	fee, err := bc.clientMgr.GetFee(int64(amt))
	if err != nil {
		return cmd.ErrorResult(err)
	}

	calcedFee := amount.Amount(fee)

	return cmd.SuccessfulResult("Sending %s will cost %s with current fee percentage."+
		"\n> Note: Consider unbond and sortition transaction fee is 0 PAC always.", amt, calcedFee.String())
}
