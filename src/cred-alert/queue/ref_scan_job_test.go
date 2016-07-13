package queue_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"archive/zip"
	"bytes"
	"log"
	"net/http"
	"net/url"

	"github.com/onsi/gomega/ghttp"
	"github.com/pivotal-golang/lager"
	"github.com/pivotal-golang/lager/lagertest"

	"cred-alert/github/githubfakes"
	"cred-alert/notifications/notificationsfakes"
	"cred-alert/queue"
	"cred-alert/scanners"
	"cred-alert/sniff"
)

var _ = Describe("RefScan Job", func() {
	var (
		client *githubfakes.FakeClient

		logger *lagertest.TestLogger

		job       *queue.RefScanJob
		server    *ghttp.Server
		sniffFunc sniff.SniffFunc
		plan      queue.RefScanPlan
		notifier  *notificationsfakes.FakeNotifier
	)

	owner := "repo-owner"
	repo := "repo-name"
	repoFullName := owner + "/" + repo
	ref := "reference"

	BeforeEach(func() {
		server = ghttp.NewServer()
		plan = queue.RefScanPlan{
			Owner:      owner,
			Repository: repo,
			Ref:        ref,
		}

		client = &githubfakes.FakeClient{}
		logger = lagertest.NewTestLogger("ref-scan-job")
		notifier = &notificationsfakes.FakeNotifier{}
	})

	JustBeforeEach(func() {
		job = queue.NewRefScanJob(plan, client, sniffFunc, notifier)
	})

	Describe("Run", func() {
		wasSniffed := false
		filePath := "some/file/path"
		fileContent := "content"

		BeforeEach(func() {
			serverUrl, _ := url.Parse(server.URL())
			client.ArchiveLinkReturns(serverUrl, nil)
			someZip := createZip()
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/"),
					ghttp.RespondWith(http.StatusOK, someZip.Bytes(), http.Header{}),
				),
			)
			sniffFunc = func(lgr lager.Logger, scanner sniff.Scanner, handleViolation func(scanners.Line)) {
				wasSniffed = true
				Expect(lgr).To(Equal(logger))
				handleViolation(scanners.Line{
					Path:       filePath,
					LineNumber: 1,
					Content:    fileContent,
				})
			}
		})

		It("fetches a link from GitHub", func() {
			err := job.Run(logger)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ArchiveLinkCallCount()).To(Equal(1))
			lgr, owner, repo := client.ArchiveLinkArgsForCall(0)
			Expect(lgr).To(Equal(logger))
			Expect(owner).To(Equal("repo-owner"))
			Expect(repo).To(Equal("repo-name"))
		})

		It("Scans the archive", func() {

			err := job.Run(logger)

			Expect(err).NotTo(HaveOccurred())
			Expect(wasSniffed).To(BeTrue())
		})

		It("sends a notification when it finds a match", func() {
			err := job.Run(logger)
			Expect(err).NotTo(HaveOccurred())

			Expect(notifier.SendNotificationCallCount()).To(Equal(3))
			lgr, repository, sha, line := notifier.SendNotificationArgsForCall(0)

			Expect(lgr).To(Equal(logger))
			Expect(repository).To(Equal(repoFullName))
			Expect(sha).To(Equal(ref))
			Expect(line).To(Equal(scanners.Line{
				Path:       filePath,
				LineNumber: 1,
				Content:    fileContent,
			}))
		})

		It("emits violations", func() {
			// TODO: finish him!
		})
	})
})

func createZip() *bytes.Buffer {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", `lolz`},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "Get animal handling licence.\nWrite more examples."},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	w.Close()

	return buf
}
