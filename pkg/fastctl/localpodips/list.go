package localpodips

import (
	"fmt"

	"github.com/cilium/ebpf"
	"github.com/gosuri/uitable"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"

	bpfmap "github.com/fast-io/fast/bpf/map"
	"github.com/fast-io/fast/pkg/util"
)

type listOptions struct {
	ioStream genericclioptions.IOStreams

	localPodIpsMap *ebpf.Map
}

func newListOptions(ipStream genericclioptions.IOStreams) *listOptions {
	return &listOptions{
		ioStream:       ipStream,
		localPodIpsMap: bpfmap.GetLocalPodIpsMap(),
	}
}

func NewListCommand(name string, ioStreaam genericclioptions.IOStreams) *cobra.Command {
	o := newListOptions(ioStreaam)
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list local pod ips",
		Long:    "list local pod ips",
		Example: fmt.Sprintf("    %s localips ls", name),
		Run: func(cmd *cobra.Command, args []string) {
			cmdutil.CheckErr(o.Complete(cmd, args))
			cmdutil.CheckErr(o.Validate(args))
			cmdutil.CheckErr(o.Run())
		},
	}
	return cmd
}

func (o *listOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

func (o *listOptions) Validate(args []string) error {
	return nil
}

func (o *listOptions) Run() error {
	var (
		key   bpfmap.LocalIpsMapKey
		value bpfmap.LocalIpsMapInfo
	)

	table := uitable.New()
	table.MaxColWidth = 80
	table.AddRow("LOCALIP", "MAC", "NODEMAC", "IFINDEX", "LXCIFINDEX")
	iter := o.localPodIpsMap.Iterate()
	for iter.Next(&key, &value) {
		table.AddRow(util.InetUint32ToIp(key.IP), value.MAC, value.NodeMAC, value.IfIndex, value.LxcIfIndex)
	}
	fmt.Println(table)
	return nil
}
