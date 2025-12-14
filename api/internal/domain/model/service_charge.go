package model

import (
	"errors"
	"time"
)

// ============================================================
// ServiceChargeStatusType - 請求ステータスタイプ
// ============================================================

type ServiceChargeStatusType string

const (
	ServiceChargeStatusReserved  ServiceChargeStatusType = "RESERVED"
	ServiceChargeStatusCompleted ServiceChargeStatusType = "COMPLETED"
	ServiceChargeStatusFailed    ServiceChargeStatusType = "FAILED"
	ServiceChargeStatusCancelled ServiceChargeStatusType = "CANCELLED"
)

// ============================================================
// ServiceChargeSource - 請求明細の発生元
// ============================================================

// ServiceChargeSource は請求明細の発生元（ワンショットまたはサブスクリプション）
type ServiceChargeSource interface {
	isServiceChargeSource()
	SourceType() string
}

// OneShotSource はワンショット課金からの請求明細
type OneShotSource struct {
	oneShotUsageID string
}

func NewOneShotSource(oneShotUsageID string) OneShotSource {
	return OneShotSource{oneShotUsageID: oneShotUsageID}
}

func (o OneShotSource) isServiceChargeSource() {}
func (o OneShotSource) SourceType() string     { return "ONE_SHOT" }
func (o OneShotSource) OneShotUsageID() string { return o.oneShotUsageID }

// SubscriptionSource はサブスクリプションからの請求明細
type SubscriptionSource struct {
	subscriptionID string
}

func NewSubscriptionSource(subscriptionID string) SubscriptionSource {
	return SubscriptionSource{subscriptionID: subscriptionID}
}

func (s SubscriptionSource) isServiceChargeSource() {}
func (s SubscriptionSource) SourceType() string     { return "SUBSCRIPTION" }
func (s SubscriptionSource) SubscriptionID() string { return s.subscriptionID }

// ManualSource は手動で追加された請求明細
type ManualSource struct {
	reason string
}

func NewManualSource(reason string) ManualSource {
	return ManualSource{reason: reason}
}

func (m ManualSource) isServiceChargeSource() {}
func (m ManualSource) SourceType() string     { return "MANUAL" }
func (m ManualSource) Reason() string         { return m.reason }

// ============================================================
// ServiceChargeItem - 請求明細（集約内のエンティティ）
// ============================================================

// ServiceChargeItem は請求明細（集約内のエンティティ）
type ServiceChargeItem struct {
	id     string
	name   string
	amount int64
	source ServiceChargeSource
}

func (sci *ServiceChargeItem) ID() string                  { return sci.id }
func (sci *ServiceChargeItem) Name() string                { return sci.name }
func (sci *ServiceChargeItem) Amount() int64               { return sci.amount }
func (sci *ServiceChargeItem) Source() ServiceChargeSource { return sci.source }

// ReconstructServiceChargeItem は永続化層から復元する
func ReconstructServiceChargeItem(
	id string,
	name string,
	amount int64,
	source ServiceChargeSource,
) *ServiceChargeItem {
	return &ServiceChargeItem{
		id:     id,
		name:   name,
		amount: amount,
		source: source,
	}
}

// ============================================================
// ServiceChargeStatus - 請求ステータス履歴（集約内のエンティティ）
// ============================================================

// ServiceChargeStatus は請求ステータス履歴（集約内のエンティティ）
type ServiceChargeStatus struct {
	id        string
	status    ServiceChargeStatusType
	createdAt time.Time
}

func (scs *ServiceChargeStatus) ID() string                      { return scs.id }
func (scs *ServiceChargeStatus) Status() ServiceChargeStatusType { return scs.status }
func (scs *ServiceChargeStatus) CreatedAt() time.Time            { return scs.createdAt }

// ReconstructServiceChargeStatus は永続化層から復元する
func ReconstructServiceChargeStatus(
	id string,
	status ServiceChargeStatusType,
	createdAt time.Time,
) *ServiceChargeStatus {
	return &ServiceChargeStatus{
		id:        id,
		status:    status,
		createdAt: createdAt,
	}
}

// ============================================================
// ServiceCharge - 請求の集約ルート
// ============================================================

// ServiceCharge は請求の集約ルート
type ServiceCharge struct {
	id        string
	accountID string
	startDate time.Time
	endDate   time.Time
	items     []*ServiceChargeItem
	statuses  []*ServiceChargeStatus
}

