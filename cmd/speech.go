package cmd

import (
	voicerss "Text-to-Speech/api"
	"Text-to-Speech/audio"

	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert text to speech",
	Long:  `Convert text to speech`,
	Run:   convertTextToSpeech,
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringP("language", "l", "en-us", "Language to convert")
	convertCmd.Flags().StringP("text", "t", "Hello World", "Text to convert")
}

func convertTextToSpeech(cmd *cobra.Command, args []string) {
	language, _ := cmd.Flags().GetString("language")
	text, _ := cmd.Flags().GetString("text")

	wav := voicerss.ConvertToSpeech(language, text)

	audio.PlayWavAudio(wav)
}
