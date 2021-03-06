package main_test

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"code.cloudfoundry.org/archiver/compressor"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Main", func() {
	var (
		cmdArgs       []string
		stdin         string
		session       *gexec.Session
		offendingText = `
			words
			AKIASOMEMORETEXTHERE
			words
		`
		offendingDiff = `
diff --git a/spec/integration/git-secrets-pattern-tests.txt b/spec/integration/git-secrets-pattern-tests.txt
index 940393e..fa5a232 100644
--- a/spec/integration/git-secrets-pattern-tests.txt
+++ b/spec/integration/git-secrets-pattern-tests.txt
@@ -28,7 +28,7 @@ header line goes here
+AKIAJDHEYSPVNSHFKSMS

 ## Suspicious Variable Names
`
	)

	BeforeEach(func() {
		stdin = ""
		cmdArgs = []string{}
	})

	JustBeforeEach(func() {
		finalArgs := append([]string{"scan"}, cmdArgs...)
		cmd := exec.Command(cliPath, finalArgs...)
		if stdin != "" {
			cmd.Stdin = strings.NewReader(stdin)
		}

		var err error
		session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).ToNot(HaveOccurred())
	})

	ItTellsPeopleHowToRemoveTheirCredentials := func() {
		It("tells people how to add example credentials to their tests or documentation", func() {
			Eventually(session.Out).Should(gbytes.Say("fake"))
		})

		It("tells people how to skip git hooks running for other false positives", func() {
			Eventually(session.Out).Should(gbytes.Say("-n"))
		})

		It("tells people how to reach us", func() {
			Eventually(session.Out).Should(gbytes.Say("Slack channel"))
		})
	}

	ItShowsTheCredentialInTheOutput := func(expectedCredential string) {
		Context("shows actual credential if show-suspected-credentials flag is set", func() {
			BeforeEach(func() {
				cmdArgs = append(cmdArgs, "--show-suspected-credentials")
			})

			It("shows credentials", func() {
				Eventually(session.Out).Should(gbytes.Say(expectedCredential))
			})
		})
	}

	ItTellsPeopleToUpdateIfTheBinaryIsOld := func() {
		Context("when the executable is not over 2 weeks old", func() {
			It("does not display a warning telling the user to upgrade", func() {
				Consistently(session.Err).ShouldNot(gbytes.Say("cred-alert-cli update"))
			})
		})

		Context("when the executable is over 2 weeks old", func() {
			JustBeforeEach(func() {
				cmdArgs = append([]string{"scan"}, cmdArgs...)
				cmd := exec.Command(oldCliPath, cmdArgs...)
				if stdin != "" {
					cmd.Stdin = strings.NewReader(stdin)
				}

				var err error
				session, err = gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
				Expect(err).ToNot(HaveOccurred())
			})

			It("displays a warning telling the user to upgrade", func() {
				Eventually(session.Err).Should(gbytes.Say("cred-alert-cli update"))
			})

			It("does not stop scanning as normal", func() {
				Eventually(session.Out).Should(gbytes.Say("[CRED]"))
				Eventually(session).Should(gexec.Exit(3))
			})
		})
	}

	Context("when given content on stdin", func() {
		BeforeEach(func() {
			stdin = offendingText
		})

		It("scans stdin", func() {
			Eventually(session.Out).Should(gbytes.Say("[CRED]"))
			Eventually(session.Out).Should(gbytes.Say("STDIN"))
		})

		It("exits with status 3", func() {
			Eventually(session).Should(gexec.Exit(3))
		})

		ItTellsPeopleHowToRemoveTheirCredentials()
		ItTellsPeopleToUpdateIfTheBinaryIsOld()
		ItShowsTheCredentialInTheOutput("AKIASOMEMORETEXTHERE")

		Context("when given a --diff flag", func() {
			BeforeEach(func() {
				cmdArgs = []string{"--diff"}
				stdin = offendingDiff
			})

			It("scans the diff", func() {
				Eventually(session.Out).Should(gbytes.Say("spec/integration/git-secrets-pattern-tests.txt:28"))
			})

			ItShowsTheCredentialInTheOutput("AKIAJDHEYSPVNSHFKSMS")
			ItTellsPeopleHowToRemoveTheirCredentials()
		})
	})

	Context("when given regexp flag with a value", func() {
		BeforeEach(func() {
			cmdArgs = []string{
				"--diff",
				"--regexp=random",
			}
			stdin = `
diff --git a/spec/integration/git-secrets-pattern-tests.txt b/spec/integration/git-secrets-pattern-tests.txt
index 940393e..fa5a232 100644
--- a/spec/integration/git-secrets-pattern-tests.txt
+++ b/spec/integration/git-secrets-pattern-tests.txt
@@ -28,7 +28,7 @@ header line goes here
+randomunsuspectedthing

 ## Suspicious Variable Names
`
		})

		It("uses the given regexp pattern", func() {
			Eventually(session.Out).Should(gbytes.Say("[CRED]"))
		})
		Context("when regex-file flags is set", func() {
			BeforeEach(func() {
				cmdArgs = append(cmdArgs, "--regexp-file=some-non-existing-file")
			})

			It("prints warning message", func() {
				Eventually(session.Out).Should(gbytes.Say("[WARN]"))
			})

			It("uses the given regexp pattern", func() {
				Eventually(session.Out).Should(gbytes.Say("[CRED]"))
			})
		})

	})

	Context("when given regex-file flag and the file reads successfully", func() {
		var (
			tmpFile *os.File
			err     error
		)

		BeforeEach(func() {
			tmpFile, err = ioutil.TempFile("", "tmp-file")
			Expect(err).NotTo(HaveOccurred())

			regexpContent := `this-does-not-match
another-pattern
does-not-match`

			err = ioutil.WriteFile(tmpFile.Name(), []byte(regexpContent), 0644)
			Expect(err).NotTo(HaveOccurred())

			cmdArgs = []string{
				fmt.Sprintf("--regexp-file=%s", tmpFile.Name()),
				"--diff",
				"--show-suspected-credentials",
			}
		})

		AfterEach(func() {
			err := tmpFile.Close()
			Expect(err).NotTo(HaveOccurred())
			os.RemoveAll(tmpFile.Name())
		})

		Context("uses the regex", func() {
			Context("and there are no matches", func() {
				BeforeEach(func() {
					stdin = offendingDiff
				})

				It("returns not match", func() {
					Consistently(session.Out).ShouldNot(gbytes.Say("[CRED]"))
				})
			})

			Context("and multiple regex matches", func() {
				BeforeEach(func() {
					stdin = `
diff --git a/spec/integration/git-secrets-pattern-tests.txt b/spec/integration/git-secrets-pattern-tests.txt
index 940393e..fa5a232 100644
--- a/spec/integration/git-secrets-pattern-tests.txt
+++ b/spec/integration/git-secrets-pattern-tests.txt
@@ -28,7 +28,7 @@ header line goes here
+this-does-not-match
+another-pattern
+pattern-another

 ## Suspicious Variable Names
`
				})

				It("scans the diff", func() {
					Eventually(session.Out).Should(gbytes.Say("[CRED]"))
				})
			})
		})
	})

	Context("when given a file flag", func() {
		var tmpFile *os.File

		BeforeEach(func() {
			var err error
			tmpFile, err = ioutil.TempFile("", "cli-main-test")
			Expect(err).NotTo(HaveOccurred())
			defer tmpFile.Close()

			ioutil.WriteFile(tmpFile.Name(), []byte(offendingText), os.ModePerm)

			cmdArgs = []string{"-f", tmpFile.Name()}
		})

		AfterEach(func() {
			os.RemoveAll(tmpFile.Name())
		})

		ItTellsPeopleHowToRemoveTheirCredentials()
		ItTellsPeopleToUpdateIfTheBinaryIsOld()

		It("scans the file", func() {
			Eventually(session.Out).Should(gbytes.Say("[CRED]"))
		})

		It("exits with status 3", func() {
			Eventually(session).Should(gexec.Exit(3))
		})

		Context("shows actual credential if show-suspected-credentials flag is set", func() {
			BeforeEach(func() {
				cmdArgs = append(cmdArgs, "--show-suspected-credentials")
			})

			It("shows credentials", func() {
				Eventually(session.Out).Should(gbytes.Say("AKIASOMEMORETEXTHERE"))
			})
		})

		var ItShowsHowLongInflationTook = func() {
			It("shows how long the inflating took", func() {
				Eventually(session.Out).Should(gbytes.Say(`Time taken \(inflating\):`))
			})
		}

		Context("when the file is a folder", func() {
			var (
				inDir, outDir string
			)

			AfterEach(func() {
				os.RemoveAll(inDir)
				os.RemoveAll(outDir)
			})

			Context("when given a folder", func() {
				BeforeEach(func() {
					var err error
					inDir, err = ioutil.TempDir("", "folder-in")
					Expect(err).NotTo(HaveOccurred())

					err = ioutil.WriteFile(path.Join(inDir, "file1"), []byte(offendingText), 0644)
					Expect(err).NotTo(HaveOccurred())

					cmdArgs = []string{"-f", inDir}
				})

				It("scans each text file in the folder", func() {
					Eventually(session.Out).Should(gbytes.Say("[CRED]"))
				})

				ItTellsPeopleHowToRemoveTheirCredentials()
				ItTellsPeopleToUpdateIfTheBinaryIsOld()
			})
		})

		Context("when the file is a zip file", func() {
			var (
				inDir, outDir, zipFilePath string
			)

			BeforeEach(func() {
				var err error
				inDir, err = ioutil.TempDir("", "zipper-unzip-in")
				Expect(err).NotTo(HaveOccurred())

				err = ioutil.WriteFile(path.Join(inDir, "file1"), []byte(offendingText), 0644)
				Expect(err).NotTo(HaveOccurred())

				outDir, err = ioutil.TempDir("", "zipper-unzip-out")
				Expect(err).NotTo(HaveOccurred())

				zipFilePath = path.Join(outDir, "out.zip")
				err = zipit(inDir, zipFilePath, "")
				Expect(err).NotTo(HaveOccurred())
			})

			AfterEach(func() {
				os.RemoveAll(inDir)
				os.RemoveAll(outDir)
			})

			Context("when given a zip without prefix bytes", func() {
				BeforeEach(func() {
					cmdArgs = []string{"-f", zipFilePath}
				})

				It("scans each text file in the zip", func() {
					Eventually(session.Out).Should(gbytes.Say("[CRED]"))
				})

				ItShowsHowLongInflationTook()
			})
		})

		Context("when the file is a tar file", func() {
			var (
				inDir, outDir string
			)

			BeforeEach(func() {
				var err error
				inDir, err = ioutil.TempDir("", "tar-in")
				Expect(err).NotTo(HaveOccurred())

				err = ioutil.WriteFile(path.Join(inDir, "file1"), []byte(offendingText), 0664)
				Expect(err).NotTo(HaveOccurred())

				outDir, err = ioutil.TempDir("", "tar-out")
				Expect(err).NotTo(HaveOccurred())

				tarFilePath := path.Join(outDir, "out.tar")
				tarFile, err := os.Create(tarFilePath)
				Expect(err).NotTo(HaveOccurred())
				defer tarFile.Close()

				err = compressor.WriteTar(inDir, tarFile)
				Expect(err).NotTo(HaveOccurred())

				cmdArgs = []string{"-f", tarFilePath}
			})

			AfterEach(func() {
				os.RemoveAll(inDir)
				os.RemoveAll(outDir)
			})

			It("scans each text file in the tar", func() {
				Eventually(session.Out).Should(gbytes.Say("[CRED]"))
			})

			ItShowsHowLongInflationTook()
			ItShowsTheCredentialInTheOutput("AKIASOMEMORETEXTHERE")
		})

		Context("when the file is a gzipped tar file", func() {
			var (
				inDir, outDir string
			)

			BeforeEach(func() {
				var err error
				inDir, err = ioutil.TempDir("", "tar-in")
				Expect(err).NotTo(HaveOccurred())

				err = ioutil.WriteFile(path.Join(inDir, "file1"), []byte(offendingText), 0664)
				Expect(err).NotTo(HaveOccurred())

				outDir, err = ioutil.TempDir("", "tar-out")
				Expect(err).NotTo(HaveOccurred())

				tarFilePath := path.Join(outDir, "out.tar")

				c := compressor.NewTgz()
				err = c.Compress(inDir, tarFilePath)
				Expect(err).NotTo(HaveOccurred())

				cmdArgs = []string{"-f", tarFilePath}
			})

			AfterEach(func() {
				os.RemoveAll(inDir)
				os.RemoveAll(outDir)
			})

			It("scans each text file in the tar", func() {
				Eventually(session.Out).Should(gbytes.Say("[CRED]"))
			})

			ItShowsHowLongInflationTook()
		})
	})

	Context("When no credentials are found", func() {
		var politeText = `
			words
			NotACredential
			words
		`
		var tmpFile *os.File

		Context("when given content on stdin", func() {
			BeforeEach(func() {
				cmdArgs = []string{}
				stdin = politeText
			})
			It("exits with status 0", func() {
				Eventually(session).Should(gexec.Exit(0))
			})
		})

		Context("when given a file flag", func() {
			BeforeEach(func() {
				var err error
				tmpFile, err = ioutil.TempFile("", "cli-main-test")
				Expect(err).NotTo(HaveOccurred())
				defer tmpFile.Close()

				ioutil.WriteFile(tmpFile.Name(), []byte(politeText), os.ModePerm)

				cmdArgs = []string{"-f", tmpFile.Name()}
			})

			AfterEach(func() {
				os.RemoveAll(tmpFile.Name())
			})

			It("exits with status 0", func() {
				Eventually(session).Should(gexec.Exit(0))
			})
		})
	})
})

// Thanks to Svett Ralchev
// http://blog.ralch.com/tutorial/golang-working-with-zip/
func zipit(source, target, prefix string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	if prefix != "" {
		_, err = io.WriteString(zipfile, prefix)
		if err != nil {
			return err
		}
	}

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		relpath, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		header.Name = strings.TrimPrefix(relpath, source)

		if info.IsDir() {
			header.Name += string(os.PathSeparator)
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
