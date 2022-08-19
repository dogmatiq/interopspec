package envelopespec_test

import (
	"time"

	. "github.com/dogmatiq/interopspec/envelopespec"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type Envelope", func() {
	Describe("func Validate()", func() {
		var env *Envelope

		BeforeEach(func() {
			createdAt := time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)
			scheduledFor := createdAt.Add(1 * time.Hour)

			env = &Envelope{
				MessageId:     "<id>",
				CausationId:   "<cause>",
				CorrelationId: "<correlation>",
				SourceSite: &Identity{
					Name: "<site-name>",
					Key:  "<site-key>",
				},
				SourceApplication: &Identity{
					Name: "<app-name>",
					Key:  "<app-key>",
				},
				SourceHandler: &Identity{
					Name: "<handler-name>",
					Key:  "<handler-key>",
				},
				SourceInstanceId: "<instance>",
				CreatedAt:        createdAt.Format(time.RFC3339Nano),
				ScheduledFor:     scheduledFor.Format(time.RFC3339Nano),
				Description:      "<description>",
				PortableName:     "<portable name>",
				MediaType:        "<media type>",
				Data:             []byte("<data>"),
			}
		})

		It("does not return an error if the envelope is well-formed", func() {
			err := env.Validate()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("does not return an error if the site is nil", func() {
			env.SourceSite = nil

			err := env.Validate()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("returns an error if the message ID is empty", func() {
			env.MessageId = ""

			err := env.Validate()
			Expect(err).Should(HaveOccurred())
		})

		It("returns an error if the causation ID is empty", func() {
			env.CausationId = ""

			err := env.Validate()
			Expect(err).Should(HaveOccurred())
		})

		It("returns an error if the correlation ID is empty", func() {
			env.CorrelationId = ""

			err := env.Validate()
			Expect(err).Should(HaveOccurred())
		})

		It("returns an error if the source site name is empty", func() {
			env.SourceSite.Name = ""

			err := env.Validate()
			Expect(err).To(MatchError("site identity is invalid: identity name must not be empty"))
		})

		It("returns an error if the source site key is empty", func() {
			env.SourceSite.Key = ""

			err := env.Validate()
			Expect(err).To(MatchError("site identity is invalid: identity key must not be empty"))
		})

		It("returns an error if the source app name is empty", func() {
			env.SourceApplication.Name = ""

			err := env.Validate()
			Expect(err).To(MatchError("application identity is invalid: identity name must not be empty"))
		})

		It("returns an error if the source app key is empty", func() {
			env.SourceApplication.Key = ""

			err := env.Validate()
			Expect(err).To(MatchError("application identity is invalid: identity key must not be empty"))
		})

		It("returns an error if the source handler name is empty", func() {
			env.SourceHandler.Name = ""

			err := env.Validate()
			Expect(err).To(MatchError("handler identity is invalid: identity name must not be empty"))
		})

		It("returns an error if the source handler key is empty", func() {
			env.SourceHandler.Key = ""

			err := env.Validate()
			Expect(err).To(MatchError("handler identity is invalid: identity key must not be empty"))
		})

		It("returns an error if the source handler is empty but the instance ID is set", func() {
			env.SourceHandler = nil

			err := env.Validate()
			Expect(err).To(MatchError("source instance ID must not be specified without providing a handler identity"))
		})

		When("there is no source handler", func() {
			BeforeEach(func() {
				env.SourceHandler = nil
				env.SourceInstanceId = ""
			})

			It("returns an error if the message is a timeout", func() {
				err := env.Validate()
				Expect(err).To(MatchError("scheduled-for time must not be specified without a providing source handler and instance ID"))
			})

			It("does not return an error if the message is not a timeout", func() {
				env.ScheduledFor = ""

				err := env.Validate()
				Expect(err).ShouldNot(HaveOccurred())
			})
		})

		It("does not return an error if the message description is empty", func() {
			env.Description = ""

			err := env.Validate()
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("returns an error if the created-at timestamp is empty", func() {
			env.CreatedAt = ""

			err := env.Validate()
			Expect(err).To(MatchError("created-at time must not be empty"))
		})

		It("returns an error if the portable name is empty", func() {
			env.PortableName = ""

			err := env.Validate()
			Expect(err).To(MatchError("portable name must not be empty"))
		})

		It("returns an error if the media-type is empty", func() {
			env.MediaType = ""

			err := env.Validate()
			Expect(err).To(MatchError("media-type must not be empty"))
		})

		It("does not return an error if the message data is empty", func() {
			env.Data = nil

			err := env.Validate()
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
