package repository

import (
	"context"
	"fmt"

	"github.com/PichuChen/go-bbs"
)

func (repo *repository) GetBoards(_ context.Context) []bbs.BoardRecord {
	return repo.boardRecords
}

func (repo *repository) GetBoardArticle(_ context.Context, boardID, filename string) ([]byte, error) {
	return repo.db.ReadBoardArticleFile(boardID, filename)
}

func (repo *repository) GetBoardArticleRecords(_ context.Context, boardID string) ([]bbs.ArticleRecord, error) {
	return repo.db.ReadBoardArticleRecordsFile(boardID)
}

func (repo *repository) GetPopularArticleRecords(_ context.Context) ([]BBSPopularArticleRecord, error) {
	// Note: go-bbs has not implemented this yet
	// TODO: delegate to repo.db when it is ready
	return make([]BBSPopularArticleRecord, 0), nil
}

// // BBSPopularArticleRecord : Mock of bbs.PopularArticleRecord
// type BBSPopularArticleRecord interface {
// 	// Note: go-bbs has not implemented this yet
// 	// TODO: use bbs.PopularArticleRecord or something when it is ready
// 	bbs.ArticleRecord
// 	BoardId() string
// }

//BBSPopularArticleRecord : Mock of bbs.PopularArticleRecord
type BBSPopularArticleRecord interface {
	// Note: go-bbs has not implemented this yet
	// TODO: use bbs.PopularArticleRecord or something when it is ready
	bbs.ArticleRecord
}

type bbsPopularArticleRecord struct {
	// Note: go-bbs has not implemented this yet
	// TODO: use bbs.PopularArticleRecord or something when it is ready
	bbs.ArticleRecord
}

// BoardId returns the board id of the popular article
func (record *bbsPopularArticleRecord) BoardId() string {
	return ""
}

func (repo *repository) GetBoardTreasureRecords(_ context.Context, boardID string, treasureIDs []string) ([]bbs.ArticleRecord, error) {
	return repo.db.ReadBoardTreasureRecordsFile(boardID, treasureIDs)
}

func loadBoardFile(db *bbs.DB) ([]bbs.BoardRecord, error) {
	boardRecords, err := db.ReadBoardRecords()
	if err != nil {
		logger.Errorf("get board header error: %v", err)
		return nil, fmt.Errorf("failed to read board records: %w", err)
	}
	for index, board := range boardRecords {
		logger.Debugf("loaded %d %v", index, board.BoardId())
	}
	return boardRecords, nil
}
