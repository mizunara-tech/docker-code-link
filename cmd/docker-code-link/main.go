package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"docker-code-link/internal/dockerutil"
	"docker-code-link/pkg/hexutil"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

const version = "0.0.1"

func main() {
	var rootCmd = &cobra.Command{Use: "docker-code-link"}

	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "List active Docker containers",
		Long:  `List all active Docker containers that can be attached to VSCode.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Dockerクライアントの初期化
			client, err := docker.NewClientFromEnv()
			if err != nil {
				log.Fatalf("Error initializing Docker client: %v", err)
			}

			// 稼働中のコンテナのリストを取得
			containers, err := client.ListContainers(docker.ListContainersOptions{All: false})
			if err != nil {
				log.Fatalf("Error listing containers: %v", err)
			}

			// コンテナの名前のリストを作成
			var containerNames []string
			for _, container := range containers {
				containerNames = append(containerNames, container.Names[0])
			}

			// promptuiを使用してコンテナを選択
			prompt := promptui.Select{
				Label: "Select Docker Container",
				Items: containerNames,
			}

			_, result, err := prompt.Run()
			if err != nil {
				log.Fatalf("Prompt failed %v\n", err)
			}
			fmt.Printf("You selected %q\n", result)

			// コンテナの作業ディレクトリを取得
			workingDir, _ := dockerutil.GetContainerWorkingDir(client, result)
			if err != nil {
				log.Fatalf("Failed to get container working directory: %v", err)
			}

			hexName := hexutil.HexEncode(result)

			// VSCodeで選択したコンテナにアタッチ
			vscodeRemoteURI := fmt.Sprintf("vscode-remote://attached-container+%s%s", hexName, workingDir)
			vscodeCmd := exec.Command("code", "--folder-uri", vscodeRemoteURI)
			err = vscodeCmd.Start()
			if err != nil {
				log.Fatalf("Failed to start VSCode: %v", err)
			}
		},
	}

	rootCmd.AddCommand(listCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getContainerWorkingDir(containerName string) (string, error) {
	// Dockerクライアントの初期化
	client, err := docker.NewClientFromEnv()
	if err != nil {
		return "", fmt.Errorf("Error initializing Docker client: %v", err)
	}

	// InspectContainerWithOptionsを使用してコンテナ情報を取得
	container, err := client.InspectContainerWithOptions(docker.InspectContainerOptions{ID: containerName})
	if err != nil {
		return "", fmt.Errorf("Error inspecting container %s: %v", containerName, err)
	}

	// コンテナのWorkingDirを返す
	return container.Config.WorkingDir, nil
}

// hexEncode 関数は文字列を16進数に変換します。
func hexEncode(str string) string {
	hexStr := ""
	for _, r := range str {
		hexStr += fmt.Sprintf("%x", r)
	}
	return hexStr
}
