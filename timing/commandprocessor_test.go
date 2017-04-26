package timing_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gitlab.com/yaotsu/core"
	"gitlab.com/yaotsu/gcn3/kernels"
	"gitlab.com/yaotsu/gcn3/timing"
)

var _ = Describe("CommandProcessor", func() {

	var (
		driver           *core.MockComponent
		dispatcher       *core.MockComponent
		commandProcessor *timing.CommandProcessor
		connection       *core.DirectConnection
	)

	BeforeEach(func() {
		connection = core.NewDirectConnection()

		driver = core.NewMockComponent("dispatcher")
		driver.AddPort("ToGPU")
		dispatcher = core.NewMockComponent("dispatcher")
		dispatcher.AddPort("ToCommandProcessor")
		commandProcessor = timing.NewCommandProcessor("commandProcessor")

		commandProcessor.Dispatcher = dispatcher

		core.PlugIn(dispatcher, "ToCommandProcessor", connection)
		core.PlugIn(commandProcessor, "ToDispatcher", connection)
		core.PlugIn(driver, "ToGPU", connection)
	})

	It("should forward kernel launching request to Dispatcher", func() {
		req := kernels.NewLaunchKernelReq()
		req.SetSrc(driver)
		req.SetDst(commandProcessor)

		dispatcher.ToReceiveReq(req, nil)

		commandProcessor.Recv(req)

		Expect(dispatcher.AllReqReceived()).To(BeTrue())
	})
})
