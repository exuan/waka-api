package service

import "github.com/exuan/waka-api/validator"


//@todo
func (s *service) BatchHeartbeat(vhs *[]*validator.Heartbeat) (int64, error) {
	return s.r.BatchHeartbeat(vhs)
}


