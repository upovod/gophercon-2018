func TestPing(t *testing.T) {
        t.Logf("Running test case: %s", name)
        storage := &mockStorager{}
        storage.
            On("Cmd", "INCR", []interface{}{"ping:count"}).
            Return(&redis.Resp{
                Err: nil,
            }).
            Once()
        h := &Handler{
            db: storage,
        }

        response, err := h.Ping()

        assert.Equal(t, nil, err)
        assert.Equal(t, "pong", response)
        storage.AssertExpectations(t)
    }
}
