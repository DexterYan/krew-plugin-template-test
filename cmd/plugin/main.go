package main

import (
	"github.com/DexterYan/krew-plugin-template-test/cmd/plugin/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp" // required for GKE
)

func main() {
	cli.InitAndExecute()
}
