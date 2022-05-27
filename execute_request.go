package gobotapi

import (
	"bytes"
	"context"
	"fmt"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
	"io"
	"mime/multipart"
	"net/http"
)

func (ctx *Client) executeRequest(url string, data map[string]string, files map[string]rawTypes.InputFile) ([]byte, error) {
	client := http.Client{}
	var body io.Reader
	var multiPartWriter *multipart.Writer
	if len(data) > 0 || len(files) > 0 {
		reader := &bytes.Buffer{}
		multiPartWriter = multipart.NewWriter(reader)
		for k, v := range data {
			_ = multiPartWriter.WriteField(k, v)
		}
		for k, v := range files {
			file, err := multiPartWriter.CreateFormFile(k, v.FileName())
			if err != nil {
				return nil, err
			}
			_, err = file.Write(v.Content())
			if err != nil {
				return nil, err
			}
		}
		_ = multiPartWriter.Close()
		body = reader
	}
	ctxConn, cancelable := context.WithCancel(context.Background())
	cancelableContext := rawTypes.CancelableContext{
		Cancel: cancelable,
	}
	cancelableContext.GenerateId()
	defer func() {
		for _, x := range ctx.requestsContext {
			if x.Id == cancelableContext.Id {
				ctx.requestsContext = append(ctx.requestsContext[:0], ctx.requestsContext[1:]...)
				break
			}
		}
	}()
	req, err := http.NewRequestWithContext(ctxConn, "POST", url, body)
	ctx.requestsContext = append(ctx.requestsContext, cancelableContext)
	if err != nil {
		return nil, err
	}
	if multiPartWriter != nil {
		req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())
	}
	do, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(do.Body)
	var buf bytes.Buffer
	_, err = io.Copy(&buf, do.Body)
	if err != nil {
		return nil, err
	}
	if do.StatusCode != http.StatusOK && do.StatusCode != http.StatusCreated && do.StatusCode != http.StatusNoContent {
		return buf.Bytes(), fmt.Errorf("error code %d", do.StatusCode)
	}
	return buf.Bytes(), nil
}