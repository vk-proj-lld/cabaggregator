package uc

import (
	"sync"
	"testing"
	"time"

	"github.com/vk-proj-lld/cabaggregator/business/repo"
	"github.com/vk-proj-lld/cabaggregator/entities/driver"
	"github.com/vk-proj-lld/cabaggregator/entities/out"
	"github.com/vk-proj-lld/cabaggregator/entities/rider"
	"github.com/vk-proj-lld/cabaggregator/interfaces/idispatcher"
)

func Test_dispatcher_broadcast(t *testing.T) {
	type fields struct {
		ridedrivers chan rideDriverResp
		disprepo    idispatcher.IDispatcherRepo
		out         out.IOout
		logger      out.IOout
	}
	type args struct {
		unitmu   *sync.Once
		drchanel chan *driver.Driver
		ride     *rider.RideRequest
		drivers  []*driver.Driver
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test-1",
			fields: fields{
				disprepo: repo.NewDispatcherRepo(),
				logger:   out.NewStdO(),
				out:      out.NewStdO(),
			},
			args: args{
				unitmu:   &sync.Once{},
				ride:     rider.NewRideRequest(1, time.Now()),
				drchanel: make(chan *driver.Driver),
				drivers: []*driver.Driver{
					driver.NewDriver("d1", driver.NewEqualChoiceStrategy(driver.AckAccept, driver.AckReject)),
					driver.NewDriver("d2", driver.NewEqualChoiceStrategy(driver.AckAccept, driver.AckReject)),
					driver.NewDriver("d3", driver.NewEqualChoiceStrategy(driver.AckAccept, driver.AckReject)),
					driver.NewDriver("d4", driver.NewEqualChoiceStrategy(driver.AckAccept, driver.AckReject)),
					driver.NewDriver("d5", driver.NewEqualChoiceStrategy(driver.AckAccept, driver.AckReject)),
					driver.NewDriver("d6", driver.NewEqualChoiceStrategy(driver.AckAccept, driver.AckReject)),
					driver.NewDriver("d7", driver.NewEqualChoiceStrategy(driver.AckAccept, driver.AckReject)),
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disp, _ := NewDispatcher(tt.fields.disprepo, tt.fields.out, tt.fields.logger).(*dispatcher)
			disp.broadcast(tt.args.drchanel, tt.args.ride, tt.args.drivers...)
			t.Log(<-tt.args.drchanel)
		})
	}
}
