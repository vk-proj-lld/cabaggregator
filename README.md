# cabaggregator

## Problem Statement
1. m drivers and n riders, m<n
2. rider request a ride, any available driver in the list can respond to the request
    with Accept/Reject.
3. Driver Acceptance/Rejection probability is 50%.
4. Driver takes sometime to process a request.(say ~5 secs)
5. All drivers (excep busy drivers) has to be informed about the new request **in parallel**.

##### Question: 
    Tell which driver will be assigned to which rider.
    all outputs are to be in console or can be to channel.

##### Note 
At the some ride requests will be accepted, and some will not be.
Driver can serve one rider at a time(no pooling).

### Project Structure
```
├── README.md
├── broadcast.go
├── business
│   ├── repo
│   │   └── dispatcherRepo.go
│   └── uc
│       └── discpatcher.go
├── entities
│   ├── driver
│   │   ├── driver.go
│   │   ├── istrategy.go
│   │   ├── signal.go
│   │   └── strategy.go
│   ├── out
│   │   ├── fileout.go
│   │   ├── iout.go
│   │   └── stdo.go
│   └── rider
│       ├── request.go
│       └── rider.go
├── go.mod
├── interfaces
│   └── idispatcher
│       ├── dispatcher.go
│       └── dispatcherrepo.go
├── main.go
└── utils
    └── seed.go
```
