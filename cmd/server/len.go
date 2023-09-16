package server

//func (s *server) TopicLength(ctx context.Context, in *proto.DequeueMsg) (*proto.Length, error) {
//	if in.GetTopicName() == "" {
//		s.l.Error("topic field is required")
//		return &proto.Length{Length: 0}, fmt.Errorf("topic name is required")
//	}
//
//	//return &proto.Length{Length: int32(s.topics.Len(in.GetTopicName()))}, nil
//}
