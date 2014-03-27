package download_step_test

import (
	"errors"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry-incubator/runtime-schema/models"
	steno "github.com/cloudfoundry/gosteno"
	"github.com/vito/gordon/fake_gordon"

	"github.com/cloudfoundry-incubator/executor/downloader/fake_downloader"
	"github.com/cloudfoundry-incubator/executor/extractor/fake_extractor"
	"github.com/cloudfoundry-incubator/executor/linux_plugin"
	"github.com/cloudfoundry-incubator/executor/sequence"
	. "github.com/cloudfoundry-incubator/executor/steps/download_step"
)

var _ = Describe("DownloadAction", func() {
	var step sequence.Step
	var result chan error

	var downloadAction models.DownloadAction
	var downloader *fake_downloader.FakeDownloader
	var extractor *fake_extractor.FakeExtractor
	var tempDir string
	var backendPlugin *linux_plugin.LinuxPlugin
	var wardenClient *fake_gordon.FakeGordon
	var logger *steno.Logger

	BeforeEach(func() {
		var err error

		result = make(chan error)

		downloadAction = models.DownloadAction{
			From:    "http://mr_jones",
			To:      "/tmp/Antarctica",
			Extract: false,
		}

		downloader = &fake_downloader.FakeDownloader{}
		extractor = &fake_extractor.FakeExtractor{}

		tempDir, err = ioutil.TempDir("", "download-action-tmpdir")
		Ω(err).ShouldNot(HaveOccurred())

		wardenClient = fake_gordon.New()

		backendPlugin = linux_plugin.New()

		logger = steno.NewLogger("test-logger")
	})

	JustBeforeEach(func() {
		step = New(
			"some-container-handle",
			downloadAction,
			downloader,
			extractor,
			tempDir,
			backendPlugin,
			wardenClient,
			logger,
		)
	})

	Describe("Perform", func() {
		It("downloads the file from the given URL", func() {
			err := step.Perform()
			Ω(err).ShouldNot(HaveOccurred())

			Ω(downloader.DownloadedUrls).ShouldNot(BeEmpty())
			Ω(downloader.DownloadedUrls[0].Host).To(ContainSubstring("mr_jones"))
		})

		It("places the file in the container", func() {
			err := step.Perform()
			Ω(err).ShouldNot(HaveOccurred())

			Ω(wardenClient.ThingsCopiedIn()).ShouldNot(BeEmpty())

			copiedFile := wardenClient.ThingsCopiedIn()[0]
			Ω(copiedFile.Handle).To(Equal("some-container-handle"))
			Ω(copiedFile.Dst).To(Equal("/tmp/Antarctica"))
		})

		Context("when there is an error copying the file in", func() {
			disaster := errors.New("oh no!")

			BeforeEach(func() {
				wardenClient.SetCopyInErr(disaster)
			})

			It("sends back the error", func() {
				err := step.Perform()
				Ω(err).Should(Equal(disaster))
			})
		})

		Context("when extract is true", func() {
			BeforeEach(func() {
				downloadAction = models.DownloadAction{
					From:    "http://mr_jones.zip",
					To:      "/tmp/Antarctica",
					Extract: true,
				}
			})

			It("uses the specified extractor", func() {
				step.Perform()
				Ω(extractor.ExtractedFilePaths).ShouldNot(BeEmpty())
				Ω(extractor.ExtractedFilePaths[0]).To(ContainSubstring(tempDir))
			})

			It("places the file in the container under the destination", func() {
				err := step.Perform()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(wardenClient.ThingsCopiedIn()).ShouldNot(BeEmpty())
				copiedFile := wardenClient.ThingsCopiedIn()[0]
				Ω(copiedFile.Handle).To(Equal("some-container-handle"))
				Ω(copiedFile.Src).To(ContainSubstring(tempDir))
				Ω(copiedFile.Dst).To(Equal("/tmp/Antarctica/"))
			})
		})
	})
})
