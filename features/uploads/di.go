package uploads

import "noteapp/di"

func init() {
	_ = di.Container.Provide(func() UploadService {
		return NewUploadService()
	})
}
