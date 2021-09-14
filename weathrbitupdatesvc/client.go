package weatherbitupdatesvc

type Client interface {
	ForceRestart() error
}
