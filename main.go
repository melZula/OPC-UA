package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gopcua/opcua"
	"github.com/gopcua/opcua/ua"
	"github.com/melZula/OPC-UA/telegram"
)

var (
	endpoint string
	bot      *telegram.Bot
)

func readValue(nodeID string) float32 {
	ctx := context.Background()

	c := opcua.NewClient(endpoint, opcua.SecurityMode(ua.MessageSecurityModeNone))
	if err := c.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	id, err := ua.ParseNodeID(nodeID)
	if err != nil {
		log.Fatalf("invalid node id: %v", err)
	}

	req := &ua.ReadRequest{
		MaxAge: 2000,
		NodesToRead: []*ua.ReadValueID{
			&ua.ReadValueID{NodeID: id},
		},
		TimestampsToReturn: ua.TimestampsToReturnBoth,
	}

	resp, err := c.Read(req)
	if err != nil {
		log.Fatalf("Read failed: %s", err)
	}
	if resp.Results[0].Status != ua.StatusOK {
		log.Fatalf("Status not OK: %v", resp.Results[0].Status)
	}
	// log.Printf("%#v", resp.Results[0].Value.Value())
	return resp.Results[0].Value.Value().(float32)
}

func readAndCheckValues(nodes []string, min float32, max float32, freq string) {
	d, err := time.ParseDuration(freq)
	if err != nil {
		log.Fatal("Invalid freq")
		return
	}
	for {
		time.Sleep(d)
		for _, node := range nodes {
			curr := readValue(node)
			if curr > max || curr < min {
				log.Printf("Out of value: %s ", node)
				bot.Notify(fmt.Sprintf("Out of value: %s ", node))
			}
		}
	}
}

func main() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := NewConfig()
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal("Invalid config")
	}
	endpoint = config.URL

	bot = telegram.NewBot(config.APIKey, config.ChanID)
	bot.Start()

	readAndCheckValues(config.Nodes, config.Min, config.Max, config.Freq)
}
