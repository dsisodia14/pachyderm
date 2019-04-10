package cmds

import (
	"fmt"
	"os"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/pachyderm/pachyderm/src/client"
	"github.com/pachyderm/pachyderm/src/client/pkg/config"
	"github.com/pachyderm/pachyderm/src/client/pkg/grpcutil"
	txnclient "github.com/pachyderm/pachyderm/src/client/transaction"
	"github.com/pachyderm/pachyderm/src/server/pkg/cmdutil"
	"github.com/pachyderm/pachyderm/src/server/pkg/tabwriter"
	"github.com/pachyderm/pachyderm/src/server/transaction/pretty"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Cmds(noMetrics *bool, noPortForwarding *bool) []*cobra.Command {
	var commands []*cobra.Command

	marshaller := &jsonpb.Marshaler{Indent: "  "}

	raw := false
	rawFlags := pflag.NewFlagSet("", pflag.ContinueOnError)
	rawFlags.BoolVar(&raw, "raw", false, "disable pretty printing, print raw json")

	fullTimestamps := false
	fullTimestampsFlags := pflag.NewFlagSet("", pflag.ContinueOnError)
	fullTimestampsFlags.BoolVar(&fullTimestamps, "full-timestamps", false, "Return absolute timestamps (as opposed to the default, relative timestamps).")

	transactionDocs := &cobra.Command{
		Short: "Docs for transactions.",
		Long:  "Do a few thing.",
	}
	cmdutil.SetDocsUsage(transactionDocs)
	commands = append(commands, cmdutil.CreateAlias(transactionDocs, "transaction"))

	listTransaction := &cobra.Command{
		Short: "List transactions.",
		Long:  "List transactions.",
		Run: cmdutil.RunFixedArgs(0, func([]string) error {
			c, err := client.NewOnUserMachine(!*noMetrics, !*noPortForwarding, "user")
			if err != nil {
				return err
			}
			defer c.Close()
			transactions, err := c.ListTransaction()
			if err != nil {
				return err
			}
			if raw {
				for _, transaction := range transactions {
					if err := marshaller.Marshal(os.Stdout, transaction); err != nil {
						return err
					}
				}
				return nil
			}
			writer := tabwriter.NewWriter(os.Stdout, pretty.TransactionHeader)
			for _, transaction := range transactions {
				pretty.PrintTransactionInfo(writer, transaction, fullTimestamps)
			}
			return writer.Flush()
		}),
	}
	listTransaction.Flags().AddFlagSet(fullTimestampsFlags)
	commands = append(commands, cmdutil.CreateAlias(listTransaction, "list transaction"))

	getActiveTransaction := func() (string, error) {
		cfg, err := config.Read()
		if err != nil {
			return "", fmt.Errorf("error reading Pachyderm config: %v", err)
		}
		if cfg.V1 == nil || cfg.V1.ActiveTransaction == "" {
			return "", fmt.Errorf("no active transaction")
		}
		return cfg.V1.ActiveTransaction, nil
	}

	setActiveTransaction := func(id string) error {
		cfg, err := config.Read()
		if err != nil {
			return fmt.Errorf("error reading Pachyderm config: %v", err)
		}
		if cfg.V1 == nil {
			cfg.V1 = &config.ConfigV1{}
		}
		cfg.V1.ActiveTransaction = id
		if err := cfg.Write(); err != nil {
			return fmt.Errorf("error writing Pachyderm config: %v", err)
		}
		return nil
	}

	startTransaction := &cobra.Command{
		Short: "Start a new transaction.",
		Long:  "Start a new transaction.",
		Run: cmdutil.RunFixedArgs(0, func([]string) error {
			c, err := client.NewOnUserMachine(!*noMetrics, !*noPortForwarding, "user")
			if err != nil {
				return err
			}
			defer c.Close()
			transaction, err := c.StartTransaction()
			if err != nil {
				return grpcutil.ScrubGRPC(err)
			}
			// TODO: use advisory locks on config so we don't have a race condition if
			// two commands are run simultaneously
			err = setActiveTransaction(transaction.ID)
			if err != nil {
				return err
			}
			fmt.Printf("Started new transaction: %s\n", transaction.ID)
			return nil
		}),
	}
	commands = append(commands, cmdutil.CreateAlias(startTransaction, "start transaction"))

	stopTransaction := &cobra.Command{
		Short: "Stop modifying the current transaction.",
		Long:  "Start modifying the current transaction.",
		Run: cmdutil.RunFixedArgs(0, func([]string) error {
			// TODO: use advisory locks on config so we don't have a race condition if
			// two commands are run simultaneously
			id, err := getActiveTransaction()
			if err != nil {
				return err
			}

			err = setActiveTransaction("")
			if err != nil {
				return err
			}

			fmt.Printf("Cleared active transaction: %s\n", id)
			return nil
		}),
	}
	commands = append(commands, cmdutil.CreateAlias(stopTransaction, "stop transaction"))

	finishTransaction := &cobra.Command{
		Use:   "{{alias}} [<transaction>]",
		Short: "Close the currently active transaction and execute it.",
		Long:  "Close the currently active transaction and execute it.",
		Run: cmdutil.RunBoundedArgs(0, 1, func(args []string) error {
			c, err := client.NewOnUserMachine(!*noMetrics, !*noPortForwarding, "user")
			if err != nil {
				return err
			}
			defer c.Close()

			// TODO: use advisory locks on config so we don't have a race condition if
			// two commands are run simultaneously
			var id string
			if len(args) > 0 {
				id = args[0]
			} else {
				id, err = getActiveTransaction()
				if err != nil {
					return err
				}
			}

			info, err := c.FinishTransaction(&txnclient.Transaction{ID: id})
			if err != nil {
				return grpcutil.ScrubGRPC(err)
			}

			err = setActiveTransaction("")
			if err != nil {
				return err
			}

			fmt.Printf("Completed transaction with %d requests: %s\n", len(info.Responses), info.Transaction.ID)
			return nil
		}),
	}
	commands = append(commands, cmdutil.CreateAlias(finishTransaction, "finish transaction"))

	deleteTransaction := &cobra.Command{
		Use:   "{{alias}} [<transaction>]",
		Short: "Cancel and delete an existing transaction.",
		Long:  "Cancel and delete an existing transaction.",
		Run: cmdutil.RunBoundedArgs(0, 1, func(args []string) error {
			c, err := client.NewOnUserMachine(!*noMetrics, !*noPortForwarding, "user")
			if err != nil {
				return err
			}
			defer c.Close()

			// TODO: use advisory locks on config so we don't have a race condition if
			// two commands are run simultaneously
			var id string
			var isActive bool
			if len(args) > 0 {
				id = args[0]

				// Don't check err here, this is just a quality-of-life check to clean
				// up the config after a successful delete
				activeTransaction, _ := getActiveTransaction()
				isActive = id == activeTransaction
			} else {
				id, err = getActiveTransaction()
				if err != nil {
					return err
				}
				isActive = true
			}

			err = c.DeleteTransaction(&txnclient.Transaction{ID: id})
			if err != nil {
				return grpcutil.ScrubGRPC(err)
			}
			if isActive {
				// The active transaction was successfully deleted, clean it up so the
				// user doesn't need to manually 'stop transaction' it.
				setActiveTransaction("")
			}
			return nil
		}),
	}
	commands = append(commands, cmdutil.CreateAlias(deleteTransaction, "delete transaction"))

	inspectTransaction := &cobra.Command{
		Use:   "{{alias}} [<transaction>]",
		Short: "Print information about an open transaction.",
		Long:  "Print information about an open transaction.",
		Run: cmdutil.RunBoundedArgs(0, 1, func(args []string) error {
			c, err := client.NewOnUserMachine(!*noMetrics, !*noPortForwarding, "user")
			if err != nil {
				return err
			}
			defer c.Close()

			var id string
			if len(args) > 0 {
				id = args[0]
			} else {
				id, err = getActiveTransaction()
				if err != nil {
					return err
				}
			}

			info, err := c.InspectTransaction(&txnclient.Transaction{ID: id})
			if err != nil {
				return grpcutil.ScrubGRPC(err)
			}
			if info == nil {
				return fmt.Errorf("transaction %s not found", id)
			}
			if raw {
				return marshaller.Marshal(os.Stdout, info)
			}
			return pretty.PrintDetailedTransactionInfo(&pretty.PrintableTransactionInfo{
				TransactionInfo: info,
				FullTimestamps:  fullTimestamps,
			})
		}),
	}
	inspectTransaction.Flags().AddFlagSet(fullTimestampsFlags)
	commands = append(commands, cmdutil.CreateAlias(inspectTransaction, "inspect transaction"))

	resumeTransaction := &cobra.Command{
		Use:   "{{alias}} <transaction>",
		Short: "Set an existing transaction as active.",
		Long:  "Set an existing transaction as active.",
		Run: cmdutil.RunFixedArgs(1, func(args []string) error {
			c, err := client.NewOnUserMachine(!*noMetrics, !*noPortForwarding, "user")
			if err != nil {
				return err
			}
			defer c.Close()
			info, err := c.InspectTransaction(&txnclient.Transaction{ID: args[0]})
			if err != nil {
				return grpcutil.ScrubGRPC(err)
			}
			if info == nil {
				return fmt.Errorf("transaction %s not found", args[0])
			}

			err = setActiveTransaction(info.Transaction.ID)
			if err != nil {
				return err
			}

			fmt.Printf("Resuming existing transaction with %d requests: %s\n", len(info.Requests), info.Transaction.ID)
			return nil
		}),
	}
	commands = append(commands, cmdutil.CreateAlias(resumeTransaction, "resume transaction"))

	return commands
}
