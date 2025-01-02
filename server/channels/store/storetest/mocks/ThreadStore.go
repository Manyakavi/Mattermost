// Code generated by mockery v2.42.2. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	store "github.com/mattermost/mattermost/server/v8/channels/store"
	mock "github.com/stretchr/testify/mock"
)

// ThreadStore is an autogenerated mock type for the ThreadStore type
type ThreadStore struct {
	mock.Mock
}

// DeleteMembershipForUser provides a mock function with given fields: userID, postID
func (_m *ThreadStore) DeleteMembershipForUser(userID string, postID string) error {
	ret := _m.Called(userID, postID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMembershipForUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userID, postID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteMembershipsForChannel provides a mock function with given fields: userID, channelID
func (_m *ThreadStore) DeleteMembershipsForChannel(userID string, channelID string) error {
	ret := _m.Called(userID, channelID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteMembershipsForChannel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userID, channelID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOrphanedRows provides a mock function with given fields: limit
func (_m *ThreadStore) DeleteOrphanedRows(limit int) (int64, error) {
	ret := _m.Called(limit)

	if len(ret) == 0 {
		panic("no return value specified for DeleteOrphanedRows")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (int64, error)); ok {
		return rf(limit)
	}
	if rf, ok := ret.Get(0).(func(int) int64); ok {
		r0 = rf(limit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *ThreadStore) Get(id string) (*model.Thread, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *model.Thread
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Thread, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Thread); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Thread)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMembershipForUser provides a mock function with given fields: userID, postID
func (_m *ThreadStore) GetMembershipForUser(userID string, postID string) (*model.ThreadMembership, error) {
	ret := _m.Called(userID, postID)

	if len(ret) == 0 {
		panic("no return value specified for GetMembershipForUser")
	}

	var r0 *model.ThreadMembership
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*model.ThreadMembership, error)); ok {
		return rf(userID, postID)
	}
	if rf, ok := ret.Get(0).(func(string, string) *model.ThreadMembership); ok {
		r0 = rf(userID, postID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ThreadMembership)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, postID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMembershipsForUser provides a mock function with given fields: userID, teamID
func (_m *ThreadStore) GetMembershipsForUser(userID string, teamID string) ([]*model.ThreadMembership, error) {
	ret := _m.Called(userID, teamID)

	if len(ret) == 0 {
		panic("no return value specified for GetMembershipsForUser")
	}

	var r0 []*model.ThreadMembership
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]*model.ThreadMembership, error)); ok {
		return rf(userID, teamID)
	}
	if rf, ok := ret.Get(0).(func(string, string) []*model.ThreadMembership); ok {
		r0 = rf(userID, teamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ThreadMembership)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userID, teamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTeamsUnreadForUser provides a mock function with given fields: userID, teamIDs, includeUrgentMentionCount
func (_m *ThreadStore) GetTeamsUnreadForUser(userID string, teamIDs []string, includeUrgentMentionCount bool) (map[string]*model.TeamUnread, error) {
	ret := _m.Called(userID, teamIDs, includeUrgentMentionCount)

	if len(ret) == 0 {
		panic("no return value specified for GetTeamsUnreadForUser")
	}

	var r0 map[string]*model.TeamUnread
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []string, bool) (map[string]*model.TeamUnread, error)); ok {
		return rf(userID, teamIDs, includeUrgentMentionCount)
	}
	if rf, ok := ret.Get(0).(func(string, []string, bool) map[string]*model.TeamUnread); ok {
		r0 = rf(userID, teamIDs, includeUrgentMentionCount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*model.TeamUnread)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []string, bool) error); ok {
		r1 = rf(userID, teamIDs, includeUrgentMentionCount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetThreadFollowers provides a mock function with given fields: threadID, fetchOnlyActive
func (_m *ThreadStore) GetThreadFollowers(threadID string, fetchOnlyActive bool) ([]string, error) {
	ret := _m.Called(threadID, fetchOnlyActive)

	if len(ret) == 0 {
		panic("no return value specified for GetThreadFollowers")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, bool) ([]string, error)); ok {
		return rf(threadID, fetchOnlyActive)
	}
	if rf, ok := ret.Get(0).(func(string, bool) []string); ok {
		r0 = rf(threadID, fetchOnlyActive)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, bool) error); ok {
		r1 = rf(threadID, fetchOnlyActive)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetThreadForUser provides a mock function with given fields: threadMembership, extended, postPriorityIsEnabled
func (_m *ThreadStore) GetThreadForUser(threadMembership *model.ThreadMembership, extended bool, postPriorityIsEnabled bool) (*model.ThreadResponse, error) {
	ret := _m.Called(threadMembership, extended, postPriorityIsEnabled)

	if len(ret) == 0 {
		panic("no return value specified for GetThreadForUser")
	}

	var r0 *model.ThreadResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.ThreadMembership, bool, bool) (*model.ThreadResponse, error)); ok {
		return rf(threadMembership, extended, postPriorityIsEnabled)
	}
	if rf, ok := ret.Get(0).(func(*model.ThreadMembership, bool, bool) *model.ThreadResponse); ok {
		r0 = rf(threadMembership, extended, postPriorityIsEnabled)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ThreadResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.ThreadMembership, bool, bool) error); ok {
		r1 = rf(threadMembership, extended, postPriorityIsEnabled)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetThreadMembershipsForExport provides a mock function with given fields: postID
func (_m *ThreadStore) GetThreadMembershipsForExport(postID string) ([]*model.ThreadMembershipForExport, error) {
	ret := _m.Called(postID)

	if len(ret) == 0 {
		panic("no return value specified for GetThreadMembershipsForExport")
	}

	var r0 []*model.ThreadMembershipForExport
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]*model.ThreadMembershipForExport, error)); ok {
		return rf(postID)
	}
	if rf, ok := ret.Get(0).(func(string) []*model.ThreadMembershipForExport); ok {
		r0 = rf(postID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ThreadMembershipForExport)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(postID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetThreadUnreadReplyCount provides a mock function with given fields: threadMembership
func (_m *ThreadStore) GetThreadUnreadReplyCount(threadMembership *model.ThreadMembership) (int64, error) {
	ret := _m.Called(threadMembership)

	if len(ret) == 0 {
		panic("no return value specified for GetThreadUnreadReplyCount")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.ThreadMembership) (int64, error)); ok {
		return rf(threadMembership)
	}
	if rf, ok := ret.Get(0).(func(*model.ThreadMembership) int64); ok {
		r0 = rf(threadMembership)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*model.ThreadMembership) error); ok {
		r1 = rf(threadMembership)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetThreadsForUser provides a mock function with given fields: userID, teamID, opts
func (_m *ThreadStore) GetThreadsForUser(userID string, teamID string, opts model.GetUserThreadsOpts) ([]*model.ThreadResponse, error) {
	ret := _m.Called(userID, teamID, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetThreadsForUser")
	}

	var r0 []*model.ThreadResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) ([]*model.ThreadResponse, error)); ok {
		return rf(userID, teamID, opts)
	}
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) []*model.ThreadResponse); ok {
		r0 = rf(userID, teamID, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ThreadResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, model.GetUserThreadsOpts) error); ok {
		r1 = rf(userID, teamID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalThreads provides a mock function with given fields: userID, teamID, opts
func (_m *ThreadStore) GetTotalThreads(userID string, teamID string, opts model.GetUserThreadsOpts) (int64, error) {
	ret := _m.Called(userID, teamID, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetTotalThreads")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) (int64, error)); ok {
		return rf(userID, teamID, opts)
	}
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) int64); ok {
		r0 = rf(userID, teamID, opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, string, model.GetUserThreadsOpts) error); ok {
		r1 = rf(userID, teamID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalUnreadMentions provides a mock function with given fields: userID, teamID, opts
func (_m *ThreadStore) GetTotalUnreadMentions(userID string, teamID string, opts model.GetUserThreadsOpts) (int64, error) {
	ret := _m.Called(userID, teamID, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetTotalUnreadMentions")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) (int64, error)); ok {
		return rf(userID, teamID, opts)
	}
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) int64); ok {
		r0 = rf(userID, teamID, opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, string, model.GetUserThreadsOpts) error); ok {
		r1 = rf(userID, teamID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalUnreadThreads provides a mock function with given fields: userID, teamID, opts
func (_m *ThreadStore) GetTotalUnreadThreads(userID string, teamID string, opts model.GetUserThreadsOpts) (int64, error) {
	ret := _m.Called(userID, teamID, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetTotalUnreadThreads")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) (int64, error)); ok {
		return rf(userID, teamID, opts)
	}
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) int64); ok {
		r0 = rf(userID, teamID, opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, string, model.GetUserThreadsOpts) error); ok {
		r1 = rf(userID, teamID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalUnreadUrgentMentions provides a mock function with given fields: userID, teamID, opts
func (_m *ThreadStore) GetTotalUnreadUrgentMentions(userID string, teamID string, opts model.GetUserThreadsOpts) (int64, error) {
	ret := _m.Called(userID, teamID, opts)

	if len(ret) == 0 {
		panic("no return value specified for GetTotalUnreadUrgentMentions")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) (int64, error)); ok {
		return rf(userID, teamID, opts)
	}
	if rf, ok := ret.Get(0).(func(string, string, model.GetUserThreadsOpts) int64); ok {
		r0 = rf(userID, teamID, opts)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string, string, model.GetUserThreadsOpts) error); ok {
		r1 = rf(userID, teamID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaintainMembership provides a mock function with given fields: userID, postID, opts
func (_m *ThreadStore) MaintainMembership(userID string, postID string, opts store.ThreadMembershipOpts) (*model.ThreadMembership, error) {
	ret := _m.Called(userID, postID, opts)

	if len(ret) == 0 {
		panic("no return value specified for MaintainMembership")
	}

	var r0 *model.ThreadMembership
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, store.ThreadMembershipOpts) (*model.ThreadMembership, error)); ok {
		return rf(userID, postID, opts)
	}
	if rf, ok := ret.Get(0).(func(string, string, store.ThreadMembershipOpts) *model.ThreadMembership); ok {
		r0 = rf(userID, postID, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ThreadMembership)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, store.ThreadMembershipOpts) error); ok {
		r1 = rf(userID, postID, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaintainMultipleFromImport provides a mock function with given fields: memberships
func (_m *ThreadStore) MaintainMultipleFromImport(memberships []*model.ThreadMembership) ([]*model.ThreadMembership, error) {
	ret := _m.Called(memberships)

	if len(ret) == 0 {
		panic("no return value specified for MaintainMultipleFromImport")
	}

	var r0 []*model.ThreadMembership
	var r1 error
	if rf, ok := ret.Get(0).(func([]*model.ThreadMembership) ([]*model.ThreadMembership, error)); ok {
		return rf(memberships)
	}
	if rf, ok := ret.Get(0).(func([]*model.ThreadMembership) []*model.ThreadMembership); ok {
		r0 = rf(memberships)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ThreadMembership)
		}
	}

	if rf, ok := ret.Get(1).(func([]*model.ThreadMembership) error); ok {
		r1 = rf(memberships)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MarkAllAsRead provides a mock function with given fields: userID, threadIds
func (_m *ThreadStore) MarkAllAsRead(userID string, threadIds []string) error {
	ret := _m.Called(userID, threadIds)

	if len(ret) == 0 {
		panic("no return value specified for MarkAllAsRead")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(userID, threadIds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkAllAsReadByChannels provides a mock function with given fields: userID, channelIDs
func (_m *ThreadStore) MarkAllAsReadByChannels(userID string, channelIDs []string) error {
	ret := _m.Called(userID, channelIDs)

	if len(ret) == 0 {
		panic("no return value specified for MarkAllAsReadByChannels")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(userID, channelIDs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkAllAsReadByTeam provides a mock function with given fields: userID, teamID
func (_m *ThreadStore) MarkAllAsReadByTeam(userID string, teamID string) error {
	ret := _m.Called(userID, teamID)

	if len(ret) == 0 {
		panic("no return value specified for MarkAllAsReadByTeam")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(userID, teamID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MarkAsRead provides a mock function with given fields: userID, threadID, timestamp
func (_m *ThreadStore) MarkAsRead(userID string, threadID string, timestamp int64) error {
	ret := _m.Called(userID, threadID, timestamp)

	if len(ret) == 0 {
		panic("no return value specified for MarkAsRead")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, int64) error); ok {
		r0 = rf(userID, threadID, timestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PermanentDeleteBatchForRetentionPolicies provides a mock function with given fields: now, globalPolicyEndTime, limit, cursor
func (_m *ThreadStore) PermanentDeleteBatchForRetentionPolicies(now int64, globalPolicyEndTime int64, limit int64, cursor model.RetentionPolicyCursor) (int64, model.RetentionPolicyCursor, error) {
	ret := _m.Called(now, globalPolicyEndTime, limit, cursor)

	if len(ret) == 0 {
		panic("no return value specified for PermanentDeleteBatchForRetentionPolicies")
	}

	var r0 int64
	var r1 model.RetentionPolicyCursor
	var r2 error
	if rf, ok := ret.Get(0).(func(int64, int64, int64, model.RetentionPolicyCursor) (int64, model.RetentionPolicyCursor, error)); ok {
		return rf(now, globalPolicyEndTime, limit, cursor)
	}
	if rf, ok := ret.Get(0).(func(int64, int64, int64, model.RetentionPolicyCursor) int64); ok {
		r0 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int64, int64, int64, model.RetentionPolicyCursor) model.RetentionPolicyCursor); ok {
		r1 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r1 = ret.Get(1).(model.RetentionPolicyCursor)
	}

	if rf, ok := ret.Get(2).(func(int64, int64, int64, model.RetentionPolicyCursor) error); ok {
		r2 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PermanentDeleteBatchThreadMembershipsForRetentionPolicies provides a mock function with given fields: now, globalPolicyEndTime, limit, cursor
func (_m *ThreadStore) PermanentDeleteBatchThreadMembershipsForRetentionPolicies(now int64, globalPolicyEndTime int64, limit int64, cursor model.RetentionPolicyCursor) (int64, model.RetentionPolicyCursor, error) {
	ret := _m.Called(now, globalPolicyEndTime, limit, cursor)

	if len(ret) == 0 {
		panic("no return value specified for PermanentDeleteBatchThreadMembershipsForRetentionPolicies")
	}

	var r0 int64
	var r1 model.RetentionPolicyCursor
	var r2 error
	if rf, ok := ret.Get(0).(func(int64, int64, int64, model.RetentionPolicyCursor) (int64, model.RetentionPolicyCursor, error)); ok {
		return rf(now, globalPolicyEndTime, limit, cursor)
	}
	if rf, ok := ret.Get(0).(func(int64, int64, int64, model.RetentionPolicyCursor) int64); ok {
		r0 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(int64, int64, int64, model.RetentionPolicyCursor) model.RetentionPolicyCursor); ok {
		r1 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r1 = ret.Get(1).(model.RetentionPolicyCursor)
	}

	if rf, ok := ret.Get(2).(func(int64, int64, int64, model.RetentionPolicyCursor) error); ok {
		r2 = rf(now, globalPolicyEndTime, limit, cursor)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SaveMultipleMemberships provides a mock function with given fields: memberships
func (_m *ThreadStore) SaveMultipleMemberships(memberships []*model.ThreadMembership) ([]*model.ThreadMembership, error) {
	ret := _m.Called(memberships)

	if len(ret) == 0 {
		panic("no return value specified for SaveMultipleMemberships")
	}

	var r0 []*model.ThreadMembership
	var r1 error
	if rf, ok := ret.Get(0).(func([]*model.ThreadMembership) ([]*model.ThreadMembership, error)); ok {
		return rf(memberships)
	}
	if rf, ok := ret.Get(0).(func([]*model.ThreadMembership) []*model.ThreadMembership); ok {
		r0 = rf(memberships)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ThreadMembership)
		}
	}

	if rf, ok := ret.Get(1).(func([]*model.ThreadMembership) error); ok {
		r1 = rf(memberships)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateMembership provides a mock function with given fields: membership
func (_m *ThreadStore) UpdateMembership(membership *model.ThreadMembership) (*model.ThreadMembership, error) {
	ret := _m.Called(membership)

	if len(ret) == 0 {
		panic("no return value specified for UpdateMembership")
	}

	var r0 *model.ThreadMembership
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.ThreadMembership) (*model.ThreadMembership, error)); ok {
		return rf(membership)
	}
	if rf, ok := ret.Get(0).(func(*model.ThreadMembership) *model.ThreadMembership); ok {
		r0 = rf(membership)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ThreadMembership)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.ThreadMembership) error); ok {
		r1 = rf(membership)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewThreadStore creates a new instance of ThreadStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewThreadStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *ThreadStore {
	mock := &ThreadStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}