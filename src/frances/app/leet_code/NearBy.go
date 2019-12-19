package main


type ClientSessionKey struct {
	uin      uint64
}

type ClientSession struct {
	key     ClientSessionKey
	geoHash string
}

type RandomArrangeMng struct {
	reqChan       chan *ClientSession
	sessions      map[ClientSessionKey]*ClientSession
}

func main() {
	
}
