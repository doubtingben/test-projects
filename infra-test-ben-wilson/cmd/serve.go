package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/INFURA/infra-test-ben-wilson/uuid"
	"github.com/fatih/color"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var infuraAPIEndpoint = "https://pmainnet.infura.io/"

type ethNumberBlock struct {
	JSON   string `json:"json"`
	Result string `json:"result"`
	ID     uint64 `json:"id"`
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a web server for infura-json-endpoint",
	Long:  `Runs a web server process designed to be fronted with a reverse proxy.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if commandLineFlags.apiKey == "" {
			color.Red("apiKey required!")
			os.Exit(100)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var router = gin.Default()
		store := persistence.NewInMemoryStore(time.Second)

		// Set up healthz endpoint
		router.GET("/healthz", func(c *gin.Context) {
			c.String(200, `{"Status": "online"}`)
		})

		// Set up endpoints
		router.GET("/eth_blockNumber", cache.CachePage(store, time.Minute, callINFURA))
		router.GET("/eth_gasPrice", cache.CachePage(store, time.Minute, callINFURA))
		router.GET("/eth_coinbase", cache.CachePage(store, time.Minute, callINFURA))

		// Start listening
		router.Run(":" + strconv.Itoa(int(commandLineFlags.ListenPort)))
	},
}

func init() {
	serveCmd.Flags().Uint16VarP(&commandLineFlags.ListenPort, "listen-port", "p", 8080, "Port to listen on for HTTP requests")

	RootCmd.AddCommand(serveCmd)
}

func callINFURAethNumberBlock(c *gin.Context) {
	var infraURI = c.Request.RequestURI
	infuraResp, err := http.Get(infuraAPIEndpoint + infraURI + "?token=" + commandLineFlags.apiKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// TODO Check/verify response
	fmt.Println("INFURA reponse: " + infuraResp.Status)
	defer infuraResp.Body.Close()

	var d ethNumberBlock

	if err := json.NewDecoder(infuraResp.Body).Decode(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("response Status:", infuraResp.Status)
	fmt.Println("response Headers:", infuraResp.Header)
	body, _ := ioutil.ReadAll(infuraResp.Body)
	fmt.Println("response Body:", string(body))

	c.JSON(200, gin.H{"result": d.Result, "id": d.ID, "timestamp": time.Now().String()})
}

func callINFURA(c *gin.Context) {
	// strip request / for method name
	url := "https://pmainnet.infura.io/"
	r := regexp.MustCompile("^/(?P<method>[a-zA-Z_]+)")
	m := r.FindStringSubmatch(c.Request.RequestURI)[1]

	uu, err := uuid.NewUUID()

	// TODO Parse and append params
	type infuraRequest struct {
		JSONRPC string      `json:"jsonrpc"`
		Method  string      `json:"method"`
		ID      string      `json:"id"`
		Params  interface{} `json:"params,omitempty"`
	}

	ir := infuraRequest{JSONRPC: "2.0", Method: m, ID: uu}
	jsonValue, _ := json.Marshal(ir)

	infuraReq, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	infuraReq.Header.Set("Content-Type", "application/json")

	//dump, err := httputil.DumpRequestOut(infuraReq, true)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%q", dump)

	client := &http.Client{}
	infuraResp, err := client.Do(infuraReq)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer infuraResp.Body.Close()

	// fmt.Println("response Status:", infuraResp.Status)
	// fmt.Println("response Headers:", infuraResp.Header)
	// body, _ := ioutil.ReadAll(infuraResp.Body)
	// fmt.Println("response Body:", string(body))

	// TODO Check/verify response
	fmt.Println("INFURA reponse: " + infuraResp.Status)

	type infuraResponse struct {
		ID      string `json:"id"`
		JSONRPC string `json:"jsonrpc"`
		Result  string `json:"result"`
		Error   *struct {
			Code    int16  `json:"code"`
			Message string `json:"message"`
			Data    string `json:"date"`
		} `json:"error,omitempty"`
	}

	var d infuraResponse

	if err := json.NewDecoder(infuraResp.Body).Decode(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if d.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": d.Error})
		return
	}

	c.JSON(200, gin.H{"result": d.Result, "id": d.ID, "timestamp": time.Now().String()})
}
