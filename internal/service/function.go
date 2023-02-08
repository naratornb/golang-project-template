package service

import (
	cm "cmd/internal/common"
	"fmt"
	"net/http"
)

type Service interface {
	Monitor(w http.ResponseWriter, r *http.Request)
	Download(w http.ResponseWriter, r *http.Request)
}

type service struct {
	sftp   cm.SFTP
	client cm.HttpClient
	// Repositories/Rules/Client defined here
	/*
		batchEventRepo repository.BatchEvent
		riskCalReqRepo repository.RiskCalReq
		riskRuleRepo   repository.RiskRule
		riskEmptyRepo  repository.RiskEmpty
		sftp           cm.SFTP
		client         cm.HttpClient
		normalRule     []repository.RiskRuleEntity
		bScoreRule     []repository.RiskRuleEntity
	*/
}

// Service Instant initiation here
func NewService(sftp cm.SFTP, client cm.HttpClient) Service {
	return &service{
		sftp:   sftp,
		client: client,
		/*
			batchEventRepo: batchEventRepo,
			riskCalReqRepo: riskCalReqRepo,
			riskRuleRepo:   riskRuleRepo,
			riskEmptyRepo:  riskEmptyRepo,
			sftp:           sftp,
			client:         client,
			normalRule:     normal,
			bScoreRule:     bScore,
		*/
	}
}

// Logic Implement here

func (s service) Monitor(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Println("{endpoint}/batch/Monitor is executed")
}

func (s service) Download(w http.ResponseWriter, r *http.Request) {
	// path := chi.URLParam(r, "path")
	w.Header().Set("Content-type", "application/json")
	fmt.Println("{endpoint}/sftp/download is executed")
	s.sftp.Download("./", "./", "file.txt")
}

/*
func (s service) Upload(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	fmt.Println("{endpoint}/sftp/upload is executed")
		zap.S().Info("..Start Upload RSKU file to sftp..")
		path := os.Getenv("LOCAL_PATH")
		lclFile := fmt.Sprintf("%v", path)
		err := s.sftp.Upload(lclFile, os.Getenv("SFTP_RSKU_FILE_PATH"), batchEvent.RskuFilename)
		if err != nil {
			zap.S().Errorf("s.sftp.Upload: %v", err)
			return err
		}
		zap.S().Infof("Upload file %v successfully", batchEvent.RskuFilename)
		return nil
	}
	zap.S().Infof("Skip uploadFile due to batchEvent not allowed: %v", batchEvent.BatchStep)
	return nil
}

*/
