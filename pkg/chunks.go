// The Chunks endpoint works together with the Memfault Firmware SDK.
// It provides a streamlined way of getting arbitrary data (coredumps, events, heartbeats, etc.) out of devices and into Memfault.

package memfault

// // GetMe Return information about the logged in User
// func (sdk mfSDK) UploadChunk(chunk string) (string, error) {
//     endpoint := "api/v0/chunks/{{device_serial}}"
//     url := fmt.Sprintf("%s/%s", sdk.chunksURL, endpoint)
//     payload := strings.NewReader(chunk)

//     req, err := http.NewRequest(http.MethodPost, url, payload)
//     if err != nil {
//         return "", err
//     }
// 	req.Header.Add("Memfault-Project-Key", "")
// 	req.Header.Add("Content-Type", "application/octet-stream")
//     resp, err := sdk.makeRequest(req)
//     if err != nil {
//         return "", err
//     }
//     return string(resp), err

// }
