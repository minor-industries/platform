package discovery

import (
	"github.com/hashicorp/serf/client"
	"github.com/minor-industries/platform/common/serf_service"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Serf struct {
	addr string
}

func NewSerf(serfAddr string) *Serf {
	return &Serf{addr: serfAddr}
}

func (s *Serf) Discover(logger *zap.Logger) ([]*Entry, error) {
	var result []*Entry

	// create a new client every time due to observed errors
	serfClient, err := client.NewRPCClient(s.addr)
	if err != nil {
		return nil, errors.Wrap(err, "new client")
	}

	defer func() {
		err := serfClient.Close()
		if err != nil {
			logger.Info("error closing serf client", zap.Error(err))
		}
	}()

	services, err := serf_service.LoadServices(logger, serfClient)
	if err != nil {
		return nil, errors.Wrap(err, "load services")
	}

	for _, s := range services {
		result = append(result, &Entry{
			Service:  s.Name,
			Instance: s.Instance,
			Hostname: s.Host,
			Port:     s.ServicePort,
			Addr:     s.Addr,
		})
	}

	return result, nil
}