func NewServiceCharge(
	accountID string,
	startDate time.Time,
	endDate time.Time,
) *ServiceCharge {
	now := time.Now()
	return &ServiceCharge{
		accountID: accountID,
		startDate: startDate,
		endDate:   endDate,
		items:     []*ServiceChargeItem{},
		statuses: []*ServiceChargeStatus{
			{
				id:        "", // リポジトリ層で採番
				status:    ServiceChargeStatusReserved,
				createdAt: now,
			},
		},
	}
}

func (sc *ServiceCharge) ID() string           { return sc.id }
func (sc *ServiceCharge) AccountID() string    { return sc.accountID }
func (sc *ServiceCharge) StartDate() time.Time { return sc.startDate }
func (sc *ServiceCharge) EndDate() time.Time   { return sc.endDate }
func (sc *ServiceCharge) Items() []*ServiceChargeItem {
	// 防御的コピー
	copied := make([]*ServiceChargeItem, len(sc.items))
	copy(copied, sc.items)
	return copied
}
func (sc *ServiceCharge) Statuses() []*ServiceChargeStatus {
	// 防御的コピー
	copied := make([]*ServiceChargeStatus, len(sc.statuses))
	copy(copied, sc.statuses)
	return copied
}

// TotalAmount は請求明細の合計金額を計算する
func (sc *ServiceCharge) TotalAmount() int64 {
	var total int64
	for _, item := range sc.items {
		total += item.amount
	}
	return total
}

// LatestStatus は最新のステータスを返す
func (sc *ServiceCharge) LatestStatus() *ServiceChargeStatus {
	if len(sc.statuses) == 0 {
		return nil
	}
	return sc.statuses[len(sc.statuses)-1]
}

// LatestStatusID は最新のステータスIDを返す
func (sc *ServiceCharge) LatestStatusID() *string {
	latest := sc.LatestStatus()
	if latest == nil {
		return nil
	}
	return &latest.id
}

// AddItem は請求明細を追加する
func (sc *ServiceCharge) AddItem(name string, amount int64, source ServiceChargeSource) error {
	if name == "" {
		return errors.New("item name must not be empty")
	}
	if amount < 0 {
		return errors.New("item amount must not be negative")
	}

	// ステータスがCOMPLETEDの場合は追加不可
	if sc.LatestStatus() != nil && sc.LatestStatus().status == ServiceChargeStatusCompleted {
		return errors.New("cannot add item to completed charge")
	}

	item := &ServiceChargeItem{
		id:     "", // リポジトリ層で採番
		name:   name,
		amount: amount,
		source: source,
	}
	sc.items = append(sc.items, item)
	return nil
}

// UpdateStatus はステータスを更新する（履歴として追加）
func (sc *ServiceCharge) UpdateStatus(statusType ServiceChargeStatusType) error {
	// ビジネスルール：既にCOMPLETEDまたはCANCELLEDなら変更不可
	latestStatus := sc.LatestStatus()
	if latestStatus != nil {
		if latestStatus.status == ServiceChargeStatusCompleted {
			return errors.New("cannot update status of completed charge")
		}
		if latestStatus.status == ServiceChargeStatusCancelled {
			return errors.New("cannot update status of cancelled charge")
		}
	}

	status := &ServiceChargeStatus{
		id:        "", // リポジトリ層で採番
		status:    statusType,
		createdAt: time.Now(),
	}
	sc.statuses = append(sc.statuses, status)
	return nil
}

// Complete は請求を完了状態にする
func (sc *ServiceCharge) Complete() error {
	if len(sc.items) == 0 {
		return errors.New("cannot complete charge without items")
	}
	return sc.UpdateStatus(ServiceChargeStatusCompleted)
}

// Cancel は請求をキャンセルする
func (sc *ServiceCharge) Cancel() error {
	return sc.UpdateStatus(ServiceChargeStatusCancelled)
}

// Fail は請求を失敗状態にする
func (sc *ServiceCharge) Fail() error {
	return sc.UpdateStatus(ServiceChargeStatusFailed)
}

// ReconstructServiceCharge は永続化層から復元する
func ReconstructServiceCharge(
	id string,
	accountID string,
	startDate time.Time,
	endDate time.Time,
	items []*ServiceChargeItem,
	statuses []*ServiceChargeStatus,
) *ServiceCharge {
	return &ServiceCharge{
		id:        id,
		accountID: accountID,
		startDate: startDate,
		endDate:   endDate,
		items:     items,
		statuses:  statuses,
	}
}
