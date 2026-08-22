package service

import (
	"log/slog"
	"net/http"
	"sync"

	"github.com/medasset/medasset/internal/constants"
	"github.com/medasset/medasset/internal/dto"
	"github.com/medasset/medasset/internal/model"
	"github.com/medasset/medasset/internal/repository"
	"github.com/medasset/medasset/internal/util"
)

func processAuditBatch(logs []model.AuditLog, write func(*model.AuditLog) error, workerGate <-chan struct{}) <-chan struct{} {
	var workers sync.WaitGroup
	done := make(chan struct{})
	for i := range logs {
		entry := logs[i].CloneForBatch()
		go func() {
			if workerGate != nil {
				<-workerGate
			}
			workers.Add(1)
			defer workers.Done()
			_ = write(&entry)
		}()
	}
	go func() {
		workers.Wait()
		close(done)
	}()
	return done
}

// AuditService 审计日志服务。
type AuditService struct {
	repo *repository.AuditRepository
	log  *slog.Logger
}

func NewAuditService(repo *repository.AuditRepository, log *slog.Logger) *AuditService {
	return &AuditService{repo: repo, log: log}
}

// Record 写入审计日志（被所有业务 service 复用）。
func (s *AuditService) Record(userID uint, username, action, module, entityID, detail, ip, requestID string) {
	a := &model.AuditLog{
		UserID:    userID,
		Username:  username,
		Action:    action,
		Module:    module,
		EntityID:  entityID,
		Detail:    detail,
		IP:        ip,
		RequestID: requestID,
	}
	if err := s.repo.Create(a); err != nil {
		s.log.Error("写入审计日志失败", "err", err)
	}
}

// List 分页查询审计日志。
func (s *AuditService) List(page, pageSize int, module, username string) (*util.PageResult, error) {
	list, total, err := s.repo.List(page, pageSize, module, username)
	if err != nil {
		s.log.Error("查询审计日志失败", "err", err)
		return nil, util.NewAppError(http.StatusInternalServerError, constants.MsgInternalError, err)
	}
	items := make([]dto.AuditListItem, 0, len(list))
	for i := range list {
		items = append(items, dto.ToAuditListItem(&list[i]))
	}
	return &util.PageResult{List: items, Total: total, Page: page, PageSize: pageSize}, nil
}
