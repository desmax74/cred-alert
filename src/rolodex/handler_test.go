package rolodex_test

import (
	"errors"
	"net"

	"code.cloudfoundry.org/lager/lagertest"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"cred-alert/metrics/metricsfakes"
	"red/redpb"
	"rolodex"
	"rolodex/rolodexfakes"
	"rolodex/rolodexpb"
)

var _ = Describe("Handler", func() {
	var (
		grpcServer *grpc.Server
		listener   net.Listener
		client     rolodexpb.RolodexClient
		connection *grpc.ClientConn
		logger     *lagertest.TestLogger
		emitter    *metricsfakes.FakeEmitter

		teamRepo *rolodexfakes.FakeTeamRepository
	)

	BeforeEach(func() {
		var err error

		logger = lagertest.NewTestLogger("handler")
		emitter = &metricsfakes.FakeEmitter{}
		fakeCounter := &metricsfakes.FakeCounter{}
		emitter.CounterReturns(fakeCounter)

		teamRepo = &rolodexfakes.FakeTeamRepository{}
		handler := rolodex.NewHandler(logger, teamRepo, emitter)

		listener, err = net.Listen("tcp", "127.0.0.1:0")
		Expect(err).NotTo(HaveOccurred())

		grpcServer = grpc.NewServer()
		rolodexpb.RegisterRolodexServer(grpcServer, handler)

		go grpcServer.Serve(listener)

		connection, err = grpc.Dial(
			listener.Addr().String(),
			grpc.WithInsecure(),
		)
		Expect(err).NotTo(HaveOccurred())

		client = rolodexpb.NewRolodexClient(connection)
	})

	AfterEach(func() {
		Expect(connection.Close()).To(Succeed())
		grpcServer.Stop()
	})

	Describe("GetOwners", func() {
		BeforeEach(func() {
			teamRepo.GetOwnersReturns([]rolodex.Team{
				{
					Name: "capi",
				},
				{
					Name: "bosh",
					SlackChannel: rolodex.SlackChannel{
						Team: "cloudfoundry",
						Name: "bosh",
					},
				},
				{
					Name: "concourse",
				},
			}, nil)
		})

		It("gets owners", func() {
			ctx := context.Background()

			resp, err := client.GetOwners(ctx, &rolodexpb.GetOwnersRequest{
				Repository: &redpb.Repository{
					Owner: "cloudfoundry",
					Name:  "cf-release",
				},
			})
			Expect(err).NotTo(HaveOccurred())

			Expect(teamRepo.GetOwnersCallCount()).To(Equal(1))
			searchedRepo := teamRepo.GetOwnersArgsForCall(0)
			Expect(searchedRepo.Name).To(Equal("cf-release"))
			Expect(searchedRepo.Owner).To(Equal("cloudfoundry"))

			teams := resp.GetTeams()
			Expect(teams).To(HaveLen(3))

			bosh := teams[1]

			Expect(bosh).To(Equal(&rolodexpb.Team{
				Name: "bosh",
				SlackChannel: &rolodexpb.SlackChannel{
					Team: "cloudfoundry",
					Name: "bosh",
				},
			}))

			concourse := teams[2]

			Expect(concourse).To(Equal(&rolodexpb.Team{
				Name: "concourse",
			}))
		})

		Context("when the request is incomplete", func() {
			It("returns an error", func() {
				ctx := context.Background()

				_, err := client.GetOwners(ctx, &rolodexpb.GetOwnersRequest{})

				Expect(err).To(MatchError(ContainSubstring("repository owner and/or name may not be empty")))
				Expect(grpc.Code(err)).To(Equal(codes.InvalidArgument))
			})
		})

		Context("when we fail to look up the team", func() {
			BeforeEach(func() {
				teamRepo.GetOwnersReturns(nil, errors.New("disaster"))
			})

			It("returns an error", func() {
				ctx := context.Background()

				_, err := client.GetOwners(ctx, &rolodexpb.GetOwnersRequest{
					Repository: &redpb.Repository{
						Owner: "cloudfoundry",
						Name:  "cf-release",
					},
				})
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when there is no team for that repository", func() {
			BeforeEach(func() {
				teamRepo.GetOwnersReturns([]rolodex.Team{}, nil)
			})

			It("returns an NotFound error", func() {
				ctx := context.Background()

				response, err := client.GetOwners(ctx, &rolodexpb.GetOwnersRequest{
					Repository: &redpb.Repository{
						Owner: "cloudfoundry",
						Name:  "cf-release",
					},
				})

				Expect(err).NotTo(HaveOccurred())
				Expect(response.Teams).To(BeEmpty())
			})
		})
	})
})
