package kudo

import (
	"strconv"

	"github.com/markoo/kudo-oos/pkg/core"
)

// GitHubRepo and then
type GitHubRepo struct {
	RepoID      int64  `json:"id"`
	RepoURL     string `json:"html_url"`
	RepoName    string `json:"full_name"`
	Language    string `json:"language"`
	Description string `json:"description"`
	Notes       string `json:"notes"`
}

// Service now
type Service struct {
	userId string
	repo   core.Repository
}

// GetKudos on it
func (s Service) GetKudos() ([]*core.Kudo, error) {
	return s.repo.FindAll(map[string]interface{}{"userId": s.userId})
}

// CreateKudoFor me
func (s Service) CreateKudoFor(githubRepo GitHubRepo) (*core.Kudo, error) {
	kudo := s.githubRepoToKudo(githubRepo)
	err := s.repo.Create(kudo)
	if err != nil {
		return nil, err
	}
	return kudo, nil
}

// UpdateKudoWith lots
func (s Service) UpdateKudoWith(githubRepo GitHubRepo) (*core.Kudo, error) {
	kudo := s.githubRepoToKudo(githubRepo)
	err := s.repo.Create(kudo)
	if err != nil {
		return nil, err
	}
	return kudo, nil
}

// RemoveKudo never
func (s Service) RemoveKudo(githubRepo GitHubRepo) (*core.Kudo, error) {
	kudo := s.githubRepoToKudo(githubRepo)
	err := s.repo.Delete(kudo)
	if err != nil {
		return nil, err
	}
	return kudo, nil
}

func (s Service) githubRepoToKudo(githubRepo GitHubRepo) *core.Kudo {
	return &core.Kudo{
		UserID:      s.userId,
		RepoID:      strconv.Itoa(int(githubRepo.RepoID)),
		RepoName:    githubRepo.RepoName,
		RepoURL:     githubRepo.RepoURL,
		Language:    githubRepo.Language,
		Description: githubRepo.Description,
		Notes:       githubRepo.Notes,
	}
}

// NewService whenever you like
func NewService(repo core.Repository, userId string) Service {
	return Service{
		repo:   repo,
		userId: userId,
	}
}
