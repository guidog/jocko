package protocol

var APIVersions = []APIVersion{
	{APIKey: ProduceKey, MinVersion: 3, MaxVersion: 5},
	{APIKey: FetchKey, MinVersion: 4, MaxVersion: 4},
	{APIKey: OffsetsKey, MinVersion: 0, MaxVersion: 2},
	{APIKey: MetadataKey, MinVersion: 0, MaxVersion: 1},
	{APIKey: LeaderAndISRKey, MinVersion: 0, MaxVersion: 1},
	{APIKey: StopReplicaKey, MinVersion: 0, MaxVersion: 0},
	{APIKey: FindCoordinatorKey, MinVersion: 0, MaxVersion: 1},
	{APIKey: JoinGroupKey, MinVersion: 0, MaxVersion: 1},
	{APIKey: HeartbeatKey, MinVersion: 0, MaxVersion: 1},
	{APIKey: LeaveGroupKey, MinVersion: 0, MaxVersion: 1},
	{APIKey: SyncGroupKey, MinVersion: 0, MaxVersion: 1},
	{APIKey: DescribeGroupsKey, MinVersion: 0, MaxVersion: 1},
	{APIKey: ListGroupsKey, MinVersion: 0, MaxVersion: 1},
	{APIKey: APIVersionsKey, MinVersion: 0, MaxVersion: 1},
	{APIKey: CreateTopicsKey, MinVersion: 0, MaxVersion: 1},
	{APIKey: DeleteTopicsKey, MinVersion: 0, MaxVersion: 1},
}
