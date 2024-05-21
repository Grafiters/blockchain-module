package lib

type LibBlockchain interface {
	FetchBlocks(height int64) ([]*TranscationConvertion, error)
}
