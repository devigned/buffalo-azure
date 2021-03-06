package eventgrid_test

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/Azure/buffalo-azure/sdk/eventgrid"
	"github.com/gobuffalo/buffalo"
)

func ExampleTypeDispatchSubscriber_Receive() {
	var mySubscriber eventgrid.Subscriber
	mySubscriber = eventgrid.BaseSubscriber{}
	mySubscriber = eventgrid.NewTypeDispatchSubscriber(mySubscriber).Bind("Microsoft.Storage.BlobCreated", func(c buffalo.Context, e eventgrid.Event) (err error) {
		_, err = fmt.Println(e.ID)
		return
	})

	req, err := http.NewRequest(http.MethodPost, "localhost", bytes.NewReader([]byte(`[{
	"topic": "/subscriptions/{subscription-id}/resourceGroups/Storage/providers/Microsoft.Storage/storageAccounts/xstoretestaccount",
	"subject": "/blobServices/default/containers/oc2d2817345i200097container/blobs/oc2d2817345i20002296blob",
	"eventType": "Microsoft.Storage.BlobCreated",
	"eventTime": "2017-06-26T18:41:00.9584103Z",
	"id": "831e1650-001e-001b-66ab-eeb76e069631",
	"data": {
		"api": "PutBlockList",
		"clientRequestId": "6d79dbfb-0e37-4fc4-981f-442c9ca65760",
		"requestId": "831e1650-001e-001b-66ab-eeb76e000000",
		"eTag": "0x8D4BCC2E4835CD0",
		"contentType": "application/octet-stream",
		"contentLength": 524288,
		"blobType": "BlockBlob",
		"url": "https://oc2d2817345i60006.blob.core.windows.net/oc2d2817345i200097container/oc2d2817345i20002296blob",
		"sequencer": "00000000000004420000000000028963",
		"storageDiagnostics": {
		"batchId": "b68529f3-68cd-4744-baa4-3c0498ec19f0"
		}
	},
	"dataVersion": "",
	"metadataVersion": "1"
}]`)))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	ctx := NewMockContext(req)

	err = mySubscriber.Receive(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Output: 831e1650-001e-001b-66ab-eeb76e069631
}
