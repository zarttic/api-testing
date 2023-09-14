/*
MIT License

Copyright (c) 2023 API Testing Authors.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package runner_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/h2non/gock"
	"github.com/linuxsuren/api-testing/pkg/runner"
	"github.com/stretchr/testify/assert"
)

func TestPutRecordToPrometheus(t *testing.T) {
	reporter := runner.NewPrometheusWriter(urlFoo, true)

	assert.Nil(t, reporter.GetAllRecords())
	_, err := reporter.ExportAllReportResults()
	assert.NoError(t, err)

	defer gock.Clean()
	gock.New(urlFoo).Put("/metrics/job/api-testing").Reply(http.StatusBadRequest)

	now := time.Now()
	reporter.PutRecord(&runner.ReportRecord{
		Group:     "foo",
		Name:      "bar",
		BeginTime: now,
		EndTime:   now.Add(time.Second * 3),
	})
}
