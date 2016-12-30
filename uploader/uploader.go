package uploader

import (
	"bytes"
	"fmt"
	"io"
	"log"

	gcs "cloud.google.com/go/storage"
	"github.com/go-microservices/resizer/option"
	"github.com/go-microservices/resizer/storage"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	opt "google.golang.org/api/option"
)

const (
	scope     = gcs.ScopeFullControl
	sixMonths = 60 * 60 * 24 * 30 * 6
)

type Uploader struct {
	context    context.Context
	bucket     *gcs.BucketHandle
	bucketName string
}

// New はアップローダーを作成する。
func New(o option.Options) (*Uploader, error) {
	ctx := context.Background()
	client, err := gcs.NewClient(ctx, opt.WithScopes(gcs.ScopeFullControl), opt.WithServiceAccountFile(o.GCServiceAccount))
	if err != nil {
		return nil, errors.Wrap(err, "can't create client for GCS")
	}
	bkt := client.Bucket(o.GCStorageBucket)
	return &Uploader{ctx, bkt, o.GCStorageBucket}, nil
}

func (u *Uploader) Upload(buf *bytes.Buffer, f storage.Image) (string, error) {
	object := u.bucket.Object(f.Filename)
	w := object.NewWriter(u.context)
	written, err := io.Copy(w, buf)
	if err != nil {
		return "", errors.Wrap(err, "can't copy buffer to GCS object writer")
	}
	if err := w.Close(); err != nil {
		return "", errors.Wrap(err, "can't close object writer")
	}
	log.Printf("Write %d bytes object '%s' in bucket '%s'", written, f.Filename, u.bucketName)

	attrs, err := object.Update(u.context, gcs.ObjectAttrsToUpdate{
		ContentType:  f.ContentType,
		CacheControl: fmt.Sprintf("max-age=%d", sixMonths),
	})
	if err != nil {
		return "", errors.Wrap(err, "can't update object attributes")
	}
	log.Printf("Attributes: %+v", *attrs)

	url := u.CreateURL(f.Filename)
	return url, nil
}

func (u *Uploader) CreateURL(path string) string {
	return fmt.Sprintf("https://%s.storage.googleapis.com/%s", u.bucketName, path)
}
