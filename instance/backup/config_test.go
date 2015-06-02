package backup_test

import (
	"path"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/cf-redis-broker/instance/backup"
	"github.com/pivotal-cf/cf-redis-broker/redisconf"
)

var _ = Describe("Config", func() {
	Describe("Load", func() {
		Context("When the file does not exist", func() {
			It("returns an error", func() {
				_, err := redisconf.Load("/this/is/an/invalid/path")
				Expect(err.Error()).To(Equal("open /this/is/an/invalid/path: no such file or directory"))
			})
		})

		Context("When the file is successfully loaded", func() {
			var config *backup.Config

			BeforeEach(func() {
				path, err := filepath.Abs(path.Join("assets", "backup.yml"))
				Expect(err).ToNot(HaveOccurred())

				config, err = backup.LoadConfig(path)
				Expect(err).ToNot(HaveOccurred())
			})

			It("Has the correct endpoint_url", func() {
				Expect(config.S3Config.EndpointUrl).To(Equal("endpointurl.com"))
			})

			It("Has the correct bucket_name", func() {
				Expect(config.S3Config.BucketName).To(Equal("some-bucket-name"))
			})

			It("Has the correct access_key_id", func() {
				Expect(config.S3Config.AccessKeyId).To(Equal("some-access-key-id"))
			})

			It("Has the correct secret_access_key", func() {
				Expect(config.S3Config.SecretAccessKey).To(Equal("secret-access-key"))
			})

			It("Has the correct path", func() {
				Expect(config.S3Config.Path).To(Equal("some-s3-path"))
			})

			It("Has the correct snapshot_timeout", func() {
				Expect(config.SnapshotTimeout).To(Equal("10s"))
			})

			It("Has the correct redis_data_root", func() {
				Expect(config.RedisDataRoot).To(Equal("/the/path/to/redis/data"))
			})

			It("Has the correct redis_config_filename", func() {
				Expect(config.RedisConfigFilename).To(Equal("/the/path/to/the/redis/config"))
			})

			It("Has the correct plan_name", func() {
				Expect(config.PlanName).To(Equal("plan-name"))
			})

			It("has the correct log file path", func() {
				Expect(config.LogFilepath).To(Equal("/log/file/path"))
			})

			It("has the correct aws cli path", func() {
				Expect(config.AwsCLIPath).To(Equal("/path/to/aws-cli"))
			})
		})
	})
})
