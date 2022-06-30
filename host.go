package host

import "github.com/adimax2953/host/internal"

func Startup(app interface{}) *Starter {
	return internal.NewStarter(app)
}

func RegisterHostService(starter *Starter, service HostService) {
	internal.RegisterHostService(starter, service)
}

func StdHostServiceInstance() HostService {
	return internal.StdHostServiceInstance()
}
