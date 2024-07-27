package websockets

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"github.com/ivang5/doodle-guessr/server/internal/util"
)

const (
	IMAGE_WIDTH  = 64
	IMAGE_HEIGHT = 64
	IMAGE_SIZE   = IMAGE_WIDTH * IMAGE_HEIGHT
	///? NOTE: Mirko A.
	/// If the received image and the last image sent to
	/// the model differ by at least this many pixels, we
	/// send the infer request again.
	IMAGE_DIFF_THRESHOLD = 10
)

type Client struct {
	conn  *websocket.Conn
	image []int
}

func NewClient(conn *websocket.Conn) *Client {
	image := make([]int, IMAGE_SIZE)
	for i := 0; i < IMAGE_SIZE; i++ {
		image[i] = 1
	}

	return &Client{
		conn,
		image,
	}
}

var errCodes = [...]int{websocket.CloseGoingAway, websocket.CloseAbnormalClosure}

func (c *Client) Run() {
	var req util.InferRequest

	log.Println("New client connected")
	defer func() {
		c.conn.Close()
		log.Println("Client disconnected")
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, errCodes[:]...) {
				log.Println("Error (ReadMessage) unexpected close error")
				log.Printf("   |_ %v\n", err.Error())
			}
			break
		}

		if err = json.Unmarshal(msg, &req); err != nil {
			log.Println("Error (ReadMessage) when unmarshalling predict request")
			log.Printf("   |_ %v\n", err.Error())
			continue
		}

		if len(req.Pixels) != IMAGE_SIZE {
			log.Println("Error (ReadMessage) incorrect pixel array length")
			continue
		}

		if countDifferentPixels(c.image, req.Pixels) < IMAGE_DIFF_THRESHOLD {
			continue
		}
		c.cacheImage(req.Pixels)

		responseBody, err := util.SendInferRequest(msg)
		if err != nil {
			log.Println("Error (PredictHandler) when sending infer request")
			log.Printf("   |_ %v\n", err.Error())
			continue
		}

		c.conn.WriteMessage(websocket.TextMessage, responseBody)
	}
}

func (c *Client) cacheImage(image []int) {
	for i := 0; i < IMAGE_SIZE; i += 1 {
		c.image[i] = image[i]
	}
}

func countDifferentPixels(cachedImage []int, incomingImage []int) int {
	numDifferent := 0
	for i := 0; i < IMAGE_SIZE; i += 1 {
		if cachedImage[i] != incomingImage[i] {
			numDifferent += 1
		}
	}

	return numDifferent
}
