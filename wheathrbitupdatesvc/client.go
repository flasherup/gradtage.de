package weatherbitsvc

type Client interface {
	ForceRestart() error
}
