package uc

import (
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
		rides    chan *rider.RideRequest
		disprepo idispatcher.IDispatcherRepo
		out      out.IOout
		logger   out.IOout
	}
	type args struct {
		ride    *rider.RideRequest
		drivers []*driver.Driver
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
				ride: rider.NewRideRequest(1, time.Now()),
				drivers: []*driver.Driver{
					driver.NewDriver("d1", driver.NewEqualChoiceStrategy(2*time.Second, driver.AckAccept, driver.AckReject)),
					driver.NewDriver("d2", driver.NewEqualChoiceStrategy(2*time.Second, driver.AckAccept, driver.AckReject)),
					driver.NewDriver("d3", driver.NewEqualChoiceStrategy(2*time.Second, driver.AckAccept, driver.AckReject)),
					driver.NewDriver("d4", driver.NewEqualChoiceStrategy(2*time.Second, driver.AckAccept, driver.AckReject)),
					driver.NewDriver("d5", driver.NewEqualChoiceStrategy(2*time.Second, driver.AckAccept, driver.AckReject)),
					driver.NewDriver("d6", driver.NewEqualChoiceStrategy(2*time.Second, driver.AckAccept, driver.AckReject)),
					driver.NewDriver("d7", driver.NewEqualChoiceStrategy(2*time.Second, driver.AckAccept, driver.AckReject)),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			disp := &dispatcher{
				rides:    tt.fields.rides,
				disprepo: tt.fields.disprepo,
				out:      tt.fields.out,
				logger:   tt.fields.logger,
			}
			disp.broadcast(tt.args.ride, tt.args.drivers)
		})
	}
}
