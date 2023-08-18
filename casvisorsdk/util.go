package casvisorsdk

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"strings"
)

func (c *Client) GetUrl(action string, queryMap map[string]string) string {
	query := ""
	for k, v := range queryMap {
		query += fmt.Sprintf("%s=%s&", k, v)
	}
	query = strings.TrimRight(query, "&")

	url := fmt.Sprintf("%s/api/%s?%s", c.Endpoint, action, query)
	return url
}

func createFormFile(formData map[string][]byte) (string, io.Reader, error) {
	// https://tonybai.com/2021/01/16/upload-and-download-file-using-multipart-form-over-http/

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	defer w.Close()

	for k, v := range formData {
		pw, err := w.CreateFormFile(k, "file")
		if err != nil {
			panic(err)
		}

		_, err = pw.Write(v)
		if err != nil {
			panic(err)
		}
	}

	return w.FormDataContentType(), body, nil
}

func createForm(formData map[string]string) (string, io.Reader, error) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range formData {
		if err := w.WriteField(k, v); err != nil {
			return "", nil, err
		}
	}
	if err := w.Close(); err != nil {
		return "", nil, err
	}

	return w.FormDataContentType(), body, nil
}
