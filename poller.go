package endure

import "go.uber.org/zap"

// poll is used to poll the errors from the vertex
// and exit from it
func (e *Endure) poll(r *result) {
	rr := r
	go func(res *result) {
		for {
			select {
			// error
			case err := <-res.errCh:
				if err != nil {
					// log error message
					e.logger.Error("vertex got an error", zap.String("vertex id", res.vertexID), zap.Error(err))

					// set the error
					res.err = err

					// send handleErrorCh signal
					e.handleErrorCh <- res
				}
			// exit from the goroutine
			case n := <-res.signal:
				if n.stop {
					e.mutex.Lock()
					e.logger.Info("vertex got exit signal", zap.String("vertex id", res.vertexID))
					err := e.internalStop(res.vertexID)
					if err != nil {
						e.logger.Error("error during exit signal", zap.String("error while stopping the vertex:", res.vertexID), zap.Error(err))
						e.mutex.Unlock()
					}
					e.mutex.Unlock()
					return
				}
				return
			}
		}
	}(rr)
}

func (e *Endure) startMainThread() {
	/*
		Main thread is the main Endure unit of work
		It used to handle errors from vertices, notify user about result, re-calculating graph according to failed vertices and sending internal_stop signals
	*/
	go func() {
		for {
			select {
			// failed Vertex
			case res, ok := <-e.handleErrorCh:
				// lock the handleErrorCh processing
				if !ok {
					e.logger.Debug("handle error channel was closed")
					return
				}

				e.logger.Debug("processing error in the main thread", zap.String("vertex id", res.vertexID))
				if e.retry {
					// TODO handle error from the retry handler
					e.retryHandler(res)
				} else {
					e.logger.Info("retry is turned off, sending exit signal to every vertex in the graph")
					// send exit signal to whole graph
					e.sendStopSignal(e.graph.Vertices)
					e.sendResultToUser(res)
				}
			}
		}
	}()
}
