package project2

import "context"

func GetResponse(client gpt3.Client, ctx context.Context, question string){
	err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{question},
		MaxTokens: gpt3.IntPtr(3000),
		Stop:      []string{"."},
		Echo:      true,
}
func main() {
	viper.SetConfigName(".env")
	viper.ReadInConfig()
	apiKey := viper.GetString("API_KEY")
	if apiKey == "" {
		log.Fatal("missing api key")
	}
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	rootcmd := &cobra.Command{
		Use:   "chatgpt",
		Short: "Chat with Chatgpt in Console",
		Run: func(cmd *cobra.Command, args []string) {
			response, err := client.Completion(ctx, gpt3.CompletionRequest{
				Prompt:    []string{"this is the first thing you have to know about golang"},
				MaxTokens: gpt3.IntPtr(30),
				Stop:      []string{"."},
				Echo:      true,
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(response.Choices[0].Text)
		},
	}
	rootcmd.Execute()
}