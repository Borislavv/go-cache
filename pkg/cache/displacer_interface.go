package displacer

import "github.com/Borislavv/go-cache/pkg/cache/storage"

type Displacer interface {
	Run(storage storage.Storage)
	Stop()
}
