package chat

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	WELCOME = `chat: talk to ChatGPT
- you can input "quit" or "exit" to quit the process.
- you can input "help" to get help.`
	LoadContextSize    = 10
	ContextHistorySize = 1000
)

type Messages []openai.ChatCompletionMessage

func NewHttpTransport() *http.Transport {
	if !DefaultSetting.UseProxy {
		return &http.Transport{}
	}
	proxyUrl, err := url.Parse(DefaultSetting.ProxyUrl)
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	return transport
}

func NewClient() *openai.Client {
	config := openai.DefaultConfig(DefaultSetting.AuthToken)
	if DefaultSetting.BaseURL != "" {
		config.BaseURL = DefaultSetting.BaseURL
	}

	config.HTTPClient = &http.Client{
		Transport: NewHttpTransport(),
	}

	return openai.NewClientWithConfig(config)
}

type Chat struct {
	client    *openai.Client
	ps1color  *color.Color
	respColor *color.Color
	hintColor *color.Color
	idx       int32
	msgs      Messages
}

func (c *Chat) AiResp(qs string) {
	c.msgs = append(c.msgs, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: qs,
	})

	msgs := c.msgs
	if len(c.msgs) > LoadContextSize {
		msgs = c.msgs[len(c.msgs)-LoadContextSize:]
	}
	req := openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: msgs,
		Stream:   true,
	}

	stream, err := c.client.CreateChatCompletionStream(context.Background(), req)

	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	content := ""
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			os.Exit(1)
		}
		deltaContent := response.Choices[0].Delta.Content
		c.respColor.Printf(deltaContent)
		content += deltaContent
	}
	c.msgs = append(c.msgs, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: content,
	})
	fmt.Println()
}

func (c *Chat) Bye() {
	c.respColor.Println("bye~")
	fmt.Println()
	os.Exit(0)
}

func (c *Chat) Help() {
	c.respColor.Println(WELCOME)
	fmt.Println()
}

func (c *Chat) Empty() {
}

func (c *Chat) Translater(raw, form, to string) string {
	res := TranslateNotCode(raw, form, to)
	c.ps1color.Printf("hint: ")
	c.hintColor.Println(res)
	fmt.Println()
	return res
}

func (c *Chat) Ai(qs string) {
	tqs := qs
	if DefaultSetting.UseTranslator && ContainsNativeLanguage(qs) {
		tqs = c.Translater(qs, DefaultSetting.NativeLanguage, "en")
	}

	c.AiResp(tqs)
	fmt.Println()

	res := c.msgs[len(c.msgs)-1].Content

	if DefaultSetting.UseTranslator && !ContainsNativeLanguage(res) {
		c.Translater(res, "en", DefaultSetting.NativeLanguage)
	}
}

func (c *Chat) Resp(qs string) {
	switch {
	case InArray[string]([]string{"quit", "exit"}, qs):
		c.Bye()
	case InArray[string]([]string{"h", "help"}, qs):
		c.Help()
	case strings.TrimSpace(qs) == "":
		c.Empty()
	default:
		c.Ai(qs)
		c.idx++
	}
}

func (c *Chat) Chat() {
	// messages := make([]openai.ChatCompletionMessage, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	c.Help()

	for {
		c.ps1color.Printf("%03d %s", c.idx, DefaultSetting.PS1)
		qs, _ := reader.ReadString('\n')
		// convert CRLF to LF
		qs = strings.Replace(qs, "\n", "", -1)
		fmt.Println()

		c.Resp(qs)
	}
}

type Once struct {
	*Chat
	contextConfig string
}

func (c *Once) restoreContext() {

	if _, err := os.Stat(c.contextConfig); os.IsNotExist(err) {
		// fmt.Println("data.json does not exist")
		return
	}
	// Read the JSON file
	file, err := os.ReadFile(c.contextConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON data
	err = json.Unmarshal(file, &c.msgs)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *Once) saveContext() {
	file, err := os.Create(c.contextConfig)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if len(c.msgs) > ContextHistorySize {
		c.msgs = c.msgs[len(c.msgs)-ContextHistorySize:]
	}
	if err := encoder.Encode(c.msgs); err != nil {
		panic(err)
	}
}

func (c *Once) ChatOnce(qs string) {
	fmt.Println()
	c.restoreContext()
	c.Resp(qs)
	c.saveContext()
}

func NewChat() *Chat {
	return &Chat{
		client:    NewClient(),
		ps1color:  color.New().Add(color.FgHiCyan),
		respColor: color.New().Add(color.FgHiGreen),
		hintColor: color.New().Add(color.FgHiMagenta),
		idx:       1,
		msgs:      make([]openai.ChatCompletionMessage, 0),
	}
}

func NewChatOnce() *Once {
	return &Once{
		NewChat(),
		filepath.Join(homeDir(), contextName),
	}
}
