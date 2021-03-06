package client

//
// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"testing"
// 	"time"
//
// 	"github.com/stretchr/testify/assert"
// )
//
// func TestQueryStreamRows(t *testing.T) {
// 	t.Run("when the Next method is called", func(t *testing.T) {
// 		var err error
// 		b := &bytes.Buffer{}
// 		enc := json.NewEncoder(b)
// 		rows := [][]interface{}{
// 			{
// 				float64(1), "1", float64(11), "alice", "home",
// 			},
// 			{
// 				float64(2), "2", float64(11), "bob", "home",
// 			},
// 			{
// 				float64(3), "3", float64(11), "james", "home",
// 			},
// 		}
// 		err = enc.Encode(rows[0])
// 		assert.NoError(t, err)
// 		err = enc.Encode(rows[1])
// 		assert.NoError(t, err)
// 		d := json.NewDecoder(b)
// 		r := &QueryStreamRows{
// 			ctx: context.Background(),
// 			dec: d,
// 			columns: columns{
// 				count: 5,
// 				names: []string{"id", "seq", "orgID", "firstName", "location"},
// 			},
// 		}
// 		dest := make([]interface{}, r.columns.count)
// 		err = r.Next(dest)
// 		assert.NoError(t, err)
// 		assert.EqualValues(t, rows[0], dest)
// 		err = r.Next(dest)
// 		assert.NoError(t, err)
// 		assert.Equal(t, rows[1], dest)
// 		go func() {
// 			time.Sleep(time.Second * 1)
// 			enc.Encode(rows[2])
// 		}()
// 		err = r.Next(dest)
// 		assert.NoError(t, err)
// 		assert.Equal(t, rows[2], dest, "it should work async")
// 	})
// }
